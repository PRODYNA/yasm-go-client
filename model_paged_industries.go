/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.8.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedIndustries struct for PagedIndustries
type PagedIndustries struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Industries *[]Industry `json:"industries,omitempty"`
}

// NewPagedIndustries instantiates a new PagedIndustries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedIndustries() *PagedIndustries {
	this := PagedIndustries{}
	return &this
}

// NewPagedIndustriesWithDefaults instantiates a new PagedIndustries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedIndustriesWithDefaults() *PagedIndustries {
	this := PagedIndustries{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *PagedIndustries) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedIndustries) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *PagedIndustries) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *PagedIndustries) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PagedIndustries) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedIndustries) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PagedIndustries) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PagedIndustries) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PagedIndustries) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedIndustries) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PagedIndustries) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PagedIndustries) SetCount(v int32) {
	o.Count = &v
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *PagedIndustries) GetIndustries() []Industry {
	if o == nil || o.Industries == nil {
		var ret []Industry
		return ret
	}
	return *o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedIndustries) GetIndustriesOk() (*[]Industry, bool) {
	if o == nil || o.Industries == nil {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *PagedIndustries) HasIndustries() bool {
	if o != nil && o.Industries != nil {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []Industry and assigns it to the Industries field.
func (o *PagedIndustries) SetIndustries(v []Industry) {
	o.Industries = &v
}

func (o PagedIndustries) MarshalJSON() ([]byte, error) {
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
	if o.Industries != nil {
		toSerialize["industries"] = o.Industries
	}
	return json.Marshal(toSerialize)
}

type NullablePagedIndustries struct {
	value *PagedIndustries
	isSet bool
}

func (v NullablePagedIndustries) Get() *PagedIndustries {
	return v.value
}

func (v *NullablePagedIndustries) Set(val *PagedIndustries) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedIndustries) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedIndustries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedIndustries(val *PagedIndustries) *NullablePagedIndustries {
	return &NullablePagedIndustries{value: val, isSet: true}
}

func (v NullablePagedIndustries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedIndustries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


