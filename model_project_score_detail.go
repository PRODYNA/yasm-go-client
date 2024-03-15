/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.16.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectScoreDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectScoreDetail{}

// ProjectScoreDetail struct for ProjectScoreDetail
type ProjectScoreDetail struct {
	ProjectDetails
	Score *float32 `json:"score,omitempty"`
	DirectHit *bool `json:"directHit,omitempty"`
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
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectScoreDetail) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ProjectScoreDetail) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *ProjectScoreDetail) SetScore(v float32) {
	o.Score = &v
}

// GetDirectHit returns the DirectHit field value if set, zero value otherwise.
func (o *ProjectScoreDetail) GetDirectHit() bool {
	if o == nil || IsNil(o.DirectHit) {
		var ret bool
		return ret
	}
	return *o.DirectHit
}

// GetDirectHitOk returns a tuple with the DirectHit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectScoreDetail) GetDirectHitOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectHit) {
		return nil, false
	}
	return o.DirectHit, true
}

// HasDirectHit returns a boolean if a field has been set.
func (o *ProjectScoreDetail) HasDirectHit() bool {
	if o != nil && !IsNil(o.DirectHit) {
		return true
	}

	return false
}

// SetDirectHit gets a reference to the given bool and assigns it to the DirectHit field.
func (o *ProjectScoreDetail) SetDirectHit(v bool) {
	o.DirectHit = &v
}

func (o ProjectScoreDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectScoreDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedProjectDetails, errProjectDetails := json.Marshal(o.ProjectDetails)
	if errProjectDetails != nil {
		return map[string]interface{}{}, errProjectDetails
	}
	errProjectDetails = json.Unmarshal([]byte(serializedProjectDetails), &toSerialize)
	if errProjectDetails != nil {
		return map[string]interface{}{}, errProjectDetails
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.DirectHit) {
		toSerialize["directHit"] = o.DirectHit
	}
	return toSerialize, nil
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


