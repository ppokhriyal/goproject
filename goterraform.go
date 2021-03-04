package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

//validate selected region
func validate_region(regionselected string) bool {
	var status bool
	status = false
	regions := map[string]bool{"us-east-1":true,"us-east-2":true,"us-west-1":true,"us-west-2":true}
	
	if regions[regionselected] {
		status = true
		return status
	} else {
		return status
	}

}
//Structure of GoTerraform.Yaml
type AWS struct{
	ProjectName string	`yaml:"projectname"`
	Region		string	`yaml:"region"`
}

func main(){

	//Vriables
	var selected_region string

	f := &AWS{}
	source,err := ioutil.ReadFile("goterraform.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Printf("error: %v", err)
	}

	//validate goterraform.yaml

	//check for project workspace exists or not
	//if not then create new and if exists, remove
	//and then create new.
	_, err = os.Stat(f.ProjectName);

	if err == nil {
		//project is already available
		//so remove it and create again
		fmt.Println("Updating Project : "+f.ProjectName)
		os.RemoveAll(f.ProjectName)
		os.MkdirAll(f.ProjectName,0755)
	} else {
		//creating new project workspace
		fmt.Println("Creating Project : "+f.ProjectName)
		os.MkdirAll(f.ProjectName,0755)
	}
	//check if region is pre-defined or it is requested for ask
	if f.Region == "ask" {
		fmt.Println("Region will be asked")
	} else {
		selected_region = f.Region
		validate_region_result := validate_region(selected_region)
		if validate_region_result == false {
			fmt.Println("\nError : Invalid region "+selected_region)
			avregions := "Available Regions \n-----------------\n"+
			"a. us-east-1\tb. us-east-2\n"+
			"c. us-west-1\td. us-west-2"
			fmt.Println(avregions)
			
		}
		//fmt.Println("Selected Region  : "+selected_region)
	}
}