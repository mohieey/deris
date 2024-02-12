# godis

**godis** is a Redis-like server implemented in Go, supporting popular Redis commands such as `SET`, `GET`, `HSET`, `HGET`, `HGETALL`, and `PING`. It can be accessed using any Redis-compatible client.

## Features

- Implements standard Redis commands:

  - `SET`: Set the value of a key.
  - `GET`: Get the value of a key.
  - `HSET`: Set the value of a field in a hash.
  - `HGET`: Get the value of a field in a hash.
  - `HGETALL`: Get all fields and values in a hash.
  - `PING`: Check if the server is running.

- Can be interacted with using any Redis-compatible client.

## Getting Started

### Installation

```bash
# Clone the repository
git clone https://github.com/mohieey/godis.git

# Navigate to the project directory
cd godis

# Build the project
go build .

# Run!
./godis
```

# Example using Redis CLI

```bash
redis-cli set mykey "Hello, godis!"

redis-cli get mykey
```
