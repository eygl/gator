# Blog Aggrigator App

This app is part of bootdotdev course.

## Setting Up Postgres

1. Install Postgres

    `sudo apt update`

    `sudo apt install postgresql postgresql-contrib`

    - Now run `psql --version` to verify it is working. 
    - Notes: You can set the password on linux with:

        `sudo passwd postgres`

2. Now start postgres service

    `sudo service postgresql start`

3. Open psql shell

    `sudo -u postgres psql`

4. Create a new database. This example I create a table named gator.

    `CREATE DATABASE gator;`

5. Change database to `gator`

    `\c gator`

5. Set password

    `ALTER USER postgres PASSWORD 'postgres'`

## Setting up Goose

Goose is a database migration tool written in Go. It runs migrations from a set of SQL files, making it a perfect fit for this project.

1. Install Goose
`go install github.com/pressly/goose/v3/cmd/goose@latest`

2. Create sql/schema directory and add 001_users.sql

    `mkdir sql/schema`

    `touch 001_users.sql`

    Then add the following to the migration file:

```sql
-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY
    created_at TIMESTAMP NOT NUL
    updated_at TIMESTAMP NOT NUL
    name TEXT UNIQUE NOT NUL
);


--+goose Down
DROP TABLE users;
```

3. Get connection string. For linux it will be something like:

    `postgres://postgres:postgres@localhost:5432/gator`

4. Connect to database directly and run migration

    `goose postgres <connection_string> up`
    `goose postgres postgres://postgres:postgres@localhost:5432/gator up`

    Check the migration by going to psql gator and entering `\dt`.


## Setting up SQLC 

1. Install SQLC

    `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

2. Configure SQLC by creating a file called sqlc.yaml in root of project:
```yaml
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```

3. Make sure the directories mentioned above are created.

    ```bash
    mkdir sql/schema
    mkdir sql/queries
    ```

## Using SQLC