.PHONY: dep build

dep-backend:
	cd backend && \
	go env -w CGO_ENABLED=1 && \
go install

dep-frontend:
	cd frontend && \
	corepack enable && \
	pnpm install

dep: dep-backend dep-frontend

run-backend:
	cd backend && \
	go run main.go

run-frontend:
	cd frontend && \
	pnpm dev --host

build-backend:
	cd backend && \
	go build -o ../bin/backend

build-frontend:
	cd frontend && \
	pnpm build

build: build-backend build-frontend
