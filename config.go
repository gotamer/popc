package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"go.hansaray.pw/lib/host"
	"go.hansaray.pw/lib/log"
)

const configFileName = "pop3.toml"

var (
	configFilePath string
	configDirPath  string
	pop3Alias      string
)

var Accounts map[string]Credentials

type Credentials struct {
	Hostname string
	Username string
	Password string
}

func (u *User) loadConfig() {
	if len(pop3Alias) == 0 {
		pop3Alias = "default"
	}
	if err := toml.DecodeFile(configFilePath, &Accounts); err != nil {
		log.F("Config file: %s", err.Error())
	}

	var account = Accounts[pop3Alias]
	if len(account.Hostname) != 0 {
		if len(account.Username) != 0 {
			if len(account.Password) != 0 {
				Hostname = account.Hostname
				Username = account.Username
				Password = account.Password
			} else {
				log.F("Could not load pop3 account Password")
			}
		} else {
			log.F("Could not load pop3 account Username")
		}
	} else {
		log.F("Could not load pop3 account Hostname")
	}
}

func UserConfigDir() string {
	configPaths()
	return configDirPath
}

func UserConfigPath() string {
	configPaths()
	return configFilePath
}

func configPaths() {
	if len(configFilePath) == 0 {

		var path string

		path = host.UserConfigDir()
		if len(path) == 0 {
			log.F("Could not find default config path")
		} else {
			configDirPath = filepath.Join(path, "mail")
			configFilePath = filepath.Join(configDirPath, configFileName)
		}
	}
}

// FileCreate creates the named file. If the file already exists an Error occurs,
// If the file does not exist, it is created with mode 0o600 (before umask).
// If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode O_RDWR.
// The directory containing the file must already exist.
// If there is an error, it will be of type *PathError.
func FileCreate(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0600)
}
