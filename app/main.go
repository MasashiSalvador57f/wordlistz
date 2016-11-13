package main

import (
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
)

// AppEnvironment is a application environment (temporary variable)
const AppEnvironment = "development"

// DBConf contains DBSetting for connecting Database
var DBConf DBSetting

// Engine is a engine that represents ...
var engine *xorm.Engine

func init() {
	dbConfigPath := fmt.Sprintf("app/config/database/%s.yaml", AppEnvironment)
	buf, err := ioutil.ReadFile(dbConfigPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, &DBConf)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@%s?charset=utf8", DBConf.User, DBConf.Password, DBConf.Database)
	engine, err = xorm.NewEngine("mysql", dsn)
}

// DBSetting is database configuration for application.
type DBSetting struct {
	User     string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:"password"`
}

func main() {
	fmt.Println("Hello world")
	fmt.Printf("%v", engine)
}
