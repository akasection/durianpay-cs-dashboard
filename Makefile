.PHONY: dep-backend dep-frontend

dep-backend:
	cd backend && \
	go env -w CGO_ENABLED=1 && \
	go install

dep-frontend:
	cd frontend && \
	corepack enable && \
	pnpm install

run-backend:
	cd backend && \
	go run main.go

run-frontend:
	cd frontend && \
	pnpm dev -- --host
