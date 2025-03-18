package main

import (
	"fmt"
	"os"
	"path"

	log "go.hansaray.pw/lib/logger"
	"go.hansaray.pw/lib/pop3"
	"go.hansaray.pw/lib/random"
)

const (
	APPNAME = "NixMail POP3 Client"
)

var (
	c        *pop3.Client
	Hostname string
	Username string
	Password string
	Filepath = "/var/mail/pop3"
)

func init() {
	flags()
	UserConfigPath()
	loadConfig()
}

func main() {
	if viewcfg {
		fmt.Printf("host: %s, user: %s, pass: %s, File: %s", Hostname, Username, Password, Filepath)
		os.Exit(0)
	}
	var err error
	if c, err = pop3.DialTLS(Hostname); err != nil {
		log.F("DialTLS: %s", err.Error())
	}
	defer Quit(c)
	if err = c.Authorization(Username, Password); err != nil {
		log.F("AUTH: %s", err.Error())
	}
	if err = c.ListCapabilities(); err != nil {
		log.I("CAPA: %s", err.Error())
	}
	if err = c.Stat(); err != nil {
		log.I("STAT: %s", err.Error())
	}
	if c.Count > 0 {
		log.D("Count: %v", c.Count)
		if err = c.ListAll(); err != nil {
			log.F("LIST: %s", err.Error())
		}
		for _, v := range c.List {
			log.D("id: %d UID: %s Size: %d", v.ID, v.UID, v.Size)
			message, err := c.RetrRaw(v.ID)
			if err != nil {
				popRset = true
				log.F("Message ID: %v ERR: %s", v.ID, err.Error())
			}
			if pipe {
				if _, err := fmt.Fprint(os.Stdout, string(message)); err != nil {
					popRset = true
					log.F(err.Error())
				}
			} else {
				var filename string
				if v.UID == "" {
					filename = random.String(10, 10, false)
				} else {
					filename = v.UID
				}
				file := path.Join(Filepath, filename)
				log.D("File: %s", v.ID, v.UID, v.Size)
				err = os.WriteFile(file, message, 0640)
				if err != nil {
					popRset = true
					log.F("Write Message ID: %v ERR: %s", v.ID, err.Error())
				}
			}
		}
	}
}

func Quit(c *pop3.Client) {
	var err error
	if popRset {
		if err = c.Rset(); err != nil {
			log.E("Rset: %s", err.Error())
		}
	}
	if err = c.Quit(); err != nil {
		log.E("QUIT: %s", err.Error())
	}
}
