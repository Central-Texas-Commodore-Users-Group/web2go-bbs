package telnet

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/Central-Texas-Commodore-Users-Group/web2go-bbs/telnet/options"
)

type (
	Command     uint8
	parserState uint8

	ConnHandler interface {
		HandleByte(uint8)
		ForwardByte(uint8)
		HandleCommand(uint8)
		SendReply([]byte) error
	}

	TelnetHandler struct {
		state       parserState
		currOption  options.Options
		conn        net.Conn
		nextHandler ConnHandler
		logger      *slog.Logger
	}

	EchoHandler struct {
		conn net.Conn
	}
)

func NewTelnetHandler(conn net.Conn, next ConnHandler, logger *slog.Logger) *TelnetHandler {
	return &TelnetHandler{
		state:       handleData,
		conn:        conn,
		nextHandler: next,
		logger:      logger,
	}
}

func NewEchoHandler(conn net.Conn) *EchoHandler {
	return &EchoHandler{
		conn: conn,
	}
}

func (c Command) ToString() string {
	var ret string

	switch c {
	case SE:
		ret = "SE"
	case NOP:
		ret = "NOP"
	case DataMark:
		ret = "Data Mark"
	case Break:
		ret = "Break"
	case InterruptProcess:
		ret = "InterruptProcess"
	case AbortOutput:
		ret = "Abort Output"
	case AreYouThere:
		ret = "Are You There"
	case EraseCharacter:
		ret = "Erase Character"
	case EraseLine:
		ret = "Erase Line"
	case GoAhead:
		ret = "Go Ahead"
	case SB:
		ret = "SB"
	case WILL:
		ret = "Will"
	case WONT:
		ret = "Won't"
	case DO:
		ret = "Do"
	case DONT:
		ret = "Don't"
	case IAC:
		ret = "IAC"
	default:
		ret = "unknown"
	}
	return ret
}

const (
	handleData   parserState = 0x00
	handleIAC    parserState = 0x01
	handleOption parserState = 0x02
	handleWill   parserState = 0x03
	handleWont   parserState = 0x04
	handleDo     parserState = 0x05
	handleDont   parserState = 0x06

	// https://www.rfc-editor.org/rfc/rfc854.html
	SE               Command = 0xF0 // End subnegotiation
	NOP              Command = 0xF1 // No Operation
	DataMark         Command = 0xF2
	Break            Command = 0xF3
	InterruptProcess Command = 0xF4
	AbortOutput      Command = 0xF5
	AreYouThere      Command = 0xF6
	EraseCharacter   Command = 0xF7
	EraseLine        Command = 0xF8
	GoAhead          Command = 0xF9
	SB               Command = 0xFA // Begin Subnegotiation
	// Option Codes
	WILL Command = 0xFB
	WONT Command = 0xFC
	DO   Command = 0xFD
	DONT Command = 0xFE
	IAC  Command = 0xFF
)

func (th *TelnetHandler) SendReply(bytes []byte) error {
	if written, err := th.conn.Write(bytes); err != nil {
		return fmt.Errorf("error when writing bytes: %w", err)
	} else if written != len(bytes) {
		return fmt.Errorf("failed to send full message")
		// TODO: attempt to finish sending message?
	}
	return nil
}

func (th *TelnetHandler) ForwardByte(b uint8) {
	th.nextHandler.HandleByte(b)
}

func (th *TelnetHandler) HandleByte(b uint8) {
	switch th.state {
	case handleData:
		if Command(b) == IAC {
			th.state = handleIAC
		} else {
			th.ForwardByte(b)
		}
	case handleIAC:
		switch Command(b) {
		case NOP:
			th.state = handleData
		case SE:
			fallthrough
		case DataMark:
			fallthrough
		case Break:
			fallthrough
		case InterruptProcess:
			fallthrough
		case AbortOutput:
			fallthrough
		case EraseLine:
			fallthrough
		case EraseCharacter:
			fallthrough
		case GoAhead:
			fallthrough
		case SB:
			th.state = handleOption
		case AreYouThere:
			if err := th.SendReply([]byte{byte(IAC), byte(NOP)}); err != nil {
				th.logger.Error("Unable to reply to AreYouThere", err)
			}
			th.state = handleData
		case WILL:
			th.state = handleWill
		case WONT:
			th.state = handleWont
		case DO:
			th.state = handleDo
		case DONT:
			th.state = handleDont
		case IAC:
			// escaped byte
			th.ForwardByte(b)
			th.state = handleData
		default:
			th.logger.Error("Unknown IAC byte", "IAC", b)
			th.state = handleData
		}
	case handleDo:
		fallthrough
	case handleDont:
		fallthrough
	case handleWill:
		switch b {
		case options.TerminalType.GetByte():
			if err := th.SendReply(options.WillTerminalType()); err != nil {
				th.logger.Error("Failed requesting terminal type list", err)
			}
		default:
			th.logger.Error("Unknown Will Option", "option", b)
		}
		th.state = handleData
	case handleWont:
		fallthrough
	case handleOption:
		// TODO: options may be interleaved, so this will not work as intended. Need to fix how an options handler
		// is managed.
		if th.currOption != nil {
			if done := th.currOption.HandleByte(b, th); done {
				th.currOption = nil
				th.state = handleData
			}
		} else {
			switch b {
			case options.TerminalType.GetByte():
				th.currOption = options.TerminalType
			default:
				th.logger.Error("Unknown SB Option", "option", b)
			}
		}
	}
}

func (eh *EchoHandler) HandleByte(b uint8) {
	eh.SendReply([]byte{b})
}

func (eh *EchoHandler) ForwardByte(_ uint8) {
	// Not Implemented
}

func (eh *EchoHandler) HandleCommand(_ uint8) {
	// Not Implemented
}

func (eh *EchoHandler) SendReply(bytes []byte) error {
	eh.conn.Write(bytes)
	return nil
}
