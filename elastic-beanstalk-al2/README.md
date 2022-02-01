# How to run multiple applications (services) on the same EC2 instance using Elastic Beanstalk Go Platform

Assume there are 3 applications we want to run on each Elastic Beanstalk instance:
1. Web app
2. Background app
3. Nginx proxy (deployed automatically by the platform)

Because of the limitations and custom setup needed for logging with multiple `Procfile` processes I prefer to use [Hivemind](https://github.com/DarthSim/hivemind) or [goreman](https://github.com/mattn/goreman) to manage our own `Procfile` and start that through the `web` process in the main `Procfile` Elastic Beanstalk will start. This however leads to messed log lines with multiple timestamps and color terminal codes, [see issue](https://github.com/DarthSim/hivemind/pull/1#issuecomment-1025004782).

## Linux Platforms on Amazon Linux 2

- [Extending Elastic Beanstalk Linux platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-linux-extend.html)

- In AL2 linux platform in Go now the default port is `5000` and it's not passed as `PORT` env variable, unless the configuration is changed.
  + In the previous platform it used to be passed in increments of `100`, e.g. `5000, 5100` for each process in the `Procfile`.

- The Cloudwatch Logs streaming only works for fixed log files, including the first process in the `Procfile` which **has to be named** `web`.
  + https://aws.amazon.com/premiumsupport/knowledge-center/elastic-beanstalk-customized-log-files/
  + Example files logged by default: `/var/log/eb-engine.log, /var/log/eb-hooks.log, /var/log/web.stdout.log, /var/log/nginx/access.log, /var/log/nginx/error.log`
    ```
    [ec2-user@ip-172-31-39-53 ~]$ ls -lha /var/log/*.stdout.log
    -rw------- 1 root root 94 Jan 29 22:46 /var/log/bgapp.stdout.log
    -rw------- 1 root root 92 Jan 29 22:46 /var/log/web.stdout.log
    ```
  + All processes from `Procfile` will have `/var/log/<name>.{stdout,stderr}.log` according to https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-linux-extend.html.
    * Opened a pull-request for the wrong documentation filenames: https://github.com/awsdocs/aws-elastic-beanstalk-developer-guide/pull/133
    * However, ONLY the `web` process will its logs streamed to CW Logs. For the other processes we need to add custom logging.
  + See here how to add more logs by using directly the CloudWatch Logs agent `awslogs`: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/AWSHowTo.cloudwatchlogs.html#AWSHowTo.cloudwatchlogs.streaming.custom
  + Sample config: https://github.com/awsdocs/elastic-beanstalk-samples/blob/master/configuration-files/aws-provided/instance-configuration/logs-streamtocloudwatch-linux.config
    * Agent config: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AgentReference.html
  + Need to setup the right permissions as well for the EC2 instance role (not the service role) if there is more than 1 Procfile process. The first one is setup by Beanstalk's service role, but the others are setup by the instance role:
  ```
  2022-01-25 02:00:28,575 - cwlogs.push.batch - INFO - 6113 - Thread-2 - Creating log group /aws/elasticbeanstalk/Test2-al2-env//var/log/web2.stdout.log.
  2022-01-25 02:00:28,597 - cwlogs.push.batch - WARNING - 6113 - Thread-2 - CreateLogGroup failed with exception An error occurred (AccessDeniedException) when calling the CreateLogGroup operation: User: arn:aws:sts::<redacted>:assumed-role/aws-elasticbeanstalk-ec2-role/i-08e1b0e2204017ae2 is not authorized to perform: logs:CreateLogGroup on resource: arn:aws:logs:eu-west-1:<redacted>:log-group:/aws/elasticbeanstalk/Test2-al2-env//var/log/web2.stdout.log:log-stream: because no identity-based policy allows the logs:CreateLogGroup action
  2022-01-25 02:00:28,598 - cwlogs.push.batch - WARNING - 6113 - Thread-2 - An error occurred (AccessDeniedException) when calling the CreateLogGroup operation: User: arn:aws:sts::<redacted>:assumed-role/aws-elasticbeanstalk-ec2-role/i-08e1b0e2204017ae2 is not authorized to perform: logs:CreateLogGroup on resource: arn:aws:logs:eu-west-1:<redacted>:log-group:/aws/elasticbeanstalk/Test2-al2-env//var/log/web2.stdout.log:log-stream: because no identity-based policy allows the logs:CreateLogGroup action
  ```

- For log rotation: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.logging.html
