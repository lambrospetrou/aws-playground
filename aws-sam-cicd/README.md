## Dev

Build the artifacts we need, including minifying and bundling the functions.

```
npm run release
```

Invoke the function locally:

```
sam local invoke HelloWorldFunction --no-event -t build/aws-iac/sam_template.yaml
```

Start the API locally:

```
sam local start-api -t build/aws-iac/sam_template.yaml
```

### E2E Tests

**Locally**

1. Start the API: `npm run start:api` (defaults to <http://localhost:3000>)
2. Run the tests: `npm run test:e2e`

**CI**

1. The API is deployed by the CodePipeline/CodeBuild. So get the endpoint from there.
2. Run the tests with custom URL: `npm run test:e2e -- --baseUrl=https://1se9reo7k9.execute-api.eu-west-1.amazonaws.com`

## References

- [AWS Serverless Application Model Developer Guide](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/)
