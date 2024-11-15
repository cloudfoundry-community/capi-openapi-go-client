/*
Cloud Controller API

Testing AppUsageEventsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package capiclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/capiclient"
)

func Test_capiclient_AppUsageEventsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppUsageEventsAPIService V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppUsageEventsAPI.V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppUsageEventsAPIService V3AppUsageEventsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppUsageEventsAPI.V3AppUsageEventsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppUsageEventsAPIService V3AppUsageEventsGuidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var guid string

		resp, httpRes, err := apiClient.AppUsageEventsAPI.V3AppUsageEventsGuidGet(context.Background(), guid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
