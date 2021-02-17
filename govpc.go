package main


import (
    "fmt"
    "os"
)


// access key and secret key setup
func setup_access_secret_key(accesskey,secretkey string) (int,int) {
	var access_result int
	var secret_result int
	access_result = len(accesskey)
	secret_result = len(secretkey)
	return access_result,secret_result
}
//function check errors
func check_err(e error){
	if e !=nil {
		panic(e)
	}
}
// start building project
func start_build_proj(projectname string) int {

	var accesskey  string
	var secretkey string

	// remove old project build
	os.Remove(projectname+"-main.tf")
	os.Remove(projectname+"-variable.tf")

	// create new project build
	project_variable_tf,err1 := os.Create(projectname+"-variable.tf")
	check_err(err1)
	defer project_variable_tf.Close()

	project_main_tf,err2 := os.Create(projectname+"-main.tf")
	check_err(err2)
	defer project_main_tf.Close()
	

	fmt.Println("Building project environment [ "+projectname+" ]")
	fmt.Println("Setup Access/Secret key")
	fmt.Println("-----------------------")
	fmt.Print("Enter access key : ")
	fmt.Scanln(&accesskey)
	fmt.Print("Enter secret key : ")
	fmt.Scanln(&secretkey)

	result_access,result_secret := setup_access_secret_key(accesskey,secretkey)
	if result_access < 20 || result_access > 20 {
		fmt.Println("\nError: Invalid Access/Secret Key")
		os.Exit(1)
	}
	if result_secret < 40 || result_secret > 40 {
		fmt.Println("\nError: Invalid Access/Secret Key")
		os.Exit(1)
	}
	fmt.Println("Done\n")
	fmt.Println("Building project ...")

	// write variable.tf
	variabletf :="#access key id \n"+
	"variable \""+projectname+"_accesskey\" {\n"+
		"default = \""+accesskey+"\"\n}\n"+
	"#secret key id \n"+
	"variable \""+projectname+"_secretkey\" {\n"+
		"default = \""+secretkey+"\"\n}\n"+
	"#vpc cidr\n"+
	"variable \""+projectname+"_vpcidr\" {\n"+
		"default = \"10.0.0.0/16\" \n}\n"+
	"#region\n"+
	"variable \""+projectname+"_region\" {\n"+
		"default = \"us-east-1\" \n}\n"+
	"#availability zones\n"+
	"data \"aws_availability_zones\" \"azs\" {}\n"+
	"#public subnet cidr\n"+
	"variable \""+projectname+"_publicsubnetcidr\" {\n"+
		"default = \"10.0.0.0/24\" \n}\n"+
	"#private subnet cidr\n"+
	"variable \""+projectname+"_privatesubnetcidr\" {\n"+
		"default = \"10.0.1.0/24\" \n}\n"


	_,variablerr := project_variable_tf.WriteString(variabletf)
	check_err(variablerr)

	// write main.tf
	maintf := "#configure aws provider\n"+
	"provider \"aws\" {\n"+
		"region = var."+projectname+"_region\n"+
		"access_key = var."+projectname+"_accesskey\n"+
		"secret_key = var."+projectname+"_secretkey\n}\n"+
	"#create custom public vpc\n"+
	"resource \"aws_vpc\" \"custom_public_vpc\" {\n"+
	" cidr_block = var."+projectname+"_vpcidr\n"+
	" enable_dns_support = true\n"+
	" enable_dns_hostnames = true\n"+
	" tags = {\n \"Name\" = \""+projectname+"_vpc\"\n}\n}\n"

	_,maintferr := project_main_tf.WriteString(maintf)
	check_err(maintferr)

	return 0
	
}

func main() {

    buildargs := os.Args
    // check argument counts,it should be 2
    if len(buildargs) > 3 {
	fmt.Println("Invalid arguments e.g build projectname")
	os.Exit(1)
    }
    // check argument name,it should be build
    if buildargs[1] != "build" {
	fmt.Println("Invalid aruguments e.g build projectname")
	os.Exit(1)
    }
    start_build_proj(buildargs[2])

}
