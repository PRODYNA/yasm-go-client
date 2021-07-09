/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.7.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectDetails struct for ProjectDetails
type ProjectDetails struct {
	Project *Project `json:"project,omitempty"`
}

// NewProjectDetails instantiates a new ProjectDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDetails() *ProjectDetails {
	this := ProjectDetails{}
	return &this
}

// NewProjectDetailsWithDefaults instantiates a new ProjectDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDetailsWithDefaults() *ProjectDetails {
	this := ProjectDetails{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectDetails) GetProject() Project {
	if o == nil || o.Project == nil {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetProjectOk() (*Project, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectDetails) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *ProjectDetails) SetProject(v Project) {
	o.Project = &v
}

func (o ProjectDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableProjectDetails struct {
	value *ProjectDetails
	isSet bool
}

func (v NullableProjectDetails) Get() *ProjectDetails {
	return v.value
}

func (v *NullableProjectDetails) Set(val *ProjectDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDetails(val *ProjectDetails) *NullableProjectDetails {
	return &NullableProjectDetails{value: val, isSet: true}
}

func (v NullableProjectDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


