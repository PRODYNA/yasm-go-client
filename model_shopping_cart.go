/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.68.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ShoppingCart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShoppingCart{}

// ShoppingCart struct for ShoppingCart
type ShoppingCart struct {
	NamedDomainModel
	Description *string `json:"description,omitempty"`
}

// NewShoppingCart instantiates a new ShoppingCart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShoppingCart(id string, name string) *ShoppingCart {
	this := ShoppingCart{}
	this.Id = id
	this.Name = name
	return &this
}

// NewShoppingCartWithDefaults instantiates a new ShoppingCart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShoppingCartWithDefaults() *ShoppingCart {
	this := ShoppingCart{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ShoppingCart) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShoppingCart) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ShoppingCart) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ShoppingCart) SetDescription(v string) {
	o.Description = &v
}

func (o ShoppingCart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShoppingCart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableShoppingCart struct {
	value *ShoppingCart
	isSet bool
}

func (v NullableShoppingCart) Get() *ShoppingCart {
	return v.value
}

func (v *NullableShoppingCart) Set(val *ShoppingCart) {
	v.value = val
	v.isSet = true
}

func (v NullableShoppingCart) IsSet() bool {
	return v.isSet
}

func (v *NullableShoppingCart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShoppingCart(val *ShoppingCart) *NullableShoppingCart {
	return &NullableShoppingCart{value: val, isSet: true}
}

func (v NullableShoppingCart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShoppingCart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


