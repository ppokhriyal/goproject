package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Engineer struct {
	Id			string		`yaml:"id"`
	Sector		int			`yaml:"sector"`
	Tasks		[]string	`yaml:"tasks"`
	DailyHours	[]int		`yaml:"dailyHours"`
	Languages	map[string]float32	`yaml:"languages"`
}

func main(){
	f := &Engineer{}
	source,err := ioutil.ReadFile("sample.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Printf("error: %v", err)
	}
	fmt.Println(f.DailyHours)
}