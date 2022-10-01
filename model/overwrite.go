package model

import (
	"bac-scraper-gui/obj"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v3"
)

func OverWriteYml(mp map[string]*widget.Entry, envar *obj.EnVariable) {
	newSlice := parseStr(mp["members"].Text)
	setEnvar(newSlice, envar)
	data, err := yaml.Marshal(envar)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile("config.yml", data, 0777)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Updated!")
}

func parseStr(nameList string) []string {
	strSlice := make([]string, 0)
	for _, s := range strings.Split(nameList, "\n") {
		if s != "" {
			strSlice = append(strSlice, s)
		}
	}

	return strSlice
}

func setEnvar(newSlice []string, envar *obj.EnVariable) {
	envar.SetMembers(newSlice)
}
