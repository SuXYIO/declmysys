package exitcode

// gonna use int, since os.Exit uses int. Saves one type conversion anyway

const (
	Success     int = 0
	InvalidArgs int = 1
	FileError   int = 2
	ConfigError int = 2
	SetupError  int = 2
	ExecError   int = 2
	Unknown     int = -1
)
