package main

import (
	"fmt"
	"github.com/cloudfoundry-community/capi-openapi-go-client/v3/capiclient"
)

func main() {
	// Test that we can import and reference the capiclient package
	fmt.Printf("capiclient package imported successfully\n")
	
	// Try to reference a type from the package (Client is commonly available)
	var _ *capiclient.Client
	fmt.Printf("capiclient.Client type is accessible\n")
}