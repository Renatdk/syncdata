package logging

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func initHandles(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func InitLogging() {
	t, _ := os.OpenFile("logs/Trace.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	i, _ := os.OpenFile("logs/Info.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	w, _ := os.OpenFile("logs/Warning.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	e, _ := os.OpenFile("logs/Error.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	initHandles(t, i, w, e)

	// Trace.Println("I have something standard to say")
	// Info.Println("Special Information")
	// Warning.Println("There is something you need to know about")
	// Error.Println("Something has failed")
}
