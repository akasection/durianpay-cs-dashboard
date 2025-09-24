# (Take Home Test) DurianPay Customer Service Dashboard
-----
This is a sample project based on interview requiremennts

## Architecture

This project consists of three parts:
1. Backend: Serves RESTful APIs and internal auth system
2. Frontend: Serves the web application and application auth
3. Infra: Docker Compose to orchestrate the runtime

### Tech Stack

- Backend: Go(1.25)+Gin, GORM+SQLite
- Frontend: Vue(3), Nuxt(3/4), TypeScript, Tailwind CSS
- Infra: Docker Compose

## Starting the project

### Dependencies

Using `Makefile:

```
make dep
```
This will instlall all dependencies for both backend and frontend.

### Build Project (Dev)

Then, you can start the project using:

#### Backend

```bash
make run-backend
```

Server will be ready at `http://localhost:8081`

#### Frontend

```bash
make run-frontend
```

Server will be ready at `http://localhost:3000`

### Workarounds
- For Apple silicon mac, you (might) need to exec this first:

```
go env -w GOARCH=arm64
```

## Extra

### Testing

TODO

### OpenAPI Spec

Read `openapi.yaml` for the OpenAPI spec

# Contribute

Well... feel free to open PRs if you want to contribute :'D
