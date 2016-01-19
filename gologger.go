package gologger

import (
	"comail.io/go/colog"
    "fmt"
	"gopkg.in/natefinch/lumberjack.v2"
    "log"
    "os"
)

// InitLogToStdout sets logging output to stdout (console)
func InitLogToStdout(env string) {
    if env == "dev"{
        colog.SetMinLevel(colog.LDebug)
    }else if env == "stage"{
        colog.SetMinLevel(colog.LInfo)
    }
    colog.SetOutput(os.Stdout)
}

// InitLogToFile sets logging error output to file (warning level or above)
func InitLogToFile(errorlog string) {
	colog.SetMinLevel(colog.LWarning)
	colog.SetOutput(&lumberjack.Logger{
        Filename:     errorlog,
		MaxSize:      500, // megabytes
		MaxBackups:   3,
		MaxAge:       14, //days
    })
}

// InitEnv configures logging settings based on environment parameter
func InitEnv(env string, errorlog string) {
    colog.SetFlags(log.LstdFlags | log.Lshortfile)
    colog.ParseFields(true)
    if env == "dev"{
        InitLogToStdout(env)
    }else if env == "stage"{
        InitLogToStdout(env)
    }else if env == "prod"{
        fmt.Println("Warning level or above outputted to",errorlog)
        InitLogToFile(errorlog)
    }else{
        log.Panic("Invalid environment argument!")
    }
}