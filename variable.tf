#accesskeyid
variable "accesskey" {
 default = "pppppppppppppppppppp"
}
#secrectaccesskey
variable "secretkey" {
 default = "pppppppppppppppppppppppppppppppppppppppp"
}
#region
variable "region" {
 default = "us-east-2"
}
#custom vpc cidr block
variable "custom_vpc_cidr" {
 default = "10.0.0.0/16"
}
#availability zone
variable "azs" {
 default = "us-east-2b"
}
#public subnet cidr block
variable "publicsubnetcidr" {
 default = "10.0.0.0/14"
}
