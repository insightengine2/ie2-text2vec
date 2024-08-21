terraform {

    required_version = "~> 1.7.0"

    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = "~> 5.65.0"
        }
    }

    // comment out the following block 
    backend "s3" {
        bucket  = "ie2-iac"
        key     = "state/lambdas/lambda-text2vec/terraform.tfstate"
        region  = "us-east-1"
    }
}

provider "aws" {

    region = var.region

    default_tags {
        tags = {
            Environment = var.tag-environment
            Name        = var.tag-project
        }
    }
}

## ROOT PROJECT STATE
##  Useful if you need to reference outputs from the root infrastructure, like referencing ARNs
data "terraform_remote_state" "lambda-state" {
    backend = "s3"
    config = {
        bucket  = "${var.iac-root-bucket}"
        key     = "state/terraform.tfstate"
        region  = "${var.region}"
    }
}