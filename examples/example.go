package main

import (
    "comail.io/go/colog"
    "flag"
    "fmt"
    "gitlab.digitas.com/media-systems/gologger"
    "log"
    "os"
    "strings"
)

func printLogs() {
	fmt.Println(strings.Repeat("=", 80))
	log.Print("debug: debug message")
	log.Print("info: Hello world!")
	log.Print("warn: has some fields key=value; one=two")
	log.Print("error: Some error")
    // Open a non-existant file test
    f, err := os.OpenFile("./nonexistantfile.csv", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        log.Print("alert: ",err)
        //panic(err)
    }
    log.Println("filename=",f)
    fmt.Println(strings.Repeat("-", 80)) 
}

// Set up the command line flags
var env, errorlog string

func init() {
    flag.StringVar(&env, "env", "dev", "Specify your environment as dev, stage, or prod. Dev writes to console at debug_level, stage writes to console at info_level, prod writes to rolling error file (warning_level or above). Required.")
    flag.StringVar(&errorlog, "errorlog", "./error.log", "The output path and file name for error log (used for prod environment only). Required.")
    // Parse the command-line flags...
	flag.Parse()
}

func main() {
    //Logging setup
    colog.Register()
    
    //Logging examples by environment
    env="dev"
    fmt.Println("ENVIRONMENT=", strings.ToUpper(env))
    gologger.InitEnv(env, errorlog)
    printLogs()
    
    env="stage"
    fmt.Println("ENVIRONMENT=", strings.ToUpper(env))
    gologger.InitEnv(env, errorlog)
    printLogs()
    
    env="prod"
    fmt.Println("ENVIRONMENT=", strings.ToUpper(env))
    gologger.InitEnv(env, errorlog)
    printLogs()    
}
