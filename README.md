# Citybike 2023

This is the pre-assignment for Solita Dev Academy 2023.

Citybike is app for displaying data of journeys made with city bikes in Helsinki Capital area.

### You can access running version of this app in [tiikewi.fi](https://tiikewi.fi/)

## Installation

Commands from app directory root.

### To run backend:

```bash
cd backend
docker-compose up
```

This starts MariaDB database and go backend with reflex, for ease of develompment.
If no .env variable given, backend should be up and running in **localhost:8080**.

### To run frontend:

```bash
cd frontend
npm install
npm start
```

Frontend should now be up and running. You can access it from **localhost:3000**.

## Technologies

### Go

Backend is build with Go, and it serves as simple API endpoint for fetching journey and station data from backend.
Backend also includes csv file handler for validating journey data.

### MariaDB

Using MariaDB.
Validated CSV files inserted to db as CSV by using DBeaver.

### Docker

Go and Mariadb are running in containers. There is a docker-compose file to start them.

### React

Frontend is made with React using Typescript.
Using TanStack Query.

### DigitalOcean

App is also up and running on web using DigitalOcean as service provider.

## TODOS

- Really improve and implement tests
- Make UI responsive for various screen sizes
- UI element for creating journeys/stations
- Improve CSV reading and validation and perhaps add ability to add csv files to db from client.
