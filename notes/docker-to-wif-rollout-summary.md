# Docker to WIF Rollout Summary

## Purpose

This is a sanitized project summary covering the work completed from the initial Docker containerization effort through the successful Workload Identity Federation deployment cutover.

It is written for external context sharing and omits secret values, provider resource strings, and other sensitive environment details.

## Goals

- Containerize the MkDocs site.
- Reduce image vulnerability exposure.
- Add automated container scanning.
- Harden GitHub Actions workflows.
- Remove dependence on legacy Firebase deploy tokens.
- Move Firebase deployment authentication to a keyless GitHub OIDC plus Workload Identity Federation model.

## Rollout Summary

1. Added a Docker-based build and runtime path for the MkDocs site.
2. Fixed an asset-path issue discovered during the first container build.
3. Replaced the initial web-server runtime approach with a smaller custom static server and a minimal runtime image to reduce security findings.
4. Pinned build-stage images by digest for repeatability.
5. Added automated container security scanning in GitHub Actions.
6. Hardened workflow supply-chain posture by pinning Actions and tightening permissions.
7. Verified the container locally with successful homepage and 404 checks.
8. Confirmed a clean vulnerability posture on the final image after the runtime redesign.
9. Investigated the failing Firebase Hosting deployment and isolated the problem to CI authentication rather than site configuration.
10. Moved away from the legacy Firebase token approach toward Google-auth-based deployment.
11. Discovered that organization policy blocked service account key creation, which ruled out the usual JSON-key workflow.
12. Switched to Workload Identity Federation using GitHub OIDC and service-account impersonation.
13. Updated the deploy workflow so the non-secret service account identity is hardwired and only the Workload Identity Provider reference remains externalized.
14. Configured the Google Cloud side trust relationship and GitHub repository secret for the provider reference.
15. Pushed the workflow change and verified that both the Firebase Hosting workflow and the fallback publishing workflow succeeded.

## Resulting State

- The repository now has a reproducible Docker build path.
- The final runtime image is minimal and security-scanned.
- The Firebase deploy path uses short-lived federated credentials instead of long-lived secrets.
- The deployment model is more compatible with stricter organization policies around service account keys.
- The workflow is simpler to operate because only the Workload Identity Provider reference remains configurable on the GitHub side.

## Key Decisions

- Prefer a minimal runtime image over a conventional web-server image when scanner results materially improve.
- Treat action pinning and permission reduction as part of the delivery work, not a later hardening pass.
- Prefer keyless federation over long-lived CI secrets when organization policy and platform support allow it.
- Hardwire non-secret identifiers in workflow files when they are stable and reduce operational complexity.

## Lessons Learned

- Build-stage images can still trigger scanner noise even when the final runtime image is small.
- Static-site repositories still benefit from container discipline if the container is treated as a first-class delivery artifact.
- Firebase Hosting failures in CI are often auth-path issues rather than site-build issues.
- Workload Identity Federation is more operationally durable than token-based CI auth in locked-down environments.

## Current Follow-Up Items

- Remove obsolete GitHub secrets that are no longer used by the deploy workflow.
- Keep the Workload Identity Provider reference current if the provider is ever recreated.
- Retain the deployment note in the repository so future auth changes start from the current operating model.