FROM golang:1.19-alpine AS build
RUN apk --update add git
WORKDIR /src/
ADD . /src
RUN CGO_ENABLED=0 go build -o /bin/db-init

FROM scratch
COPY --from=build /bin/db-init /bin/db-init
USER nonroot:nonroot
ENTRYPOINT ["/bin/db-init"]
