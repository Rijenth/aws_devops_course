#!/bin/bash
# 🛠️ Mettre à jour le système
sudo apt update -y
sudo apt upgrade -y

# 🐳 Installer Docker
sudo apt install -y docker.io git
sudo systemctl start docker
sudo systemctl enable docker

# 🔧 Ajouter l'utilisateur ubuntu à Docker
sudo usermod -aG docker ubuntu

cd /home/ubuntu

# 🌍 Récupérer l'IP de MySQL via Terraform
MYSQL_HOST=${mysql_instance_ip}  # Injecté par Terraform

# 📌 Cloner le repo (ajoute `--depth 1` pour éviter de cloner tout l'historique Git)
if [ -d "aws_devops_course" ]; then
    sudo rm -rf aws_devops_course
fi

GITHUB_AUTH_TOKEN="${github_auth_token}"
git clone https://Rijenth:$GITHUB_AUTH_TOKEN@github.com/Rijenth/aws_devops_course.git

cd aws_devops_course

# Modification du DB host dans le .env
mv .env.example .env

echo $MYSQL_HOST

sed -i "s/^DB_HOST=.*/DB_HOST=\"$MYSQL_HOST\"/" /home/ubuntu/aws_devops_course/.env

# 🚀 Build puis démarrer l'API Go

sudo docker build -t go-api -f Dockerfile.api.prod .

sudo docker run -d --name go-container -p 12345:12345 go-api

# 🚀 Démarrer Envoy
# sudo docker run -d \
#   --name envoy \
#   -v /home/ubuntu/envoy.yaml:/etc/envoy.yaml \
#   -p 8000:8000 \
#   envoyproxy/envoy:v1.25.0