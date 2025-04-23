
# Go-Vue Full-Stack Starter

A production-ready boilerplate demonstrating a **Go**-powered RESTful backend with a **Vue 3** Single-Page Application front end. Features:

- **CRUD API** with Gin & GORM  
- **Vue 3 + Vue Router + Vuex** (or Pinia) front-end  
- **Axios** for HTTP calls  
- Environment-driven configuration  
- CORS, logging, graceful shutdown  
- Docker + Docker Compose orchestration  

---

## Table of Contents

1. [Overview](#overview)  
2. [Features](#features)  
3. [Tech Stack](#tech-stack)  
4. [Getting Started](#getting-started)  
   - [Prerequisites](#prerequisites)  
   - [Installation](#installation)  
   - [Configuration](#configuration)  
5. [Project Structure](#project-structure)  
6. [Backend API](#backend-api)  
7. [Frontend Usage](#frontend-usage)  
8. [Docker & Deployment](#docker--deployment)  
9. [Testing](#testing)  
10. [Contributing](#contributing)  
11. [License](#license)  

---

## Overview

`Go-Vue` is a template for building modern web applications. The backend (in `app/server`) exposes a JSON API; the frontend (in `app/client_tab`) consumes it as a Vue SPA. Everything can run locally or in Docker.

---

## Features

- ✅ **CRUD endpoints** for a sample `Item` or `User` entity  
- 🔒 **CORS** and **Graceful Shutdown**  
- 📦 **GORM** for database ORM (PostgreSQL or SQLite)  
- 📝 **Structured Logging** with Zap  
- ⚡️ **Vue 3** SPA with Vue Router & Vuex/Pinia  
- 🌐 **Axios**-powered API client  
- 🐳 **Docker Compose** for local orchestration  

---

## Tech Stack

| Layer       | Technology                             |
| ----------- | -------------------------------------- |
| Backend     | Go 1.18+, [GIN](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/) |
| DB          | PostgreSQL (Prod) / SQLite (Dev)       |
| Frontend    | Vue 3, Vue Router, Vuex/Pinia, Axios   |
| DevOps      | Docker, Docker Compose, Make           |

---

## Getting Started

### Prerequisites

- [Go 1.18+](https://golang.org/dl/)  
- [Node.js 16+ & npm](https://nodejs.org/)  
- (Optionally) Docker & Docker Compose  

---

### Installation

1. **Clone the repo**  
   ```bash
   git clone https://github.com/nicktretyakov/Go-Vue.git
   cd Go-Vue
   ```

2. **Configure environment**  
   Copy the example env file and adjust:
   ```bash
   cp .env.example .env
   ```
   Edit `.env` to set:
   ```ini
   # Server
   SERVER_PORT=8080

   # Database (Postgres)
   DB_DRIVER=postgres
   DB_DSN=host=localhost port=5432 user=postgres password=secret dbname=go_vue sslmode=disable

   # OR for SQLite (dev)
   # DB_DRIVER=sqlite
   # DB_DSN=./go_vue.db

   # CORS
   CORS_ORIGINS=http://localhost:3000
   ```

3. **Run the Backend**  
   ```bash
   cd app/server
   go mod download
   go run main.go
   ```
   By default it listens on `:8080`.

4. **Run the Frontend**  
   ```bash
   cd ../client_tab
   npm install
   npm run serve
   ```
   This starts a development server on `http://localhost:3000` with hot-reload.

---

## Project Structure

```
Go-Vue/
├── .env.example            # Sample environment settings
├── docker-compose.yml      # Compose file for backend + db + frontend
├── app/
│   ├── server/             # Go/Gin backend
│   │   ├── main.go         # Entrypoint
│   │   ├── router.go       # Route definitions & handlers
│   │   ├── models/         # GORM model definitions
│   │   ├── controllers/    # Handler functions
│   │   ├── middleware/     # CORS, logging, error handling
│   │   └── utils/          # Config loader, DB init, logger
│   └── client_tab/         # Vue 3 front-end
│       ├── package.json
│       ├── public/         # index.html, favicon
│       ├── src/
│       │   ├── main.js     # Vue app bootstrap
│       │   ├── App.vue
│       │   ├── router.js   # Vue Router setup
│       │   ├── store.js    # Vuex/Pinia store
│       │   └── components/ # Reusable Vue components
│       └── vite.config.js  # Vite build config (if using Vite)
└── README.md
```

---

## Backend API

Base URL: `http://localhost:8080/api/items`

| Method | Path         | Description            |
| ------ | ------------ | ---------------------- |
| GET    | `/api/items` | List all items         |
| GET    | `/api/items/:id` | Get item by ID     |
| POST   | `/api/items` | Create a new item      |
| PUT    | `/api/items/:id` | Update item by ID  |
| DELETE | `/api/items/:id` | Delete item by ID  |

_All responses are JSON with standard HTTP status codes._

---

## Frontend Usage

- The SPA runs at `http://localhost:3000`.  
- It calls the above API via Axios in `src/api.js`.  
- You can view, add, edit, and delete items in a dynamic Vue UI.

---

## Docker & Deployment

A `docker-compose.yml` is provided to stand up:

- **Postgres** DB  
- **Go** backend  
- **Vue** frontend (served by Nginx in production mode)

```bash
docker-compose up --build -d
```

- Backend will be on port **8080**.  
- Frontend on port **80** (or remapped as you prefer).

---

## Testing

### Backend

```bash
cd app/server
go test ./... -cover
```

### Frontend

```bash
cd app/client_tab
npm test
```

---

## Contributing

1. Fork the repository  
2. Create a feature branch (`git checkout -b feature/XYZ`)  
3. Commit your changes & tests  
4. Push to your fork & open a PR  

Please follow Go and Vue style conventions.

---

## License

Released under the **MIT License**. See [LICENSE](LICENSE) for details.
```
