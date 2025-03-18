package main

import (
	"fmt"
	"os"
	"strings"

	log "go.hansaray.pw/lib/logger"
	"go.hansaray.pw/lib/version"
)

var (
	ERR_NO_NEXT  = "Option %s needs an argument"
	ERR_NO_ALIAS = "No alias given for --pop3-alias"
)

var (
	popRset bool
)

func flags() {
	log.APPNAME = "pop3"

	if len(os.Args) == 1 {
		fmt.Print(help)
		os.Exit(1)
	}

	for no, arg := range os.Args {
		switch arg {

		case "-c", "--config":
			if mustNext(no) {
				configFilePath = getNext(no)
			}
			continue

		case "--pop3-alias":
			if mustNext(no) {
				pop3Alias = getNext(no)
			}
			continue

		case "--Reset", "--reset":
			// send pop3 Rset, which will unmark any messages that have being marked for deletion in the current session.
			popRset = true
			continue

		case "--config-template":
			log.Info("Not yet implemented")
			os.Exit(0)

		case "--default-config-path":
			fmt.Printf("Default config path: %s\n", UserConfigPath())
			os.Exit(0)

		case "-h", "--help":
			fmt.Print(help)
			os.Exit(0)

		case "-v", "--version":
			fmt.Print(version.Info())
			os.Exit(0)

		case "-d", "--debug":
			log.Verbose(2)
			continue

		case "--verbose":
			log.Verbose(3)
			continue
		}
	}
}

func hasNext(no int) bool {
	if len(os.Args) > no {
		return true
	}
	return false
}

func mustNext(no int) bool {
	var next = os.Args[no+1]
	if hasNext(no) && !strings.HasPrefix(next, "-") {
		return true
	}
	log.F(ERR_NO_NEXT, os.Args[no])
	return false
}

func getNext(no int) string {
	return os.Args[no+1]
}

var help = `
pop3 - NixMail POP3 command line client

DESCRIPTION
The POP3 (Post Office Protocol Version 3) is an application-layer Internet standard protocol used by local Email clients to retrieve eMail from a remote server over a TCP/IP connection.

The NixMail POP3 client is a simple, stand alone tool to retrieve email from a server using the command line or possibly automate the download via a scheduled cron job. The client can be used independent from the rest of the NixMail system.

SYNOPSIS
	pop3 [OPTIONS]

OPTIONS
	-c, --config [string]
		Custom config file path

	--pop3-alias [string]
		Pop3 account to use from config file.

	--reset  [bool] 
		Send pop3 Rset, which will unmark any messages that have being marked for deletion in the current session.

	--default-config-path [bool]
		Will print the default config path, and exit.

	-h, --help [bool]
		Report usage information and exit.

	-v, --version [bool]
		Print version tag and exit.

	-d, --debug [bool]
		Print debug information about the process.

	--verbose [bool]
		Print extra verbose debug information about the process.

REPORTING BUGS
	npub12jjczvd2mzstyhr468fyas7vzmsm5d2x3tv5l9tev6q0jakk9djqx4uk7x

COPYRIGHT
	© 2025 MIT 
`
