package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"
	"strconv"
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
//validate access key and secret key
func validate_access_secret_key(accesskeys,secretkeys string) bool {
	var status bool
	status = false

	if len(accesskeys) == 20 && len(secretkeys) == 40 {
		status = true
		return status
	} else {
		return status
	}
}
//function check errors
func check_err(e error){
	if e !=nil {
		panic(e)
	}
}
//Structure of GoTerraform.Yaml
type AWS struct {
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
		Type             string `yaml:"type"`
		Name             string `yaml:"name"`
		CidrBlock        string `yaml:"cidr_block"`
		AvailabilityZone string `yaml:"availability_zone"`
	} `yaml:"buildsubnet"`
}

func main(){

	//Vriables
	var selected_region string
	var selected_vpcname string
	var selected_accesskey string
	var selected_secretkey string
	

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
	_, err = os.Stat(f.Projectname);

	if err == nil {
		//project is already available
		//so remove it and create again
		fmt.Println("Updating Project : "+f.Projectname)
		os.RemoveAll(f.Projectname)
		os.MkdirAll(f.Projectname,0755)
		//create project.tf and project_variable.tf file
		pwd,_ := os.Getwd()
		filemain_tf,err2 := os.Create(pwd+"/"+f.Projectname+"/"+f.Projectname+"_main.tf")
		check_err(err2)
		defer filemain_tf.Close()
	} else {
		//creating new project workspace
		fmt.Println("Creating Project : "+f.Projectname)
		os.MkdirAll(f.Projectname,0755)
		//create project.tf and project_variable.tf file
		pwd,_ := os.Getwd()
		filemain_tf,err2 := os.Create(pwd+"/"+f.Projectname+"/"+f.Projectname+"_main.tf")
		check_err(err2)
		defer filemain_tf.Close()
	}
	//build awskey/region
	// 1. check for aws region 'ask'/valid_regionname
	selected_region = f.Awskey.Region

	if selected_region == "ask" {
		fmt.Println("region need to asked")
	} else {

		validate_region_result := validate_region(selected_region)
		if validate_region_result == true {
			fmt.Println("Selected Region  : "+selected_region)
		} else {
			fmt.Println("\nError: invalid US region "+selected_region)
			valid_regions := "\nValid Regions in US\n-------------------\n"+
			"a. us-east-1\t b. us-east-2\n"+
			"c. us-west-1\t c. us-west-2"
			fmt.Println(valid_regions)
			os.Exit(1)
		}
	}
	// 2. check access/secret key
	selected_accesskey = f.Awskey.AccessKey
	selected_secretkey = f.Awskey.SecretKey

	selected_accessecretkey_result :=  validate_access_secret_key(selected_accesskey,selected_secretkey)
	if selected_accessecretkey_result == true {
		fmt.Println("Selected AWS Key : "+selected_accesskey)
		//write project main.tf
		maintf := "#configure AWS provider\n"+
		"provider \"aws\" {\n"+
		" region = \""+selected_region+"\"\n"+
		" access_key = \""+selected_accesskey+"\"\n"+
		" secret_key = \""+selected_secretkey+"\"\n}\n"
		
		pwd,_ := os.Getwd()
		fil,err := os.OpenFile(pwd+"/"+f.Projectname+"/"+f.Projectname+"_main.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		check_err(err)
		if _, err := fil.Write([]byte(maintf)); err != nil {
			log.Fatal(err)
		}
		if err := fil.Close(); err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("\nError: invalid Access/Secret key")
		os.Exit(1)
	}
	//build vpc
	selected_vpcname = f.Buildvpc.Name
	fmt.Println("VPC Name	 : "+selected_vpcname)
	maintf := "#Create VPC\n"+
	"resource \"aws_vpc\" \"custom_vpc\" {\n"+
	 "cidr_block = \""+f.Buildvpc.CidrBlock+"\"\n"+
	 "enable_dns_support = "+f.Buildvpc.EnableDNSSupport+"\n"+
	 "enable_dns_hostnames = "+f.Buildvpc.EnableDNSHostnames+"\n"+
	 "tags = {\n"+
	  " Name = \""+f.Projectname+"_vpc\"\n}\n}\n"

	pwd,_ := os.Getwd()
	fil,err := os.OpenFile(pwd+"/"+f.Projectname+"/"+f.Projectname+"_main.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	check_err(err)
	if _, err := fil.Write([]byte(maintf)); err != nil {
		log.Fatal(err)
	}
	if err := fil.Close(); err != nil {
		log.Fatal(err)
	}

	//build subnet
	buildsubnet_count := len(f.Buildsubnet)
	count := 0
	for count < buildsubnet_count {
		
		if f.Buildsubnet[count].Type == "public" {
			maintf := "#Configure Public Subnet\n"+
			"resource \"aws_subnet\" \"custom_publicsubnet_"+strconv.Itoa(count)+"\" {\n"+
			" vpc_id = aws_vpc.custom_vpc.id\n"+
			" cidr_block = \""+f.Buildsubnet[count].CidrBlock+"\"\n"+
			" availability_zone = \""+f.Buildsubnet[count].AvailabilityZone+"\"\n"+
			" tags = {\n \"Name\" = \""+f.Projectname+"_publicsubet_"+strconv.Itoa(count)+"\"\n}\n}\n"
			
			pwd,_ := os.Getwd()
			fil,err := os.OpenFile(pwd+"/"+f.Projectname+"/"+f.Projectname+"_main.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			check_err(err)
			if _, err := fil.Write([]byte(maintf)); err != nil {
				log.Fatal(err)
			}
			if err := fil.Close(); err != nil {
				log.Fatal(err)
			}
			
		} else {
			fmt.Println("Private it is")
		}
		count += 1
	}
	
}	