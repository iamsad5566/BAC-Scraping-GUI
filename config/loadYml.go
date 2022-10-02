package config

import (
	"bac-scraper-gui/obj"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func LoadYml() *obj.EnVariable {
	data, err := ioutil.ReadFile("./config.yml")
	errDealer(err)

	yml := obj.EnVariable{}
	err = yaml.Unmarshal(data, &yml)
	errDealer(err)

	return &yml
}
func errDealer(err error) {
	if err != nil {
		log.Println(err)
	}
}
