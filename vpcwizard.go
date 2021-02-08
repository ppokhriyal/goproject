/* This Program enables you to configure VPC configuration
	1. VPC with Single Public Subnet
	2. VPC with Public and Private Subnet
*/
package main

import(
	"fmt"
	"os"
	"os/exec"
	"time"
)

//clear screen
func clear_screen(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
//function count the accesskey id length.valid key length is 20
func validate_accesskey_count(accesskey string) int{
	var result int
	result = len(accesskey)
	return result
}

//function count the secretkey id length.valid key length is 40
func validate_secretkey_count(secretkey string) int{
	var result int
	result = len(secretkey)
	return result
}
//function check errors
func check_err(e error){
	if e !=nil {
		panic(e)
	}
}

//single public subnet
func vpc_single_public_subnet(){
	
	//color code
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorBackground := "\033[100m"
	var azselected string
	var custom_vpc string
	var vpc_region int
	var selected_region string
	var custom_vpc_cidr string
	var custom_public_subnet string
	var custom_public_subnet_cidr string
	var ec2_name string
	var ec2_ami string
	var ec2_type string


	fmt.Println("--------------------------")
	fmt.Println("VPC CLI Management Console")
	fmt.Println("--------------------------")
	fmt.Println("VPC with Single Public Subnet")
	fmt.Println("=============================")
	intro := `Your instances run in a private, isolated section of the AWS cloud with
direct access to the Internet. Network access control lists and security
groups can be used to provide strict control over inbound and outbound
network traffic to your instances.`
	fmt.Println(intro)
	fmt.Println("\n[ 1 ] Continue")
	fmt.Println("[ 2 ] Go back to main menu")
	fmt.Println("[ 0 ] Exit wizard\n")
	fmt.Print("Enter an option : ")
	var option int
	fmt.Scanln(&option)

	switch {
	case option == 1:
		clear_screen()
		fmt.Println("--------------------------")
		fmt.Println("VPC CLI Management Console")
		fmt.Println("--------------------------")
		fmt.Println("VPC with Single Public Subnet")
		fmt.Println("=============================")
		//select region
		fmt.Println("\nSelect Region\n")
		select_region := "US Region\n********* \n1) US East (N. Virginia) us-east-1\n"+
		"2) US East (Ohio) us-east-2\n"+
		"3) US West (N. California) us-west-1\n"+
		"4) US West (Oregon) us-west-2\n\n"+
		"Africa Region\n*************\n"+"5) Cape Town af-south-1\n\n"+
		"Asia Pacfic Region\n******************\n"+
		"6) Hong Kong ap-east-1\n"+
		"7) Mumbai ap-south-1\n"+
		"8) Seoul ap-northeast-2\n"+
		"9) Singapore ap-southeast-1\n"+
		"10) Sydney ap-southeast-2\n"+
		"11) Tokyo ap-northeast-1\n\n"+
		"Canada Region\n*************\n"+
		"12) Central ca-central-1\n\n"+
		"Europe Region\n*************\n"+
		"13) Frankfurt eu-central-1\n"+
		"14) Ireland eu-west-1\n"+
		"15) London eu-west-2\n"+
		"16) Milan eu-south-1\n"+
		"17) Paris eu-west-3\n"+
		"18) Stockholm eu-north-1\n\n"+
		"Middle East\n***********\n"+
		"19) Bahrain me-south-1\n\n"+
		"South America Region\n********************\n"+
		"20) Sao Paulo sa-east-1\n\n"

		fmt.Println(select_region)

		fmt.Print("Enter an option : ")
		fmt.Scanln(&vpc_region)
		

		switch {
		case vpc_region == 1:
			selected_region = "us-east-1"
		case vpc_region == 2:
			selected_region = "us-east-2"
		case vpc_region == 3:
			selected_region = "us-west-1"
		case vpc_region == 4:
			selected_region = "us-west-2"
		case vpc_region == 5:
			selected_region = "af-south-1"
		case vpc_region == 6:
			selected_region = "ap-east-1"
		case vpc_region == 7:
			selected_region = "ap-south-1"
		case vpc_region == 8:
			selected_region = "ap-northeast-2"
		case vpc_region == 9:
			selected_region = "ap-southeast-2"
		case vpc_region == 10:
			selected_region = "ap-southeast-2"
		case vpc_region == 11:
			selected_region = "ap-northeast-1"
		case vpc_region == 12:
			selected_region = "ca-central-1"
		case vpc_region == 13:
			selected_region = "eu-central-1"
		case vpc_region == 14:
			selected_region = "eu-west-1"
		case vpc_region == 15:
			selected_region = "eu-west-2"
		case vpc_region == 16:
			selected_region = "eu-south-1"
		case vpc_region == 17:
			selected_region = "eu-west-3"
		case vpc_region == 18:
			selected_region = "eu-north-1"
		case vpc_region == 19:
			selected_region = "me-south-1"
		case vpc_region == 20:
			selected_region = "sa-east-1"
		
		}

		//get custom VPC name
		clear_screen()
		fmt.Println("--------------------------")
		fmt.Println("VPC CLI Management Console")
		fmt.Println("--------------------------")
		fmt.Println("VPC with Single Public Subnet")
		fmt.Println("=============================")
		
		fmt.Print("VPC name : ")
		
		fmt.Scanln(&custom_vpc)

		//get custom VPC CIDR block
		fmt.Print("IPv4 CIDR block (e.g 10.0.0.0/16) : ")
		
		fmt.Scanln(&custom_vpc_cidr)

		//get public subnet name
		fmt.Print("\nPublic Subnet name : ")
		
		fmt.Scanln(&custom_public_subnet)

		//get public subnet cidr
		fmt.Print("Public subnet's IPv4 CIDR (e.g 10.0.0.0/24) : ")
		
		fmt.Scanln(&custom_public_subnet_cidr)

		//get availability zone
		fmt.Println("\nSelect Availability Zone available in "+selected_region+" region\n")
		
		//check for selected regions
		var azs string

		switch {

		case selected_region == "us-east-1":
			azs="1) us-east-1a\t4) us-east-1d\n"+
			"2) us-east-1b\t5) us-east-1e\n"+
			"3) us-east-1c\t6) us-east-1f\n"
			fmt.Println(azs)
			var azoption int
			fmt.Print("Enter an option : ")
			fmt.Scanln(&azoption)
			switch {
			case azoption == 1:
				azselected = "us-east-1a"
			case azoption == 2:
				azselected = "us-east-1b"
			case azoption == 3:
				azselected = "us-east-1c"
			case azoption == 4:
				azselected = "us-east-1d"
			case azoption == 5:
				azselected = "us-east-1e"
			case azoption == 6:
				azselected = "us-east-1f"
			case azoption != 1 || azoption != 2 || azoption != 3 || azoption != 4 || azoption != 5 || azoption != 6	:
				fmt.Println(string(colorRed),"\nError: Invalid option",string(colorReset))
			}
		case selected_region == "us-east-2":
			azs="1) us-east-2a\n"+
			"2) us-east-2b\n"+
			"3) us-east-2c\n"
			fmt.Println(azs)
			var azoption int
			fmt.Print("Enter an option : ")
			fmt.Scanln(&azoption)
			switch {
			case azoption == 1:
				azselected = "us-east-2a"
			case azoption == 2:
				azselected = "us-east-2b"
			case azoption == 3:
				azselected = "us-east-2c"
			case azoption != 1 || azoption != 2 || azoption != 3 :
				fmt.Println(string(colorRed),"\nError: Invalid option",string(colorReset))	
			}
		case selected_region == "us-west-1":
			azs="1) us-west-1a\n"+
			"2) us-west-1c"
			fmt.Println(azs)
			var azoption int
			fmt.Print("Enter an option : ")
			fmt.Scanln(&azoption)
			switch{
			case azoption == 1:
				azselected = "us-west-1a"
			case azoption == 2:
				azselected = "us-east-1c"
			case azoption != 1 || azoption != 2 :
				fmt.Println(string(colorRed),"\nError: Invalid option",string(colorReset))
			}
		case selected_region == "us-west-2":
			azs="1) us-west-2a\n"+
			"2) us-west-2b\n"+
			"3) us-west-2c\n"+
			"4) us-west-2d\n"
			fmt.Println(azs)
			var azoption int
			fmt.Print("Enter an option : ")
			fmt.Scanln(&azoption)
			switch{
			case azoption == 1:
				azselected = "us-west-2a"
			case azoption == 2:
				azselected = "us-west-2b"
			case azoption == 3:
				azselected = "us-west-2c"
			case azoption == 4:
				azselected = "us-west-2d"
			case azoption != 1 || azoption != 2 || azoption != 3 || azoption != 4 :
				fmt.Println(string(colorRed),"\nError: Invalid option",string(colorReset))	
			}
		}
	case option == 2:
		clear_screen()
		main()
	case option == 0:
		fmt.Println("Exiting wizard\n")
		os.Exit(0)
	case option != 1 || option != 2 || option != 0:
		fmt.Println(string(colorRed),"\nError: Invalid option",string(colorReset))
	}
	//build ec2 instance
	fmt.Print("\nEnter EC2 instance name : ")
	fmt.Scanln(&ec2_name)
	fmt.Print("Enter ami : ")
	fmt.Scanln(&ec2_ami)
	fmt.Print("Enter instance type : ")
	fmt.Scanln(&ec2_type)


	//review vpc configuration
	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("VPC CLI Management Console")
	fmt.Println("--------------------------")
	fmt.Println("VPC with Single Public Subnet")
	fmt.Println("=============================")
	fmt.Println("Review your VPC configuration\n")
	
	fmt.Println(string(colorBackground),"VPC : "+custom_vpc,string(colorReset))
	fmt.Println(string(colorBackground),"Region : "+selected_region,string(colorReset))
	fmt.Println(string(colorBackground),"Availability Zone : "+azselected,string(colorReset))
	fmt.Println(string(colorBackground),"IPv4 CIDR block : "+custom_vpc_cidr,string(colorReset))
	fmt.Println(string(colorBackground),"Public Subnet : "+custom_public_subnet,string(colorReset))
	fmt.Println(string(colorBackground),"Public Subnet CIDR block : "+custom_public_subnet_cidr,string(colorReset))
	fmt.Println(string(colorBackground),"Ec2 Instance  : "+ec2_name,string(colorReset))
	fmt.Println(string(colorBackground),"Ec2 Instance ami : "+ec2_ami,string(colorReset))
	fmt.Println(string(colorBackground),"Ec2 Instace type  : "+ec2_type,string(colorReset))
	fmt.Println("\n")
	var reviewoption int
	fmt.Println("[ 1 ] Continue")
	fmt.Println("[ 2 ] Go back to main menu")
	fmt.Println("[ 0 ] Exit\n")
	fmt.Print("Enter an option : ")
	fmt.Scanln(&reviewoption)
	switch {
	case reviewoption == 1:
		//get access key
		fmt.Println("\nAuthentication")
		fmt.Println("==============")
		fmt.Print("Enter AccessKey : ")
		var accesskey string
		fmt.Scanln(&accesskey)
		//pass access key to validate func
		result_accesskey := validate_accesskey_count(accesskey)
		if result_accesskey < 20 || result_accesskey > 20 {
			fmt.Println(string(colorRed),"\nError: Invalid Access Key", string(colorReset))
			os.Exit(1)
		}
		//get secretkeyid
		fmt.Print("Enter SecretKey : ")
		var secretkey string
		fmt.Scanln(&secretkey)
		//pass secret key to validate func
		result_secretkey := validate_secretkey_count(secretkey)
		if result_secretkey < 40 || result_secretkey > 40 {
			fmt.Println(string(colorRed),"\nError: Invalid Secret Key", string(colorReset))
			os.Exit(1)
		}
		//create-write variable.tf and main.tf
		file_variable_tf,err1 := os.Create("variable.tf")
		check_err(err1)
		defer file_variable_tf.Close()

		main_tf,err2 := os.Create("main.tf")
		check_err(err2)
		defer main_tf.Close()

		//write variable.tf file
		variabletf :="#accesskeyid\n"+
		"variable \"accesskey\" {\n default = \""+accesskey+"\"\n}\n"+
		"#secrectaccesskey\n"+
		"variable \"secretkey\" {\n default = \""+secretkey+"\"\n}\n"+
		"#region\n"+
		"variable \"region\" {\n default = \""+selected_region+"\"\n}\n"+
		"#custom vpc cidr block\n"+
		"variable \"custom_vpc_cidr\" {\n default = \""+custom_vpc_cidr+"\"\n}\n"+
		"#availability zone\n"+
		"variable \"azs\" {\n default = \""+azselected+"\"\n}\n"+
		"#public subnet cidr block\n"+
		"variable \"publicsubnetcidr\" {\n default = \""+custom_public_subnet_cidr+"\"\n}\n"

		_,variablerr := file_variable_tf.WriteString(variabletf)
		check_err(variablerr)

		//write main.tf
		maintf := "#configure AWS provider\n"+
		"provider \"aws\" {\n region = var.region\n access_key = var.accesskey\n secret_key = var.secretkey\n}\n"+
		"#create custom public vpc\n"+
		"resource \"aws_vpc\" \"custom_public_vpc\" {\n"+
		" cidr_block = var.custom_vpc_cidr\n"+
		" enable_dns_support = true\n"+
		" enable_dns_hostnames = true\n"+
		" tags = {\n \"Name\" = \""+custom_vpc+"\"\n}\n}\n"+
		"#create public subnet\n"+
		"resource \"aws_subnet\" \""+custom_public_subnet+"\" {\n"+
		" vpc_id = aws_vpc.custom_public_vpc.id\n"+
		" cidr_block = var.custom_vpc_cidr\n"+
		" availability_zone = var.azs\n"+
		" tags = {\n \"Name\" = \""+custom_public_subnet+"\"\n}\n}\n"+
		"#create custom internet gateway\n"+
		"resource \"aws_internet_gateway\" \"customigw\" {\n"+
		" vpc_id = aws_vpc.custom_public_vpc.id\n"+
		" tags = {\n\"Name\" = \"customigw\"\n}\n}\n"+
		"#create public route table\n"+
		"resource \"aws_route_table\" \"publicroute\" {\n"+
		" vpc_id = aws_vpc.custom_public_vpc.id\n"+
		" route {\n cidr_block = \"0.0.0.0/0\" \n gateway_id = aws_internet_gateway.customigw.id\n }\n tags = {\n \"Name\" = \"publicroute\"\n}\n}\n"+
		"#associate public subnet to public route\n"+
		"resource \"aws_route_table_association\" \"publicsubacc\" {\n"+
		" subnet_id = aws_subnet."+custom_public_subnet+".id\n"+
		" route_table_id = aws_route_table.publicroute.id \n}\n"+
		"#create custom ec2 instance\n"+
		"resource \"aws_instance\" \""+ec2_name+"\" {\n"+
		" ami = \""+ec2_ami+"\"\n"+
		" instance_type = \""+ec2_type+"\"\n"+
		" availability_zone = var.azs\n"+
		" subnet_id = aws_subnet."+custom_public_subnet+".id\n"+
		" associate_public_ip_address = true\n"+
		" tags = {\n \"Name\" = \""+ec2_name+"\"\n}\n}\n"


		_,maintferr := main_tf.WriteString(maintf)
		check_err(maintferr)

		//prepare terraform configuration file
		fmt.Println("\nPreparing your Terraform Configuration ...")
		time.Sleep(3 * time.Second)
		fmt.Println("\nTerraform Configuration main.tf and variable.tf is ready.")
		
		
	case reviewoption == 2:
		clear_screen()
		main()
	case reviewoption == 0:
		os.Exit(0)	
	case reviewoption != 1 || reviewoption != 2 || reviewoption != 0 :
		fmt.Println(string(colorRed),"\nError: Invalid option",string(colorReset))
		os.Exit(0)	
	}
}
func main(){

	//color code
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	//remove main.tf and variable.tf
	os.Remove("main.tf")
	os.Remove("variable.tf")

	//vpc configuration wizard
	fmt.Println("--------------------------")
	fmt.Println("VPC CLI Management Console")
	fmt.Println("--------------------------")
	fmt.Println("Select a VPC Configuration")
	fmt.Println("==========================")
	fmt.Println("[ 1 ] VPC with Single Public Subnet")
	fmt.Println("[ 2 ] VPC with Public and Private Subnet")
	fmt.Println("[ 0 ] Exit wizard\n")
	fmt.Print("Enter an option : ")
	var option int
	fmt.Scanln(&option)

	switch {
	case option == 1:
		clear_screen()
		vpc_single_public_subnet()
	case option == 2:
		fmt.Println("second option selected")
	case option == 0:
		fmt.Println("Exiting wizard\n")
		os.Exit(0)
	case option != 1 || option != 2 || option != 0:
		fmt.Println(string(colorRed),"\nError: Invalid option",string(colorReset))
	}

}

