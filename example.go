package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type AutoGenerated struct {
	Projectname string `yaml:"projectname"`
	Awskey      struct {
		Region    string `yaml:"region"`
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
	} `yaml:"awskey"`
	Buildvpc struct {
		Name               string `yaml:"name"`
		CidrBlock          string `yaml:"cidr_block"`
		EnableDNSSupport   string `yaml:"enable_dns_support"`
		EnableDNSHostnames string `yaml:"enable_dns_hostnames"`
	} `yaml:"buildvpc"`
	Buildsubnet []struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
	} `yaml:"buildsubnet"`
}

func main(){

	f := &AutoGenerated{}
	source,err := ioutil.ReadFile("goterraform.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Printf("error: %v", err)
	}
	fmt.Println(len(f.Buildsubnet))
}