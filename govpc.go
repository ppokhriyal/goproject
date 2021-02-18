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
		"default = \"10.0.1.0/24\" \n}\n"+
	"#ami for reverse proxy\n"+
	"variable \"ami1\" {\n"+
		"default = \"\" \n}\n"+
	"#ami for FE-ELB\n"+
	"variable \"ami2\" {\n"+
		"default = \"\" \n}\n"+
	"#ami for BE-ELB\n"+
	"variable \"ami3\" {\n"+
		"default = \"\" \n}\n"+
	"#ami for FE-1 FE-2\n"+
	"variable \"ami4\" {\n"+
		"default = \"\" \n}\n"+
	"#ami for MICRO-1 MICRO-2\n"+
	"variable \"ami5\" {\n"+
		"default = \"\" \n}\n"+
	"#ami for MySQL\n"+
	"variable \"ami6\" {\n"+
		"default = \"\" \n}\n"


	_,variablerr := project_variable_tf.WriteString(variabletf)
	check_err(variablerr)

	// write security group for public
	// Reverse-Proxy-SG 
	reverse_proxy_sg := "resource \"aws_security_group\" \"reverse_proxy_sg\" {\n"+
					   " name = \"reverse-proxy-sg\"\n"+
					   " description = \"security group for reverse proxy\"\n"+
					   " vpc_id = aws_vpc.custom_vpc.id\n"+
					   " ingress {\n"+
					   "   description = \"allow 22 port\"\n"+
					   "   from_port = 22\n"+
					   "   to_port = 22\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 80 port\"\n"+
					   "   from_port = 80\n"+
					   "   to_port = 80\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 443 port\"\n"+
					   "   from_port = 443\n"+
					   "   to_port = 443\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " tags = {\n"+
				 	   " 	\"Name\" = \"Reverse-Proxy-SG\"\n}\n}\n"
	// ELB-SG
	elb_sg := "resource \"aws_security_group\" \"elb_sg\" {\n"+
					   " name = \"elb-sg\"\n"+
					   " description = \"security group for ELB\"\n"+
					   " vpc_id = aws_vpc.custom_vpc.id\n"+
					   " ingress {\n"+
					   "   description = \"allow 22 port\"\n"+
					   "   from_port = 22\n"+
					   "   to_port = 22\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 80 port\"\n"+
					   "   from_port = 80\n"+
					   "   to_port = 80\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 443 port\"\n"+
					   "   from_port = 443\n"+
					   "   to_port = 443\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " tags = {\n"+
				 	   " 	\"Name\" = \"ELB-SG\"\n}\n}\n"
	// FE-SG
	fe_sg := "resource \"aws_security_group\" \"fe_sg\" {\n"+
					   " name = \"fe-sg\"\n"+
					   " description = \"security group for FE\"\n"+
					   " vpc_id = aws_vpc.custom_vpc.id\n"+
					   " ingress {\n"+
					   "   description = \"allow 22 port\"\n"+
					   "   from_port = 22\n"+
					   "   to_port = 22\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 80 port\"\n"+
					   "   from_port = 80\n"+
					   "   to_port = 80\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 443 port\"\n"+
					   "   from_port = 443\n"+
					   "   to_port = 443\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " tags = {\n"+
				 	   " 	\"Name\" = \"FE-SG\"\n}\n}\n"

	micro_sg := "resource \"aws_security_group\" \"micro_sg\" {\n"+
					   " name = \"micro-sg\"\n"+
					   " description = \"security group for micro-services\"\n"+
					   " vpc_id = aws_vpc.custom_vpc.id\n"+
					   " ingress {\n"+
					   "   description = \"allow 22 port\"\n"+
					   "   from_port = 22\n"+
					   "   to_port = 22\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 80 port\"\n"+
					   "   from_port = 80\n"+
					   "   to_port = 80\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 443 port\"\n"+
					   "   from_port = 443\n"+
					   "   to_port = 443\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " tags = {\n"+
				 	   " 	\"Name\" = \"MICRO-SG\"\n}\n}\n"

	mysql_sg := "resource \"aws_security_group\" \"mysql_sg\" {\n"+
					   " name = \"mysql-sg\"\n"+
					   " description = \"security group for MYSQL\"\n"+
					   " vpc_id = aws_vpc.custom_vpc.id\n"+
					   " ingress {\n"+
					   "   description = \"allow 22 port\"\n"+
					   "   from_port = 22\n"+
					   "   to_port = 22\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 80 port\"\n"+
					   "   from_port = 80\n"+
					   "   to_port = 80\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 443 port\"\n"+
					   "   from_port = 443\n"+
					   "   to_port = 443\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " ingress {\n"+
					   "   description = \"allow 3306 port\"\n"+
					   "   from_port = 3306\n"+
					   "   to_port = 3306\n"+
					   "   protocol = \"tcp\"\n"+
					   "   cidr_blocks = [\"0.0.0.0/0\"]\n}\n"+
					   " tags = {\n"+
				 	   " 	\"Name\" = \"MYSQL-SG\"\n}\n}\n"

	// write main.tf
	maintf := "#configure aws provider\n"+
	"provider \"aws\" {\n"+
		"region = var."+projectname+"_region\n"+
		"access_key = var."+projectname+"_accesskey\n"+
		"secret_key = var."+projectname+"_secretkey\n}\n"+
	"#create custom public vpc\n"+
	"resource \"aws_vpc\" \"custom_vpc\" {\n"+
	" cidr_block = var."+projectname+"_vpcidr\n"+
	" enable_dns_support = true\n"+
	" enable_dns_hostnames = true\n"+
	" tags = {\n \"Name\" = \""+projectname+"_vpc\"\n}\n}\n"+
	"#create custom public subnet\n"+
	"resource \"aws_subnet\" \"custom_public_subnet\" {\n"+
	" vpc_id = aws_vpc.custom_vpc.id\n"+
	" cidr_block = var."+projectname+"_publicsubnetcidr\n"+
	" availability_zone = data.aws_availability_zones.azs.names[0]\n"+
	" tags = {\n \"Name\" = \""+projectname+"_publicsubnet\"\n}\n}\n"+
	"#create custom internet gateway\n"+
	"resource \"aws_internet_gateway\" \"customigw\" {\n"+
	" vpc_id = aws_vpc.custom_vpc.id\n"+
	" tags = {\n\"Name\" = \""+projectname+"_customigw\"\n}\n}\n"+
	"#create public route table\n"+
	"resource \"aws_route_table\" \"publicroute\" {\n"+
	" vpc_id = aws_vpc.custom_vpc.id\n"+
	" route {\n cidr_block = \"0.0.0.0/0\" \n gateway_id = aws_internet_gateway.customigw.id\n }\n tags = {\n \"Name\" = \""+projectname+"_publicroute\"\n}\n}\n"+
	"#associate public subnet to public route\n"+
	"resource \"aws_route_table_association\" \"publicsubacc\" {\n"+
	" subnet_id = aws_subnet.custom_public_subnet.id\n"+
	" route_table_id = aws_route_table.publicroute.id \n}\n"+
	"#create custom security group for public subnet\n"+reverse_proxy_sg+
	"#create custom private subnet\n"+
	"resource \"aws_subnet\" \"custom_private_subnet\" {\n"+
	" vpc_id = aws_vpc.custom_vpc.id\n"+
	" cidr_block = var."+projectname+"_privatesubnetcidr\n"+
	" availability_zone = data.aws_availability_zones.azs.names[1]\n"+
	" tags = {\n \"Name\" = \""+projectname+"_privatesubnet\"\n}\n}\n"+
	"#create private route table\n"+
	"resource \"aws_route_table\" \"privateroute\" {\n"+
	" vpc_id = aws_vpc.custom_vpc.id\n"+
	" route {\n cidr_block = \"0.0.0.0/0\" \n nat_gateway_id = aws_nat_gateway.awsnat.id\n }\n tags = {\n \"Name\" = \""+projectname+"_privateroute\"\n}\n}\n"+
	"#associate private subnet to private route\n"+
	"resource \"aws_route_table_association\" \"privatesubacc\" {\n"+
	" subnet_id = aws_subnet.custom_private_subnet.id\n"+
	" route_table_id = aws_route_table.privateroute.id \n}\n"+
	"#elastic ip\n"+
	"resource \"aws_eip\" \"awseip\" {\n"+
	 "vpc = true \n}\n"+
	"#create security group for ELB,FE,MICRO and MYSQL\n"+elb_sg+fe_sg+micro_sg+mysql_sg+
	"#create Reverse Proxy EC2 in public subnet\n"+
	"resource \"aws_instance\" \"reverse_proxy\" {\n"+
	" ami = var.ami1\n"+
	" instance_type = \"t2.micro\"\n"+
	" availability_zone = data.aws_availability_zones.azs.names[0]\n"+
	" security_groups = [ aws_security_group.reverse_proxy_sg.id ]\n"+
	" subnet_id = aws_subnet.custom_public_subnet.id\n"+
	" associate_public_ip_address = true\n"+
	" tags = {\n \"Name\" = \""+projectname+"_reverseproxy\"\n}\n}\n"+
	"#create FE-ELB in private subnet\n"+
	"resource \"aws_instance\" \"fe-elb\" {\n"+
	" ami = var.ami2\n"+
	" instance_type = \"t2.micro\"\n"+
	" availability_zone = data.aws_availability_zones.azs.names[1]\n"+
	" security_groups = [ aws_security_group.elb_sg.id ]\n"+
	" subnet_id = aws_subnet.custom_private_subnet.id\n"+
	" tags = {\n \"Name\" = \""+projectname+"_fe-elb\"\n}\n}\n"


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
