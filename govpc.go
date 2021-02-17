package main


import (
    "fmt"
    "os"
)


// access key and secret key setup
func setup_access_secret_key(accesskey,secretey string) init {

	return 0
}
// start building project
func start_build_proj(projectname string) int {

	var accesskey  string
	var secretkey string

	// remove old project build
	os.Remove(projectname+"-main.tf")
	os.Remove(projectname+"-variable.tf")

	// create new project build
	os.Create(projectname+"-main.tf")
	os.Create(projectname+"-variable.tf")

	fmt.Println("Building project environment [ "+projectname+" ]")
	fmt.Println("Done")

	fmt.Println("Setup Access/Secret key")
	fmt.Print("Enter access key : ")
	fmt.Scanln(&accesskey)
	fmt.Print("Enter secret key : ")
	fmt.Scanln(&secretkey)

	setup_access_secret_key(accesskey,secretkey)
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
