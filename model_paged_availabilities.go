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

// PagedAvailabilities struct for PagedAvailabilities
type PagedAvailabilities struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Availabilities []AvailabilityDetail `json:"availabilities,omitempty"`
}

// NewPagedAvailabilities instantiates a new PagedAvailabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedAvailabilities() *PagedAvailabilities {
	this := PagedAvailabilities{}
	return &this
}

// NewPagedAvailabilitiesWithDefaults instantiates a new PagedAvailabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedAvailabilitiesWithDefaults() *PagedAvailabilities {
	this := PagedAvailabilities{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *PagedAvailabilities) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedAvailabilities) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *PagedAvailabilities) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *PagedAvailabilities) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PagedAvailabilities) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedAvailabilities) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PagedAvailabilities) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PagedAvailabilities) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PagedAvailabilities) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedAvailabilities) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PagedAvailabilities) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PagedAvailabilities) SetCount(v int32) {
	o.Count = &v
}

// GetAvailabilities returns the Availabilities field value if set, zero value otherwise.
func (o *PagedAvailabilities) GetAvailabilities() []AvailabilityDetail {
	if o == nil || o.Availabilities == nil {
		var ret []AvailabilityDetail
		return ret
	}
	return o.Availabilities
}

// GetAvailabilitiesOk returns a tuple with the Availabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedAvailabilities) GetAvailabilitiesOk() ([]AvailabilityDetail, bool) {
	if o == nil || o.Availabilities == nil {
		return nil, false
	}
	return o.Availabilities, true
}

// HasAvailabilities returns a boolean if a field has been set.
func (o *PagedAvailabilities) HasAvailabilities() bool {
	if o != nil && o.Availabilities != nil {
		return true
	}

	return false
}

// SetAvailabilities gets a reference to the given []AvailabilityDetail and assigns it to the Availabilities field.
func (o *PagedAvailabilities) SetAvailabilities(v []AvailabilityDetail) {
	o.Availabilities = v
}

func (o PagedAvailabilities) MarshalJSON() ([]byte, error) {
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
	if o.Availabilities != nil {
		toSerialize["availabilities"] = o.Availabilities
	}
	return json.Marshal(toSerialize)
}

type NullablePagedAvailabilities struct {
	value *PagedAvailabilities
	isSet bool
}

func (v NullablePagedAvailabilities) Get() *PagedAvailabilities {
	return v.value
}

func (v *NullablePagedAvailabilities) Set(val *PagedAvailabilities) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedAvailabilities) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedAvailabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedAvailabilities(val *PagedAvailabilities) *NullablePagedAvailabilities {
	return &NullablePagedAvailabilities{value: val, isSet: true}
}

func (v NullablePagedAvailabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedAvailabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


