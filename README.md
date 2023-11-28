# Shopidel API Backend

Shopidel is a lightweight e-commerce platform built with Golang and Gin, replicating the functionality of Shopee Indonesia Marketplace. This API backend handles various aspects of the marketplace system, including user registration, product selection, order management, and payment processing.

## Tech Stack

-   **Golang**: The programming language used to develop the backend.
-   **Gin**: A web framework for Golang, providing a minimalistic yet powerful API.
-   **PostgreSQL**: A powerful, open-source relational database system.
-   **Gorm**: An Object Relational Mapping (ORM) library for Golang, used with PostgreSQL for database interactions.
-   **Redis**: An in-memory data structure store, utilized for caching and improving performance.
-   **Docker**: Containerization technology for packaging the application and its dependencies into a container.

## Features

-   **User Registration**: Allow users to register and create accounts.
-   **Product Selection**: Facilitate users in browsing and selecting products from the marketplace.
-   **Order Management**: Manage the end-to-end process of placing and tracking orders.
-   **Payment Processing**: Ensure secure and seamless payment transactions.
-   **Seller Management**: Enable sellers to register and manage their seller accounts. Sellers can view and process orders related to their products. Provide APIs for seller-specific functionalities.

## Getting Started

### Prerequisites

-   Golang installed on your machine.
-   PostgreSQL and Redis set up and running.
-   Docker installed

### Installation

1. Clone the repository
2. Set up the configuration by copying `.env.example` to `.env` and adjusting the values as needed.
3. Install dependencies and run the application:
    ```
     go mod download
     go run main.go
    ```
    Or using docker
    ```
    docker-compose up
    ```

### Documentation

For detailed information on using the Shopidel API, refer to the [API documentation.](https://documenter.getpostman.com/view/16000432/2s9YXcc4WC)