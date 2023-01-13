FROM golang:1.19-alpine AS build
WORKDIR /src/
ADD . /src
RUN CGO_ENABLED=0 go build -o /bin/db-init

FROM scratch
COPY --from=build /bin/db-init /bin/db-init
COPY ./people.csv /data/people.csv
USER nonroot:nonroot
ENTRYPOINT ["/bin/db-init"]
