
provider "aws" {
  profile = "prod"
}

module "iam" {
  source = "../modules/iam"
}
