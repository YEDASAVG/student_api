# student_api
A REST API to manage student records. You can create, read, update and delete student records.

## Overview
It provides REST API to manage student records.

It is for developers or users who need to create and manage student data programmatically.

Right now it supports a healthcheck endpoint and student CRUD endpoints will be added next.

## Tech-Stack
It is Built with Go, Gin- Go HTTP framework, PostgreSQL, golang-migrate, Make

## Prerequisites
- Go
- PostgreSQL
- Make

## Project Structure

- cmd/ contains the application entry point.
- internal/handlers/ contains the HTTP handlers for API routes
- internal/models/ contains the data structure used in application
- internal/db/ contains the database connection and database-related logic
- migrations/ contains SQL migration files used to create and update the database

## Getting Started

1. Clone the repo
    - git clone https://github.com/YEDASAVG/student_api.git
2. Move into project directory
    - cd student_api
3. Download Go dependencies
    - go mod tidy
4. Set required environment variables
5. Run the app with make
    - make run

## Available Make commands
- make run - starts application locally.
- make build - builds the application binary
- make test - runs all tests in the project

## Running the Application
1. Run the application with make run
2. By default the API runs on port 8080
3. You can verify that the service is working by calling the healthcheck endpoint at /healthcheck.

## API Endpoints
- GET/healthcheck - returns the helath status of the service
