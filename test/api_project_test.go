/*
YASM (Yet Another Skill Management) API

Testing ProjectAPIService

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

func Test_client_ProjectAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectAPIService AddPersonProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.AddPersonProject(context.Background(), personId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService AddPersonProjectSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var projectId string
		var skillId string

		resp, httpRes, err := apiClient.ProjectAPI.AddPersonProjectSkill(context.Background(), personId, projectId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService ConfirmSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var projectId string
		var skillId string
		var confirmingPersonId string

		resp, httpRes, err := apiClient.ProjectAPI.ConfirmSkill(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService CreateProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ProjectAPI.CreateProject(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService DeleteConfirmation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var projectId string
		var skillId string
		var confirmingPersonId string

		resp, httpRes, err := apiClient.ProjectAPI.DeleteConfirmation(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService DeletePersonProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.DeletePersonProject(context.Background(), personId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService DeletePersonProjectSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var projectId string
		var skillId string

		resp, httpRes, err := apiClient.ProjectAPI.DeletePersonProjectSkill(context.Background(), personId, projectId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService DeleteProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.DeleteProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService GetOrganizationProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ProjectAPI.GetOrganizationProjects(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService GetProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.GetProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService MergeProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var otherProjectId string

		resp, httpRes, err := apiClient.ProjectAPI.MergeProjects(context.Background(), projectId, otherProjectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService ReadPersonProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.ReadPersonProject(context.Background(), personId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService SearchProjectParticipations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectAPI.SearchProjectParticipations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService SearchProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectAPI.SearchProjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService UpdatePersonProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.UpdatePersonProject(context.Background(), personId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService UpdatePersonProjectSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var projectId string
		var skillId string

		resp, httpRes, err := apiClient.ProjectAPI.UpdatePersonProjectSkill(context.Background(), personId, projectId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService UpdateProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.UpdateProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService UpdateProjectOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.UpdateProjectOrganization(context.Background(), organizationId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
