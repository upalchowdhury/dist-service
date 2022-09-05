# START: beginning
FROM golang:1.18-alpine AS build
WORKDIR /go/src/dist-service
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/dist-service ./cmd/distlog
# END: beginning
# START_HIGHLIGHT
# RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
#     wget -qO/go/bin/grpc_health_probe \
#     https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/\
#     ${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
#     chmod +x /go/bin/grpc_health_probe
# END_HIGHLIGHT
# START: beginning

FROM scratch
COPY --from=build /go/bin/dist-service /bin/dist-service
# END: beginning
# START_HIGHLIGHT
# COPY --from=build /go/bin/grpc_health_probe /bin/grpc_health_probe
# END_HIGHLIGHT
# START: beginning
ENTRYPOINT ["/bin/dist-service"]
# END: beginning
