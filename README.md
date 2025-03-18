NixMail POP3 Command Line Client
================================

popc --help
-----------
```
popc - NixMail POP3 Command Line Client

DESCRIPTION
The POP3 (Post Office Protocol Version 3) is an application-layer Internet standard protocol used by local Email clients to retrieve eMail from a remote server over a TCP/IP connection.

The NixMail POP3 client is a simple, stand alone tool to retrieve email from a server using the command line or possibly automate the download via a scheduled cron job. The client can be used independent from the rest of the NixMail system.

SYNOPSIS
	popc [OPTIONS]

OPTIONS
	-c, --config [string]
		Custom config file path

	-a, --pop3-alias [string]
		Pop3 account to use from config file

	--reset  [bool] 
		Send pop3 Rset, which will unmark any messages that have being marked for deletion in the current session

	--pipe [bool]
		Pipe out mail, rather then saving to mail quene

 	--print-config
		Print POP3 configuration, and exit

	--default-config-path [bool]
		Print the default config path, and exit

	-h, --help [bool]
		Report usage information and exit

	-v, --version [bool]
		Print version tag and exit

	-d, --debug [bool]
		Print debug information about the process

	--verbose [bool]
		Print extra verbose debug information about the process

REPORTING BUGS
	npub12jjczvd2mzstyhr468fyas7vzmsm5d2x3tv5l9tev6q0jakk9djqx4uk7x

COPYRIGHT
	© 2025 MIT
```

--------------------------------------------------------------------------------
## Note:

#### Some email servers let you login to a specific folder.

You need to specify the folder in the Username, such as:

	Username = user@excample.com#foldername
	Username = foldername#user@excample.com

More info on folders at [afterlogic.com](https://afterlogic.com/docs/mailbee-net-tutorials/advanced/pop3-folders)
