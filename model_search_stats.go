/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SearchStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchStats{}

// SearchStats struct for SearchStats
type SearchStats struct {
	DirectHits *int32 `json:"directHits,omitempty"`
}

// NewSearchStats instantiates a new SearchStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchStats() *SearchStats {
	this := SearchStats{}
	return &this
}

// NewSearchStatsWithDefaults instantiates a new SearchStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchStatsWithDefaults() *SearchStats {
	this := SearchStats{}
	return &this
}

// GetDirectHits returns the DirectHits field value if set, zero value otherwise.
func (o *SearchStats) GetDirectHits() int32 {
	if o == nil || IsNil(o.DirectHits) {
		var ret int32
		return ret
	}
	return *o.DirectHits
}

// GetDirectHitsOk returns a tuple with the DirectHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStats) GetDirectHitsOk() (*int32, bool) {
	if o == nil || IsNil(o.DirectHits) {
		return nil, false
	}
	return o.DirectHits, true
}

// HasDirectHits returns a boolean if a field has been set.
func (o *SearchStats) HasDirectHits() bool {
	if o != nil && !IsNil(o.DirectHits) {
		return true
	}

	return false
}

// SetDirectHits gets a reference to the given int32 and assigns it to the DirectHits field.
func (o *SearchStats) SetDirectHits(v int32) {
	o.DirectHits = &v
}

func (o SearchStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DirectHits) {
		toSerialize["directHits"] = o.DirectHits
	}
	return toSerialize, nil
}

type NullableSearchStats struct {
	value *SearchStats
	isSet bool
}

func (v NullableSearchStats) Get() *SearchStats {
	return v.value
}

func (v *NullableSearchStats) Set(val *SearchStats) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchStats) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchStats(val *SearchStats) *NullableSearchStats {
	return &NullableSearchStats{value: val, isSet: true}
}

func (v NullableSearchStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


