/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedProjects struct for PagedProjects
type PagedProjects struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Count *int32 `json:"count,omitempty"`
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

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *PagedProjects) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjects) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *PagedProjects) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *PagedProjects) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PagedProjects) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjects) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PagedProjects) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PagedProjects) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PagedProjects) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjects) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PagedProjects) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PagedProjects) SetCount(v int32) {
	o.Count = &v
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
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
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


