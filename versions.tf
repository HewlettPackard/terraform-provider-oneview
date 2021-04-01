terraform {
  required_version = ">= 0.13"
  required_providers {
    oneview = {
      source = "HewlettPackard/oneview"
      version = "6.0.0-13" #Check https://registry.terraform.io/v1/providers/hewlettpackard/oneview/versions
    }
  }
}
