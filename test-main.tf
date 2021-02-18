#configure aws provider
provider "aws" {
region = var.test_region
access_key = var.test_accesskey
secret_key = var.test_secretkey
}
#create custom public vpc
resource "aws_vpc" "custom_vpc" {
 cidr_block = var.test_vpcidr
 enable_dns_support = true
 enable_dns_hostnames = true
 tags = {
 "Name" = "test_vpc"
}
}
#create custom public subnet
resource "aws_subnet" "custom_public_subnet" {
 vpc_id = aws_vpc.custom_vpc.id
 cidr_block = var.test_publicsubnetcidr
 availability_zone = data.aws_availability_zones.azs.names[0]
 tags = {
 "Name" = "test_publicsubnet"
}
}
#create custom internet gateway
resource "aws_internet_gateway" "customigw" {
 vpc_id = aws_vpc.custom_vpc.id
 tags = {
"Name" = "test_customigw"
}
}
#create public route table
resource "aws_route_table" "publicroute" {
 vpc_id = aws_vpc.custom_vpc.id
 route {
 cidr_block = "0.0.0.0/0" 
 gateway_id = aws_internet_gateway.customigw.id
 }
 tags = {
 "Name" = "test_publicroute"
}
}
#associate public subnet to public route
resource "aws_route_table_association" "publicsubacc" {
 subnet_id = aws_subnet.custom_public_subnet.id
 route_table_id = aws_route_table.publicroute.id 
}
#create custom security group for public subnet
resource "aws_security_group" "reverse_proxy_sg" {
 name = "reverse-proxy-sg"
 description = "security group for reverse proxy"
 vpc_id = aws_vpc.custom_vpc.id
 ingress {
   description = "allow 22 port"
   from_port = 22
   to_port = 22
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 80 port"
   from_port = 80
   to_port = 80
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 443 port"
   from_port = 443
   to_port = 443
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 tags = {
 	"Name" = "Reverse-Proxy-SG"
}
}
#create custom private subnet
resource "aws_subnet" "custom_private_subnet" {
 vpc_id = aws_vpc.custom_vpc.id
 cidr_block = var.test_privatesubnetcidr
 availability_zone = data.aws_availability_zones.azs.names[1]
 tags = {
 "Name" = "test_privatesubnet"
}
}
#create private route table
resource "aws_route_table" "privateroute" {
 vpc_id = aws_vpc.custom_vpc.id
 route {
 cidr_block = "0.0.0.0/0" 
 nat_gateway_id = aws_nat_gateway.awsnat.id
 }
 tags = {
 "Name" = "test_privateroute"
}
}
#associate private subnet to private route
resource "aws_route_table_association" "privatesubacc" {
 subnet_id = aws_subnet.custom_private_subnet.id
 route_table_id = aws_route_table.privateroute.id 
}
#elastic ip
resource "aws_eip" "awseip" {
vpc = true 
}
#create security group for ELB,FE,MICRO and MYSQL
resource "aws_security_group" "elb_sg" {
 name = "elb-sg"
 description = "security group for ELB"
 vpc_id = aws_vpc.custom_vpc.id
 ingress {
   description = "allow 22 port"
   from_port = 22
   to_port = 22
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 80 port"
   from_port = 80
   to_port = 80
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 443 port"
   from_port = 443
   to_port = 443
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 tags = {
 	"Name" = "ELB-SG"
}
}
resource "aws_security_group" "fe_sg" {
 name = "fe-sg"
 description = "security group for FE"
 vpc_id = aws_vpc.custom_vpc.id
 ingress {
   description = "allow 22 port"
   from_port = 22
   to_port = 22
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 80 port"
   from_port = 80
   to_port = 80
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 443 port"
   from_port = 443
   to_port = 443
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 tags = {
 	"Name" = "FE-SG"
}
}
resource "aws_security_group" "micro_sg" {
 name = "micro-sg"
 description = "security group for micro-services"
 vpc_id = aws_vpc.custom_vpc.id
 ingress {
   description = "allow 22 port"
   from_port = 22
   to_port = 22
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 80 port"
   from_port = 80
   to_port = 80
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 443 port"
   from_port = 443
   to_port = 443
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 tags = {
 	"Name" = "MICRO-SG"
}
}
resource "aws_security_group" "mysql_sg" {
 name = "mysql-sg"
 description = "security group for MYSQL"
 vpc_id = aws_vpc.custom_vpc.id
 ingress {
   description = "allow 22 port"
   from_port = 22
   to_port = 22
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 80 port"
   from_port = 80
   to_port = 80
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 443 port"
   from_port = 443
   to_port = 443
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 ingress {
   description = "allow 3306 port"
   from_port = 3306
   to_port = 3306
   protocol = "tcp"
   cidr_blocks = ["0.0.0.0/0"]
}
 tags = {
 	"Name" = "MYSQL-SG"
}
}
#create network interface for reverse proxy
resource "aws_network_interface" "reverse_proxy_nic" {
 subnet_id = aws_subnet.custom_public_subnet.id
 associate_with_private_ip = [var.reverse_proxy_private_ip]
 security_groups = [ aws_security_group.reverse_proxy_sg.id ]
}
#create Reverse Proxy EC2 in public subnet
resource "aws_instance" "reverse_proxy" {
 ami = var.ami1
 instance_type = "t2.micro"
 availability_zone = data.aws_availability_zones.azs.names[0]
 security_groups = [ aws_security_group.reverse_proxy_sg.id ]
 subnet_id = aws_subnet.custom_public_subnet.id
 associate_public_ip_address = true
 tags = {
 "Name" = "test_reverseproxy"
}
}
#create FE-ELB in private subnet
resource "aws_instance" "fe-elb" {
 ami = var.ami2
 instance_type = "t2.micro"
 availability_zone = data.aws_availability_zones.azs.names[1]
 security_groups = [ aws_security_group.elb_sg.id ]
 subnet_id = aws_subnet.custom_private_subnet.id
 tags = {
 "Name" = "test_fe-elb"
}
}
#nat setup
resource "aws_nat_gateway" "awsnat" {
 allocation_id = aws_eip.awseip.id
 subnet_id = aws_subnet.custom_public_subnet.id
 tags = {
 	"Name" = "nat"
}
