# VTT Odisseia

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![WebSocket](https://img.shields.io/badge/WebSocket-supported-brightgreen)
![REST API](https://img.shields.io/badge/REST_API-supported-blue)

This project is a MVP for a Virtual Table Top application.

## Running Project

### Development mode

Starting envirioment

```bash
DOCKERFILE=./deployments/dev.dockerfile docker-compose up
```

Starting application

```bash
go run main.go
```

## What is VTT Odisseia?

VTT Odisseia is a project of a virtual table top, with features like:

- User authentication
- Rooms management
- Real-time chat
- User presence in rooms

## Why?

I started this project because I wanted to learn Go and I needed a project to do so. I play tabletop RPGs with my friends, but we lived too far away, so I decided to create a platform to make it happen. I believe that platforms of the same type that currently exist do not offer the features we would like.

## Features

- User authentication with JWT
- Room management, with creation, deletion and update
- Real-time chat with WebSockets
- User presence in rooms

## Tech Stack

- Go 1.17
- Gorm
- Gorilla WebSocket
- Redis
- Docker
- Docker Compose

## How can I help?

You can help me by:

- Giving me feedback on the project
- Opening issues for bugs or features
- Creating pull requests with new features or bug fixes
