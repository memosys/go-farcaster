package neynar

// Temporary Logger interface
type Logger interface {
	Log(message string)
}

func Log(logger Logger, message string) {
	logger.Log(message)
}
