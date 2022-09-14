FROM alpine:3.16
COPY freebsd_exporter /bin/freebsd_exporter
RUN chmod 0700 /bin/freebsd_exporter
RUN apk --update add ca-certificates tzdata
ENTRYPOINT ["/bin/freebsd_exporter"]
