# Firebase Hosting

## Purpose

Firebase Hosting is a managed hosting service for serving static websites, single-page applications, and web app front ends through Firebase's hosting platform.

It is useful when a team needs a simple, managed way to publish web content without operating web servers, configuring infrastructure from scratch, or managing TLS certificates manually.

## Definition

Firebase Hosting is a production-oriented web hosting service from Firebase that serves static and dynamic web content over Google's hosting infrastructure.

At its simplest, Firebase Hosting takes a local build output directory, uploads those files to Firebase, and serves them from a Firebase-generated domain or a custom domain. For static documentation sites, landing pages, portfolios, and frontend applications, this means the deployment unit is usually the generated static site output rather than a running server.

In simple terms:

> Firebase Hosting is a managed way to deploy and serve web assets without managing the web server layer yourself.

Firebase Hosting can serve static files directly, support custom domains, provision SSL certificates, and integrate with Firebase CLI-based deployment workflows. It can also be extended with rewrites to backend services, but the core hosting model is especially strong for static sites and frontend applications.

## What Problem It Solves

Hosting a web front end sounds simple until the operational details appear.

A team still has to decide:

- where the files are hosted,
- how deployments are performed,
- how HTTPS is handled,
- how custom domains are connected,
- how old versions are managed,
- how preview environments are created,
- how redirects and 404 behavior work,
- and how public content delivery is separated from write/deployment permissions.

Firebase Hosting packages many of those concerns into a managed service. It gives teams a straightforward hosting target while still supporting production-oriented features such as custom domains, SSL, preview channels, versioned releases, redirects, rewrites, and CLI-based deployment.

## How It Is Commonly Used

Firebase Hosting is commonly used for:

- static documentation sites,
- marketing or portfolio sites,
- single-page applications,
- frontend assets for Firebase-backed applications,
- web applications that need simple hosting and managed HTTPS,
- preview deployments for review before release,
- lightweight public sites that do not need a dedicated web server.

It is often used with static site generators, frontend frameworks, or build tools that output HTML, CSS, JavaScript, images, and other static assets into a build directory.

## When to Use It

Use Firebase Hosting when:

- the site can be served primarily as static files,
- the team wants managed HTTPS and custom domain support,
- the deployment process should be simple and repeatable,
- the workload does not justify maintaining a web server or container runtime,
- preview channels would help review changes before production,
- the front end may later integrate with Firebase or Google Cloud services.

Firebase Hosting is a strong first choice for public documentation, lightweight web applications, and frontend-heavy projects where operational simplicity matters.

## When Not to Use It

Do not use Firebase Hosting as the default answer for every web workload.

It may not be the best fit when:

- the application requires long-running server-side processes,
- the workload needs detailed network placement inside a private VPC,
- the team needs full control over load balancer architecture,
- the site must run custom server software directly,
- the backend and frontend are tightly coupled into a single server runtime,
- compliance requirements demand a hosting architecture that Firebase Hosting cannot satisfy.

In those cases, services such as Cloud Run, App Engine, Google Kubernetes Engine, or a custom load-balancer-backed architecture may be more appropriate.

## Common Mistakes

Common Firebase Hosting mistakes include:

- confusing the generated build output with the source files,
- committing generated build folders to Git,
- configuring a static documentation site as a single-page application when it is not one,
- pointing DNS records at the wrong Firebase target,
- forgetting that custom domain and SSL provisioning can take time,
- deploying manually without validating the build first,
- granting broad project access when only deployment access is needed,
- using Firebase Hosting as a workaround for unclear architecture decisions.

A simple hosting service still needs disciplined source control, build validation, IAM review, and deployment habits.

## Cloud Engineering Considerations

### Identity and Access

Firebase Hosting deployment requires permissions to manage Firebase Hosting resources and deploy content to the target Firebase project.

For individual learning projects, a human administrator may deploy manually at first. For production or portfolio-grade workflows, deployment should eventually move toward a narrower identity model, such as a CI/CD workflow using an appropriate service account or federated identity.

A cloud engineer should separate:

- who can edit source content,
- who can change hosting configuration,
- who can deploy to the live site,
- and who can administer the Firebase or Google Cloud project.

Those are related permissions, but they are not the same responsibility.

### Project and Organization Policy

Firebase Hosting depends on Firebase being added to a Google Cloud project. In managed Google Cloud organizations, this can be affected by organization policies, API enablement, service usage permissions, and Firebase Terms of Service acceptance.

Important setup blockers can include:

