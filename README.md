# Golang Webserver Project

This project is a Go-based web server utilizing the `chi` router for routing, `sqlc` for SQL query generation, and `goose` for database migrations. The API provides endpoints for user authentication and role-based access to student, teacher, and admin functionalities.

## Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [API Endpoints](#api-endpoints)
  - [User](#user)
  - [Student](#student)
  - [Teacher](#teacher)
  - [Admin](#admin)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

This web server project aims to provide a robust and scalable API service with role-based access control. It includes user authentication and separate endpoints for students, teachers, and admins. The server is built with Go, leveraging `chi` for routing, `sqlc` for type-safe SQL queries, and `goose` for database migrations.

## Features

- User authentication (login and registration)
- Role-based access control
- RESTful API design
- Database migrations and query generation
- Scalable and modular code structure

## Technologies Used

- [Go](https://golang.org/)
- [chi router](https://github.com/go-chi/chi)
- [sqlc](https://github.com/kyleconroy/sqlc)
- [goose](https://github.com/pressly/goose)
- [PostgreSQL](https://www.postgresql.org/)

## API Endpoints

### User

- `POST /v1/user/login`: User login endpoint.
- `POST /v1/user/register`: User registration endpoint.

### Student

- `GET /v1/student`: Get student information.

### Teacher

- `GET /v1/teacher`: Get teacher information.

### Admin

- `GET /v1/admin`: Get admin information.

