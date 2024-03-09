package main

type (
	Command uint8
	Options uint8
)

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

func (o Options) ToString() string {
	var ret string
	switch o {
	case BinaryTransmission:
		ret = "Binary Transmission"
	case Echo:
		ret = "Echo"
	case Reconnection:
		ret = "Reconnection"
	case SuppressGoAhead:
		ret = "Supporess Go Ahead"
	case ApproxMessageSizeNegotiation:
		ret = "Approx Message Size Negotiation"
	case Status:
		ret = "Status"
	case TimingMark:
		ret = "Timing Mark"
	case RemoteControlledTransAndEcho:
		ret = "Remote Controlled Trans and Echo"
	case OutputLineWidth:
		ret = "Output Line Width"
	case OutputPageSize:
		ret = "Output Page Size"
	case OutputCarriageReturnDisposition:
		ret = "Output Carriage-Return Disposition"
	case OutputHorizontalTabStops:
		ret = "Output Horizontal Tab Stops"
	case OutputHorizontalTabDisposition:
		ret = "Output Horizontal Tab Disposition"
	case OutputFormfeedDisposition:
		ret = "Output Formfeed Disposition"
	case OutputVerticalTabStops:
		ret = "Output Vertical Tab Stops"
	case OutputVerticalTabDisposition:
		ret = "Output Vertical Tab Disposition"
	case OutputLinefeedDisposition:
		ret = "Output Linefeed Disposition"
	case ExtendedASCII:
		ret = "Extended ASCII"
	case Logout:
		ret = "Logout"
	case ByteMacro:
		ret = "Byte Macro"
	case DataEntryTerminal:
		ret = "Data Entry Terminal"
	case SUPDUP:
		ret = "SUPDUP"
	case SUPDUPOutput:
		ret = "SUPDUP Output"
	case SendLocation:
		ret = "Send Location"
	case TerminalType:
		ret = "Terminal Type"
	case EndOfRecord:
		ret = "End of Record"
	case TACACSUserIdentification:
		ret = "TACACS User Identification"
	case OutputMarking:
		ret = "Output Marking"
	case TerminalLocationNumber:
		ret = "Terminal Location Number"
	case Telnet3270Regime:
		ret = "Telnet 3270 Regime"
	case X3PAD:
		ret = "X.3 PAD"
	case NegotiateAboutWindowSize:
		ret = "Negotiate About Window Size"
	case TerminalSpeed:
		ret = "Terminal Speed"
	case RemoteFlowControl:
		ret = "Remote Flow Control"
	case Linemode:
		ret = "Linemode"
	case XDisplayLocation:
		ret = "X Display Location"
	case EnvironmentOption:
		ret = "Environment Option"
	case AuthenticationOption:
		ret = "Authentication Option"
	case EncryptionOption:
		ret = "Encryption Option"
	case NewEnvironmentOption:
		ret = "New Environment Option"
	case TN3270E:
		ret = "TN3270E"
	case XAUTH:
		ret = "XAUTH"
	case CHARSET:
		ret = "CHARSET"
	case TelnetRemoteSerialPort:
		ret = "Telnet Remote Serial Port (RSP)"
	case ComPortControlOption:
		ret = "Com Port Control Option"
	case TelnetSuppressLocalEcho:
		ret = "Telnet Suppress Local Echo"
	case TelnetStartTLS:
		ret = "Telnet Start TLS"
	case KERMET:
		ret = "KERMIT"
	case SENDURL:
		ret = "SEND-URL"
	case FORWARDX:
		ret = "FORWARD_X"
	case TELOPTPRAGMALOGON:
		ret = "TELOPT PRAGMA LOGON"
	case TELOPTSSPILOGON:
		ret = "TELOPT SSPI LOGON"
	case TELOPTPRAGMAHEARTBEAT:
		ret = "TELOPT PRAGMA HEARTBEAT"
	case ExtendedOptionsList:
		ret = "Extended-Options-List"
	default:
		ret = "Unassigned"
	}
	return ret
}

