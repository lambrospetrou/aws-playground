# Socket activation on Elastic Beanstalk

**This is not really working as needed (see the `Why it's not ideal` section below).**

The systemd socket activation works perfectly, but after doing it I realized that it's hopeless since the front-end Nginx proxy gets stopped and then started again during the deployment, which makes the whole point void.

The whole downtime by default is roughly 1 second, just restarting nginx and the server basically, so works fine for most use-cases.

Anyway, good to keep for reference.

## Overview

Elastic Beastalk starts the service named [`web` from the `Procfile`](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/go-procfile.html) and gives it the port `5000`. 
This is the port that the nginx proxy forwards the web requests.

In our case, we want to handle port `5000` from the server started through systemd, and not `Procfile`, since systemd has to start the service itself.

The approach is quite hacky. We start a dummy service using the `web` name in Profile which does not listen to port `5000` as expected.
We instead just listen on a different port, `5001`, but we could just as well not do anything other than keep running (i.e. sleep).

Then, [using platform hooks](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-linux-extend.html) we manually call systemd to enable and start the actual service and socket listening at port `5000`.

**References**
- https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-linux-extend.html
- http://0pointer.de/blog/projects/socket-activation.html
- https://vincent.bernat.ch/en/blog/2018-systemd-golang-socket-activation
- https://github.com/coreos/go-systemd/blob/master/examples/activation/httpserver/httpserver.go

## Why it's not ideal

After inspecting the `eb-engine.log` it's clear that `nginx` is stopped during the **deploy/flip app and proxy** stage.

This means that even though our socket activation with systemd works, it has no effect since the proxy in front of it is down.

The good thing (which I did not expect) is that the flip actually takes **roughly 1 second** , see the log starting at `18:34:05.719900`, and finishing the health check for `web.service` at `18:34:06.495089`. So, maybe the socket activation is not needed after all.

