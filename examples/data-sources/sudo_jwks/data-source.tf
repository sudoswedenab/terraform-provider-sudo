resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

data "sudo_jwks" "example" {
  rsa_private_key_pem = tls_private_key.example.private_key_pem
}
