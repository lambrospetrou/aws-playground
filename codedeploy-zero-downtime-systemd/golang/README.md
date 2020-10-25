# Socket activation on EC2 with CodeDeploy

## Overview

AWS CodeDeploy allows us to use anything we want during deployment on standard EC2 hosts.

This sample project contains everything you need to deploy a sample application that uses `systemd` socket activation to achieve zero downtime during deployments. The only CodeDeploy specific file is the `appspec.yml` which specifies what files to copy, and what commands to run at each lifecycle step, pretty straightforward.

The key scripts that call `systemd` to start and stop the application properly are `start-application.sh` and `stop-application.sh`, under the `deploy-tools/codedeploy-hooks/` directory.

This sample does not necessarily follow best-practices in terms of deployment assets since each deployment directly overrides previous assets.
In a proper production deployment we should have separate deployment folders to allow robust rollbacks. Also, graceful server shutdown is something left as an exercise to the reader.

**References**
- https://docs.aws.amazon.com/codedeploy/latest/userguide/application-revisions-appspec-file.html
- https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html
- http://0pointer.de/blog/projects/socket-activation.html
- https://vincent.bernat.ch/en/blog/2018-systemd-golang-socket-activation
- https://github.com/coreos/go-systemd/blob/master/examples/activation/httpserver/httpserver.go

## Try it

To prepare the bundle `.zip` file to upload you just need to run `make`. You need Docker and the Rust toolchain installed.

Then follow the instructions at https://docs.aws.amazon.com/codedeploy/latest/userguide/applications-create-in-place.html to create a CodeDeploy application and deploy the created bundle.

## Test

In order to test that there is no downtime I use the [`ab`](https://httpd.apache.org/docs/2.4/programs/ab.html) network tool to send a continuous stream of requests to the EC2 instance. While the requests are flowing, I trigger a deployment and make sure there are not failed requests (it will throw with `connection reset` or `connection refused` if the server goes offline).

```
ab -c 10 -n 10000 http://<public-ip-of-your-instance>/lambros
```
