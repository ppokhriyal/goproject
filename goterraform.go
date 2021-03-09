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
	Securitygroups []struct {
		Name         string `yaml:"name"`
		Description  string `yaml:"description"`
		InboundPorts []int  `yaml:"inbound_ports"`
	} `yaml:"securitygroups"`
	Buildec2Instances []struct {
		Type         string   `yaml:"type"`
		Name         string   `yaml:"name"`
		Instancetype string   `yaml:"instancetype"`
		Ami          string   `yaml:"ami"`
		Security     []string `yaml:"security"`
		Key          string   `yaml:"key"`
		Hddsize      string   `yaml:"hddsize"`
	} `yaml:"buildec2instances"`
}

func main(){

	//Global Vriables
	var selected_region string
	var selected_accesskey string
	var selected_secretkey string
	var selected_vpcname string

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
	} else {
		//creating new project workspace
		fmt.Println("Creating Project : "+f.Projectname)
		os.MkdirAll(f.Projectname,0755)
	}

	//Build AWS Provider
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
		providertf := "#configure AWS provider\n"+
		"provider \"aws\" {\n"+
		" region = \""+selected_region+"\"\n"+
		" access_key = \""+selected_accesskey+"\"\n"+
		" secret_key = \""+selected_secretkey+"\"\n}\n"
		
		pwd1,_ := os.Getwd()
		fil1,err := os.OpenFile(pwd1+"/"+f.Projectname+"/"+f.Projectname+"_provider.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		check_err(err)
		if _, err := fil1.Write([]byte(providertf)); err != nil {
			log.Fatal(err)
		}
		if err := fil1.Close(); err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("\nError: invalid Access/Secret key")
		os.Exit(1)
	}
	
	//Build VPC
	selected_vpcname = f.Buildvpc.Name
	fmt.Println("VPC Name	 : "+selected_vpcname)
	vpctf := "#Create VPC\n"+
	"resource \"aws_vpc\" \"custom_vpc\" {\n"+
	 "cidr_block = \""+f.Buildvpc.CidrBlock+"\"\n"+
	 "enable_dns_support = "+f.Buildvpc.EnableDNSSupport+"\n"+
	 "enable_dns_hostnames = "+f.Buildvpc.EnableDNSHostnames+"\n"+
	 "tags = {\n"+
	  " Name = \""+f.Projectname+"_vpc\"\n}\n}\n"
	
	pwd2,_ := os.Getwd()
	fil2,err := os.OpenFile(pwd2+"/"+f.Projectname+"/"+f.Projectname+"_vpc.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	check_err(err)
	if _, err := fil2.Write([]byte(vpctf)); err != nil {
		  log.Fatal(err)
	}
	if err := fil2.Close(); err != nil {
	  log.Fatal(err)
	}
	  
	//Build Subnet
	buildsubnet_count := len(f.Buildsubnet)
	count := 0
	for count < buildsubnet_count {
		if f.Buildsubnet[count].Type == "public" {
			//a. create public subnet
			//b. create IGW
			//c. create route table
			//d. subnet association with route table
			subnetf := "#Configure Public Subnet\n"+
			"resource \"aws_subnet\" \"custom_publicsubnet_"+strconv.Itoa(count)+"\" {\n"+
			" vpc_id = aws_vpc.custom_vpc.id\n"+
			" cidr_block = \""+f.Buildsubnet[count].CidrBlock+"\"\n"+
			" availability_zone = \""+f.Buildsubnet[count].AvailabilityZone+"\"\n"+
			" tags = {\n Name = \""+f.Projectname+"_publicsubet_"+strconv.Itoa(count)+"\"\n}\n}\n"+
			"#create IGW for public subnet\n"+
			"resource \"aws_internet_gateway\" \"custom_publicigw_"+strconv.Itoa(count)+"\" {\n"+
			" vpc_id = aws_vpc.custom_vpc.id\n"+
			" tags = {\n Name = \""+f.Projectname+"_publicigw_"+strconv.Itoa(count)+"\"\n}\n}\n"+
			"#create Route for public subnet\n"+
			"resource \"aws_route_table\" \"custom_publicroute_"+strconv.Itoa(count)+"\" {\n"+
			" vpc_id = aws_vpc.custom_vpc.id\n"+
			" route {\n cidr_block = \"0.0.0.0/0\"\n gateway_id = aws_internet_gateway.custom_publicigw_"+strconv.Itoa(count)+".id\n}\n"+
			" tags = {\n Name = \""+f.Projectname+"_publicroute_"+strconv.Itoa(count)+"\"\n}\n}\n"+
			"#associate public subnet to public route\n"+
			"resource \"aws_route_table_association\" \"custom_routepubassociation_"+strconv.Itoa(count)+"\" {\n"+
			" subnet_id = aws_subnet.custom_publicsubnet_"+strconv.Itoa(count)+".id \n"+
			" route_table_id = aws_route_table.custom_publicroute_"+strconv.Itoa(count)+".id \n}\n"
			

			pwd3,_ := os.Getwd()
			fil3,err := os.OpenFile(pwd3+"/"+f.Projectname+"/"+f.Projectname+"_subnet.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			check_err(err)
			if _, err := fil3.Write([]byte(subnetf)); err != nil {
				log.Fatal(err)
			}
			if err := fil3.Close(); err != nil {
				log.Fatal(err)
			}
             
            //Build EIP/NAT
			eipnatf := "#Create EIP\n"+
			" resource \"aws_eip\" \"awseip\" {\n"+
			" vpc = true\n}\n"+
			"#Create NAT\n"+
			"resource \"aws_nat_gateway\" \"awsnat\" {\n"+
			" allocation_id = aws_eip.awseip.id\n"+
			" subnet_id = aws_subnet.custom_publicsubnet_"+strconv.Itoa(count)+".id \n"+
			" tags = {\n Name = \""+f.Projectname+"_nat\"\n}\n}\n"

			pwdd3,_ := os.Getwd()
			fill3,err := os.OpenFile(pwdd3+"/"+f.Projectname+"/"+f.Projectname+"_eipnat.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			check_err(err)
			if _, err := fill3.Write([]byte(eipnatf)); err != nil {
				log.Fatal(err)
			}
			if err := fill3.Close(); err != nil {
				log.Fatal(err)
			}

		} else {

			//a. create private subnet
			//b. create IGW
			//c. create route table
			//d. subnet association with route table

			subnetf := "#Configure Private Subnet\n"+
			"resource \"aws_subnet\" \"custom_privatesubnet_"+strconv.Itoa(count)+"\" {\n"+
			" vpc_id = aws_vpc.custom_vpc.id\n"+
			" cidr_block = \""+f.Buildsubnet[count].CidrBlock+"\"\n"+
			" availability_zone = \""+f.Buildsubnet[count].AvailabilityZone+"\"\n"+
			" tags = {\n Name = \""+f.Projectname+"_privatesubet_"+strconv.Itoa(count)+"\"\n}\n}\n"+
			"#create Route for private subnet\n"+
			"resource \"aws_route_table\" \"custom_privateroute_"+strconv.Itoa(count)+"\" {\n"+
			" vpc_id = aws_vpc.custom_vpc.id\n"+
			" route {\n cidr_block = \"0.0.0.0/0\"\n nat_gateway_id = aws_nat_gateway.awsnat.id \n}\n"+
			" tags = {\n Name = \""+f.Projectname+"_privateroute_"+strconv.Itoa(count)+"\"\n}\n}\n"+
			"#associate private subnet to private route\n"+
			"resource \"aws_route_table_association\" \"custom_routeprivassociation_"+strconv.Itoa(count)+"\" {\n"+
			" subnet_id = aws_subnet.custom_privatesubnet_"+strconv.Itoa(count)+".id \n"+
			" route_table_id = aws_route_table.custom_privateroute_"+strconv.Itoa(count)+".id \n}\n"

			pwd3,_ := os.Getwd()
			fil3,err := os.OpenFile(pwd3+"/"+f.Projectname+"/"+f.Projectname+"_subnet.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			check_err(err)
			if _, err := fil3.Write([]byte(subnetf)); err != nil {
				log.Fatal(err)
			}
			if err := fil3.Close(); err != nil {
				log.Fatal(err)
			}	

		}
		count += 1
	}
	//Build Security Groups
	securitygroups_count := len(f.Securitygroups)
	scount := 0
	for scount < securitygroups_count {
		securitygtf := "#Security Group-"+strconv.Itoa(scount)+"\n"+
					   "resource \"aws_security_group\" \""+f.Securitygroups[scount].Name+"\" {\n"+
		               " name =  \""+f.Securitygroups[scount].Name+"\"\n"+
					   " description = \""+f.Securitygroups[scount].Description+"\"\n"+
					   " vpc_id = aws_vpc.custom_vpc.id\n"

		pwd4,_ := os.Getwd()
		fil4,err := os.OpenFile(pwd4+"/"+f.Projectname+"/"+f.Projectname+"_security.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		check_err(err)
		if _, err := fil4.Write([]byte(securitygtf)); err != nil {
			log.Fatal(err)
		}
		if err := fil4.Close(); err != nil {
		  log.Fatal(err)
		}
		//add ingress ports
		ingressport_count := len(f.Securitygroups[scount].InboundPorts)
		pcount := 0
		for pcount < ingressport_count{
			addingress := " ingress {\n"+
			              " description = \"allow "+strconv.Itoa(f.Securitygroups[scount].InboundPorts[pcount])+" port\"\n"+
						  " from_port = "+strconv.Itoa(f.Securitygroups[scount].InboundPorts[pcount])+"\n"+
						  " to_port = "+strconv.Itoa(f.Securitygroups[scount].InboundPorts[pcount])+"\n"+
						  " protocol = \"tcp\"\n"+
						  " cidr_blocks = [\"0.0.0.0/0\"]\n}\n"

			pwdd4,_ := os.Getwd()
			fill4,err := os.OpenFile(pwdd4+"/"+f.Projectname+"/"+f.Projectname+"_security.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			check_err(err)
			if _, err := fill4.Write([]byte(addingress)); err != nil {
				log.Fatal(err)
			}
			if err := fill4.Close(); err != nil {
				log.Fatal(err)
			}
			pcount += 1
		}
		//add end curly-engress
		curly := " egress {\n"+
		         " from_port = 0\n"+
				 " to_port = 0\n"+
				 " protocol = \"-1\"\n"+
				 " cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+ 
		         " tags = {\n"+
		         "  Name = \""+f.Securitygroups[scount].Name+"\"\n}\n}\n"
		
		pwdd5,_ := os.Getwd()
		fill5,err := os.OpenFile(pwdd5+"/"+f.Projectname+"/"+f.Projectname+"_security.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		check_err(err)
		if _, err := fill5.Write([]byte(curly)); err != nil {
			log.Fatal(err)
		}
		if err := fill5.Close(); err != nil {
			log.Fatal(err)
		}
		
		scount += 1

	}	
}	