```
2020/10/24 18:34:05.563826 [INFO] Executing instruction: RunAppDeployPreDeployHooks
2020/10/24 18:34:05.563848 [INFO] Executing platform hooks in .platform/hooks/predeploy/
2020/10/24 18:34:05.563888 [INFO] Following platform hooks will be executed in order: [01-echo-world.sh]
2020/10/24 18:34:05.563892 [INFO] Running platform hook: .platform/hooks/predeploy/01-echo-world.sh
2020/10/24 18:34:05.708848 [INFO] Finished running the platform hooks in .platform/hooks/predeploy/
2020/10/24 18:34:05.708857 [INFO] Executing instruction: stop X-Ray
2020/10/24 18:34:05.708860 [INFO] stop X-Ray ...
2020/10/24 18:34:05.708876 [INFO] Running command /bin/sh -c systemctl show -p PartOf xray.service
2020/10/24 18:34:05.713886 [WARN] stopProcess Warning: process xray is not registered 
2020/10/24 18:34:05.713912 [INFO] Running command /bin/sh -c systemctl stop xray.service
2020/10/24 18:34:05.719900 [INFO] Executing instruction: stop proxy
2020/10/24 18:34:05.719923 [INFO] Running command /bin/sh -c systemctl show -p PartOf httpd.service
2020/10/24 18:34:05.723752 [WARN] deregisterProcess Warning: process httpd is not registered, skipping...

2020/10/24 18:34:05.723779 [INFO] Running command /bin/sh -c systemctl show -p PartOf nginx.service
2020/10/24 18:34:05.727800 [INFO] Running command /bin/sh -c systemctl is-active nginx.service
2020/10/24 18:34:05.731122 [INFO] Running command /bin/sh -c systemctl show -p PartOf nginx.service
2020/10/24 18:34:05.734879 [INFO] Running command /bin/sh -c systemctl stop nginx.service
2020/10/24 18:34:05.819856 [INFO] Running command /bin/sh -c systemctl disable nginx.service
2020/10/24 18:34:05.879137 [INFO] Running command /bin/sh -c systemctl daemon-reload
2020/10/24 18:34:05.937741 [INFO] Running command /bin/sh -c systemctl reset-failed
2020/10/24 18:34:05.941129 [INFO] Executing instruction: FlipApplication
2020/10/24 18:34:05.941145 [INFO] fetching environment variables...
2020/10/24 18:34:05.941158 [INFO] no go.mod file, setting GOPATH=/var/app/current
2020/10/24 18:34:05.941279 [INFO] Purge old process...
2020/10/24 18:34:05.941296 [INFO] Running command /bin/sh -c systemctl stop eb-app.target
2020/10/24 18:34:05.947481 [INFO] Running command /bin/sh -c systemctl show -p ConsistsOf eb-app.target | cut -d= -f2
2020/10/24 18:34:05.954185 [INFO] web.service

2020/10/24 18:34:05.954213 [INFO] Running command /bin/sh -c systemctl show -p PartOf web.service
2020/10/24 18:34:05.957839 [INFO] Running command /bin/sh -c systemctl is-active web.service
2020/10/24 18:34:05.960891 [INFO] Running command /bin/sh -c systemctl disable web.service
2020/10/24 18:34:06.020661 [INFO] Running command /bin/sh -c systemctl daemon-reload
2020/10/24 18:34:06.085450 [INFO] Running command /bin/sh -c systemctl reset-failed
2020/10/24 18:34:06.089300 [INFO] Register application processes...
2020/10/24 18:34:06.089314 [INFO] Registering the proc: web

2020/10/24 18:34:06.089328 [INFO] Running command /bin/sh -c systemctl show -p PartOf web.service
2020/10/24 18:34:06.094785 [INFO] Running command /bin/sh -c systemctl daemon-reload
2020/10/24 18:34:06.158914 [INFO] Running command /bin/sh -c systemctl reset-failed
2020/10/24 18:34:06.162411 [INFO] Running command /bin/sh -c systemctl is-enabled eb-app.target
2020/10/24 18:34:06.165660 [INFO] Running command /bin/sh -c systemctl enable eb-app.target
2020/10/24 18:34:06.232812 [INFO] Running command /bin/sh -c systemctl start eb-app.target
2020/10/24 18:34:06.236977 [INFO] Running command /bin/sh -c systemctl enable web.service
2020/10/24 18:34:06.300665 [INFO] Running command /bin/sh -c systemctl show -p PartOf web.service
2020/10/24 18:34:06.304998 [INFO] Running command /bin/sh -c systemctl is-active web.service
2020/10/24 18:34:06.308097 [INFO] Running command /bin/sh -c systemctl start web.service
2020/10/24 18:34:06.320351 [INFO] Executing instruction: start X-Ray
2020/10/24 18:34:06.320370 [INFO] X-Ray is not enabled.
2020/10/24 18:34:06.320374 [INFO] Executing instruction: start proxy with new configuration
2020/10/24 18:34:06.320400 [INFO] Running command /bin/sh -c /usr/sbin/nginx -t -c /var/proxy/staging/nginx/nginx.conf
2020/10/24 18:34:06.328124 [INFO] Running command /bin/sh -c cp -rp /var/proxy/staging/nginx/* /etc/nginx
2020/10/24 18:34:06.330223 [INFO] Running command /bin/sh -c systemctl show -p PartOf nginx.service
2020/10/24 18:34:06.336525 [INFO] Running command /bin/sh -c systemctl daemon-reload
2020/10/24 18:34:06.393204 [INFO] Running command /bin/sh -c systemctl reset-failed
2020/10/24 18:34:06.396605 [INFO] Running command /bin/sh -c systemctl show -p PartOf nginx.service
2020/10/24 18:34:06.400971 [INFO] Running command /bin/sh -c systemctl is-active nginx.service
2020/10/24 18:34:06.403859 [INFO] Running command /bin/sh -c systemctl start nginx.service
2020/10/24 18:34:06.486137 [INFO] Executing instruction: configureSqsd
2020/10/24 18:34:06.486151 [INFO] This is a web server environment instance, skip configure sqsd daemon ...
2020/10/24 18:34:06.486155 [INFO] Executing instruction: startSqsd
2020/10/24 18:34:06.486160 [INFO] This is a web server environment instance, skip start sqsd daemon ...
2020/10/24 18:34:06.486163 [INFO] Executing instruction: Track pids in healthd
2020/10/24 18:34:06.486167 [INFO] This is an enhanced health env...
2020/10/24 18:34:06.486184 [INFO] Running command /bin/sh -c systemctl show -p ConsistsOf aws-eb.target | cut -d= -f2
2020/10/24 18:34:06.490957 [INFO] nginx.service cfn-hup.service healthd.service

2020/10/24 18:34:06.490990 [INFO] Running command /bin/sh -c systemctl show -p ConsistsOf eb-app.target | cut -d= -f2
2020/10/24 18:34:06.495089 [INFO] web.service

```
