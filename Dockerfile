FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.20-openshift-4.15 AS builder
WORKDIR /go/src/github.com/openshift/csi-driver-shared-resource
COPY . .
RUN make build

FROM registry.ci.openshift.org/ocp/4.15:base
COPY --from=builder /go/src/github.com/openshift/csi-driver-shared-resource/_output/csi-driver-shared-resource /usr/bin/
ENTRYPOINT []
CMD ["/usr/bin/csi-driver-shared-resource"]