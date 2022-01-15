## Dev

Build the artifacts we need, including minifying and bundling the functions.

```
npm run release
```

Invoke the function locally:

```
sam local invoke HelloWorldFunction --no-event -t aws-iac/sam_template.yaml
```

Start the API locally:

```
sam local start-api -t aws-iac/sam_template.yaml
```

## References

- [AWS Serverless Application Model Developer Guide](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/)
