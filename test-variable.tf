#access key id 
variable "test_accesskey" {
default = "pppppppppppppppppppp"
}
#secret key id 
variable "test_secretkey" {
default = "pppppppppppppppppppppppppppppppppppppppp"
}
#vpc cidr
variable "test_vpcidr" {
default = "10.0.0.0/16" 
}
#region
variable "test_region" {
default = "us-east-1" 
}
#availability zones
data "aws_availability_zones" "azs" {}
#public subnet cidr
variable "test_publicsubnetcidr" {
default = "10.0.0.0/24" 
}
#private subnet cidr
variable "test_privatesubnetcidr" {
default = "10.0.1.0/24" 
}
#ami for reverse proxy
variable "ami1" {
default = "" 
}
#ami for FE-ELB
variable "ami2" {
default = "" 
}
#ami for BE-ELB
variable "ami3" {
default = "" 
}
#ami for FE-1 FE-2
variable "ami4" {
default = "" 
}
#ami for MICRO-1 MICRO-2
variable "ami5" {
default = "" 
}
#ami for MySQL
variable "ami6" {
default = "" 
}
#reverse_proxy private_ip
variable "reverse_proxy_private_ip" {
 default = "10.0.0.40" 
}
#fe-elb private ip
variable "fe-elb_private_ip" {
 default = "10.0.1.41" 
}
#be-elb private ip
variable "be-elb_private_ip" {
 default = "10.0.1.42" 
}
#fe-1 private ip
variable "fe-1_private_ip" {
 default = "10.0.1.51" 
}
#fe-2 private ip
variable "fe-2_private_ip" {
 default = "10.0.1.52" 
}
#micro-1 private ip
variable "micro-1_private_ip" {
 default = "10.0.1.61" 
}
#micro-2 private ip
variable "micro-2_private_ip" {
 default = "10.0.1.62" 
}
#mysql private ip
variable "mysql_private_ip" {
 default = "10.0.1.71" 
}
