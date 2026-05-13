# Project 02: Serverless Contact Form on AWS

## Purpose

Build a serverless API-backed contact form so you practice API design, event-driven compute, data persistence, secrets handling, and monitoring together.

## Scenario

Assume you need a simple public contact workflow for a website or internal team, but you do not want to run a traditional web server just to accept and store submissions. The workload needs validation, storage, secure configuration, and enough telemetry to troubleshoot failures.

This project is useful because it turns several AWS fundamentals into one coherent application path.

## Architecture

```text
User request
-> Amazon API Gateway
-> AWS Lambda
-> Amazon DynamoDB
-> AWS Secrets Manager and Amazon CloudWatch
```

## What You Will Build

- An API endpoint for form submission.
- A Lambda function that validates and stores submissions.
- A durable data store for the submitted records.
- Logging and monitoring for the full request path.

## Why This Architecture Works

API Gateway gives the workload a clean public entry point. Lambda keeps the compute layer lightweight and event-driven. DynamoDB fits well for request-driven application data, and Secrets Manager plus CloudWatch complete the security and operations story.

The project is small, but it models the structure of many real serverless APIs.

## Services Used

- [Amazon API Gateway](../services/api-gateway.md)
- [AWS Lambda](../services/lambda.md)
- [Amazon DynamoDB](../services/dynamodb.md)
- [AWS Secrets Manager](../services/secrets-manager.md)
- [Amazon CloudWatch](../services/cloudwatch.md)

## Skills Practiced

- Serverless API design
- Function-based integration
- Application data modeling
- Secret handling and monitoring
- Explaining end-to-end request flow

## Implementation Steps

1. Define the API contract, including request validation rules and the basic error model.
2. Create the API Gateway endpoint and connect it to a Lambda handler.
3. Design the DynamoDB table around the access pattern you need for stored submissions.
4. Store any sensitive runtime configuration in Secrets Manager and keep the Lambda execution role narrow.
5. Add logging, metrics, and alarms so request failures, validation issues, and storage problems are visible.
6. Document how the request enters the system, how data is stored, and how you would operate or extend the workload.

## Security and Operations Considerations

Validate request input carefully, restrict API access as needed, keep any downstream credentials in Secrets Manager, and ensure each service role only has the permissions it needs. Think about abuse handling as well: rate limits, spam protection, and clear error behavior matter even for a small form API.

Operationally, the important question is whether you can tell when the API is failing, storing bad data, or costing more than expected.

## Cost Considerations

This project is usually low cost at small traffic levels, but repeated testing, high request volume, DynamoDB access patterns, and verbose logging can increase spend.

## How to Extend This Project

- Add spam protection or approval workflow.
- Add notification delivery for new submissions.
- Add a small admin view for reviewing stored submissions.
- Add authentication and an internal review workflow.

## Portfolio Value

This project demonstrates that you can connect a public API, serverless compute, data storage, and operational controls into one coherent AWS system instead of presenting isolated service knowledge.

## Official References

- [Amazon API Gateway documentation](https://docs.aws.amazon.com/apigateway/)
- [AWS Lambda documentation](https://docs.aws.amazon.com/lambda/)
- [Amazon DynamoDB documentation](https://docs.aws.amazon.com/dynamodb/)
- [AWS Secrets Manager documentation](https://docs.aws.amazon.com/secretsmanager/)
