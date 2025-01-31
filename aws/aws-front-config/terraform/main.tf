# Fournisseur AWS
provider "aws" {
  region = var.aws_region  # Utilisation de la variable pour la région
}

# Génération d'une clé SSH
resource "tls_private_key" "ssh_key" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

# Création d'un Key Pair dans AWS
resource "aws_key_pair" "generated_key" {
  key_name   = var.key_name  # Utilisation de la variable pour le nom de la clé
  public_key = tls_private_key.ssh_key.public_key_openssh
}

# Security Group pour SSH et Vue.js
resource "aws_security_group" "ssh_access" {
  name        = var.security_group_name  # Utilisation de la variable pour le nom du groupe de sécurité
  description = "Allow SSH traffic"

  # Ouvrir les ports SSH
  dynamic "ingress" {
    for_each = var.ssh_ports
    content {
      from_port   = ingress.value
      to_port     = ingress.value
      protocol    = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  # Ouvrir les ports Vue.js ou autres
  dynamic "ingress" {
    for_each = var.vue_ports
    content {
      from_port   = ingress.value
      to_port     = ingress.value
      protocol    = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


# Instance EC2 Debian
resource "aws_instance" "debian_instance" {
  ami           = var.ami_id  # Utilisation de la variable pour l'AMI
  instance_type = var.instance_type  # Utilisation de la variable pour le type d'instance

  key_name      = aws_key_pair.generated_key.key_name
  security_groups = [aws_security_group.ssh_access.name]

  tags = {
    Name = var.instance_name  # Utilisation de la variable pour le nom de l'instance
  }
}

# 📌 Générer l'inventaire pour Ansible
resource "local_file" "ansible_inventory" {
  content  = <<-EOT
  [web]
  ${aws_instance.debian_instance.public_ip} ansible_user=${var.ansible_user} ansible_ssh_private_key_file=../ssh_key.pem
  EOT
  filename = "../ansible/inventory"
}

# 📌 Exécution d’Ansible après Terraform
resource "null_resource" "ansible_provision" {
  depends_on = [aws_instance.debian_instance]

  provisioner "local-exec" {
    command = "ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i ../ansible/inventory ../ansible/playbook.yml"
  }
}

# 📌 Exporter la clé privée localement
resource "local_file" "ssh_private_key" {
  content        = tls_private_key.ssh_key.private_key_pem
  filename       = "../ssh_key.pem"
  file_permission = "0600"  # Sécurise la clé SSH
}

# 📌 Exporter l'IP publique de l'instance dans un fichier
resource "local_file" "instance_ip" {
  content  = aws_instance.debian_instance.public_ip
  filename = "../instance_ip.txt"
}
