# Project 02: Serverless Contact Form on Google Cloud

## Purpose

Build a serverless contact form so you practice managed container hosting, application data handling, secrets, and observability on Google Cloud.

## Scenario

Assume you need a small public form workflow, but you want to avoid operating a traditional application server. The workload still needs request validation, a place to store submissions, runtime identity, and enough monitoring to make failures visible.

This project is a strong Google Cloud exercise because it combines Cloud Run, Secret Manager, service-account design, and application data in one small but realistic system.

## Architecture

```text
User request
-> Cloud Run
-> Firestore or another application data store
-> Secret Manager
-> Cloud Monitoring
```

## What You Will Build

- A user-facing contact form endpoint.
- A backend service that validates and stores submissions.
- A secure secret and configuration path for the runtime.
- Basic monitoring, alerting, and deployment controls.

## Why This Architecture Works

Cloud Run is a good fit because the workload is small but still needs a real HTTP service and controlled runtime identity. Firestore or another simple application store can hold the submissions. Secret Manager and Cloud Monitoring round out the security and operational side of the design.

## Services Used

- [Cloud Run](../services/cloud-run.md)
- [Secret Manager](../services/secret-manager.md)
- [Cloud Monitoring](../services/cloud-monitoring.md)
- [IAM and Service Accounts](../services/iam-and-service-accounts.md)
- Firestore

## Skills Practiced

- Serverless application delivery
- Service account design
- Runtime secret handling
- Basic application data integration
- Explaining an HTTP-to-data request path clearly

## Implementation Steps

1. Define the request contract, validation behavior, and error model for the form submissions.
2. Build and deploy the Cloud Run service that receives and validates the requests.
3. Choose the application data store shape and store submissions in a way you can query or review later.
4. Use a narrow service account and Secret Manager for any sensitive runtime configuration.
5. Add monitoring and useful logs so request failures, latency issues, and dependency problems are visible.
6. Document the runtime identity model, data flow, and cost watchpoints.

## Security and Operations Considerations

Use service accounts and Secret Manager, validate form input carefully, and keep data store permissions scoped to the application runtime. Think about public exposure, spam or abuse patterns, and how administrators would review submissions safely.

Operationally, the project becomes credible when you can explain how the service is deployed, what identity it runs as, and how a failed request would be investigated.

## Cost Considerations

Cloud Run and supporting services should remain low cost at small scale, but request volume, runtime configuration, verbose logging, and data-store growth can increase spend.

## How to Extend This Project

- Add spam controls or moderation.
- Add a notification workflow.
- Add an admin review view for submissions.
- Add authentication for internal-only review functionality.

## Portfolio Value

This project demonstrates that you can connect a public endpoint, a managed runtime, secrets, service accounts, and operational telemetry into one coherent Google Cloud application path.

## Official References

- [Cloud Run documentation](https://cloud.google.com/run/docs)
- [Secret Manager documentation](https://cloud.google.com/secret-manager/docs)
- [Firestore documentation](https://cloud.google.com/firestore/docs)
- [Cloud Monitoring documentation](https://cloud.google.com/monitoring/docs)
