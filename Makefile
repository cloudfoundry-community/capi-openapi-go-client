.PHONY: validate
validate:
	@echo "Validating that the module can be imported..."
	cd tests/import-test && go mod tidy && go build -o /dev/null .
	@echo "✓ Import validation passed"

.PHONY: clean
clean:
	rm -f tests/import-test/go.sum
	go clean -modcache

.PHONY: test-local
test-local:
	@echo "Testing with local module replacement..."
	cd tests/import-test && go mod edit -replace github.com/cloudfoundry-community/capi-openapi-go-client/v3=../.. && go mod tidy && go build -o /dev/null .
	@echo "✓ Local validation passed"
	cd tests/import-test && go mod edit -dropreplace github.com/cloudfoundry-community/capi-openapi-go-client/v3