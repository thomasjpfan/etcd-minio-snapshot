FROM golang:1.11.0-alpine3.8 as builder

WORKDIR /develop
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o download-etcd-snapshot -ldflags '-w'

FROM gcr.io/etcd-development/etcd:v3.3.9

ENV ETCDCTL_API=3

COPY entrypoint.sh /usr/local/bin/entrypoint
COPY --from=builder /develop/download-etcd-snapshot /usr/local/bin/download-etcd-snapshot
RUN chmod +x /usr/local/bin/download-etcd-snapshot /usr/local/bin/entrypoint

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=10 CMD [ "etcdctl", "endpoint" , "health"]

CMD ["entrypoint"]
