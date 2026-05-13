# Project 01: Static Site on Google Cloud

## Purpose

Build and deploy the **Cloud Engineer** documentation site as a static website hosted on Google Cloud using **Firebase Hosting**.

This project turns a Markdown-based documentation site into a real cloud-hosted portfolio asset. The goal is not just to publish pages. The goal is to practice the work behind cloud delivery: local builds, hosting configuration, custom DNS, organization-policy troubleshooting, deployment discipline, and fallback planning.

## Scenario

Assume you need to publish a lightweight documentation, portfolio, or informational site without maintaining servers. The site is built from Markdown using MkDocs Material, deployed as static files, and served through a managed Google Cloud hosting service.

The final public site is intended to be available at:

```text
https://cloudengineer.mcqueencloud.com/
```

A default Firebase Hosting URL remains available as a fallback.

This is a good first Google Cloud project because it introduces cloud hosting, DNS, IAM, organization policy, deployment artifacts, and public delivery without requiring application servers, containers, or Kubernetes.

## Final Architecture

```text
Markdown source files
-> MkDocs Material build
-> generated site/ directory
-> Firebase Hosting
-> Squarespace DNS CNAME
-> cloudengineer.mcqueencloud.com
```

## Services and Tools Used

- [Firebase Hosting](../services/firebase-hosting.md)
- [IAM and Service Accounts](../services/iam-and-service-accounts.md)
- [Cloud Monitoring](../services/cloud-monitoring.md)
- MkDocs Material
- GitHub
- Squarespace DNS

If `firebase-hosting.md` does not exist yet in the Google Cloud services section, add it later. Firebase Hosting is the actual Google Cloud-backed hosting service used in this implementation.

## What Was Built

- A MkDocs Material documentation site.
- A separate GCP-hosted site repository.
- A Firebase project connected to a Google Cloud project.
- Firebase Hosting configured to serve the generated `site/` directory.
- A custom subdomain: `cloudengineer.mcqueencloud.com`.
- A Squarespace DNS CNAME record pointing the subdomain to Firebase Hosting.
- A GitHub Pages fallback strategy for future cost or availability concerns.
- A cleaner `.gitignore` to avoid committing generated build artifacts and Firebase local cache files.

## Why Firebase Hosting Was Used

Firebase Hosting was chosen because it is a managed static hosting option backed by Google Cloud. For a static documentation site, it provides a simpler path than manually building a Cloud Storage bucket, HTTPS load balancer, managed certificate, and CDN configuration.

Firebase Hosting is a strong first implementation because it handles several production-facing concerns with less operational overhead:

- static asset hosting,
- HTTPS support,
- custom domain support,
- simple CLI deployment,
- predictable deployment workflow,
- low infrastructure complexity.

This project may later be extended with a more infrastructure-heavy implementation using Cloud Storage, Cloud Load Balancing, Cloud CDN, and managed certificates. The Firebase version is the practical first deployment path.

## Implementation Steps

### 1. Create the GCP-hosted site repository

A separate repository was created for the GCP-hosted implementation.

The original documentation site remained useful as a starting point, but the new repository became the source for the Firebase-hosted version.

The new repo copied the MkDocs project structure:

```text
docs/
mkdocs.yml
requirements.txt
README.md
.gitignore
LICENSE
```

The generated MkDocs output folder was intentionally not treated as source:

```text
site/
```

### 2. Update MkDocs metadata

The `mkdocs.yml` file was updated so the site points to the planned custom domain:

```yaml
site_url: https://cloudengineer.mcqueencloud.com/
```

The repository metadata was also updated to reflect the GCP-hosted repository.

### 3. Build the MkDocs site locally

The local build process used a Python virtual environment.

```powershell
.venv\Scripts\activate
mkdocs build --strict
```

The strict build confirmed that the documentation source could be converted into static site output without broken navigation or missing files.

### 4. Install and authenticate the Firebase CLI

Firebase Hosting required the Firebase CLI.

The first issue was that `npm` was not recognized, which meant Node.js was not installed or was not available on the Windows PATH.

The resolution was to install Node.js LTS, reopen the terminal, and confirm:

```powershell
node -v
npm -v
```

Then Firebase CLI was installed globally:

```powershell
npm install -g firebase-tools
```

