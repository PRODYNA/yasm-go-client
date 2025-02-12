/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.64.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the AwardedProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwardedProject{}

// AwardedProject struct for AwardedProject
type AwardedProject struct {
	AwardId *string `json:"awardId,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
}

// NewAwardedProject instantiates a new AwardedProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwardedProject() *AwardedProject {
	this := AwardedProject{}
	return &this
}

// NewAwardedProjectWithDefaults instantiates a new AwardedProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwardedProjectWithDefaults() *AwardedProject {
	this := AwardedProject{}
	return &this
}

// GetAwardId returns the AwardId field value if set, zero value otherwise.
func (o *AwardedProject) GetAwardId() string {
	if o == nil || IsNil(o.AwardId) {
		var ret string
		return ret
	}
	return *o.AwardId
}

// GetAwardIdOk returns a tuple with the AwardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwardedProject) GetAwardIdOk() (*string, bool) {
	if o == nil || IsNil(o.AwardId) {
		return nil, false
	}
	return o.AwardId, true
}

// HasAwardId returns a boolean if a field has been set.
func (o *AwardedProject) HasAwardId() bool {
	if o != nil && !IsNil(o.AwardId) {
		return true
	}

	return false
}

// SetAwardId gets a reference to the given string and assigns it to the AwardId field.
func (o *AwardedProject) SetAwardId(v string) {
	o.AwardId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AwardedProject) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwardedProject) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AwardedProject) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AwardedProject) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o AwardedProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwardedProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwardId) {
		toSerialize["awardId"] = o.AwardId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableAwardedProject struct {
	value *AwardedProject
	isSet bool
}

func (v NullableAwardedProject) Get() *AwardedProject {
	return v.value
}

func (v *NullableAwardedProject) Set(val *AwardedProject) {
	v.value = val
	v.isSet = true
}

func (v NullableAwardedProject) IsSet() bool {
	return v.isSet
}

func (v *NullableAwardedProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwardedProject(val *AwardedProject) *NullableAwardedProject {
	return &NullableAwardedProject{value: val, isSet: true}
}

func (v NullableAwardedProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwardedProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


