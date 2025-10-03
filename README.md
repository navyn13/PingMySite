# PingMySite

PingMySite is a lightweight Go application that lets you **check the response time and HTTP status of multiple websites** concurrently. It uses a worker pool pattern to efficiently handle multiple URLs at the same time, making it fast and reliable for website monitoring.

---

## Features

- Check multiple URLs in one request.
- Measures **response time** for each URL.
- Returns **HTTP status** (OK, Not Found, etc.).
- Concurrent execution with configurable number of workers.
- Easy to run locally or in Docker.

---

## Tech Stack

- **Language:** Go 1.25+
- **Concurrency:** Goroutines + channels (worker pool)
- **HTTP client:** Standard `net/http`
- **Containerization:** Docker

---

## Installation

### Prerequisites

- Go 1.25+ installed (for local development)
- Docker (optional, for containerized setup)

### Run Locally

1. Clone the repository:

```bash
git clone https://github.com/navyn13/PingMySite.git
cd PingMySite
