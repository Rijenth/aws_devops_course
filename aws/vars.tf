variable "AWS_ACCESS_KEY" {}

variable "AWS_SECRET_KEY" {}

variable "SSH_PUBLIC_KEY" {}

variable "AWS_REGION" {
  default = "eu-west-3"
}

variable "AWS_AMI" {
  type = map(any)
  default = {
    "eu-west-3" = "ami-06e02ae7bdac6b938"
  }
}
