# UpdateHub

UpdateHub is a server application written in Go that manages device updates. It features a web interface for uploading updates, managing devices, and automated responses to device update requests. The project uses PostgreSQL for data storage and supports database migrations, unit testing, and automated CI/CD using GitHub Actions.

## Features

- **Upload Updates:** Seamlessly upload new software updates.
- **Manage Devices:** Add, remove, and manage devices.
- **Automatic Updates:** Devices receive the latest software versions automatically.
- **Database Management:** Track devices and their software versions.

## Installation

### Prerequisites

- Go (version 1.22 or higher)
- Docker
- golang-migrate (for database migrations)
- sqlc (for database interaction)

### Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/AmirhoseinMasoumi/UpdateHub.git
    cd UpdateHub
    ```

2. Start the PostgreSQL database using Docker:
    ```sh
    make postgres
    ```

3. Create the database:
    ```sh
    make createdb
    ```

4. Run database migrations:
    ```sh
    make migrateup
    ```

5. Generate SQL code:
    ```sh
    make sqlc
    ```

6. Run the server:
    ```sh
    make server
    ```

## Makefile Commands

- **postgres**: Start PostgreSQL in a Docker container.
- **createdb**: Create the database.
- **dropdb**: Drop the database.
- **migrateup**: Apply all up migrations.
- **migrateup1**: Apply the next up migration.
- **migratedown**: Apply all down migrations.
- **migratedown1**: Apply the next down migration.
- **new_migration**: Create a new migration file.
- **sqlc**: Generate Go code from SQL schema.
- **mock**: Generate mock implementations for tests.
- **server**: Run the server.
- **test**: Run unit tests.

## GitHub Actions CI/CD

This project uses GitHub Actions for continuous integration. The workflow file is located in `.github/workflows`.

### Workflow Steps

1. **Checkout Code**: Uses `actions/checkout@v4` to checkout the code.
2. **Set up Go**: Uses `actions/setup-go@v4` to set up the Go environment.
3. **Install golang-migrate**: Installs the migration tool.
4. **Run Migrations**: Applies database migrations.
5. **Run Tests**: Executes unit tests.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request. For major changes, open an issue to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
