FROM golang:1.19-alpine AS build
RUN apk --update add git
WORKDIR /src/
ADD . /src
RUN CGO_ENABLED=0 go build -o /bin/server

FROM scratch
COPY --from=build /bin/server /bin/server
USER nonroot:nonroot
ENTRYPOINT ["/bin/server"]