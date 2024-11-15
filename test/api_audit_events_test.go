/*
Cloud Controller API

Testing AuditEventsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package capiclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func Test_capiclient_AuditEventsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuditEventsAPIService V3AuditEventsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuditEventsAPI.V3AuditEventsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditEventsAPIService V3AuditEventsGuidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var guid string

		resp, httpRes, err := apiClient.AuditEventsAPI.V3AuditEventsGuidGet(context.Background(), guid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
