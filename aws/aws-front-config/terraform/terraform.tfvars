# Valeurs des variables pour la configuration Terraform

# Région AWS où les ressources seront déployées
aws_region = "eu-west-3"

# Nom de la clé SSH à générer dans AWS
key_name = "my-ssh-key"

security_group_name = "allow_ssh"
ssh_ports           = [22, 80, 443]  # Ouvrir SSH + HTTP (80) et HTTPS (443)
vue_ports           = [5173]          # Port Vue.js

# Identifiant de l'AMI Debian à utiliser pour l'instance EC2
ami_id = "ami-0359cb6c0c97c6607"

# Type de l'instance EC2
instance_type = "t3.micro"

# Nom de l'instance EC2
instance_name = "VueJS-Instance"

# Utilisateur pour se connecter via Ansible
ansible_user = "admin"

# Chemin du fichier de clé privée pour l'authentification SSH via Ansible
ansible_ssh_private_key_file = "../ssh_key.pem"
