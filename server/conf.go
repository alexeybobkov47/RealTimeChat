package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"realtimechat/structs"
)

func config() *structs.Conf {
	conf := new(structs.Conf)
	bytes, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(bytes, conf); err != nil {
		log.Fatal(err)
	}
	return conf
}
