# Go Poker League Backend
A backend application built with **Go** for managing poker player scores and league rankings.

This project demonstrates practical Go backend development through a combination of **HTTP API development, JSON serialization, dependency injection, and command-line applications.**

The application was developed incrementally using **Test-Driven Development (TDD)**.

## Overview

The system provides a simple poker league where players can record wins and retrieve their current scores.

It exposes a REST-style HTTP API and supports persistent storage using MariaDB.

### Main Features

- REST-style HTTP API
- Player score management
- League ranking
- JSON API responses
- Automatic ranking based on wins
- Command-line interface
- Interface-based storage abstraction
- Dependency injection
- Integration testing
- Unit testing
- Multiple executable applications sharing the same business logic

---

## Architecture

The application separates the HTTP layer, application logic, and persistence layer.

```text
                     HTTP Client       CLI Application 
                         │                 │   
                         ▼                 │
                ┌─────────────────┐        │
                │   HTTP Server   │        │
                │                 │        │
                │  Routing        │        │
                │  HTTP Handlers  │        │
                │  JSON Encoding  │        │
                └────────┬────────┘        │
                         │                 │
                         ▼                 │
                ┌─────────────────┐        │
                │   PlayerStore   │        │
                │    Interface    │        │
                └────────┬────────┘        │
                         │                 │
              ┌──────────┴──────────┐      │
              │                     │      │
              ▼                     ▼      ▼
     ┌─────────────────┐   ┌─────────────────────┐
     │ In-Memory Store │   │ Database Store      │
     │                 │   │                     │
     │ Used for tests  │   │     MariaDB         │
     └─────────────────┘   └─────────────────────┘
                        
                         
```

## Running the Application
Store DB configuration in .env in project root

Build the executables
```bash
(windows)
go build .\cmd\webserver\ .\cmd\webserver
go build .\cmd\cli .\cmd\cli
```
Start the HTTP server:
```bash
(windows)
cd cmd\webserver
webserver.exe
```

The server listens on port 5000 and can be accessed locally at:
```text
http://localhost:5000
```

## API Endpoints
| Method | Endpoint         | Description |
|--------|------------------|-------------|
| GET    | /players/{name}  |Retrieve a player's total wins|
| POST   | /players/{name}  |Record one additional win|
| GET    | /league  |Retrieve all players sorted by wins|

## API Usage
:::spoiler Details
### Get Player Score
```http
GET /players/{name}
```
**Example**
```bash
curl http://localhost:5000/players/Ruth
```
**Response**
```text
3
```
This indicates that Ruth currently has 3 wins.

### Record a Player Win
```http
POST /players/{name}
```
**Example**
```bash
curl -X POST http://localhost:5000/players/Ruth
```
Each request increments the player's win count by one. If the player is new, a new player record will be created.
**For example**
```text
Before:

Ruth = 3 wins

POST /players/Ruth

After:

Ruth = 4 wins
```

### Get League Rankings
```http
GET /league
```
**Example**
```bash
curl http://localhost:5000/league
```
**Response**
```json
[
    {
        "Name": "Ruth",
        "Wins": 10
    },
    {
        "Name": "David",
        "Wins": 7
    },
    {
        "Name": "Alice",
        "Wins": 3
    }
]
```
Players are sorted by their number of wins.
::: 