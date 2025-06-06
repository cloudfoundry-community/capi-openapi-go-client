# CAPI OpenAPI Go Client

This repository contains the auto-generated Go client for the Cloud Foundry CAPI (Cloud Controller API) v3.

## Version

This client is generated from CAPI version **3.195.0**.

## Installation

```bash
go get github.com/cloudfoundry-community/capi-openapi-go-client/capiclient/v3@v3.195.0
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    
    "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient/v3"
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

## Generation

This client is automatically generated from the OpenAPI specification using [oapi-codegen](https://github.com/deepmap/oapi-codegen).

To regenerate this client:
1. Clone [capi-openapi-spec](https://github.com/cloudfoundry-community/capi-openapi-spec)
2. Run `make gen-go-client VERSION=3.195.0`
3. Run `./bin/publish --version=3.195.0`

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
