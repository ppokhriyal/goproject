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
	fmt.Println("VPC with Single Public Subnet")
	fmt.Println("=============================")
	fmt.Println("Your instances run in a private, isolated section of the AWS cloud with\ndirect access to the Internet. Network access control lists and security\ngroups can be used to provide strict control over inbound and outbound\nnetwork traffic to your instances.")
}
func main(){

	//color code
	colorReset := "\033[0m"
	colorRed := "\033[31m"

	//vpc configuration wizard
	fmt.Println("Select a VPC Configuration")
	fmt.Println("==========================")
	fmt.Println("[ 1 ] VPC with Single Public Subnet")
	fmt.Println("[ 2 ] VPC with Public and Private Subnet\n")
	fmt.Print("Enter an option : ")
	var option int
	fmt.Scanln(&option)

	switch {
	case option == 1:
		clear_screen()
		vpc_single_public_subnet()
	case option == 2:
		fmt.Println("second option selected")
	case option != 1 || option != 2:
		fmt.Println(string(colorRed),"\nError: Invalid option",string(colorReset))
	}

}

