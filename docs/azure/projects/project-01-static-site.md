# Project 01: Static Site on Azure

## Purpose

Build and deploy a simple static site on Azure so you practice object storage hosting, access control, public delivery decisions, and baseline monitoring.

## Scenario

Assume you need to publish a small documentation, portfolio, or informational site without maintaining an application server. The site is simple, but it still needs a repeatable deployment path, clear Azure access control, and enough monitoring to make operational ownership visible.

This is a useful first Azure project because it introduces the provider through a lightweight workload while still forcing you to think about storage configuration, RBAC, and public exposure.

## Architecture

```text
Source content
-> Deployment workflow
-> Azure Blob Storage static website hosting
-> Optional CDN or custom domain
-> Azure Monitor
```

## What You Will Build

- A static site deployed to Blob Storage.
- A repeatable deployment path for updates.
- Basic monitoring and access review around the delivery path.
- A clear explanation of how public access and deployment access differ.

## Why This Architecture Works

Azure Blob Storage is a good fit because the workload is just static content. Microsoft Entra ID and Azure RBAC let you separate who can deploy the site from who can administer the broader subscription. Azure Monitor gives the project at least a baseline operational story instead of leaving it as an unmanaged public asset.

## Services Used

- [Azure Blob Storage](../services/blob-storage.md)
- [Microsoft Entra ID](../services/entra-id.md)
- [Azure Role-Based Access Control](../services/role-based-access-control.md)
- [Azure Monitor](../services/monitor.md)

## Skills Practiced

- Static site hosting
- Azure access control basics
- Object storage delivery patterns
- Explaining a simple public Azure architecture

## Implementation Steps

1. Create the storage account and enable static website hosting for the site content.
2. Decide how the public content will be organized and whether the first version needs only storage hosting or an edge layer later.
3. Create a narrow deployment path using Azure identity and RBAC so publishing rights stay limited.
4. Deploy the site content and verify public behavior, URL structure, and update flow.
5. Add baseline monitoring and activity visibility so changes and delivery issues are easier to inspect.
6. Document the deployment path, access model, and main cost considerations.

## Security and Operations Considerations

Keep write access narrow, review any public site exposure carefully, and use RBAC so deployment rights are not broader than necessary. If you later add CDN, a custom domain, or pipeline automation, revisit the trust boundaries and logging expectations.

Operationally, the project is valuable because it proves you can reason about simple public delivery in Azure without pretending the workload has no security or ownership model.

## Cost Considerations

The base project should stay low cost, but storage, request volume, data transfer, and optional CDN or custom domain services can increase spend.

## How to Extend This Project

- Add a custom domain and CDN.
- Add deployment automation.
- Add access logs and change rollback steps.
- Add a staging and production publishing process.

## Portfolio Value

This project demonstrates that you can deliver public content on a cloud platform while still thinking about access control, deployment discipline, and operational hygiene.

## Official References

- [Azure Blob Storage documentation](https://learn.microsoft.com/en-us/azure/storage/blobs/)
- [Static website hosting in Azure Storage](https://learn.microsoft.com/en-us/azure/storage/blobs/storage-blob-static-website)
- [Azure Monitor documentation](https://learn.microsoft.com/en-us/azure/azure-monitor/)
