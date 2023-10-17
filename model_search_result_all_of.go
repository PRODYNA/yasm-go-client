/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SearchResultAllOf struct for SearchResultAllOf
type SearchResultAllOf struct {
	Items []SearchResultItem `json:"items,omitempty"`
}

// NewSearchResultAllOf instantiates a new SearchResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResultAllOf() *SearchResultAllOf {
	this := SearchResultAllOf{}
	return &this
}

// NewSearchResultAllOfWithDefaults instantiates a new SearchResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultAllOfWithDefaults() *SearchResultAllOf {
	this := SearchResultAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchResultAllOf) GetItems() []SearchResultItem {
	if o == nil || o.Items == nil {
		var ret []SearchResultItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultAllOf) GetItemsOk() ([]SearchResultItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchResultAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SearchResultItem and assigns it to the Items field.
func (o *SearchResultAllOf) SetItems(v []SearchResultItem) {
	o.Items = v
}

func (o SearchResultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSearchResultAllOf struct {
	value *SearchResultAllOf
	isSet bool
}

func (v NullableSearchResultAllOf) Get() *SearchResultAllOf {
	return v.value
}

func (v *NullableSearchResultAllOf) Set(val *SearchResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResultAllOf(val *SearchResultAllOf) *NullableSearchResultAllOf {
	return &NullableSearchResultAllOf{value: val, isSet: true}
}

func (v NullableSearchResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