const (
	// https://www.rfc-editor.org/rfc/rfc854.html
	SE               Command = 0xF0
	NOP              Command = 0xF1
	DataMark         Command = 0xF2
	Break            Command = 0xF3
	InterruptProcess Command = 0xF4
	AbortOutput      Command = 0xF5
	AreYouThere      Command = 0xF6
	EraseCharacter   Command = 0xF7
	EraseLine        Command = 0xF8
	GoAhead          Command = 0xF9
	SB               Command = 0xFA
	// Option Codes
	WILL Command = 0xFB
	WONT Command = 0xFC
	DO   Command = 0xFD
	DONT Command = 0xFE
	IAC  Command = 0xFF

	// https://www.iana.org/assignments/telnet-options/telnet-options.xhtml
	BinaryTransmission              Options = 0x00 // https://www.iana.org/go/rfc856
	Echo                            Options = 0x01 // https://www.iana.org/go/rfc857
	Reconnection                    Options = 0x02
	SuppressGoAhead                 Options = 0x03 // https://www.iana.org/go/rfc858
	ApproxMessageSizeNegotiation    Options = 0x04
	Status                          Options = 0x05 // https://www.iana.org/go/rfc859
	TimingMark                      Options = 0x06 // https://www.iana.org/go/rfc860
	RemoteControlledTransAndEcho    Options = 0x07 // https://www.iana.org/go/rfc726
	OutputLineWidth                 Options = 0x08
	OutputPageSize                  Options = 0x09
	OutputCarriageReturnDisposition Options = 0x0A // https://www.iana.org/go/rfc652
	OutputHorizontalTabStops        Options = 0x0B // https://www.iana.org/go/rfc653
	OutputHorizontalTabDisposition  Options = 0x0C // https://www.iana.org/go/rfc654
	OutputFormfeedDisposition       Options = 0x0D // https://www.iana.org/go/rfc655
	OutputVerticalTabStops          Options = 0x0E // https://www.iana.org/go/rfc656
	OutputVerticalTabDisposition    Options = 0x0F // https://www.iana.org/go/rfc657
	OutputLinefeedDisposition       Options = 0x10 // https://www.iana.org/go/rfc658
	ExtendedASCII                   Options = 0x11 // https://www.iana.org/go/rfc698
	Logout                          Options = 0x12 // https://www.iana.org/go/rfc727
	ByteMacro                       Options = 0x13 // https://www.iana.org/go/rfc735
	DataEntryTerminal               Options = 0x14 // https://www.iana.org/go/rfc1043 https://www.iana.org/go/rfc732
	SUPDUP                          Options = 0x15 // https://www.iana.org/go/rfc736 https://www.iana.org/go/rfc734
	SUPDUPOutput                    Options = 0x16 // https://www.iana.org/go/rfc749
	SendLocation                    Options = 0x17 // https://www.iana.org/go/rfc779
	TerminalType                    Options = 0x18 // https://www.iana.org/go/rfc1091
	EndOfRecord                     Options = 0x19 // https://www.iana.org/go/rfc885
	TACACSUserIdentification        Options = 0x1A // https://www.iana.org/go/rfc927
	OutputMarking                   Options = 0x1B // https://www.iana.org/go/rfc933
	TerminalLocationNumber          Options = 0x1C // https://www.iana.org/go/rfc946
	Telnet3270Regime                Options = 0x1D // https://www.iana.org/go/rfc1041
	X3PAD                           Options = 0x1E // https://www.iana.org/go/rfc1053
	NegotiateAboutWindowSize        Options = 0x1F // https://www.iana.org/go/rfc1073
	TerminalSpeed                   Options = 0x20 // https://www.iana.org/go/rfc1079
	RemoteFlowControl               Options = 0x21 // https://www.iana.org/go/rfc1372
	Linemode                        Options = 0x22 // https://www.iana.org/go/rfc1184
	XDisplayLocation                Options = 0x23 // https://www.iana.org/go/rfc1096
	EnvironmentOption               Options = 0x24 // https://www.iana.org/go/rfc1408
	AuthenticationOption            Options = 0x25 // https://www.iana.org/go/rfc2941
	EncryptionOption                Options = 0x26 // https://www.iana.org/go/rfc2946
	NewEnvironmentOption            Options = 0x27 // https://www.iana.org/go/rfc1572
	TN3270E                         Options = 0x28 // https://www.iana.org/go/rfc2355
	XAUTH                           Options = 0x29
	CHARSET                         Options = 0x2A // https://www.iana.org/go/rfc2066
	TelnetRemoteSerialPort          Options = 0x2B
	ComPortControlOption            Options = 0x2C // https://www.iana.org/go/rfc2217
	TelnetSuppressLocalEcho         Options = 0x2D
	TelnetStartTLS                  Options = 0x2E
	KERMET                          Options = 0x2F // https://www.iana.org/go/rfc2840
	SENDURL                         Options = 0x30
	FORWARDX                        Options = 0x31 // https://www.iana.org/go/rfc861
	// 0x32 - 0x89 Unassigned
	TELOPTPRAGMALOGON     Options = 0x8A
	TELOPTSSPILOGON       Options = 0x8B
	TELOPTPRAGMAHEARTBEAT Options = 0x8C
	// 0x8D - 0xFE
	ExtendedOptionsList Options = 0xFF
)
