# Project 02: Serverless Contact Form on Azure

## Purpose

Build a serverless contact form so you practice API exposure, function integration, application data storage, secrets handling, and telemetry.

## Scenario

Assume you need a simple public form workflow for a website or team site, but you want to keep the backend lightweight and managed. The workload still needs validation, secure access to configuration, durable storage, and application-level telemetry.

This project is a good Azure application exercise because it connects API Management, Functions, Cosmos DB, managed identities, and Application Insights into one request path.

## Architecture

```text
User request
-> Azure API Management
-> Azure Functions
-> Azure Cosmos DB
-> Azure Key Vault and Application Insights
```

## What You Will Build

- An API endpoint for form submission.
- A function app that validates and stores submitted data.
- A secure configuration path for secrets or connection data.
- Operational telemetry for requests, failures, and dependencies.

## Why This Architecture Works

Azure API Management provides a governed API front door. Azure Functions keeps the compute layer serverless and event-driven. Cosmos DB gives the application a managed data store, while Key Vault, managed identities, and Application Insights complete the security and observability model.

## Services Used

- [Azure API Management](../services/api-management.md)
- [Azure Functions](../services/functions.md)
- [Azure Cosmos DB](../services/cosmos-db.md)
- [Azure Key Vault](../services/key-vault.md)
- [Application Insights](../services/application-insights.md)
- [Managed Identities](../services/managed-identities.md)

## Skills Practiced

- Serverless API delivery
- Managed identity usage
- Application data modeling
- Operational telemetry
- Explaining request flow and backend dependencies

## Implementation Steps

1. Define the API contract, validation rules, and error behavior for the form submissions.
2. Create the API Management surface and route requests to Azure Functions.
3. Design the Cosmos DB data model for how submissions will be written and reviewed.
4. Use managed identities and Key Vault so the runtime can access secrets or configuration without embedded credentials.
5. Add Application Insights telemetry and useful alerts for failures, latency, and dependency issues.
6. Document how identity, storage, and monitoring work together across the request path.

## Security and Operations Considerations

Use managed identities where possible, keep secrets in Key Vault, and make sure backend access is scoped to the minimum required resources. Think about request abuse, validation errors, and how admins would safely review stored submissions later.

Operationally, the project becomes stronger when you can show how failures appear in telemetry and how permissions are separated between deployment and runtime.

## Cost Considerations

This project is typically low cost at small scale, but request volume, Cosmos DB throughput choices, API Management tier, and logging volume can increase costs if left unchecked.

## How to Extend This Project

- Add authentication for an admin view.
- Add notifications or queue-based processing.
- Add spam controls or abuse monitoring.
- Add approval or workflow routing for submissions.

## Portfolio Value

This project shows that you can connect an API front door, serverless compute, storage, identity, and observability into a coherent Azure application path.

## Official References

- [Azure API Management documentation](https://learn.microsoft.com/en-us/azure/api-management/)
- [Azure Functions documentation](https://learn.microsoft.com/en-us/azure/azure-functions/)
- [Azure Cosmos DB documentation](https://learn.microsoft.com/en-us/azure/cosmos-db/)
- [Azure Key Vault documentation](https://learn.microsoft.com/en-us/azure/key-vault/)
