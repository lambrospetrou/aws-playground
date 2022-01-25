# Docker Linux Platform on Amazon Linux 2

- [Using the Docker platform](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/docker.html)
- [Docker configuration](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/single-container-docker-configuration.html)
- [Extending Elastic Beanstalk Linux platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-linux-extend.html)

- In AL2 linux platform Docker there is no `proxy` running, hence no `nginx`. This means that one of the services running in our Docker needs to expose port `80` to accept the traffic.

- For logs there is some configuration needed:
  + [Generating logs for enhanced health reporting (Docker Compose)](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/create_deploy_docker.container.console.html#docker-env-cfg.healthd-logging)
  + [Docker container customized logging (Docker Compose)](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/create_deploy_docker.container.console.html#docker-env-cfg.dc-customized-logging)

- For log rotation: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.logging.html
