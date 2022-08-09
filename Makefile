check_install:
	which swagger || (dir=$(mktemp -d); git clone https://github.com/go-swagger/go-swagger "$dir"; cd "$dir"; go install ./cmd/swagger)

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models