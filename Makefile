generate-swagger:
	swag init -g cmd/app/main.go -o docs/api/
	
run: generate-swagger
	go run cmd/app/main.go

build:
	go build -o app cmd/app/main.go