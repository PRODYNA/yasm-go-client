/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Page type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Page{}

// Page struct for Page
type Page struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Count *int32 `json:"count,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewPage instantiates a new Page object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPage() *Page {
	this := Page{}
	return &this
}

// NewPageWithDefaults instantiates a new Page object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageWithDefaults() *Page {
	this := Page{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *Page) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *Page) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *Page) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *Page) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Page) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *Page) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Page) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Page) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Page) SetCount(v int32) {
	o.Count = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *Page) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *Page) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *Page) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o Page) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Page) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	return toSerialize, nil
}

type NullablePage struct {
	value *Page
	isSet bool
}

func (v NullablePage) Get() *Page {
	return v.value
}

func (v *NullablePage) Set(val *Page) {
	v.value = val
	v.isSet = true
}

func (v NullablePage) IsSet() bool {
	return v.isSet
}

func (v *NullablePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePage(val *Page) *NullablePage {
	return &NullablePage{value: val, isSet: true}
}

func (v NullablePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


