package main


import (
    "fmt"
    "os"
)

// start building project
func start_build_proj(projectname string) string {
	return projectname

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
