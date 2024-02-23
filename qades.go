package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alecthomas/kong"
)

type (
	// CLIConfig loads command line configuration with defaults.
	CLIConfig struct {
		Addr string `help:"Server listen address." default:"127.0.0.1"`
		Port int    `help:"Server port." default:"8000"`
	}

	Qades struct {
		listener net.Listener
		wg       sync.WaitGroup
		close    chan struct{}
	}
)

func NewServer(cfg *CLIConfig) *Qades {
	q := &Qades{}
	if l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)); err != nil {
		log.Fatal(err)
	} else {
		q.listener = l
	}
	q.wg.Add(1)
	go q.Serve()
	return q
}

func (q *Qades) Serve() {
	defer q.wg.Done()

	for {
		if conn, err := q.listener.Accept(); err != nil {
			select {
			case <-q.close:
				return
			default:
				fmt.Printf("connection error: %v\n", err)
			}
		} else {
			fmt.Println("connection opened")
			q.wg.Add(1)
			go q.echo(conn)
		}

	}
}

func (q *Qades) Stop() {
	close(q.close)
	q.listener.Close()
	q.wg.Wait()
}

func (q *Qades) echo(conn net.Conn) {
	defer q.wg.Done()
	defer conn.Close()

	var nErr *net.OpError

	readBuff := make([]byte, 1024)

	for {
		select {
		// when channel is closed, this will trigger, force the connection closed
		case <-q.close:
			return
		default:
			conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
			if bytes, err := conn.Read(readBuff); err != nil {
				if errors.As(err, &nErr) && nErr.Timeout() {
					continue
				}
				fmt.Printf("read error: %v\n", err)
				return
			} else {
				if bytes == 0 {
					fmt.Println("empty read, exiting")
					return
				}
				fmt.Printf("read data: %v | %s\n", readBuff[:bytes], string(readBuff[:bytes]))
				conn.Write(readBuff[:bytes])
			}
		}

	}

}

func main() {
	// parse CLI config
	cfg := &CLIConfig{}
	_ = kong.Parse(cfg)

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	q := NewServer(cfg)

	<-shutdownSignal

	q.Stop()
}
