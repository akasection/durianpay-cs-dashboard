# (Take Home Test) Durianpay Customer Service Dashboard
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

### Run Project (Dev)

Then, you can start the project using:

#### Backend

```bash
make run-backend
```

Backend server will be ready at `http://localhost:8081`

#### Frontend

```bash
make run-frontend
```

Nuxt server will be ready at `http://localhost:3000`

### Build Project (Prod)

```bash
make build
```
The files will be generated into:
- Backend: `./bin`
- Frontend: `./frontend/.output/public`

### Workarounds
- For Apple silicon mac, you (might) need to exec this first:

```bash
go env -w GOARCH=arm64
```

## Extra

### Testing

#### Manual (Frontend)Testing

There are two users pre-created in the database:
- username: `cs_1`, password: `1234` (`cs` role)
- username: `operator`, password: `4321` (`operation` role)

And there are some pre-created transactions in the database. Feel free to login and commit some changes.

#### New user generation

New user can be directly inserted into `users` table, with password hashed using `bcrypt` (https://bcrypt-generator.com/)

### OpenAPI Spec

Read `openapi.yaml` for the OpenAPI spec

# Contribute

Well... feel free to open PRs if you want to contribute :'D
