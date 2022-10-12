# Upstash Redis + Go + AWS Lambda = 🚀❤️

This code runs both locally, and in AWS Lambda.

## Requirements

- You need an [Upstash Redis account](https://upstash.com/redis), and the URL which would look something like this `rediss://:********@eu1-actual-dragon-11111.upstash.io:11111`. You need to set the environment variable `UPSTASH_REDIS_URL` to this value.

## Run the server locally

- Build the server:

```sh
make build-server
```

- Run the server:

```sh
./build/server
```

## Run on AWS Lambda

- Setup the [SAM cli](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) and make sure your user/role has the [right permissions](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-permissions.html).

- Build the handler bundle for AWS Lambda and deploy:

```sh
make sam-deploy
```

You can now go to your AWS console, and find the URL of the newly created function which will be in the outputs of the CloudFormation stack `UpstashRedisGoArticleStackDemo`.

You should also set the `UPSTASH_REDIS_URL` environment variable on your Lambda as well, otherwise it will just crash every time.

### Local test

You can also test it locally by updating the `sample-event.json` with the right path/cookie/query parameters/etc, and then running:

```sh
make sam-invoke
```
