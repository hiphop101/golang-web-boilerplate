dev:
	go build httpd/main.go && "./main.exe"

test:
	go test -cover ./...