# Overview

Gologger incorporates the elements of the best user-designed logging tools for Go. It primarily makes use of [CoLog](https://github.com/comail/colog), a prefix-based execution logging tool. 

The goal of this tool is to provide a simple, easy-to-use, customizable framework for implementing leveled, environment-based logging for new Go applications.

![Image gologger example program](/gologger.png)

# Usage
This tool is setup to customize logging outputs based on the environment command line flag. It uses CoLog to read and parse the standard log output, effectively acting as an extension of the default log library. It also uses [Lumberjack](https://github.com/natefinch/lumberjack) as a rolling error log for prod environment outputs. This errorlog (name and path determined by command line flag) has been configured to rotate once it exceeds 500 MB or 14 days, for a maximum of 3 backups. 

You'll need to install the following Go packages to start using this:
<pre>
	go get github.com/andrewrgoss/gologger
	go get comail.io/go/colog
	go get gopkg.in/natefinch/lumberjack.v2
</pre>
 
# Command line flags
The following command line flags should be included in your main program:
<pre>
  -env string
    	Specify your environment as dev, stage, or prod. Dev writes to console at debug_level, stage writes to console at info_level, prod writes to rolling error file (warning_level or above). Required. (default "dev")
  -errorlog string
    	The output path and file name for error log (used for prod environment only). Required. (default "./error.log")
</pre>
