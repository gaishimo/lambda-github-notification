output "lambda_function_arn" {
  value = "${aws_lambda_function.tool_02_github_notification.arn}"
}
output "event_rule_github_notification_arn" {
  value = "${aws_cloudwatch_event_rule.tool_02_github_notification.arn}"
}
