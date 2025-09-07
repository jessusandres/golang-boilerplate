# <center> Golang Boilerplate (Gin + GCP Functions + GORM) </center>

<div align="center" style="margin-bottom: 20px; margin-top: 20px; font-size: 50px; display: flex; align-items: center; justify-content: center;">
    <img width="159px" src="https://go.dev/images/gophers/motorcycle.svg">
    <span style="margin: 0 20px;">+</span>
    <img width="159px" src="https://gorm.io/gorm.svg">
    <span style="margin: 0 20px;">+</span>
    <img width="125px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">
</div>

---

### This repository provides a practical starter for building HTTP APIs that can run both:
- As a local server (Gin) for development
- As a Google Cloud Function via functions-framework-go

It also includes:
- Postgres integration via GORM
- Pub/Sub listener scaffold
- Structured logging (zap)
- Request validation middleware
- Clear layering: routes, controllers, services, tasks, models, etc.

---

## Quick start
- Requirements
  - Go 1.25+
  - Docker (optional, for Postgres)
  - make (optional)
- Clone and install
  - git clone <this-repo>
  - cd golang-boilerplate
  - go mod download
- Configure environment
  - Copy and adjust the following .env example in the project root:

    FUNCTION_TARGET=HttpEntrypoint
    PORT=8080
    GO_ENV=development
    API_PREFIX=api
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=looker
    DB_USERNAME=postgres
    DB_PASSWORD=postgres
    DB_MIN_CONNECTIONS=0
    DB_MAX_CONNECTIONS=10
    DB_SSL_MODE=disable
    DB_LOGGER=true
    USE_SQL_CONNECTOR=false

  - Note: FUNCTION_TARGET must match the name you want to expose as an HTTP function (see function.go). Any string is accepted as long as it stays consistent across .env and deployment.
- Start Postgres (optional, via Docker Compose)
  - Ensure the .env contains the DB_* variables
  - docker compose up -d
  - This will also apply db_init/ddl.sql on first run (includes demo data)
- Run locally (Gin server via Functions Framework)
  - go run ./cmd
  - Optionally, set LOCAL_ONLY=true to bind to 127.0.0.1 only
  - Server listens on http://localhost:8080 by default
- Run tests
  - go test ./...


## API overview
- Base routes
  - GET /        → {"message": "Hello from Gin!"}
  - GET /ping    → pong
  - GET /fail    → returns a 400 error (sample error path)
- Incidents endpoints
  - Base path is prefixed with API_PREFIX (default "/api"). Full path: POST {API_PREFIX}/incidents
  - POST {API_PREFIX}/incidents
    - Headers: Content-Type: application/json
    - Body:
      {
        "title": "string (required)",
        "description": "string (required)",
        "incidentType": "string (required)",
        "location": "string (required)",
        "image": "base64 string (optional)",
        "eventDate": "ISO 8601 timestamp (required)"
      }
    - Response: 200 OK
      {
        "data": {
          "incident": {
            "id": number,
            "title": string,
            "description": string,
            "incidentType": string,
            "location": string,
            "image": string,
            "eventDate": string,
            "createdAt": string
          }
        }
      }
    - Example:
      curl -X POST "http://localhost:8080/api/incidents" \
           -H "Content-Type: application/json" \
           -d '{
                 "title": "Power outage",
                 "description": "Area wide blackout",
                 "incidentType": "Power",
                 "location": "Sector 7",
                 "image": "",
                 "eventDate": "2025-09-06T20:00:00Z"
               }'


## Project structure
- Root
  - README.md
  - docker-compose.yml            → Local Postgres with seed data (db_init/ddl.sql)
  - function.go                   → Registers the HTTP entrypoint for Cloud Functions
  - go.mod, go.sum
  - runDeploy.sh, sonar.properties (if used)
- cmd/
  - main.go                       → Local entrypoint; starts functions framework server
  - config/config.go              → Loads .env and validates against schema
  - controllers/                  → HTTP controllers (e.g., incidents.controller.go)
  - dto/                          → Data transfer objects (request/response shapes)
  - errors/
    - api/                        → API error transport objects
    - app/                        → Application-level error helpers
  - instances/
    - gorm.instance.go            → GORM DB connection (Postgres)
    - pubsub.instance.go          → Pub/Sub listener scaffold
  - interfaces/                   → Interfaces for services, validators, etc.
  - logger/logger.go              → zap SugaredLogger initialization
  - middlewares/                  → Gin middlewares (errors, state, validation)
  - models/                       → GORM models (e.g., Incident)
  - routes/
    - router.go                   → Base routes and API group registration
    - incidents.router.go         → Incidents route wiring
    - not_foud.route.go           → 404 handler
  - services/                     → Business logic (e.g., IncidentsService)
  - tasks/                        → Unit-of-work operations (e.g., SaveIncident)
  - types/                        → Shared types (env config, API responses, validation)
  - utils/                        → Helpers (env parsing, error helpers, API response)
- db_init/
  - ddl.sql                       → Schema and demo data for local DB


## Configuration and environment
- See cmd/types/env.types.go for the full environment schema
  - PORT: server port (default 8080)
  - FUNCTION_TARGET: required; exported HTTP function name
  - API_PREFIX: default /api
  - GO_ENV: development|production (affects logger)
  - DB_*: connection settings
  - USE_SQL_CONNECTOR: set true to use Cloud SQL connector driver
  - DB_LOGGER: toggle GORM logging
- config/config.go loads .env via godotenv (skips when testing)


## Local development notes
- The app uses Google Functions Framework under the hood for parity with Cloud Functions. main.go starts the framework and binds Gin to it.
- For 127.0.0.1 binding, set LOCAL_ONLY=true.
- 404 handler returns a JSON error with a request UUID when available.


## Database
- docker-compose.yml starts a Postgres 17 container and mounts db_init/ddl.sql to auto-create the schema and seed demo data at first startup.
- Default DB credentials are read from .env. Adjust as needed.


## Pub/Sub listener
- cmd/instances/pubsub.instance.go contains a scaffold to receive messages and persist incidents using the same task as the HTTP flow.
- It is not started by default. Wire it in as needed (e.g., from an init or main) and fill projectID/subscriptionName.


## Testing
- Unit tests exist under cmd/utils (api, errors, service helpers)
- Run: go test ./...


## Deployment (Cloud Functions HTTP)
- Ensure FUNCTION_TARGET in env matches the function name you register.
- function.go dynamically registers functions.HTTP(functionTarget, handler).
- Follow gcloud CLI steps to deploy an HTTP Cloud Function with Go and set the same FUNCTION_TARGET.

## License
- MIT or as appropriate for your project.
