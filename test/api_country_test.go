/*
YASM (Yet Another Skill Management) API

Testing CountryAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func Test_client_CountryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CountryAPIService AddLanguageToCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var countryId string
		var languageId string

		resp, httpRes, err := apiClient.CountryAPI.AddLanguageToCountry(context.Background(), countryId, languageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CountryAPIService CreateCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CountryAPI.CreateCountry(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CountryAPIService DeleteCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var countryId string

		resp, httpRes, err := apiClient.CountryAPI.DeleteCountry(context.Background(), countryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CountryAPIService GetCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var countryId string

		resp, httpRes, err := apiClient.CountryAPI.GetCountry(context.Background(), countryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CountryAPIService RemoveLanguageFromCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var countryId string
		var languageId string

		resp, httpRes, err := apiClient.CountryAPI.RemoveLanguageFromCountry(context.Background(), countryId, languageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CountryAPIService SearchCountries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CountryAPI.SearchCountries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CountryAPIService UpdateCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var countryId string

		resp, httpRes, err := apiClient.CountryAPI.UpdateCountry(context.Background(), countryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CountryAPIService UpdateOfficeCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId string
		var countryId string

		resp, httpRes, err := apiClient.CountryAPI.UpdateOfficeCountry(context.Background(), officeId, countryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
