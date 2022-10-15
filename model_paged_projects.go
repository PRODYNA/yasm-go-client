/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedProjects struct for PagedProjects
type PagedProjects struct {
	Projects []ProjectScoreDetail `json:"projects,omitempty"`
}

// NewPagedProjects instantiates a new PagedProjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedProjects() *PagedProjects {
	this := PagedProjects{}
	return &this
}

// NewPagedProjectsWithDefaults instantiates a new PagedProjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedProjectsWithDefaults() *PagedProjects {
	this := PagedProjects{}
	return &this
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PagedProjects) GetProjects() []ProjectScoreDetail {
	if o == nil || o.Projects == nil {
		var ret []ProjectScoreDetail
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjects) GetProjectsOk() ([]ProjectScoreDetail, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PagedProjects) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []ProjectScoreDetail and assigns it to the Projects field.
func (o *PagedProjects) SetProjects(v []ProjectScoreDetail) {
	o.Projects = v
}

func (o PagedProjects) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return json.Marshal(toSerialize)
}

type NullablePagedProjects struct {
	value *PagedProjects
	isSet bool
}

func (v NullablePagedProjects) Get() *PagedProjects {
	return v.value
}

func (v *NullablePagedProjects) Set(val *PagedProjects) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedProjects) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedProjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedProjects(val *PagedProjects) *NullablePagedProjects {
	return &NullablePagedProjects{value: val, isSet: true}
}

func (v NullablePagedProjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedProjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


