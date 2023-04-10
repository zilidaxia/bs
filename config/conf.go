package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"strings"
)

type Mysql struct {
	Host      string
	Port      string
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
	Loc       string
}

type Server struct {
	IP   string
	Port int
}

type Redis struct {
	IP       string
	Port     int
	Database int
}

type Path struct {
	StaticSourcePath string `toml:"static_source_path"`
}

type Config struct {
	DB     Mysql `toml:"mysql"`
	RDB    Redis `toml:"redis"`
	Server `toml:"server"`
	Path   `toml:"path"`
}

var Info Config

func init() {
	if _, err := toml.DecodeFile("\\Users\\caozili\\MyCode\\bs\\config\\config.toml", &Info); err != nil {
		panic(err)
	}
	strings.Trim(Info.Server.IP, " ")
	strings.Trim(Info.RDB.IP, " ")
	strings.Trim(Info.DB.Host, " ")
}

func DBConnectString() string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Info.DB.Username, Info.DB.Password, Info.DB.Host, Info.DB.Port, Info.DB.Database,
		Info.DB.Charset, Info.DB.ParseTime, Info.DB.Loc)
	log.Println(arg)
	return arg
}
