# SMSC Gateway

A high-performance SMSC Gateway software designed for direct operator connectivity and bulk SMS services, with full Sigtran protocol support.

## Features

- Full Sigtran Protocol Support (SCTP, M3UA, M2PA, MTP3, SCCP, MAP)
- Multi-operator support with dynamic traffic routing
- High-performance bulk SMS processing
- Web-based admin panel
- Real-time monitoring and analytics
- Intelligent routing optimization
- Multi-tenancy support
- Comprehensive API with SDK support
- Advanced security features
- Billing and payment integration

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- PostgreSQL 14 or later

## Installation

### 1. Install Go

For macOS, using Homebrew:
```bash
brew install go
```

For Linux:
```bash
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 2. Install Docker

For macOS:
```bash
brew install --cask docker
```

For Linux:
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
```

### 3. Build and Run

1. Clone the repository:
```bash
git clone <repository-url>
cd smsc
```

2. Start the services using Docker Compose:
```bash
docker-compose up -d
```

The application will be available at:
- Admin Panel: http://localhost:8080
- SMPP Port: 2775
- SMPP TLS Port: 2776

## Configuration

Configuration files are located in the `config` directory. Copy the example configuration and modify as needed:

```bash
cp config/config.example.yaml config/config.yaml
```

## Development

### Project Structure

```
.
├── cmd/
│   └── smsc/
│       └── main.go
├── internal/
│   ├── api/
│   ├── core/
│   ├── db/
│   ├── models/
│   ├── protocols/
│   │   └── sigtran/
│   └── services/
├── pkg/
│   ├── logger/
│   └── utils/
├── web/
│   ├── admin/
│   └── api/
├── config/
├── docker-compose.yml
└── Dockerfile
```

### Building from Source

```bash
go build -o smsc-gateway ./cmd/smsc/main.go
```

## API Documentation

API documentation is available at `/api/docs` when running the server.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support, please contact me.
