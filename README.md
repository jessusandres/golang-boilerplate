# <center> Golang Boilerplate (Gin + GORM + CQRS) </center>

<div align="center" style="margin-bottom: 20px; margin-top: 20px; font-size: 50px; display: flex; align-items: center; justify-content: center;">
    <img width="159px" src="https://go.dev/images/gophers/motorcycle.svg">
    <span style="margin: 0 20px;">+</span>
    <img width="159px" src="https://gorm.io/gorm.svg">
    <span style="margin: 0 20px;">+</span>
    <img width="125px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">
</div>

---

### A practical starter for building HTTP APIs with Gin, GORM, and a clean CQRS modular architecture.

It includes:
- PostgreSQL integration via GORM
- CQRS setup (commands, queries, buses, handlers)
- Structured logging with slog
- Request validation middleware
- Clear layering: internal/modules (feature modules), shared utilities, HTTP routes/middlewares

---

## Quick start
- Requirements
  - Go 1.25+
  - Docker (optional, for Postgres)
- Clone and install
  - git clone <this-repo>
  - cd golang-boilerplate
  - go mod download
- Configure environment
  - Create a .env file in the project root. Example:

    PORT=8080
    LOCAL_ONLY=true
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

  - For Docker Compose database startup, also add:

    POSTGRES_PORT=5432
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    POSTGRES_DB=looker

- Start Postgres (optional, via Docker Compose)
  - docker compose up -d
  - This will also apply scripts/db/ddl.sql on first run (includes demo data)
- Run locally
  - go run ./cmd/server
  - Optionally, set LOCAL_ONLY=true to bind to 127.0.0.1 only
  - Server listens on http://localhost:8080 by default
- Run tests
  - go test ./...


## API overview
- Base routes
  - GET /         → {"message": "Hello from Gin CQRS API!", "status": "healthy"}
  - GET /ping     → pong
  - GET /health   → {"status": "healthy", "service": "incidents-api"}
  - GET /fail     → returns a 400 error (sample error path)
- Incidents endpoints
  - Base path is prefixed with API_PREFIX and versioned: {API_PREFIX}/v1
  - GET {API_PREFIX}/v1/incidents
    - Query params:
      - description: string (optional, max 50)
      - limit: number (1..100, default 10)
      - offset: number (>= 0, optional)
  - POST {API_PREFIX}/v1/incidents
    - Headers: Content-Type: application/json
    - Body:
      {
        "title": "string (required)",
        "description": "string (required)",
        "incidentType": "string (required, one of: emergency | warning | info)",
        "location": "string (required)",
        "image": "base64 string (optional)",
        "eventDate": "ISO 8601 timestamp (required)"
      }
    - Response example (list shape may include):
      {
        "incidents": [
          {
            "id": number,
            "title": string,
            "description": string,
            "incidentType": string,
            "location": string,
            "image": string,
            "eventDate": string,
            "createdAt": string
          }
        ],
        "total": number
      }
    - POST example:
      curl -X POST "http://localhost:8080/api/v1/incidents" \
           -H "Content-Type: application/json" \
           -d '{
                 "title": "Power outage",
                 "description": "Area wide blackout",
                 "incidentType": "emergency",
                 "location": "Sector 7",
                 "image": "",
                 "eventDate": "2025-09-06T20:00:00Z"
               }'


## Project structure
- Root
  - README.md
  - docker-compose.yml            → Local Postgres with seed data (scripts/db/ddl.sql)
  - Dockerfile
  - go.mod, go.sum
  - runDeploy.sh, sonar.properties (optional)
- cmd/
  - server/
    - main.go                     → App entrypoint; sets up middlewares, routes, and CQRS handlers
- internal/
  - config/
    - config.go                   → Loads .env and validates against schema (godotenv + custom parser)
  - modules/
    - incident/
      - adapters/
        - impl/                   → GORM implementation
        - models/                 → GORM models
      - commands/                 → Commands (create/update) + handlers
      - queries/                  → Queries (find) + handlers
      - http/
        - controllers/            → Gin controllers
        - dto/                    → HTTP DTOs (req/res)
        - routes/                 → Module routes
      - infrastructure/
        - repository/             → Repository abstraction
      - interfaces/               → Interfaces for controller/service
      - mappers/                  → Entity ↔ DTO mappers
      - services/                 → Domain services (uses command/query buses)
  - shared/
    - cqrs/                       → Command/Query buses and setup
    - errors/                     → API and App error helpers
    - http/
      - middlewares/              → Error handling, state, validation
      - routes/                   → Base router and not_found handler
      - types/                    → Router dependency wiring
    - instances/
      - gorm.instance.go          → GORM DB connection (Postgres)
      - pubsub.instance.go        → Pub/Sub listener scaffold (optional)
    - slog/                       → Structured logger setup
    - types/                      → Env, API, HTTP error, validation types
    - utils/                      → API/Env/Error/Service utilities (with tests)
- scripts/
  - db/
    - ddl.sql                     → Schema and demo data for local DB


## Configuration and environment
- See internal/shared/types/env.types.go for the full environment schema
  - PORT: server port (default 8080)
  - LOCAL_ONLY: when true binds to 127.0.0.1, else 0.0.0.0
  - API_PREFIX: default /api
  - GO_ENV: development|production (affects logger)
  - DB_*: connection settings
  - USE_SQL_CONNECTOR: set true to use Cloud SQL connector driver
  - DB_LOGGER: toggle GORM logging
- internal/config/config.go loads .env via godotenv (skips when testing)


## Local development notes
- Server is a plain Gin app. CQRS is wired in cmd/server/main.go.
- 404 handler returns a JSON error with a request UUID when available.


## Database
- docker-compose.yml starts a Postgres 17 container and mounts scripts/db/ for automatic init (ddl.sql runs on first startup).
- Default DB credentials are read from .env. Adjust as needed.


## Pub/Sub listener
- internal/shared/instances/pubsub.instance.go contains a scaffold to receive messages and persist incidents using the same flow.
- It is not started by default. Wire it in as needed and configure project/subscription.


## Testing
- Unit tests exist under internal/shared/utils (api, errors, service helpers)
- Run: go test ./...


## License
- MIT or as appropriate for your project.