- missing Firebase or Service Usage permissions,
- disabled Firebase-related APIs,
- organization policies that restrict which principals can be added to IAM policies,
- stale CLI authentication after permissions or terms have changed,
- or Firebase not yet being attached to the Google Cloud project.

This is a useful reminder that managed services still depend on the surrounding governance model. A service can be simple to use and still be blocked by identity, policy, or project setup.

### Deployment

Firebase Hosting deploys files from a configured public directory. Static site generators usually produce this directory during a build step.

A typical deployment flow looks like this:

```text
Source files
-> Static site generator or frontend build
-> Generated output directory
-> Firebase Hosting deploy
-> Firebase domain or custom domain
```

For documentation sites, the source files should usually remain in version control while generated build output should be ignored and regenerated as needed.

### Custom Domains and DNS

Firebase Hosting can serve content from Firebase-generated domains or from a custom domain.

For a custom domain, the usual process is:

1. Add the custom domain in Firebase Hosting.
2. Add the DNS records that Firebase provides through the domain's DNS provider.
3. Wait for DNS propagation and certificate provisioning.
4. Verify that the custom domain resolves to the hosted site.

For subdomains, Firebase often uses a CNAME record. The DNS provider may ask for only the subdomain label rather than the full domain name. For example, a record for `docs.example.com` may use `docs` as the host/name value, depending on the DNS provider's interface.

### Security

Firebase Hosting serves public web content by default. The security model should focus on controlling who can deploy, who can alter configuration, and what content becomes public.

A cloud engineer should review:

- project-level IAM,
- deployment permissions,
- custom domain ownership,
- DNS access,
- whether generated output contains sensitive files,
- and whether redirects or rewrites expose unintended routes.

Static hosting does not eliminate security concerns. It changes where those concerns appear.

### Observability

Firebase Hosting provides release and hosting management features, and Google Cloud/Firebase project activity can help trace deployment activity. For simple static sites, the main operational concerns are usually deployment status, domain status, certificate status, and whether the expected content is being served.

For more advanced sites, observability may include:

- release history,
- deploy logs,
- custom monitoring,
- web analytics,
- uptime checks,
- and alerting for domain or availability issues.

### Cost

Firebase Hosting has a free tier and usage-based limits, but cost should still be reviewed. Static sites are usually inexpensive, but bandwidth, storage, and traffic patterns can matter as usage grows.

Cost risk is generally lower than a compute-based hosting model, but it is not zero. A public site should still have basic cost awareness and billing visibility.

## Firebase Hosting vs. Cloud Storage Static Hosting

Firebase Hosting and Cloud Storage can both serve static web content, but they are not the same operational experience.

Firebase Hosting is usually easier when the team wants:

- a simple deploy command,
- managed custom domain flow,
- managed SSL certificate provisioning,
- preview channels,
- versioned hosting releases,
- and a web-app-oriented hosting workflow.

Cloud Storage static website hosting is useful when the team wants to work directly with object storage, bucket policies, and lower-level hosting architecture. It may also be part of a more explicit design involving Cloud Load Balancing and Cloud CDN.

For a first static web project, Firebase Hosting is often the simpler managed path. For a deeper infrastructure project, Cloud Storage plus a load balancer can expose more architecture details.

## How This Fits Into Cloud Engineering

Firebase Hosting is a good example of cloud engineering tradeoff thinking.

The service hides much of the web-serving infrastructure, but the engineer is still responsible for the surrounding decisions:

- source versus build output,
- deployment process,
- IAM boundaries,
- DNS records,
- custom domain ownership,
- SSL status,
- release discipline,
- and fallback or rollback strategy.

That makes Firebase Hosting useful for learning because the service is simple enough to deploy quickly, but still real enough to expose practical cloud engineering issues.

## Related Projects

- [Project 01: Static Site on Google Cloud](../projects/project-01-static-site.md)

## Related Patterns

- [Static Site](../patterns/static-site.md)

## Official References

- [Firebase Hosting documentation](https://firebase.google.com/docs/hosting)
- [Get started with Firebase Hosting](https://firebase.google.com/docs/hosting/quickstart)
- [Configure Hosting behavior](https://firebase.google.com/docs/hosting/full-config)
- [Connect a custom domain](https://firebase.google.com/docs/hosting/custom-domain)
- [Test locally, preview changes, and deploy](https://firebase.google.com/docs/hosting/test-preview-deploy)
- [Manage live and preview channels, releases, and versions](https://firebase.google.com/docs/hosting/manage-hosting-resources)
