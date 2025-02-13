/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.65.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectImageSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectImageSearch{}

// ProjectImageSearch struct for ProjectImageSearch
type ProjectImageSearch struct {
	Search
	ProjectIds []string `json:"projectIds,omitempty"`
	LayoutIds []string `json:"layoutIds,omitempty"`
}

// NewProjectImageSearch instantiates a new ProjectImageSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectImageSearch(skip int32, limit int32) *ProjectImageSearch {
	this := ProjectImageSearch{}
	this.Skip = skip
	this.Limit = limit
	return &this
}

// NewProjectImageSearchWithDefaults instantiates a new ProjectImageSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectImageSearchWithDefaults() *ProjectImageSearch {
	this := ProjectImageSearch{}
	return &this
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *ProjectImageSearch) GetProjectIds() []string {
	if o == nil || IsNil(o.ProjectIds) {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectImageSearch) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ProjectImageSearch) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *ProjectImageSearch) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetLayoutIds returns the LayoutIds field value if set, zero value otherwise.
func (o *ProjectImageSearch) GetLayoutIds() []string {
	if o == nil || IsNil(o.LayoutIds) {
		var ret []string
		return ret
	}
	return o.LayoutIds
}

// GetLayoutIdsOk returns a tuple with the LayoutIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectImageSearch) GetLayoutIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.LayoutIds) {
		return nil, false
	}
	return o.LayoutIds, true
}

// HasLayoutIds returns a boolean if a field has been set.
func (o *ProjectImageSearch) HasLayoutIds() bool {
	if o != nil && !IsNil(o.LayoutIds) {
		return true
	}

	return false
}

// SetLayoutIds gets a reference to the given []string and assigns it to the LayoutIds field.
func (o *ProjectImageSearch) SetLayoutIds(v []string) {
	o.LayoutIds = v
}

func (o ProjectImageSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectImageSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSearch, errSearch := json.Marshal(o.Search)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	errSearch = json.Unmarshal([]byte(serializedSearch), &toSerialize)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	if !IsNil(o.ProjectIds) {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if !IsNil(o.LayoutIds) {
		toSerialize["layoutIds"] = o.LayoutIds
	}
	return toSerialize, nil
}

type NullableProjectImageSearch struct {
	value *ProjectImageSearch
	isSet bool
}

func (v NullableProjectImageSearch) Get() *ProjectImageSearch {
	return v.value
}

func (v *NullableProjectImageSearch) Set(val *ProjectImageSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectImageSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectImageSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectImageSearch(val *ProjectImageSearch) *NullableProjectImageSearch {
	return &NullableProjectImageSearch{value: val, isSet: true}
}

func (v NullableProjectImageSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectImageSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


