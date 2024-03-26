package main

import (
	"bufio"
	"log/slog"
	"net"
	"os"
	"testing"

	"github.com/Central-Texas-Commodore-Users-Group/web2go-bbs/telnet"
	"github.com/Central-Texas-Commodore-Users-Group/web2go-bbs/telnet/options"
)

var tests = []string{"hello", "お早う", "☀️"}

func TestNewServer(t *testing.T) {
	programLevel := new(slog.LevelVar)
	programLevel.Set(slog.LevelDebug)

	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})

	s := NewServer(&CLIConfig{Addr: "127.0.0.1", Port: 8000}, handler)

	if conn, err := net.Dial("tcp", "127.0.0.1:8000"); err != nil {
		t.Errorf("error connecting %v", err)
	} else {
		in := bufio.NewReader(conn)
		readBuff := make([]byte, 1024)

		for _, test := range tests {
			if _, err := conn.Write([]byte(test)); err != nil {
				t.Errorf("error writing to server")
			} else {
				if res, err := in.Read(readBuff); err != nil {
					t.Errorf("error on test %s reading: %v", test, err)
				} else if string(readBuff[:res]) != test {
					t.Errorf("expected receive %s, got %s", test, string(readBuff[:res]))
				}
			}
		}
		conn.Close()
	}
	s.Stop()
}

func TestCommand(t *testing.T) {
	if telnet.IAC.ToString() != "IAC" {
		t.Errorf("expected IAC, got %s", telnet.IAC.ToString())
	}
}

func TestOptions(t *testing.T) {
	if options.TerminalType.ToString() != "Terminal Type" {
		t.Errorf("expected Extended-Options-List, got %s", options.TerminalType.ToString())
	}
}
