#!/bin/bash
# Mettre à jour le système
sudo apt update -y
sudo apt upgrade -y

# Installer les dépendances pour ajouter le dépôt Docker
sudo apt install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# Ajouter la clé GPG officielle de Docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Ajouter le dépôt Docker à APT sources
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Mettre à jour à nouveau pour inclure Docker
sudo apt update -y

# Installer Docker CE (Community Edition)
sudo apt install -y docker-ce docker-ce-cli containerd.io

# Activer et démarrer Docker
sudo systemctl enable docker
sudo systemctl start docker

# Ajouter l'utilisateur ubuntu au groupe Docker (évite le besoin de sudo pour exécuter Docker)
sudo usermod -aG docker ubuntu

# Lancer MySQL avec la dernière image
sudo docker run -d \
  --name mysql-container \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=database \
  -p 3306:3306 \
  --restart always \
  mysql:latest
  