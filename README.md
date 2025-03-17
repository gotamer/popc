# NixMail POP3 command line client

## DESCRIPTION
The POP3 (Post Office Protocol Version 3) is an application-layer Internet standard protocol used by local Email clients to retrieve Email from a remote server over a TCP/IP connection.

The NixMail POP3 client is a simple, stand alone tool to retrieve email from a server using the command line or possibly automate the download via a scheduled cron job. The client can be used independent from the rest of the NixMail system.

Following enviroment variables must be set:
(example shown is for the fish shell)

	set -gx POP3_HOSTNAME "mail.example.net:995"
	set -gx POP3_USERNAME "alias@example.net"
	set -gx POP3_USERPASS "pop3 password"
	set -gx POP3_FILEPATH "/var/mail/quene"

--------------------------------------------------------------------------------

	pop3 [OPTIONS]

	The following options are available:

	-h or --help
	       Print this help and exit.

	-v or --version
	       Print version and exit.

--------------------------------------------------------------------------------
## Note:

#### Some email servers let you login to a specific folder.

You need to specify the folder in the Username, such as:

	Username = user@excample.com#foldername
	Username = foldername#user@excample.com

More info on folders at [afterlogic.com](https://afterlogic.com/docs/mailbee-net-tutorials/advanced/pop3-folders)
