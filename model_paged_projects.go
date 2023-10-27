/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedProjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedProjects{}

// PagedProjects struct for PagedProjects
type PagedProjects struct {
	Page
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
	if o == nil || IsNil(o.Projects) {
		var ret []ProjectScoreDetail
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjects) GetProjectsOk() ([]ProjectScoreDetail, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PagedProjects) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []ProjectScoreDetail and assigns it to the Projects field.
func (o *PagedProjects) SetProjects(v []ProjectScoreDetail) {
	o.Projects = v
}

func (o PagedProjects) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedProjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	return toSerialize, nil
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


