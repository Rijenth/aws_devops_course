# outputs.tf

output "instance_ip" {
  description = "L'adresse IP publique de l'instance EC2"
  value       = aws_instance.debian_instance.public_ip
}

output "private_key" {
  description = "Clé privée SSH générée"
  value       = tls_private_key.ssh_key.private_key_pem
  sensitive   = true
}

output "security_group_name" {
  description = "Nom du groupe de sécurité créé"
  value       = aws_security_group.ssh_access.name
}

output "instance_name" {
  description = "Nom de l'instance EC2"
  value       = aws_instance.debian_instance.tags["Name"]
}

output "ansible_inventory" {
  description = "Inventaire Ansible pour la connexion SSH"
  value       = "[web]\n${aws_instance.debian_instance.public_ip} ansible_user=${var.ansible_user} ansible_ssh_private_key_file=../ssh_key.pem"
}
