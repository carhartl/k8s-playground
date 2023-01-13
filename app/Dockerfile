FROM golang:1.19-alpine AS build
WORKDIR /src/
ADD . /src
RUN CGO_ENABLED=0 go build -o /bin/people

FROM scratch
COPY --from=build /bin/people /bin/people
USER nonroot:nonroot
ENTRYPOINT ["/bin/people"]
