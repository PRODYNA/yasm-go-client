/*
YASM (Yet Another Skill Management) API

Testing IndustryApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_client_IndustryApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndustryApiService AttachOrganizationToIndustry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var industryId string

        resp, httpRes, err := apiClient.IndustryApi.AttachOrganizationToIndustry(context.Background(), organizationId, industryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndustryApiService CreateIndustry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.IndustryApi.CreateIndustry(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndustryApiService DeleteIndustry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var industryId string

        resp, httpRes, err := apiClient.IndustryApi.DeleteIndustry(context.Background(), industryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndustryApiService DetachOrganizationFromIndustry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var industryId string

        resp, httpRes, err := apiClient.IndustryApi.DetachOrganizationFromIndustry(context.Background(), organizationId, industryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndustryApiService GetIndustries", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.IndustryApi.GetIndustries(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndustryApiService GetIndustry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var industryId string

        resp, httpRes, err := apiClient.IndustryApi.GetIndustry(context.Background(), industryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndustryApiService UpdateIndustry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var industryId string

        resp, httpRes, err := apiClient.IndustryApi.UpdateIndustry(context.Background(), industryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
