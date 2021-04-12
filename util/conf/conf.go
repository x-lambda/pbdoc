package conf

import (
	"github.com/go-ini/ini"
)

type SqliteConf struct {
	Name string `json:"name"`
}

var DB SqliteConf

func init() {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		panic(err)
	}

	cfg.Section("sqlite").MapTo(&DB)
}
