package bridge

const (

	//cmd|length|token
	CmdAuth   = 1
	// cmd|length|from|to|bytes
	CmdData   = 2
	// cmd|length|string
	CmdMsg = 3
	// cmd|length|address
	CmdConnect = 4
	// cmd|length|address
	CmdClose = 5
)

const (
	FlagStart    = 0xE0
	FlagEscStart = 0xEF

	FlagEsc    = 0xF0
	FlagEscEsc = 0xFF
)

const (
	EmptyAddress = "0.0.0.0:0"
)