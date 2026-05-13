# Project 01: Static Site on AWS

## Purpose

Build and deploy a simple static site on AWS so you practice object storage hosting, deployment discipline, public access design, and baseline operational controls.

## Scenario

Assume you need to publish a small documentation, portfolio, or marketing site without running application servers. The workload is simple, but it still needs a safe deployment path, clear ownership, and enough observability to detect accidental changes or delivery problems.

This is a good first AWS project because it introduces the idea that even a "simple website" is still a cloud system with identity, storage, exposure, and operations concerns.

## Architecture

```text
Source content
-> Deployment workflow
-> Amazon S3 static website hosting
-> Optional CDN or custom domain
-> Amazon CloudWatch and access review
```

## What You Will Build

- A static site deployed to an S3 bucket.
- A controlled deployment path for updates.
- Basic logging or monitoring around site delivery and changes.
- An architecture explanation that distinguishes storage, deployment, and public access.

## Why This Architecture Works

Amazon S3 is a good fit because the workload is static content, not dynamic server-side rendering. IAM gives you a clean way to separate deployment access from public delivery. CloudWatch helps turn a public bucket into something that still has operational visibility.

The project is intentionally small, but it teaches an important lesson: simple workloads should still be designed deliberately.

## Services Used

- [Amazon S3](../services/s3.md)
- [AWS Identity and Access Management (IAM)](../services/iam.md)
- [Amazon CloudWatch](../services/cloudwatch.md)

## Skills Practiced

- Object storage hosting
- Least-privilege deployment access
- Basic monitoring and lifecycle control
- Explaining a simple public architecture clearly

## Implementation Steps

1. Create an S3 bucket structure for the site and decide whether the first version will use only S3 website hosting or include a CDN later.
2. Configure the hosting settings, content layout, and any versioning or lifecycle basics you want from the start.
3. Create a narrow deployment identity so publishing content is separate from broader administrative access.
4. Deploy the site content and verify public behavior, broken links, and update flow.
5. Add basic monitoring or logging so content changes, failed deployments, or unusual traffic patterns are easier to investigate.
6. Document the architecture and explain the public access model, deployment path, and main cost drivers.

## Security and Operations Considerations

Allow only the minimum public access required for serving site assets, keep write access restricted to the deployment path, and review bucket policies carefully. If you later add CloudFront, custom domains, or CI/CD, revisit how the trust boundaries change.

Operationally, the important question is not whether the site is complex. It is whether you can explain how content is updated, how bad changes are detected, and how access is controlled.

## Cost Considerations

This project should be low cost, but storage, request volume, data transfer, and optional CDN usage can add up if the site receives meaningful traffic or if content is redeployed frequently.

## How to Extend This Project

- Add a custom domain and CDN.
- Add CI/CD deployment automation.
- Add access logs and versioned content rollback.
- Add a simple staging-to-production publishing workflow.

## Portfolio Value

This project shows you can deploy a basic web asset, control access paths, and explain the difference between storage, delivery, deployment, and monitoring concerns. That is a stronger story than "I uploaded files to a bucket."

## Official References

- [Hosting a static website using Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/WebsiteHosting.html)
- [AWS IAM documentation](https://docs.aws.amazon.com/iam/)
- [Amazon CloudWatch documentation](https://docs.aws.amazon.com/cloudwatch/)
