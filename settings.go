package auth

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var Settings SettingsStruct

type SettingsStruct struct {
	RandomLevel int //0 -> math.rand (faster but unsafe)| 1 -> crypto.rand
	//(safer but slower)
	DebugLevel    int    //0-> only errors | 1 -> logs any main task done
	CheckIp       bool   //true -> check the user ip on login | false -> doesn't
	AccountsFile  string //path the where the accounst should be stored
	TokenValidity int64  //time in seconds a token is valid
	LogsFile      string //The file in which the logs are stored
}

//creates the settings from a yaml file storing theme.
//This must be called before calling any other functionfs from this package
func SetSettingsFromFile(path string) {
	set, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s := SettingsStruct{}
	err = yaml.Unmarshal(set, &s)
	if err != nil {
		log.Fatal(err)
	}
	Settings = s

	f, err := os.OpenFile(Settings.LogsFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
}
