# Amazon API Gateway

## Purpose

Amazon API Gateway provides a managed front door for HTTP, REST, and WebSocket APIs on AWS.

## Definition

Amazon API Gateway is a managed service for publishing, securing, and operating APIs. It sits between clients and backend services, giving teams a place to define routes, apply authorization, shape requests, and expose stable endpoints without building a custom gateway layer from scratch.

That makes it more than just an HTTP entry point. It is part of how API-facing systems express boundaries, security controls, deployment stages, and operational visibility.

In simple terms:

> API Gateway is the managed edge layer that lets AWS backends behave like a structured API instead of a collection of raw service endpoints.

## What Problem It Solves

Without a gateway layer, teams often end up duplicating authentication, throttling, routing, and request-management logic inside application code or infrastructure they do not want to own. API Gateway gives them a managed place to handle those concerns consistently.

## How It Is Commonly Used

API Gateway is commonly used for:

- serverless APIs backed by Lambda,
- managed HTTP front doors in front of AWS services,
- request validation, throttling, and staged rollout behavior,
- authenticated public or partner-facing APIs,
- lightweight backend-for-frontend patterns.

## When to Use It

- Use it to expose serverless APIs backed by Lambda or other AWS services.
- Use it when you need request validation, authorization, throttling, or staged deployment behavior.
- Use it when an application needs a managed API entry point that integrates with AWS identity and monitoring tools.
- Use it when you want a clear boundary between external callers and backend implementation details.

## When Not to Use It

- Do not use it if a simpler direct integration already meets the requirement.
- Do not treat the gateway as a replacement for application-level validation and security checks.
- Do not push so much transformation logic into the gateway that the API becomes hard to reason about.

## Common Mistakes

- Creating overly chatty APIs that drive unnecessary request count and latency.
- Treating the gateway as the only security layer instead of combining it with backend validation and authorization.
- Neglecting API versioning and staged rollout design.
- Publishing public endpoints without a clear authentication and abuse-control model.
- Failing to correlate gateway metrics with downstream Lambda or service behavior.

## Cloud Engineering Considerations

### Identity and Access

Decide whether the API should use IAM, tokens, custom authorizers, or another front-door identity model. The caller identity model should make sense for the users and systems involved, not just for the easiest initial setup.

### Networking

Review private versus public exposure, custom domains, and how downstream services are reached. An API boundary is also a network boundary.

### Security

Validate inputs, enforce authentication, rate-limit high-risk endpoints, and review how headers, payloads, and identity context reach downstream code.

### Observability

Track request volume, status codes, latency, and downstream failure patterns. Gateway metrics alone are not enough if the backend is still opaque.

### Cost

Request volume and data transfer drive cost, so chatty clients, overly fine-grained APIs, and unnecessary retries can become expensive.

## How This Fits Into Cloud Engineering

API Gateway matters because cloud engineering is not only about running backend code. It is also about deciding how external systems enter the platform, how requests are governed, and how the interface stays stable as the implementation evolves.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Serverless API](../patterns/serverless-api.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Amazon API Gateway documentation](https://docs.aws.amazon.com/apigateway/)
- [API Gateway developer guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html)
