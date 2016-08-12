variable "aws_region" {
  default = "ap-northeast-1"
}

variable "target_name" {
  default = "tf-example-cloudwatch-event-target-for-sns"
}

variable "arn_lambda_function_weather" {
  default = "arn:aws:lambda:ap-northeast-1:109572732475:function:tool_weather"
}
