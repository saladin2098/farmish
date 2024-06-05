# Farmish

Farmish is a comprehensive farm management application designed to streamline various aspects of animal husbandry, including animal tracking, health monitoring, and feeding schedules.

## Features

- Animal tracking with detailed attributes
- Health condition monitoring
- Feeding schedule management

## Prerequisites

- Go 1.16 or highera
- PostgreSQL

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/saladin2098/farmish.git
    cd farmish
    ```

2. Set up the PostgreSQL database using the provided migrations:

    ```bash
    psql -U youruser -d yourdb -f migrations/initial.sql
    ```

3. Copy the `.env.example` to `.env` and configure your environment variables:

    ```bash
    cp .env.example .env
    ```

4. Install the dependencies:

    ```bash
    go mod tidy
    ```

## Usage

To run the application:

```bash
go run main.go
