# Farmish

**Project Description:**


## Installation

1. Initialize a git repository and clone the project:
    ```sh
    git init
    git clone git@github.com:saladin2098/farmish.**git**
    ```
2. Create a database named `farmish` on port `5432`.
3. Update the `.env` file with the appropriate configuration.
4. Use the following Makefile commands to manage the database migrations and set up the project:
    ```makefile
    # Set the database URL
    exp:
        export DBURL="postgres://mrbek:QodirovCoder@localhost:5432/farmish?sslmode=disable"

    # Run migrations
    mig-up:
        migrate -path migrations -database ${DBURL} -verbose up

    # Rollback migrations
    mig-down:
        migrate -path migrations -database ${DBURL} -verbose down

    # Create a new migration
    mig-create:
        migrate create -ext sql -dir migrations -seq create_table

    # Create an insert migration
    mig-insert:
        migrate create -ext sql -dir migrations -seq insert_table

    # Generate Swagger documentation
    swag:
        swag init -g api/handler.go -o api/docs

    # Clean up migrations (commented out by default)
    # mig-delete:
    #   rm -r db/migrations
    ```
5. Set the environment variable and run the project:
    ```sh
    make exp
    make mig-up
    go run main.go
    ```
6. Open the following URL to access the Swagger documentation:
    ```
    http://localhost:8080/api/swagger/index.html#/
    ```

## Features
1. Your text
2. Your text
3. Your text
4. Your text

## Dependencies

- **Scheduling**: [github.com/go-co-op/gocron](https://github.com/go-co-op/gocron)
- **Swagger**: [github.com/swaggo/swag](https://github.com/swaggo/swag)
- **Database**:
    - [database/sql](https://golang.org/pkg/database/sql/)
    - [github.com/lib/pq](https://github.com/lib/pq)
- **Environment Variables**: [github.com/joho/godotenv](https://github.com/joho/godotenv)
- **API Framework**: [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- **Migrations**: [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
****
## Acknowledgments

- Azizbek Qodirov
- Shamsiddin Okilov
- Feruza Mirjalilova

## Known Issues
Your text

## Special Features
Your text
