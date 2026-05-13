# Serverless API on Azure

## Purpose

This pattern explains how to build a managed API path on Azure using a gateway, serverless compute, and supporting platform services.

## Pattern Summary

A serverless API pattern on Azure often uses API Management as the front door and Azure Functions as the compute layer. Data, secrets, and telemetry are handled by supporting services such as Cosmos DB, Key Vault, Application Insights, and Azure Monitor.

This pattern matters because it shows how cloud engineers connect identity, request handling, runtime logic, monitoring, and data storage into a coherent application path without managing servers directly.

## When This Pattern Fits

Use this pattern when:

- the API is lightweight or event-driven,
- the team wants managed runtime operations,
- Azure identity and API governance are important,
- and the system benefits from clear separation between the gateway and runtime.

## When Not to Use It

Do not use this pattern when the workload needs deeper runtime control, more persistent application behavior, or a platform shape that Functions and APIM do not fit well.

## Common Use Cases

- Forms and lightweight application backends
- Internal APIs
- Event-driven service endpoints

## Reference Architecture

```text
Client
-> Azure API Management
-> Azure Functions
-> Data store and secrets
-> Monitoring and telemetry
```

## Why This Pattern Works

It works because Azure separates the API boundary, serverless logic, identity model, and telemetry clearly. API Management governs entry, Functions runs the code, managed identities and Key Vault secure the runtime, and Application Insights plus Monitor help operators see what is happening.

## Provider Services

- Azure API Management
- Azure Functions
- Azure Cosmos DB
- Azure Key Vault
- Managed Identities
- Application Insights

## Design Considerations

### Security

Validate inputs, restrict API exposure, and use managed identities and least-privilege roles for downstream access.

### Reliability

Design for retries, clear failure handling, and predictable dependency behavior so the API remains understandable under load or error conditions.

### Observability

Correlate request telemetry, function execution, and dependency calls so issues can be traced end to end.

### Cost

Requests, function execution, data store settings, and telemetry volume are the main cost drivers.

### Deployment

Deploy the gateway, function app, identities, and configuration together so environments stay aligned.

## Common Mistakes

- Treating the gateway as the only security layer.
- Letting managed identities accumulate broad roles over time.
- Ignoring retry and timeout behavior in the function layer.
- Measuring only function health instead of full request health.
- Choosing serverless by default without checking whether the application shape really fits.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)

## How This Fits Into Cloud Engineering

This pattern is useful because it teaches how Azure identity, API exposure, runtime logic, data access, and telemetry come together. That is much closer to real platform engineering than simply listing the services involved.

## Official References

- [Azure API Management documentation](https://learn.microsoft.com/en-us/azure/api-management/)
- [Azure Functions documentation](https://learn.microsoft.com/en-us/azure/azure-functions/)
- [Azure Cosmos DB documentation](https://learn.microsoft.com/en-us/azure/cosmos-db/)
