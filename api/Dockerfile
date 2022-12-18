FROM gcr.io/distroless/static
LABEL maintainer="nandor-magyar"
COPY --chown=nonroot:nonroot ./cmd/server/api.bin /api.bin
ENTRYPOINT ["/api.bin"]