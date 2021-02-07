#configure AWS provider
provider "aws" {
 region = var.region
 access_key = var.accesskey
 secret_key = var.secretkey
}
#create custom public vpc
resource "aws_vpc" "custom_public_vpc" {
 cidr_block = var.custom_vpc_cidr
 enable_dns_support = true
 enable_dns_hostnames = true
 tags = {
 "Name" = "myvpc"
}
}
