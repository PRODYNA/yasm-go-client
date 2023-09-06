/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectScoreDetail struct for ProjectScoreDetail
type ProjectScoreDetail struct {
	ProjectDetails
	Score *float32 `json:"score,omitempty"`
}

// NewProjectScoreDetail instantiates a new ProjectScoreDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectScoreDetail() *ProjectScoreDetail {
	this := ProjectScoreDetail{}
	return &this
}

// NewProjectScoreDetailWithDefaults instantiates a new ProjectScoreDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectScoreDetailWithDefaults() *ProjectScoreDetail {
	this := ProjectScoreDetail{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ProjectScoreDetail) GetScore() float32 {
	if o == nil || o.Score == nil {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectScoreDetail) GetScoreOk() (*float32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ProjectScoreDetail) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *ProjectScoreDetail) SetScore(v float32) {
	o.Score = &v
}

func (o ProjectScoreDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedProjectDetails, errProjectDetails := json.Marshal(o.ProjectDetails)
	if errProjectDetails != nil {
		return []byte{}, errProjectDetails
	}
	errProjectDetails = json.Unmarshal([]byte(serializedProjectDetails), &toSerialize)
	if errProjectDetails != nil {
		return []byte{}, errProjectDetails
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableProjectScoreDetail struct {
	value *ProjectScoreDetail
	isSet bool
}

func (v NullableProjectScoreDetail) Get() *ProjectScoreDetail {
	return v.value
}

func (v *NullableProjectScoreDetail) Set(val *ProjectScoreDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectScoreDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectScoreDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectScoreDetail(val *ProjectScoreDetail) *NullableProjectScoreDetail {
	return &NullableProjectScoreDetail{value: val, isSet: true}
}

func (v NullableProjectScoreDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectScoreDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


