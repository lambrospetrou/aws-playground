# How to run multiple applications (services) on the same EC2 instance using Elastic Beanstalk Go Platform

Assume there are 3 applications we want to run on each Elastic Beanstalk instance:
1. Web service 1
2. Web service 2
3. Nginx proxy (deployed automatically by the platform)

## Linux Platforms on Amazon Linux 2

- [Extending Elastic Beanstalk Linux platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-linux-extend.html)

- In AL2 linux platform in Go now the default port is `5000` and it's not passed as `PORT` env variable, unless the configuration is changed.
  + In the previous platform it used to be passed in increments of `100`, e.g. `5000, 5100` for each process in the `Procfile`.

- Also now the Cloudwatch Logs streaming only works for fixed log files, and the first process in the `Procfile` that should be named `web`!
  + https://aws.amazon.com/premiumsupport/knowledge-center/elastic-beanstalk-customized-log-files/
  + See here how to add more logs by using directly the CloudWatch Logs agent `awslogs`: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/AWSHowTo.cloudwatchlogs.html#AWSHowTo.cloudwatchlogs.streaming.custom
  + Sample config: https://github.com/awsdocs/elastic-beanstalk-samples/blob/master/configuration-files/aws-provided/instance-configuration/logs-streamtocloudwatch-linux.config
    * Agent config: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AgentReference.html
  + Files logged by default: `/var/log/eb-engine.log, /var/log/eb-hooks.log, /var/log/web.stdout.log, /var/log/nginx/access.log, /var/log/nginx/error.log`
  + All processes from `Procfile` will have `/var/log/<name>.{stdout,stderr}.log` according to https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-linux-extend.html.
  + Need to setup the right permissions as well:
  ```
  2022-01-25 02:00:28,575 - cwlogs.push.batch - INFO - 6113 - Thread-2 - Creating log group /aws/elasticbeanstalk/Test2-al2-env//var/log/web2.stdout.log.
  2022-01-25 02:00:28,597 - cwlogs.push.batch - WARNING - 6113 - Thread-2 - CreateLogGroup failed with exception An error occurred (AccessDeniedException) when calling the CreateLogGroup operation: User: arn:aws:sts::717437904155:assumed-role/aws-elasticbeanstalk-ec2-role/i-08e1b0e2204017ae2 is not authorized to perform: logs:CreateLogGroup on resource: arn:aws:logs:eu-west-1:717437904155:log-group:/aws/elasticbeanstalk/Test2-al2-env//var/log/web2.stdout.log:log-stream: because no identity-based policy allows the logs:CreateLogGroup action
  2022-01-25 02:00:28,598 - cwlogs.push.batch - WARNING - 6113 - Thread-2 - An error occurred (AccessDeniedException) when calling the CreateLogGroup operation: User: arn:aws:sts::717437904155:assumed-role/aws-elasticbeanstalk-ec2-role/i-08e1b0e2204017ae2 is not authorized to perform: logs:CreateLogGroup on resource: arn:aws:logs:eu-west-1:717437904155:log-group:/aws/elasticbeanstalk/Test2-al2-env//var/log/web2.stdout.log:log-stream: because no identity-based policy allows the logs:CreateLogGroup action
  ```

- For log rotation: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.logging.html
