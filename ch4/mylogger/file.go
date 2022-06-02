package mylogger

type FIleLogger struct {
	Levle       LogLevel
	filePath    string //日志文件保存路径
	fileName    string //日志文件保存的文件名
	errFileName string
	maxFileSize int64
}

//
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FIleLogger {
	logLevel, err := parseLevelstr(levelStr)
	if err != nil {
		panic(err)
	}
	return &FIleLogger{
		Levle:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
}
