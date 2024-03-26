package options

// TODO: This design is really bad and needs to be revisited. It is brittle and not thread safe.
// I am thinking maybe have an option registry using a map[string]Options. When an option is
// received, a new handler can be instantiated using new(Options) and act as the specific handler
// for that connection. Still a bit of memory overhead but no use pre-optimizing.

type (
	// TODO: Options is missing any meaningful way to manage the connection. What is the proper level?
	// I think the connection handler needs to know what type of terminal with which is being communicated,
	// as well as column count/etc. There may be, however, specific flags for telnet like binary mode?
	// I think for now we might need to add some sort of callbacks.
	Options interface {
		GetByte() uint8
		HandleByte(uint8, ConnHandler) bool
		ToString() string
	}

	ConnHandler interface {
		SendReply([]byte) error
	}

	baseOption struct {
		getByte    func() uint8
		toString   func() string
		handleByte func(uint8, ConnHandler) bool
	}
)

func (bo baseOption) GetByte() uint8 {
	return bo.getByte()
}

func (bo baseOption) ToString() string {
	return bo.toString()
}

func (bo baseOption) HandleByte(b uint8, c ConnHandler) bool {
	return bo.handleByte(b, c)
}

/*
const (
	// https://www.iana.org/assignments/telnet-options/telnet-options.xhtml
	BinaryTransmission              Option = 0x00 // https://www.iana.org/go/rfc856
	Echo                            Option = 0x01 // https://www.iana.org/go/rfc857
	Reconnection                    Option = 0x02
	SuppressGoAhead                 Option = 0x03 // https://www.iana.org/go/rfc858
	ApproxMessageSizeNegotiation    Option = 0x04
	Status                          Option = 0x05 // https://www.iana.org/go/rfc859
	TimingMark                      Option = 0x06 // https://www.iana.org/go/rfc860
	RemoteControlledTransAndEcho    Option = 0x07 // https://www.iana.org/go/rfc726
	OutputLineWidth                 Option = 0x08
	OutputPageSize                  Option = 0x09
	OutputCarriageReturnDisposition Option = 0x0A // https://www.iana.org/go/rfc652
	OutputHorizontalTabStops        Option = 0x0B // https://www.iana.org/go/rfc653
	OutputHorizontalTabDisposition  Option = 0x0C // https://www.iana.org/go/rfc654
	OutputFormfeedDisposition       Option = 0x0D // https://www.iana.org/go/rfc655
	OutputVerticalTabStops          Option = 0x0E // https://www.iana.org/go/rfc656
	OutputVerticalTabDisposition    Option = 0x0F // https://www.iana.org/go/rfc657
	OutputLinefeedDisposition       Option = 0x10 // https://www.iana.org/go/rfc658
	ExtendedASCII                   Option = 0x11 // https://www.iana.org/go/rfc698
	Logout                          Option = 0x12 // https://www.iana.org/go/rfc727
	ByteMacro                       Option = 0x13 // https://www.iana.org/go/rfc735
	DataEntryTerminal               Option = 0x14 // https://www.iana.org/go/rfc1043 https://www.iana.org/go/rfc732
	SUPDUP                          Option = 0x15 // https://www.iana.org/go/rfc736 https://www.iana.org/go/rfc734
	SUPDUPOutput                    Option = 0x16 // https://www.iana.org/go/rfc749
	SendLocation                    Option = 0x17 // https://www.iana.org/go/rfc779
	TerminalType                    Option = 0x18 // https://www.iana.org/go/rfc1091
	EndOfRecord                     Option = 0x19 // https://www.iana.org/go/rfc885
	TACACSUserIdentification        Option = 0x1A // https://www.iana.org/go/rfc927
	OutputMarking                   Option = 0x1B // https://www.iana.org/go/rfc933
	TerminalLocationNumber          Option = 0x1C // https://www.iana.org/go/rfc946
	Telnet3270Regime                Option = 0x1D // https://www.iana.org/go/rfc1041
	X3PAD                           Option = 0x1E // https://www.iana.org/go/rfc1053
	NegotiateAboutWindowSize        Option = 0x1F // https://www.iana.org/go/rfc1073
	TerminalSpeed                   Option = 0x20 // https://www.iana.org/go/rfc1079
	RemoteFlowControl               Option = 0x21 // https://www.iana.org/go/rfc1372
	Linemode                        Option = 0x22 // https://www.iana.org/go/rfc1184
	XDisplayLocation                Option = 0x23 // https://www.iana.org/go/rfc1096
	EnvironmentOption               Option = 0x24 // https://www.iana.org/go/rfc1408
	AuthenticationOption            Option = 0x25 // https://www.iana.org/go/rfc2941
	EncryptionOption                Option = 0x26 // https://www.iana.org/go/rfc2946
	NewEnvironmentOption            Option = 0x27 // https://www.iana.org/go/rfc1572
	TN3270E                         Option = 0x28 // https://www.iana.org/go/rfc2355
	XAUTH                           Option = 0x29
	CHARSET                         Option = 0x2A // https://www.iana.org/go/rfc2066
	TelnetRemoteSerialPort          Option = 0x2B
	ComPortControlOption            Option = 0x2C // https://www.iana.org/go/rfc2217
	TelnetSuppressLocalEcho         Option = 0x2D
	TelnetStartTLS                  Option = 0x2E
	KERMET                          Option = 0x2F // https://www.iana.org/go/rfc2840
	SENDURL                         Option = 0x30
	FORWARDX                        Option = 0x31 // https://www.iana.org/go/rfc861
	// 0x32 - 0x89 Unassigned
	TELOPTPRAGMALOGON     Option = 0x8A
	TELOPTSSPILOGON       Option = 0x8B
	TELOPTPRAGMAHEARTBEAT Option = 0x8C
	// 0x8D - 0xFE
	ExtendedOptionsList Option = 0xFF
)

func (o Option) ToString() string {
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
*/
