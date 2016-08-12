# Lambda Function
resource "aws_lambda_function" "tool_02_github_notification" {
  filename = "empty_lambda_func.zip"
  function_name = "tool_02_github_notification"
  role = "${module.iam.lambda_function_role_id}"
  handler = "_apex_index.handle"
  lifecycle {
    ignore_changes = ["filename", "description", "timeout"]
  }
}

# CloudWatch Event Rules
resource "aws_cloudwatch_event_rule" "tool_02_github_notification" {
  name = "tool_02_github_notification"
  description = "Check github notifications on a regular basis"
  schedule_expression = "rate(3 minutes)"
  is_enabled = true
}

# CloudWwatch Event Targets
resource "aws_cloudwatch_event_target" "github_notification" {
  rule = "${aws_cloudwatch_event_rule.tool_02_github_notification.name}"
  target_id = "tool-github-notification"
  arn = "${aws_lambda_function.tool_02_github_notification.arn}"
}

# Lambda Permission
resource "aws_lambda_permission" "allow_cloudwatch_github_notification" {
  statement_id = "AllowExecutionFromCloudWatchEventGithubNotification"
  action = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.tool_02_github_notification.arn}"
  principal = "events.amazonaws.com"
  source_arn = "${aws_cloudwatch_event_rule.tool_02_github_notification.arn}"
}
