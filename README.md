# RabbitMQ Microservices decoupling Example

This project demonstrates a simple microservices architecture using RabbitMQ for communication between services in Go.

## Services

1. **Inventory Service** - Publishes inventory updates.
2. **Payment Service** - Subscribes to inventory updates to process payments.
3. **Order Service** - Subscribes to inventory updates to process orders.

## Prerequisites

- Docker installed on your machine.
