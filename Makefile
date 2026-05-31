certs:
	@mkdir -p tls
	@cd tls && go run $(shell go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
