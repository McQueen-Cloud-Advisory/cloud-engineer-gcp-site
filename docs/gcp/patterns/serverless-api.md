# Serverless API on Google Cloud

## Purpose

This pattern explains how to expose an API on Google Cloud using managed runtimes and supporting platform services.

## Pattern Summary

A serverless API pattern on Google Cloud often uses Cloud Run as the main request-handling runtime, with Secret Manager, service accounts, and monitoring services supporting the application path. Cloud Functions can also support event-driven pieces around the API.

This pattern matters because it shows how to build a production-minded API path without operating servers directly.

## When This Pattern Fits

Use this pattern when:

- the workload is HTTP-driven,
- the team wants a managed runtime with container packaging flexibility,
- service-account-based runtime identity is a good fit,
- and the API is better modeled as a stateless service boundary than as a traditional server fleet.

## When Not to Use It

Do not use this pattern when the workload needs deep host control, complex persistent runtime state, or a more specialized runtime topology.

## Common Use Cases

- Contact forms
- Internal APIs
- Lightweight application backends

## Reference Architecture

```text
Client
-> Cloud Run
-> Application logic and secrets
-> Application data store
-> Monitoring
```

## Why This Pattern Works

It works because Google Cloud separates the runtime, secret handling, service-account identity, and monitoring model cleanly. Cloud Run keeps deployment and scaling simple, while Secret Manager and Cloud Monitoring help the service behave like an operable system instead of a container-only demo.

## Provider Services

- Cloud Run
- Secret Manager
- IAM and Service Accounts
- Cloud Monitoring
- Optional Firestore or another application data store

## Design Considerations

### Security

Validate inputs, scope service accounts carefully, and keep secrets out of code and deployment artifacts.

### Reliability

Design for retries, downstream failure handling, and predictable deployment rollouts.

### Observability

Capture request latency, errors, and dependency behavior so issues can be traced end to end.

### Cost

Runtime requests, scale behavior, and telemetry volume are the main cost drivers.

### Deployment

Ship the runtime, configuration, permissions, and monitoring together so environments stay consistent.

## Common Mistakes

- Using a broad default service account for production traffic.
- Treating container packaging as the whole deployment story.
- Ignoring how secrets and configuration are loaded at runtime.
- Measuring only container health instead of request health.
- Forgetting that serverless still needs clear rollout and failure-handling design.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)

## How This Fits Into Cloud Engineering

This pattern matters because it connects runtime identity, HTTP exposure, secrets, data access, and monitoring into one small but real application architecture. That is a core cloud-engineering skill, not just a deployment trick.

## Official References

- [Cloud Run documentation](https://cloud.google.com/run/docs)
- [Secret Manager documentation](https://cloud.google.com/secret-manager/docs)
- [Cloud Monitoring documentation](https://cloud.google.com/monitoring/docs)
