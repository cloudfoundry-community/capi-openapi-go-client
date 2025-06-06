# CAPI OpenAPI Go Client

This repository contains the auto-generated Go client for the Cloud Foundry CAPI (Cloud Controller API) v3.

## Version

This client is generated from CAPI version **3.195.0**.

## Installation

```bash
go get github.com/cloudfoundry-community/capi-openapi-go-client/v3@v3.195.1
```

Or add this to your `go.mod` file:

```go
require github.com/cloudfoundry-community/capi-openapi-go-client/v3 v3.195.1
```

## Module Structure

This module uses Go's versioned module structure:
- Module path: `github.com/cloudfoundry-community/capi-openapi-go-client/v3`
- Package: `capiclient` (located in the `capiclient/` subdirectory)

## Usage

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    
    "github.com/cloudfoundry-community/capi-openapi-go-client/v3/capiclient"
)

func main() {
    // Create a new client
    client, err := capiclient.NewClientWithResponses("https://api.example.com")
    if err != nil {
        panic(err)
    }
    
    // Use the client to interact with CAPI
    ctx := context.Background()
    
    // Example: List organizations
    resp, err := client.V3OrganizationsGetWithResponse(ctx, &capiclient.V3OrganizationsGetParams{})
    if err != nil {
        panic(err)
    }
    
    if resp.StatusCode() == http.StatusOK {
        fmt.Printf("Found %d organizations\n", len(*resp.JSON200.Resources))
    }
}
```

## Authentication

The client supports various authentication methods. Here's an example using Bearer token authentication:

```go
import (
    "context"
    "net/http"
)

// Create a custom HTTP client with authentication
httpClient := &http.Client{
    Transport: &authTransport{
        Token: "your-bearer-token",
        Base:  http.DefaultTransport,
    },
}

// Use the custom HTTP client
client, err := capiclient.NewClientWithResponses(
    "https://api.example.com",
    capiclient.WithHTTPClient(httpClient),
)
```

Where `authTransport` is a custom RoundTripper that adds the Authorization header:

```go
type authTransport struct {
    Token string
    Base  http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    req.Header.Set("Authorization", "Bearer "+t.Token)
    return t.Base.RoundTrip(req)
}
```

## Documentation

- [CAPI Documentation](https://v3-apidocs.cloudfoundry.org/version/3.195.0/)
- [OpenAPI Specification Source](https://github.com/cloudfoundry-community/capi-openapi-spec)
- For detailed API documentation, see the generated code in the [capiclient directory](capiclient/)

## Development and Validation

### Testing Import Compatibility

This repository includes a validation test to ensure the module can be properly imported by other Go projects:

```bash
# Validate that the module can be imported
make validate

# Test with local changes (before pushing)
make test-local
```

### Project Structure

```
.
├── capiclient/           # The actual capiclient package
│   └── client.go        # Generated client code
├── tests/               # Import validation tests
│   └── import-test/     # Test module that imports capiclient
├── go.mod               # Module definition for v3
├── Makefile             # Build and validation commands
└── README.md            # This file
```

### Validation Process

The validation process ensures:
1. The module structure is correct for Go's versioned modules
2. The package can be imported using the expected import path
3. The types and functions are accessible to importing code

This is particularly important for v2+ Go modules which require specific directory structures.

## Generation

This client is automatically generated from the OpenAPI specification using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen).

To regenerate this client:
1. Clone [capi-openapi-spec](https://github.com/cloudfoundry-community/capi-openapi-spec)
2. Run `make gen-go-client VERSION=3.195.0`
3. Run `./bin/publish --version=3.195.0`

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! This client is auto-generated from the Cloud Foundry OpenAPI specification. To update the client:

1. Update the OpenAPI specification in the [capi-openapi-spec repository](https://github.com/cloudfoundry-community/capi-openapi-spec)
2. Regenerate the client using the generation process described above
3. Validate the changes using `make test-local`
4. Submit a pull request with your changes

### Release Process

When creating a new release:

1. Ensure all changes are committed and pushed
2. Run `make validate` to ensure the module works correctly
3. Tag the release with the appropriate version (e.g., `v3.196.0`)
4. Push the tag to GitHub
5. The Go module proxy will automatically pick up the new version

## Troubleshooting

### Import Errors

If you encounter import errors like "module found but does not contain package":

1. Ensure you're using the correct import path: `github.com/cloudfoundry-community/capi-openapi-go-client/v3/capiclient`
2. Check that your Go version supports modules (Go 1.11+)
3. Try clearing your module cache: `go clean -modcache`
4. Ensure you're using the latest version of the module

### Version Issues

For v2+ modules in Go:
- The module path must include the major version suffix (e.g., `/v3`)
- The actual code must be properly organized (package in subdirectory)
- Tags must match the module version (e.g., `v3.195.1`)

## Support

For issues and questions, please open an issue in the [GitHub repository](https://github.com/cloudfoundry-community/capi-openapi-go-client).