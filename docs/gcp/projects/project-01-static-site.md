# Project 01: Static Site on Google Cloud

## Purpose

Build and deploy a simple static site on Google Cloud so you practice object storage hosting, deployment discipline, public delivery design, and baseline monitoring.

## Scenario

Assume you need to publish a lightweight documentation, portfolio, or informational site without maintaining a server fleet. The workload is straightforward, but it still needs a safe deployment path, controlled permissions, and enough monitoring to prove it is being operated deliberately.

This is a good first Google Cloud project because it introduces Cloud Storage, service-account thinking, and basic operational discipline without a large runtime footprint.

## Architecture

```text
Source content
-> Deployment workflow
-> Cloud Storage static website hosting
-> Optional load balancer or CDN
-> Cloud Monitoring
```

## What You Will Build

- A static site deployed to a Cloud Storage bucket.
- A controlled deployment path for updates.
- Baseline monitoring and access review around the delivery path.
- A clear explanation of how public content delivery differs from deployment access.

## Why This Architecture Works

Cloud Storage is a strong fit because the workload is static content rather than application compute. IAM and service accounts let you define a narrow publishing path instead of relying on broad human permissions. Cloud Monitoring helps turn a simple public site into a system that still has visible operational behavior.

## Services Used

- [Cloud Storage](../services/cloud-storage.md)
- [IAM and Service Accounts](../services/iam-and-service-accounts.md)
- [Cloud Monitoring](../services/cloud-monitoring.md)

## Skills Practiced

- Static site delivery
- Google Cloud access control basics
- Object storage hosting
- Explaining a lightweight public architecture well

## Implementation Steps

1. Create the bucket and configure static website hosting behavior for the site.
2. Decide how the content will be published and whether the first version needs only bucket hosting or an edge layer later.
3. Configure narrow deployment access using IAM and service accounts where appropriate.
4. Publish the site and verify public behavior, URL structure, and update flow.
5. Add baseline monitoring or activity visibility so changes and delivery issues are easier to inspect.
6. Document the architecture, access model, and main cost drivers.

## Security and Operations Considerations

Keep write access narrow, review public content exposure carefully, and use service accounts deliberately for deployment automation. If you later add edge delivery, custom domains, or CI/CD, revisit how the identity and logging model should change.

Operationally, the main value of the project is proving that even a simple object-storage site still has a deployment boundary, an access model, and an operations story.

## Cost Considerations

This project should be low cost, but storage, egress, request volume, and optional edge services can increase spend as traffic grows.

## How to Extend This Project

- Add a custom domain and CDN.
- Add CI/CD deployment automation.
- Add access logs and rollback procedures.
- Add separate staging and production publishing paths.

## Portfolio Value

This project shows that you can deliver a simple web asset on Google Cloud while still thinking about access control, deployment, and operations.

## Official References

- [Cloud Storage documentation](https://cloud.google.com/storage/docs)
- [Static website hosting on Cloud Storage](https://cloud.google.com/storage/docs/hosting-static-website)
- [Cloud Monitoring documentation](https://cloud.google.com/monitoring/docs)
