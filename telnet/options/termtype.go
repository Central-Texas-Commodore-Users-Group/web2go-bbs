package options

import "fmt"

// https://www.iana.org/go/rfc1091
type (
	TerminalTypeCommand uint8
)

var (
	TerminalType = struct {
		baseOption
	}{}

	TermList []string
	currTerm []byte
)

const (
	TTIS   TerminalTypeCommand = 0x00
	TTSEND TerminalTypeCommand = 0x01
)

// DoTerminalType - IAC DO TerminalType
func DoTerminalType() []byte {
	return []byte{0xFF, 0xFD, 0x18}
}

// WillTerminalType - IAC SB TerminalType SEND IAC SE
func WillTerminalType() []byte {
	return []byte{0xFF, 0xFA, 0x18, byte(TTSEND), 0xFF, 0xF0}
}

func (ttc TerminalTypeCommand) ToString() string {
	var ret string
	switch ttc {
	case TTIS:
		ret = "IS"
	case TTSEND:
		ret = "SEND"
	default:
		ret = "unknown"
	}
	return ret
}

func init() {
	TerminalType.baseOption.toString = func() string {
		return "Terminal Type"
	}

	TerminalType.baseOption.getByte = func() uint8 {
		return 0x18
	}

	TerminalType.baseOption.handleByte = func(b uint8, c ConnHandler) bool {
		switch b {
		case uint8(TTIS):
			// TODO: only reason this works right now is we wait for the IS to reset the temp terminal buffer
			currTerm = nil
		case 0xff: // IAC
			// Assume only good data for now
			if TermList == nil || string(currTerm) != TermList[len(TermList)-1] {
				TermList = append(TermList, string(currTerm))
				fmt.Printf("Found terminal type: %s\n", string(currTerm))
				c.SendReply(WillTerminalType())
				// continue requesting until the same value is received twice in a row
			}
			currTerm = nil
		case 0xF0: // SE
			return true
		default:
			currTerm = append(currTerm, b)
		}
		return false
	}
}
