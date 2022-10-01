package config

import (
	"bac-scraper-gui/obj"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func LoadYml() *obj.EnVariable {
	pathArr := strings.Split(os.Getenv("PATH"), ":")
	data, err := combinePath(pathArr)
	errDealer(err)

	yml := obj.EnVariable{}
	err = yaml.Unmarshal(data, &yml)
	errDealer(err)

	return &yml
}

func combinePath(pathArr []string) ([]byte, error) {
	err := errors.New("")
	for _, path := range pathArr {
		data, err := ioutil.ReadFile(path + "/config.yml")
		if err == nil {
			return data, nil
		}
	}

	return nil, err
}

func errDealer(err error) {
	if err != nil {
		log.Println(err)
	}
}
