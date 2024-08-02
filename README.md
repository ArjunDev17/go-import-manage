# Go Import Manage

## Overview
`go-import-manage` is a Golang program designed to import data from an Excel file, store it into MySQL, and cache the data in Redis. It provides a simple CRUD system to view, edit, and update imported data in both the database and cache.

## Features
- Import Excel data
- Store data in MySQL
- Cache data in Redis
- View imported data
- Edit and update records
- Error handling and validation
- Asynchronous processing
- Scalable architecture

## Technologies Used
- Golang
- Gin Framework
- MySQL
- Redis

## Prerequisites
- Go 1.19+
- MySQL 5.7+
- Redis 6.0+
- Docker (optional, for easier setup)

## Getting Started

### Configuration
Create a `.env` file in the root directory and add the following environment variables:
