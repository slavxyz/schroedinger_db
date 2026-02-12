# schroedinger_db

Quirky Store is a small experimental key–value store written in Go that intentionally behaves incorrectly some of the time.  
Its goal is to simulate unpredictable and chaotic system behavior for debugging, testing and fault-tolerance experiments.

The store uses a real MySQL database as its backend and exposes a minimal API:

- Put(key, value)
- Get(key)
- Delete(key)
- Dump()

In addition to normal behavior, the store randomly introduces incorrect and unexpected results.

---

## Features

- Persistent key–value storage backed by MySQL
- Layered architecture (service, repository, chaos)
- Chaos engine injected via interfaces
- Random failures and incorrect operations
- Internal state inspection through `Dump()`
- Optional data mutation over time (values may change without `Put`)
- Deterministic test doubles for chaos behavior
- Integration tests using a real database

---

## Requirements

- Go 1.22 or newer
- MySQL or MariaDB
- A running database instance

---

## Database setup

Create the database and table:

```sql
CREATE DATABASE quirky_store;
USE quirky_store;
CREATE TABLE store (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    `key` VARCHAR(255) NOT NULL UNIQUE,
    `value` TEXT
);
INSERT INTO store (`key`, `value`)
VALUES ('cat', 'meow'),
    ('dog', 'woof'),
    ('cow', 'moo'),
    ('duck', 'quack'),
    ('lion', 'roar'),
    ('apple', 'fruit'),
    ('banana', 'yellow'),
    ('car', 'vehicle'),
    ('sky', 'blue'),
    ('water', 'wet');
```

## Installation

Clone the repository and move into the project directory:

git clone <your-repository-url>
cd QuirkyStore


Download dependencies:

go mod tidy

## Configuration

Edit the database connection string in:

cmd/app/main.go

and in:

tests/quirky_test.go

Update the following line with your credentials:

"username:password@tcp(localhost:3306)/quirky_store"

## Running the application

From the project root:

go run ./cmd/app

## Running tests

go test ./
