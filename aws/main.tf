provider "aws" {
  region     = var.AWS_REGION
  access_key = var.AWS_ACCESS_KEY
  secret_key = var.AWS_SECRET_KEY
}

resource "aws_instance" "mysql_server" {
  ami                    = var.AWS_AMI[var.AWS_REGION]
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.mysql_sg.id]

  key_name = "grpc_project"

  #security_groups = [aws_security_group.mysql_sg.name]

  user_data = file("${path.module}/userdata.sh")

  tags = {
    Name = "MySQL-EC2"
  }
}

resource "aws_security_group" "mysql_sg" {
  name        = "mysql-sg"
  description = "configuration des connections"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # ⚠️ À restreindre avec ton IP pour la sécurité
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_key_pair" "grpc_project" {
  key_name   = "grpc_project"
  public_key = var.SSH_PUBLIC_KEY
}

output "mysql_instance_ip" {
  value = aws_instance.mysql_server.public_ip
}



# mysql_instance_ip = "35.180.97.228"
# mysql_public_ip = "35.180.97.228"