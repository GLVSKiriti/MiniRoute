test:
	cd ./Backend/ && go test -v -cover ./...

frontend:
	cd ./Frontend/ && npm run dev

backend:
	cd ./Backend/ && go run main.go
