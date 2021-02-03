/* This Program enables you to configure VPC configuration
	1. VPC with Single Public Subnet
	2. VPC with Public and Private Subnet
*/
package main

import(
	"fmt"
	"os"
	"os/exec"
)
//clear screen
func clear_screen(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
//single public subnet
func vpc_single_public_subnet(){
	
	//color code
	colorReset := "\033[0m"
	colorRed := "\033[31m"

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
		select_region := "US Region\n========= \n1) US East (N. Virginia) us-east-1\n"+
		"2) US East (Ohio) us-east-2\n"+
		"3) US West (N. California) us-west-1\n"+
		"4) US West (Oregon) us-west-2\n\n"+
		"Africa Region\n=============\n"+"5) Cape Town af-south-1\n\n"+
		"Asia Pacfic Region\n==================\n"+
		"6) Hong Kong ap-east-1\n"+
		"7) Mumbai ap-south-1\n"+
		"8) Seoul ap-northeast-2\n"+
		"9) Singapore ap-southeast-1\n"+
		"10) Sydney ap-southeast-2\n"+
		"11) Tokyo ap-northeast-1\n\n"+
		"Canada Region\n=============\n"+
		"12) Central ca-central-1\n\n"+
		"Europe Region\n=============\n"+
		"13) Frankfurt eu-central-1\n"+
		"14) Ireland eu-west-1\n"+
		"15) London eu-west-2\n"+
		"16) Milan eu-south-1\n"+
		"17) Paris eu-west-3\n"+
		"18) Stockholm eu-north-1\n\n"+
		"Middle East\n===========\n"+
		"19) Bahrain me-south-1\n\n"+
		"South America Region\n====================\n"+
		"20) Sao Paulo sa-east-1\n\n"

		fmt.Println(select_region)

		fmt.Print("Enter an option : ")
		var vpc_region int
		fmt.Scanln(&vpc_region)
		var selected_region string
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
		var custom_vpc string
		fmt.Scanln(&custom_vpc)

		//get custom VPC CIDR block
		fmt.Print("IPv4 CIDR block (e.g 10.0.0.0/16) : ")
		var custom_vpc_cidr string
		fmt.Scanln(&custom_vpc_cidr)

		//get public subnet name
		fmt.Print("\nPublic Subnet name : ")
		var custom_public_subnet string
		fmt.Scanln(&custom_public_subnet)

		//get public subnet cidr
		fmt.Print("Public subnet's IPv4 CIDR (e.g 10.0.0.0/24) : ")
		var custom_public_subnet_cidr string
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
}
func main(){

	//color code
	colorReset := "\033[0m"
	colorRed := "\033[31m"

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

