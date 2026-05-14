FROM python:3.12-slim@sha256:401f6e1a67dad31a1bd78e9ad22d0ee0a3b52154e6bd30e90be696bb6a3d7461 AS site-builder

WORKDIR /app

ENV PIP_DISABLE_PIP_VERSION_CHECK=1 \
    PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    NO_MKDOCS_2_WARNING=true

COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY mkdocs.yml ./
COPY docs ./docs

RUN mkdocs build --strict

FROM cgr.dev/chainguard/go:latest@sha256:1c82b6aaf8e34ac560c77048cc7e1bd1a774b3ab2a8b5d60ceaa55663a7aeb35 AS server-builder

WORKDIR /src

COPY docker/server/main.go ./main.go

ENV CGO_ENABLED=0

RUN ["/usr/bin/go", "build", "-trimpath", "-ldflags=-s -w", "-o", "/server", "./main.go"]

FROM scratch AS runtime

ENV PORT=8080

COPY --from=server-builder /server /server
COPY --from=site-builder /app/site /site

USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/server"]
