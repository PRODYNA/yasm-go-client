/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Score type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Score{}

// Score struct for Score
type Score struct {
	Score *float32 `json:"score,omitempty"`
	// Specifies the type of score
	ScoreType *string `json:"scoreType,omitempty"`
}

// NewScore instantiates a new Score object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScore() *Score {
	this := Score{}
	var scoreType string = "skillScore"
	this.ScoreType = &scoreType
	return &this
}

// NewScoreWithDefaults instantiates a new Score object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScoreWithDefaults() *Score {
	this := Score{}
	var scoreType string = "skillScore"
	this.ScoreType = &scoreType
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *Score) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Score) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *Score) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *Score) SetScore(v float32) {
	o.Score = &v
}

// GetScoreType returns the ScoreType field value if set, zero value otherwise.
func (o *Score) GetScoreType() string {
	if o == nil || IsNil(o.ScoreType) {
		var ret string
		return ret
	}
	return *o.ScoreType
}

// GetScoreTypeOk returns a tuple with the ScoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Score) GetScoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScoreType) {
		return nil, false
	}
	return o.ScoreType, true
}

// HasScoreType returns a boolean if a field has been set.
func (o *Score) HasScoreType() bool {
	if o != nil && !IsNil(o.ScoreType) {
		return true
	}

	return false
}

// SetScoreType gets a reference to the given string and assigns it to the ScoreType field.
func (o *Score) SetScoreType(v string) {
	o.ScoreType = &v
}

func (o Score) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Score) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.ScoreType) {
		toSerialize["scoreType"] = o.ScoreType
	}
	return toSerialize, nil
}

type NullableScore struct {
	value *Score
	isSet bool
}

func (v NullableScore) Get() *Score {
	return v.value
}

func (v *NullableScore) Set(val *Score) {
	v.value = val
	v.isSet = true
}

func (v NullableScore) IsSet() bool {
	return v.isSet
}

func (v *NullableScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScore(val *Score) *NullableScore {
	return &NullableScore{value: val, isSet: true}
}

func (v NullableScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


