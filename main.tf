# This is required for Terraform 0.13+
terraform {
  required_providers {
    pets = {
      version = "~> 0.0.1"
      source  = "drehnstrom/providers/pets"
    }
  }
}
resource "pets_dog" "my-dog" {
  name     = "Noir"
  breed    = "Schnoodle"
}
