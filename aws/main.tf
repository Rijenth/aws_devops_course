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

  user_data = file("${path.module}/database.sh")

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
    cidr_blocks = ["0.0.0.0/0"] # ⚠️ À restreindre car accessible à tous
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

# ---------------------------------------------------- #

resource "aws_instance" "golang_server" {
  ami                    = var.AWS_AMI[var.AWS_REGION]
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.go_sg.id, aws_security_group.envoy_sg.id]
  key_name               = aws_key_pair.grpc_project.key_name
  user_data = templatefile("${path.module}/golang.sh", {
    mysql_instance_ip = aws_instance.mysql_server.public_ip
    github_auth_token = var.GITHUB_AUTH_TOKEN
  })

  tags = {
    Name = "Golang-EC2"
  }

  depends_on = [aws_instance.mysql_server]
}

# 🔐 Security Group pour l'instance Golang
resource "aws_security_group" "go_sg" {
  name = "go-sg"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # ⚠️ À restreindre car accessible à tous
  }

  ingress {
    from_port   = 12345
    to_port     = 12345
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# 🔐 Ajouter un Security Group pour Envoy
resource "aws_security_group" "envoy_sg" {
  name        = "envoy-sg"
  description = "Securite pour Envoy Proxy"

  ingress {
    from_port   = 8000
    to_port     = 8000
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # ⚠️ À restreindre pour la sécurité
  }

  ingress {
    from_port   = 9901
    to_port     = 9901
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
