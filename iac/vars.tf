## Lambda Component
variable "lambda-name" {
    default = "lambda-text2vec"
    type    = string
}

## Tags
variable "tag-project" {
    default = "ie2"
    type    = string
}

variable "tag-environment" {
    default = "prod"
    type    = string
}

## Region
variable "region" {
    default = "us-east-1"
    type    = string
}

## Project Variables
variable "filename" {
    default = "bootstrap.zip"
    type    = string
}

variable "iac-root-bucket" {
    default = "ie2-iac"
    type    = string
}

variable "aws-runtime" {
    default = "provided.al2023"
    type    = string
}

variable "architecture" {
    default = "arm64"
    type    = string
}

variable "handler" {
    default = "bootstrap"
    type    = string
}