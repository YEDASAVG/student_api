# student_api
A REST API to manage student records. You can create, read, update and delete student records.

## Overview
It provides REST API to manage student records.

It is for developers or users who need to create and manage student data programmatically.

Right now it supports a healthcheck endpoint and student CRUD endpoints will be added next.

## Tech-Stack
It is Built with Go, Gin- Go HTTP framework, PostgreSQL, GORM, Make

## Prerequisites
- Go
- PostgreSQL
- Make

## Project Structure

- cmd/ contains the application entry point.
- internal/handlers/ contains the HTTP handlers for API routes
- internal/models/ contains the data structure used in application
- internal/db/ contains the database connection and database-related logic
- internal/config/ contains the application configuration and environment varibake loading
## Getting Started

1. Clone the repo
    - git clone https://github.com/YEDASAVG/student_api.git
2. Move into project directory
    - cd student_api
3. Download Go dependencies
    - go mod tidy
4. Copy the example env file and update values
    - Open .env and set your DATABASE_URL and PORT
5. Create the database in PostgreSQL
    - createdb student_api
6. Run the app
    - make run

## Environment Variables 
This project uses .env file to load configuration. See .env.example for the template.
| Variable | Description | Example |
|----------|-------------|---------|
| DATABASE_URL | PostgreSQL connection string | postgresql://user@localhost:5432/student_api?sslmode=disable |
| PORT | Port the API server listens on | 8080 |

## Available Make commands
- make run - starts application locally.
- make build - builds the application binary
- make test - runs all tests in the project

## Running the Application
1. Run the application with make run
2. By default the API runs on port 8080
3. You can verify that the service is working by calling the healthcheck endpoint at /healthcheck.

## API Endpoints
- GET /healthcheck - returns the health status of the service
