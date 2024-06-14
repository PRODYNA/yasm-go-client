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

	t.Run("Test ProjectAPIService AddExecutiveOrganizationToProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var organizationId string

		resp, httpRes, err := apiClient.ProjectAPI.AddExecutiveOrganizationToProject(context.Background(), projectId, organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService AddProjectParticipation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectAPI.AddProjectParticipation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService AddSkillConfirmation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string
		var skillId string
		var confirmingPersonId string

		resp, httpRes, err := apiClient.ProjectAPI.AddSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService CreateProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectAPI.CreateProject(context.Background()).Execute()

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

	t.Run("Test ProjectAPIService DeleteProjectParticipation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string

		resp, httpRes, err := apiClient.ProjectAPI.DeleteProjectParticipation(context.Background(), projectParticipationId).Execute()

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

	t.Run("Test ProjectAPIService MoveProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var organizationId string

		resp, httpRes, err := apiClient.ProjectAPI.MoveProject(context.Background(), projectId, organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService ReadProjectParticipation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string

		resp, httpRes, err := apiClient.ProjectAPI.ReadProjectParticipation(context.Background(), projectParticipationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService RemoveExecutiveOrganizationFromProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var organizationId string

		resp, httpRes, err := apiClient.ProjectAPI.RemoveExecutiveOrganizationFromProject(context.Background(), projectId, organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService RemoveSkillConfirmation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string
		var skillId string
		var confirmingPersonId string

		resp, httpRes, err := apiClient.ProjectAPI.RemoveSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()

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

	t.Run("Test ProjectAPIService SearchProjectsStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectAPI.SearchProjectsStats(context.Background()).Execute()

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

	t.Run("Test ProjectAPIService UpdateProjectParticipation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string

		resp, httpRes, err := apiClient.ProjectAPI.UpdateProjectParticipation(context.Background(), projectParticipationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
