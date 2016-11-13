package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// AppEnvironment is a application environment (temporary variable)
const AppEnvironment = "development"

// DBConf contains DBSetting for connecting Database
var DBConf DBSetting

func init() {
	buf, err := ioutil.ReadFile("config/database/development.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, &DBConf)
	if err != nil {
		panic(err)
	}
}

// DBSetting is database configuration for application.
type DBSetting struct {
	User     string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:"password"`
}

func main() {
	fmt.Println("Hello world")
}
