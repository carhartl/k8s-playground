module db-init

go 1.19

require (
	github.com/go-faker/faker/v4 v4.1.0
	github.com/yugabyte/pgx/v4 v4.14.3
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.11.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.10.0 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
)

replace golang.org/x/crypto v0.0.0-20220214200702-86341886e292 => golang.org/x/crypto v0.0.0-20220314234659-1baeb1ce4c0b
