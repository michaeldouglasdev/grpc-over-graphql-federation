# gRPC over GraphQL Federation

This project demonstrates a GraphQL Federation architecture where gRPC services are used as subgraphs. The federation router aggregates multiple gRPC-based microservices into a unified GraphQL API.

## Architecture Overview

The project consists of:

- **Router**: A GraphQL Federation router that coordinates requests across multiple subgraphs
- **Subgraphs**: Multiple gRPC services that expose GraphQL schemas and act as federation subgraphs:
  - `categories`: Product categories service
  - `customers`: Customer management service
  - `offers`: Product offers and promotions service
  - `products`: Product catalog service
  - `services`: General services management

## Technology Stack

- **GraphQL Federation**: For API composition, providing a unified GraphQL API
- **gRPC**: For inter-service communication and service definitions
- **Go**: Programming language for services
- **Code Generation**: GraphQL schema-first approach with automatic protobuf generation
- **Protocol Buffers**: Generated from GraphQL schemas for service contracts

## Project Structure

```
├── router/                 # Federation router configuration
└── subgraphs/             # Individual gRPC services
    ├── categories/        # Categories subgraph
    ├── customers/         # Customers subgraph
    ├── offers/           # Offers subgraph
    ├── products/         # Products subgraph
    └── services/         # Services subgraph
```

Each subgraph follows a consistent structure with:

- Protocol buffer definitions (`proto/`)
- Generated gRPC code (`generated/`)
- GraphQL schema (`graph/`)
- Service implementation (`service/`)
- Data layer (`data/`)

## Getting Started

1. Navigate to the router directory and start the federation router
2. Start individual subgraph services
3. The router will automatically discover and federate the subgraphs

This architecture allows for independent development and deployment of microservices while providing a unified GraphQL interface to clients.