The CLI authentication was completed with:

```powershell
firebase login
```

### 5. Create and prepare the Google Cloud project

A Google Cloud project was created for the site.

The working project ID became:

```text
cloud-engineer-site-496219
```

The project was intended to support the Firebase-hosted static site.

### 6. Troubleshoot Firebase project attachment

The first major deployment blocker occurred when trying to add Firebase to the Google Cloud project.

The Firebase CLI returned:

```text
Error: Failed to add Firebase to Google Cloud Platform project.
```

The debug log showed that the Firebase `addFirebase` API call failed with:

```text
403 PERMISSION_DENIED
The caller does not have permission
```

This required several checks and fixes.

#### Issue: Domain Restricted Sharing organization policy

The Google Cloud organization had the legacy **Domain Restricted Sharing** policy enabled:

```text
constraints/iam.allowedPolicyMemberDomains
```

This policy can block IAM principals outside the allowed domain from being added. That matters because Firebase may need to add Google-managed service agents or service identities to the project.

The project-level policy was reviewed and overridden for the specific `Cloud Engineer Site` project. The effective policy was changed to:

```text
Allowed: All
```

This kept the exception scoped to the project rather than weakening the entire organization.

#### Issue: Required APIs were not enabled

The following APIs were enabled on the project:

```text
Firebase Management API
Firebase Hosting API
Cloud Resource Manager API
Service Usage API
```

These APIs were needed for Firebase setup and hosting management.

#### Issue: Firebase Terms of Service

Firebase setup still failed after IAM and policy changes. A dummy Firebase project was created from the Firebase Console so the Firebase Terms of Service could be accepted in the browser.

This mattered because Firebase Terms acceptance cannot always be completed through the CLI flow.

#### Issue: stale authentication/session state

After the policy changes, API enablement, and Terms acceptance, the Firebase CLI still failed until the Firebase CLI session was refreshed.

The resolution was:

```powershell
firebase logout
firebase login
```

After logging out and logging back in, Firebase setup succeeded.

### 7. Initialize Firebase Hosting

Firebase Hosting was initialized from the repository root:

```powershell
firebase init hosting
```

The selected setup was:

```text
Hosting only
Use an existing Firebase project
Public directory: site
Configure as a single-page app: No
Set up automatic GitHub deploys: No
Overwrite index.html: No
```

Firebase created the hosting configuration files:

```text
firebase.json
.firebaserc
```

The important hosting setting is that Firebase serves the generated MkDocs output from:

```text
site
```

### 8. Deploy to Firebase Hosting

The site was built and deployed manually:

```powershell
mkdocs build --strict
firebase deploy --only hosting
```

Firebase provided a default hosting URL first. That confirmed the static site could be served successfully before custom domain setup.

### 9. Configure the custom domain

The preferred domain was:

```text
cloudengineer.mcqueencloud.com
```

In Firebase Hosting, a custom domain was added for the subdomain.

Firebase provided this DNS record:

```text
Type: CNAME
Domain name: cloudengineer.mcqueencloud.com
Value: cloud-engineer-site-496219.web.app
```

Because the DNS provider was Squarespace, the custom DNS record was entered as:

```text
Type: CNAME
Host: cloudengineer
Value: cloud-engineer-site-496219.web.app
```

After DNS propagation and Firebase verification, the custom domain worked.

### 10. Update `.gitignore`

The generated site files and Firebase local cache files should not be committed.

The `.gitignore` was updated to include:

```gitignore
# MkDocs build output
site/

# Firebase local/cache files
.firebase/

# Firebase debug logs
firebase-debug.log
firebase-debug.*.log
```

This prevents accidental commits of generated files such as:

```text
site/index.html
site/sitemap.xml
site/assets/
.firebase/hosting.*.cache
```

### 11. Preserve a GitHub Pages fallback strategy

The original GitHub Pages deployment was unpublished to prevent two public versions from drifting apart.

The plan is to make the GCP-hosted repository the source of truth and eventually deploy the same MkDocs source to two targets:

```text
Firebase Hosting -> primary custom domain
GitHub Pages -> free fallback URL
```

That keeps one documentation source while preserving a no-cost fallback if Firebase hosting is paused later.

## Issues and Resolutions

