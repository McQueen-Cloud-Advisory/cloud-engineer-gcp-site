# Static Site on Google Cloud

## Purpose

This pattern explains how to host and deliver static content on Google Cloud with a simple, low-operations architecture.

## Pattern Summary

A static site pattern on Google Cloud commonly starts with Cloud Storage website hosting and may later add a load balancer, CDN, or custom domain. The design separates content delivery from server-side application logic and is a useful first production pattern.

This pattern matters because it introduces public content delivery, deployment discipline, and access control without requiring an application server.

## When This Pattern Fits

Use this pattern when:

- the workload is mostly static files,
- the team wants low runtime overhead,
- the publishing path can be kept simple,
- and the main architectural concern is public delivery rather than dynamic logic.

## When Not to Use It

Do not use this pattern when the workload needs complex server-side logic, deep per-user state, or backend processing on every request.

## Common Use Cases

- Documentation sites
- Portfolio sites
- Simple public pages

## Reference Architecture

```text
Client
-> Optional load balancer or CDN
-> Cloud Storage static website hosting
-> Monitoring
```

## Why This Pattern Works

It works because static content is fundamentally a storage and delivery problem. Google Cloud lets the team keep storage, deployment identity, and monitoring separate enough to operate the site cleanly without introducing a full application runtime.

## Provider Services

- Cloud Storage
- IAM and Service Accounts
- Cloud Monitoring

## Design Considerations

### Security

Keep write access narrow, review public content exposure carefully, and treat bucket permissions as part of the architecture.

### Reliability

Static site delivery is reliable when deployment is simple and rollback is easy.

### Observability

Use monitoring and access insights so delivery problems and unusual traffic are visible.

### Cost

This pattern is usually inexpensive, but data transfer and optional edge services can increase cost.

### Deployment

Deploy content through a repeatable workflow rather than ad hoc manual uploads wherever possible.

## Common Mistakes

- Treating a public bucket as if it needs no deployment or access story.
- Leaving publishing access too broad.
- Skipping rollback planning.
- Adding edge services without revisiting logging and exposure controls.

## Related Projects

- [Project 01: Static Site](../projects/project-01-static-site.md)

## How This Fits Into Cloud Engineering

This pattern is valuable because it shows that a simple public site still needs identity boundaries, deployment discipline, and basic observability. That is a small but real cloud-engineering workload.

## Official References

- [Static website hosting on Cloud Storage](https://cloud.google.com/storage/docs/hosting-static-website)
- [Cloud Storage documentation](https://cloud.google.com/storage/docs)
