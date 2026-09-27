# Two stages, the same shape for every language: the build image assembles the app with ap-build,
# the runtime image runs it. The app serves HTTP on $PORT.
FROM ghcr.io/actionplatform/build-go:0 AS build
ARG TARGETARCH
WORKDIR /w
COPY . .
RUN AP_ARCH=$([ "$TARGETARCH" = "amd64" ] && echo x86_64 || echo arm64) AP_ARTIFACTS=/out ap-build package

FROM ghcr.io/actionplatform/runtime-go:0
WORKDIR /app
COPY --from=build /out /app
ENV PORT=8000
EXPOSE 8000
ENTRYPOINT ["/app/bootstrap"]
