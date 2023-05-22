# Satang Test Project (Etheruem Transaction Monitor by Kittipol)

This is a Golang system for monitoring incoming and outgoing transactions of specified Ethereum addresses and storing the transaction data in a PostgreSQL database. It is designed to be scalable and handle a large number of Ethereum addresses.

## Features

- Read incoming and outgoing transactions of specified Ethereum addresses.
- Store transaction data in a PostgreSQL database.
- Scalable architecture to handle a large number of Ethereum addresses.

## Prerequisites

Before running the application, make sure you have the following prerequisites installed:

- Go programming language (version 1.16 or later)
- PostgreSQL database
- Ethereum node or an API service to retrieve Ethereum transaction data

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/seedchamp45/satang-go.git
   ```

2. Change into the project directory:

   ```shell
   cd satang-go
   ```

3. Install the project dependencies:

   ```shell
   go mod download
   ```

4. Update the `config.yml` file:

   - Specify your Ethereum node or API service endpoint.
   - Configure your PostgreSQL database connection details.

5. Initialize the PostgreSQL database:

   ```shell
   go run script/db.go
   ```

## Usage

1. Update the the Ethereum addresses you want to monitor in `config/config.yml`. Each address should be on a new line.

2. Start the Ethereum transaction monitor:

   ```shell
   go run cmd/main.go
   ```

3. The application will start monitoring the transactions for the specified Ethereum addresses and store them in the PostgreSQL database.

## Configuration

The configuration for the application is stored in the `condig/config.yml` file. You can modify the following settings:

- `ethereumNodeURL`: The URL of the Ethereum node or API service to retrieve transaction data.
- `database`:
  - `host`: The hostname or IP address of the PostgreSQL database server.
  - `port`: The port number on which the PostgreSQL database server is running.
  - `name`: The name of the PostgreSQL database.
  - `user`: The username to authenticate with the PostgreSQL database server.
  - `password`: The password to authenticate with the PostgreSQL database server.
- `address` : The ethereum address

## Result
<img width="750" alt="Screen Shot 2566-05-22 at 21 11 39" src="https://github.com/seedchamp45/satang-go/assets/8091233/e82b9b6a-d976-43de-a101-4245cd32d62e">
<img width="1083" alt="Screen Shot 2566-05-22 at 21 11 59" src="https://github.com/seedchamp45/satang-go/assets/8091233/48b11905-1670-4faa-a016-b7d4be857223">



