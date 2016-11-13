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
	buf, err := ioutil.ReadFile("config/database.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, &DBConf)
	if err != nil {
		panic(err)
	}
	fmt.Println("----")
	fmt.Println(DBConf.Database)
	fmt.Println(DBConf.Password)
	fmt.Println(DBConf.User)
	fmt.Println("----")
	fmt.Printf("%v", DBConf)
}

// DBSetting is database configuration for application.
type DBSetting struct {
	User     string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:" password"`
}

func main() {
	fmt.Println("Hello world")
}
