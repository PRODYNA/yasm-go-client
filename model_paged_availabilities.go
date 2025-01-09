/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedAvailabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedAvailabilities{}

// PagedAvailabilities struct for PagedAvailabilities
type PagedAvailabilities struct {
	Page
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

// GetAvailabilities returns the Availabilities field value if set, zero value otherwise.
func (o *PagedAvailabilities) GetAvailabilities() []AvailabilityDetail {
	if o == nil || IsNil(o.Availabilities) {
		var ret []AvailabilityDetail
		return ret
	}
	return o.Availabilities
}

// GetAvailabilitiesOk returns a tuple with the Availabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedAvailabilities) GetAvailabilitiesOk() ([]AvailabilityDetail, bool) {
	if o == nil || IsNil(o.Availabilities) {
		return nil, false
	}
	return o.Availabilities, true
}

// HasAvailabilities returns a boolean if a field has been set.
func (o *PagedAvailabilities) HasAvailabilities() bool {
	if o != nil && !IsNil(o.Availabilities) {
		return true
	}

	return false
}

// SetAvailabilities gets a reference to the given []AvailabilityDetail and assigns it to the Availabilities field.
func (o *PagedAvailabilities) SetAvailabilities(v []AvailabilityDetail) {
	o.Availabilities = v
}

func (o PagedAvailabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedAvailabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.Availabilities) {
		toSerialize["availabilities"] = o.Availabilities
	}
	return toSerialize, nil
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


