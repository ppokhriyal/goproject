//Write to a file
package main

import (
	"fmt"
	"os"
	
)

//function check errors
func check_err(e error){
	if e !=nil {
		panic(e)
	}
}
func main() {
	
	//remove main.tf and variable.tf
	os.Remove("main.tf")
	os.Remove("variable.tf")

	//get accesskeyid
	fmt.Print("Enter AccessKey : ")
	var accesskey string
	fmt.Scanln(&accesskey)

	//get secretkeyid
	fmt.Print("Enter SecretKey : ")
	var secretkey string
	fmt.Scanln(&secretkey)

	//get custom VPC name
	fmt.Print("Enter new VPC name : ")
	var custom_vpc string
	fmt.Scanln(&custom_vpc)

	//get custom VPC CIDR block
	fmt.Print("Enter VPC CIDR block e.g 10.0.0.0/16 : ")
	var custom_vpc_cidr string
	fmt.Scanln(&custom_vpc_cidr)

	//get custom Publc Subnet name
	fmt.Print("Enter new Public Subnet name : ")
	var custom_subnet_public string
	fmt.Scanln(&custom_subnet_public)

	//get custom Public Subnet CIDR
	fmt.Print("Enter Public Subnet CIDR block e.g 10.0.0.0/24  : ")
	var custom_subnet_public_cidr string
	fmt.Scanln(&custom_subnet_public_cidr)

	//create variable.tf file
	file_variable_tf,err1 := os.Create("variable.tf")
	check_err(err1)
	defer file_variable_tf.Close()

	//create main.tf file
	main_tf,err2 := os.Create("main.tf")
	check_err(err2)
	defer main_tf.Close()

	//write variable.tf file
	variabletf := "#accesskeyid\n"+
				  "variable \"accesskey\" {\n default = \""+accesskey+"\"\n}\n"+
				  "#secrectaccesskey\n"+
				  "variable \"secretkey\" {\n default = \""+secretkey+"\"\n}\n"+
				  "#region\n"+
				  "variable \"region\" {\n default = \"us-east-1\"\n}\n"+
				  "#custom vpc cidr block\n"+
				  "variable \"custom_vpc_cidr\" {\n default = \""+custom_vpc_cidr+"\"\n}\n"+
				  "#availability zone\n"+
				  "data \"aws_availability_zones\" \"azs\" {}\n"+
				  "#public subnet cidr block\n"+
				  "variable \"publicsubnetcidr\" {\n default = \""+custom_subnet_public_cidr+"\"\n}\n"

	 _,err3 := file_variable_tf.WriteString(variabletf)
	 check_err(err3)

	//write main.tf file
	maintf := "#configure AWS provider\n"+
			  "provider \"aws\" {\n region = var.region\n access_key = var.accesskey\n secret_key = var.secretkey\n}\n"+
			  "#create custom vpc\n"+
			  "resource \"aws_vpc\" \"custom_vpc\" {\n"+
			  " cidr_block = var.custom_vpc_cidr\n"+
			  " enable_dns_support = true\n"+
			  " enable_dns_hostnames = true\n"+
			  " tags = {\n \"Name\" = \""+custom_vpc+"\"\n}\n}\n"+
			  "#create public subnet\n"+
			  "resource \"aws_subnet\" \""+custom_subnet_public+"\"{\n"+
			  " vpc_id = aws_vpc.custom_vpc.id\n"+
			  " cidr_block = var.publicsubnetcidr\n"+
			  " availability_zone = data.aws_availability_zones.azs.names[0]\n"+
			  " tags = {\n \"Name\" = \""+custom_subnet_public+"\"\n}\n}\n"
			  
	_,err4 := main_tf.WriteString(maintf)
	check_err(err4)			  
}