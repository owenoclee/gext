# gext-server
An anonymous textboard backend written in golang. Uses MySQL.

See [gext-client](https://github.com/owenoclee/gext-client) for the client.

## Prerequisites
gext-server requires go1.10 or above to compile. Dependencies can be resolved by running `go get`.

You need to have a MySQL server with a database set up for gext-server to use.

## Configuration
Modify the `DataSourceName` value in the `cfg/database.json` configuration file to point to your
database. Please see the
[Go-MySQL-Driver documentation](https://github.com/go-sql-driver/mysql/#dsn-data-source-name) for
more information on DSN.

The default server configuration will run on localhost:8080. To change this, set the `Address` value
in `cfg/server.json`.

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
