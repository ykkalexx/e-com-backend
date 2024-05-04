APP_NAME = e-com-backend

build:
	@go build -o .\bin\$(APP_NAME).exe .\cmd\main.go

test:
	@go test -v .\...

run: build
	@.\bin\$(APP_NAME).exe