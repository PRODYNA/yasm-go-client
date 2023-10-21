/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ScoreResult struct for ScoreResult
type ScoreResult struct {
	Score *float32 `json:"score,omitempty"`
	DirectHit *bool `json:"directHit,omitempty"`
}

// NewScoreResult instantiates a new ScoreResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScoreResult() *ScoreResult {
	this := ScoreResult{}
	return &this
}

// NewScoreResultWithDefaults instantiates a new ScoreResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScoreResultWithDefaults() *ScoreResult {
	this := ScoreResult{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ScoreResult) GetScore() float32 {
	if o == nil || o.Score == nil {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScoreResult) GetScoreOk() (*float32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ScoreResult) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *ScoreResult) SetScore(v float32) {
	o.Score = &v
}

// GetDirectHit returns the DirectHit field value if set, zero value otherwise.
func (o *ScoreResult) GetDirectHit() bool {
	if o == nil || o.DirectHit == nil {
		var ret bool
		return ret
	}
	return *o.DirectHit
}

// GetDirectHitOk returns a tuple with the DirectHit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScoreResult) GetDirectHitOk() (*bool, bool) {
	if o == nil || o.DirectHit == nil {
		return nil, false
	}
	return o.DirectHit, true
}

// HasDirectHit returns a boolean if a field has been set.
func (o *ScoreResult) HasDirectHit() bool {
	if o != nil && o.DirectHit != nil {
		return true
	}

	return false
}

// SetDirectHit gets a reference to the given bool and assigns it to the DirectHit field.
func (o *ScoreResult) SetDirectHit(v bool) {
	o.DirectHit = &v
}

func (o ScoreResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.DirectHit != nil {
		toSerialize["directHit"] = o.DirectHit
	}
	return json.Marshal(toSerialize)
}

type NullableScoreResult struct {
	value *ScoreResult
	isSet bool
}

func (v NullableScoreResult) Get() *ScoreResult {
	return v.value
}

func (v *NullableScoreResult) Set(val *ScoreResult) {
	v.value = val
	v.isSet = true
}

func (v NullableScoreResult) IsSet() bool {
	return v.isSet
}

func (v *NullableScoreResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScoreResult(val *ScoreResult) *NullableScoreResult {
	return &NullableScoreResult{value: val, isSet: true}
}

func (v NullableScoreResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScoreResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


