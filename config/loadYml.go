package config

import (
	"bac-scraper-gui/obj"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"gopkg.in/yaml.v3"
)

func LoadYml() *obj.EnVariable {
	cmd := exec.Command("pwd")
	path, _ := cmd.Output()
	fmt.Print(string(path))
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
