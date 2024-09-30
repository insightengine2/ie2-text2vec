resource "aws_lambda_function" "text2vec" {
  architectures     = ["${var.architecture}"]
  filename          = "../dist/${var.filename}"
  function_name     = "${var.lambda-name}"
  role              = data.terraform_remote_state.lambda-state.outputs.lambda-s3readerwriter-role-arn
  source_code_hash  = filebase64sha256("../dist/${var.filename}")
  runtime           = var.aws-runtime
  handler           = var.handler
  timeout           = 180

  environment {
    variables = {
      "ENV_METADATA_FILENAME" = "metadata.json",
      "ENV_EMBEDDING_URL" = "https://not/a/real/url",
      "ENV_EMBEDDING_API_KEY" = "10010101001010",
      "ENV_EMBEDDING_API_HEADER" = "x-api-key"
    }
  }
}