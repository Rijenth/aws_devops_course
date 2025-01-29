output "mysql_public_ip" {
  value = aws_instance.mysql_server.public_ip
}

output "golang_public_ip" {
  value = aws_instance.golang_server.public_ip
}