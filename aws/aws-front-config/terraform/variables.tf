# variables.tf

variable "aws_region" {
  description = "La région AWS où les ressources seront déployées"
  default     = "eu-west-3"  # Valeur par défaut
}

variable "key_name" {
  description = "Nom de la clé SSH générée"
  default     = "my-ssh-key"  # Valeur par défaut
}

variable "security_group_name" {
  description = "Nom du groupe de sécurité"
  default     = "allow_ssh"  # Valeur par défaut
}

# Ports à ouvrir pour SSH
variable "ssh_ports" {
  description = "Liste des ports SSH à ouvrir"
  type        = list(number)
  default     = [22]  # Par défaut, ouvrir le port 22 pour SSH
}

# Ports à ouvrir pour Vue.js ou autres services
variable "vue_ports" {
  description = "Liste des ports Vue.js à ouvrir"
  type        = list(number)
}


variable "ansible_ssh_private_key_file" {
  description = "Chemin du fichier de clé privée pour Ansible"
  default     = "../ssh_key.pem"
}


variable "ami_id" {
  description = "ID de l'AMI Debian"
  default     = "ami-01427dce5d2537266"  # Valeur par défaut
}

variable "instance_type" {
  description = "Type de l'instance EC2"
  default     = "t3.micro"  # Valeur par défaut
}

variable "instance_name" {
  description = "Nom de l'instance EC2"
}

variable "ansible_user" {
  description = "Nom de l'utilisateur Ansible"
}
