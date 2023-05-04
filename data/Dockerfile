FROM golang:1.19-alpine AS build
RUN apk --update add git
WORKDIR /src/
ADD . /src
RUN CGO_ENABLED=0 go build -o /bin/db-init

FROM scratch
LABEL org.opencontainers.image.source=https://github.com/carhartl/k8s-playground
COPY --from=build /bin/db-init /bin/db-init
USER nonroot:nonroot
ENTRYPOINT ["/bin/db-init"]
