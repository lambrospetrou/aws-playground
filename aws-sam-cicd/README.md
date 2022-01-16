## Dev

Build the artifacts we need, including minifying and bundling the functions.

```
npm run release
```

### Local testing and debugging

We use [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) so make sure it's installed in your system (not available in NPM anymore).

Invoke the function locally:

```
sam local invoke HelloWorldFunction --no-event -t build/aws-iac/sam_template.yaml
```

Start the API locally:

```
sam local start-api -t build/aws-iac/sam_template.yaml
```

### Cypress E2E-Integration testing

1. Start the API: `npm run start:api` (defaults to `http://localhost:3000`)
2. Open the Cypress test GUI: `npm run cy:watch`

Example how to run this on CI with a custom endpoint, e.g. for our dev/prod APIs:
```
npm run cy:run -- --config baseUrl=https://1se9reo7k9.execute-api.eu-west-1.amazonaws.com
```

## References

- [AWS Serverless Application Model Developer Guide](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/)
