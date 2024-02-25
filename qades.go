package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type (
	// CLIConfig loads command line configuration with defaults.
	CLIConfig struct {
		Addr string
		Port int
	}

	Qades struct {
		listener net.Listener
		wg       sync.WaitGroup
		close    chan struct{}
		logger   *slog.Logger
	}
)

func NewServer(cfg *CLIConfig, h slog.Handler) *Qades {
	q := &Qades{logger: slog.New(h), close: make(chan struct{})}
	if l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)); err != nil {
		q.logger.Error("Error creating listener", err)
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
				q.logger.Error("connection error", err)
			}
		} else {
			q.logger.Info("connection opened")
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
				q.logger.Error("read error", err)
				return
			} else {
				if bytes == 0 {
					q.logger.Error("empty read, exiting. is this unnecessary?")
					return
				}
				//slog will marshall the bytes into base64
				q.logger.Info("read data", "bytes", fmt.Sprintf("%v", readBuff[:bytes]), "string", string(readBuff[:bytes]))
				fmt.Printf("read data: %v | %s\n", readBuff[:bytes], string(readBuff[:bytes]))
				conn.Write(readBuff[:bytes])
			}
		}
	}

}

func main() {
	// parse CLI config
	cfg := new(CLIConfig)
	flag.StringVar(&cfg.Addr, "addr", "127.0.0.1", "Server listen address.")
	flag.IntVar(&cfg.Port, "port", 8000, "Server port.")
	flag.Parse()

	programLevel := new(slog.LevelVar)
	programLevel.Set(slog.LevelDebug)

	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel, AddSource: true})

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	q := NewServer(cfg, h)

	<-shutdownSignal
	q.logger.Info("shutting down")

	q.Stop()
}
