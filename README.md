# gext-server
An anonymous textboard backend written in golang. Uses MySQL.

See [gext-client](https://github.com/owenoclee/gext-client) for the client.

## Dependencies
Dependencies can be resolved by running `go get`.

## Setup
You need to have a MySQL server with a database set up for gext-server to use.

Modify the `DataSourceName` value in the `cfg/database.json` configuration file to point to your
database. Please see the
[Go-MySQL-Driver documentation](https://github.com/go-sql-driver/mysql/#dsn-data-source-name) for
more information on DSN.

## Build & run
Compile the protobuf files to Go source:
```
npm run protos
```

Build the server:
```
go build
```

Run the server:
```
./gext-server
```
