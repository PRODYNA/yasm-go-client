/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedAvailabilitiesAllOf struct for PagedAvailabilitiesAllOf
type PagedAvailabilitiesAllOf struct {
	Availabilities []AvailabilityDetail `json:"availabilities,omitempty"`
}

// NewPagedAvailabilitiesAllOf instantiates a new PagedAvailabilitiesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedAvailabilitiesAllOf() *PagedAvailabilitiesAllOf {
	this := PagedAvailabilitiesAllOf{}
	return &this
}

// NewPagedAvailabilitiesAllOfWithDefaults instantiates a new PagedAvailabilitiesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedAvailabilitiesAllOfWithDefaults() *PagedAvailabilitiesAllOf {
	this := PagedAvailabilitiesAllOf{}
	return &this
}

// GetAvailabilities returns the Availabilities field value if set, zero value otherwise.
func (o *PagedAvailabilitiesAllOf) GetAvailabilities() []AvailabilityDetail {
	if o == nil || o.Availabilities == nil {
		var ret []AvailabilityDetail
		return ret
	}
	return o.Availabilities
}

// GetAvailabilitiesOk returns a tuple with the Availabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedAvailabilitiesAllOf) GetAvailabilitiesOk() ([]AvailabilityDetail, bool) {
	if o == nil || o.Availabilities == nil {
		return nil, false
	}
	return o.Availabilities, true
}

// HasAvailabilities returns a boolean if a field has been set.
func (o *PagedAvailabilitiesAllOf) HasAvailabilities() bool {
	if o != nil && o.Availabilities != nil {
		return true
	}

	return false
}

// SetAvailabilities gets a reference to the given []AvailabilityDetail and assigns it to the Availabilities field.
func (o *PagedAvailabilitiesAllOf) SetAvailabilities(v []AvailabilityDetail) {
	o.Availabilities = v
}

func (o PagedAvailabilitiesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Availabilities != nil {
		toSerialize["availabilities"] = o.Availabilities
	}
	return json.Marshal(toSerialize)
}

type NullablePagedAvailabilitiesAllOf struct {
	value *PagedAvailabilitiesAllOf
	isSet bool
}

func (v NullablePagedAvailabilitiesAllOf) Get() *PagedAvailabilitiesAllOf {
	return v.value
}

func (v *NullablePagedAvailabilitiesAllOf) Set(val *PagedAvailabilitiesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedAvailabilitiesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedAvailabilitiesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedAvailabilitiesAllOf(val *PagedAvailabilitiesAllOf) *NullablePagedAvailabilitiesAllOf {
	return &NullablePagedAvailabilitiesAllOf{value: val, isSet: true}
}

func (v NullablePagedAvailabilitiesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedAvailabilitiesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


