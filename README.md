# gext
An anonymous textboard forum written in golang. Uses MySQL.

## Use Docker (recommended)
```
git clone https://github.com/owenoclee/gext.git
cd gext
docker compose up
```

After a few moments you should see gext at `localhost` on port 80.

## Manual Installation (not recommended)
### Prerequisites
* golang 1.7+
* MySQL 5.7 (other versions may work but have not been tested)

### Configuration
You need to export the following environment variables (and adjust them as necessary) before running gext manually:

```bash
export GEXT_ADDRESS=localhost
export GEXT_PORT=8080

export GEXT_DATASTORE=mysql
export GEXT_DATASTORE_MYSQL_DSN=root@/gext # see link below

# path to html templates (views) and public folder (i.e. for CSS)
export GEXT_VIEWS_PATH=/go/src/github.com/owenoclee/gext/views/
export GEXT_PUBLIC_PATH=/go/src/github.com/owenoclee/gext/public/
```

The format for the MySQL data source name (DSN) is described in detail [here](https://github.com/go-sql-driver/mysql#dsn-data-source-name).

### Install & Run
Ensure your `GOPATH` environment variable is set properly before installing. To make running gext easier you should also put the go bin folder in your `PATH` if it is not already there (i.e. `PATH=$PATH:$GOPATH/bin`).

```bash
go get github.com/owenoclee/gext
go install github.com/owenoclee/gext
gext
```
