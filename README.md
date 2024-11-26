# {{project_name}}

{{project_description}}

## Using this template

1. Click the "Use this template" button at the top of the repository
2. Choose a name for your new repository
3. Clone your new repository locally
4. Run the template setup script:
   ```bash
   bash ./.github/template-setup.sh

{{project_description}}

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go
- Node.js
- pnpm
- Protocol Buffers compiler (will be installed automatically if using Homebrew)

## Available Commands

### Development

Start both API and web servers in development mode:
```bash
make dev
```

Start only the API server:
```bash
make dev-api
```

Start only the web server:
```bash
make dev-web
```

### Protocol Buffers

Generate protocol buffer code:
```bash
make proto
```

Install protocol buffer dependencies:
```bash
make install-proto-deps
```

### Building and Testing

Build the application:
```bash
make build
```

Run the test suite:
```bash
make test
```

### Other Commands

Watch mode for development (uses Air for live reload):
```bash
make watch
```

Clean generated files:
```bash
make clean
```

## Environment Variables

The project uses environment variables for configuration. Create a `.env` file at the root of the project based on `.env.example`.

### Server Configuration
- `NODE_ENV=development`: Node environment (development/production)
- `PORT=8080`: API server port
- `ALLOWED_ORIGINS=http://localhost:5173`: Comma-separated list of allowed CORS origins

### gRPC Configuration
- `GRPC_HOST=0.0.0.0`: gRPC server host
- `GRPC_PORT=50051`: gRPC server port

### Frontend Configuration (Vite)
All frontend environment variables are prefixed with `VITE_`.

#### API Settings
- `VITE_API_HOST=localhost`: API server host
- `VITE_API_PORT=8080`: API server port (should match `PORT`)

#### gRPC Web Settings
- `VITE_GRPC_HOST=localhost`: gRPC-Web client host (should match `GRPC_HOST` for development)
- `VITE_GRPC_PORT=50051`: gRPC-Web client port (should match `GRPC_PORT`)

Copy `.env.example` to `.env` to get started:
```bash
cp .env.example .env
```