| Issue | Cause | Resolution |
|---|---|---|
| `npm` was not recognized | Node.js was not installed or not on PATH | Installed Node.js LTS and reopened the terminal |
| Firebase project setup failed | Firebase could not attach to the GCP project | Reviewed IAM, org policy, APIs, Terms, and CLI auth |
| `403 PERMISSION_DENIED` on `addFirebase` | Initially appeared to be IAM, but also involved org policy/API/session state | Granted appropriate project roles, enabled APIs, accepted Firebase Terms, refreshed CLI login |
| Domain Restricted Sharing blocked setup | Organization policy restricted allowed IAM principals | Added a project-level override with `Allowed: All` for this specific project |
| Required Firebase APIs were missing | APIs were not enabled on the project | Enabled Firebase Management API, Firebase Hosting API, Cloud Resource Manager API, and Service Usage API |
| Firebase setup still failed after fixes | Firebase CLI auth/session state was stale | Ran `firebase logout` and `firebase login` |
| Generated files appeared in Git changes | `mkdocs build` and Firebase deploy created local artifacts | Updated `.gitignore` to exclude `site/`, `.firebase/`, and Firebase debug logs |
| Default Firebase URL was not the desired public URL | Firebase Hosting creates a default URL first | Added `cloudengineer.mcqueencloud.com` as a custom domain and configured Squarespace DNS |

## Security and Operations Considerations

This project is simple, but it still includes real cloud engineering concerns.

### Access control

Write access to the repository and Firebase project should remain limited. Deployment rights should not be granted broadly just because the site is public.

### Organization policy

The Domain Restricted Sharing exception should remain scoped to the static site project. The organization-wide policy should not be weakened globally just to support one hosting project.

### Build artifacts

The `site/` directory is generated output. It should be rebuilt during deployment rather than committed to source control.

### Custom domain ownership

DNS records should be controlled through the domain provider. For this project, Squarespace manages DNS for `mcqueencloud.com`.

### Deployment model

Deployment is currently manual:

```powershell
mkdocs build --strict
firebase deploy --only hosting
```

A future improvement is to automate Firebase deployment with GitHub Actions after the manual process is stable.

## Cost Considerations

Firebase Hosting should be low cost for a small static documentation site, but the project should still be monitored.

Potential cost drivers include:

- storage usage,
- bandwidth/egress,
- build or deployment automation,
- future use of additional Firebase or Google Cloud services.

A no-cost GitHub Pages fallback is useful if the hosted custom-domain version is paused later.

## How to Extend This Project

- Add a GitHub Actions workflow for Firebase Hosting deployment.
- Add a GitHub Pages fallback workflow from the same repository.
- Add a staging channel using Firebase Hosting preview channels.
- Add uptime monitoring for the custom domain.
- Add a project architecture diagram.
- Add a short architecture decision record explaining why Firebase Hosting was chosen over Cloud Storage static hosting.
- Add a Cloud Storage + Load Balancer version as an advanced alternative.

## Portfolio Value

This project shows more than basic static hosting.

It demonstrates the ability to:

- build a technical documentation site from Markdown,
- configure a managed Google Cloud hosting service,
- troubleshoot IAM and organization policy issues,
- enable required APIs,
- configure custom DNS through an external registrar/DNS provider,
- separate source files from generated build artifacts,
- document the deployment process clearly,
- preserve a fallback hosting strategy.

That makes it a useful first cloud engineering portfolio project because it combines content, deployment, governance, DNS, and operational discipline.

## Official References

- [Firebase Hosting documentation](https://firebase.google.com/docs/hosting)
- [Firebase Hosting custom domains](https://firebase.google.com/docs/hosting/custom-domain)
- [Firebase CLI reference](https://firebase.google.com/docs/cli)
- [Use Firebase with existing Google Cloud projects](https://firebase.google.com/docs/projects/use-firebase-with-existing-cloud-project)
- [Google Cloud Domain Restricted Sharing](https://cloud.google.com/resource-manager/docs/organization-policy/domain-restricted-sharing)
- [Google Cloud Organization Policy Service](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
- [Google Cloud Service Usage documentation](https://cloud.google.com/service-usage/docs)
- [MkDocs documentation](https://www.mkdocs.org/)
- [MkDocs Material documentation](https://squidfunk.github.io/mkdocs-material/)
