# (Take Home Test) DurianPay Customer Service Dashboard
-----
This is a sample project based on test requiremennts

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

### Workarounds
- For Apple silicon mac, you (might) need to exec this first:

```
go env -w GOARCH=arm64
```

## Extra

### Testing

### OpenAPI Spec

### Contribution

Well... feel free to open PRs if you want to contribute :'D
