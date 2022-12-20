build: frontend
	go build -o bin/pocketKit main.go

frontend:
	cd frontend; npm run build

run:
	go run main.go serve

alias m := migrtation

migrtation name:
	go run main.go migrate create {{name}}