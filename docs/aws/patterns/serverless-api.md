# Serverless API on AWS

## Purpose

This pattern explains how to expose an API on AWS using managed services for the front door, compute, data, and operations layers.

## Pattern Summary

A serverless API pattern typically uses Amazon API Gateway as the entry point and AWS Lambda for request handling. Application data often sits in DynamoDB or another managed store, while secrets and telemetry are handled by supporting platform services.

This pattern matters because it shows how to connect public request handling, business logic, access control, and monitoring without running web servers.

## When This Pattern Fits

Use this pattern when:

- the API is request-driven and event-friendly,
- the team wants low host-management overhead,
- the workload does not require complex long-running sessions,
- and independent deployment of API components is useful.

## When Not to Use It

Do not use this pattern when the workload needs deep runtime control, long-lived stateful connections, or a heavier application platform by default.

## Common Use Cases

- Contact forms
- Internal tools
- Simple application backends

## Reference Architecture

```text
Client
-> Amazon API Gateway
-> AWS Lambda
-> Application data store
-> Logging and monitoring
```

## Why This Pattern Works

It works because the gateway, runtime, data, and operations concerns are separated cleanly. API Gateway handles entry, Lambda handles logic, managed storage holds state, and CloudWatch plus supporting services make the system observable.

## Provider Services

- Amazon API Gateway
- AWS Lambda
- Amazon DynamoDB
- AWS Secrets Manager
- Amazon CloudWatch

## Design Considerations

### Security

Validate inputs, control who can call the API, and keep downstream permissions scoped tightly to the function runtime.

### Reliability

Design for retries, idempotency, and downstream failure handling so the API behaves predictably under load or transient errors.

### Observability

Capture API status codes, function errors, latency, and data store behavior together.

### Cost

Request volume, function duration, and logging volume are the main cost drivers.

### Deployment

Deploy the gateway, function, permissions, and configuration together so environments stay consistent.

## Common Mistakes

- Letting the API surface grow without clear boundaries.
- Giving many functions the same broad execution role.
- Ignoring idempotency or retry behavior.
- Treating logs as the whole monitoring strategy.
- Choosing serverless for every API shape without checking fit.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)

## How This Fits Into Cloud Engineering

This pattern matters because it teaches how to connect public access, runtime identity, secret handling, storage, and observability into one managed application design. That is a core cloud-engineering skill.

## Official References

- [Amazon API Gateway documentation](https://docs.aws.amazon.com/apigateway/)
- [AWS Lambda documentation](https://docs.aws.amazon.com/lambda/)
- [Amazon DynamoDB documentation](https://docs.aws.amazon.com/dynamodb/)
