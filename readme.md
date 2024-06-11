# Eth Storage Music Project - Second Test Activity

Welcome to the Eth Storage Music Project, an innovative platform designed to store and manage music files using blockchain technology. This project utilizes Go microservices and integrates with decentralized storage solutions to provide a secure and scalable infrastructure for music management.

## Overview

This project is part of the second test activity for Eth Storage, focusing on the seamless integration of music management with blockchain technology. It consists of three main microservices:

1. **Gateway Service**
2. **User Service**
3. **Music Service**

### Technology Stack

- **Programming Language:** Go
- **Framework:** Gin
- **Microservices:** Go-Micro
- **Storage Solutions:** Arweave (via the jrys library)
- **Database:** MySQL (via GORM)
- **Configuration Management:** INI
- **Message Broker:** RabbitMQ (via AMQP)
- **Authentication:** JWT
- **Registry:** Etcd
- **Caching:** Redis
- **Serialization:** JSON (via goccy/go-json)
- **Protobuf:** Google Protocol Buffers

## Microservices

### 1. Gateway Service

The Gateway Service acts as the entry point for all client requests. It handles routing and forwarding of requests to the appropriate services (User Service and Music Service).

### 2. User Service

The User Service manages all user-related functionalities, including:

- **Login and Registration:** Users can sign up and log in using their wallet addresses.
- **Post Management:** Users can create and manage posts.

### 3. Music Service

The Music Service handles all music-related functionalities, including:

- **Music Upload:** Upload audio files to Arweave using the jrys library.
- **Music Download:** Support downloading of audio files.
- **NFT Management:** Store and manage NFT information corresponding to each music file.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Docker (optional, for containerized deployment)
- An Arweave wallet and access to the jrys library