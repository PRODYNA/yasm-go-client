/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.2.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SearchResultItem struct for SearchResultItem
type SearchResultItem struct {
	Score *float32 `json:"score,omitempty"`
	Type *string `json:"type,omitempty"`
	Item *NamedDomainModel `json:"item,omitempty"`
}

// NewSearchResultItem instantiates a new SearchResultItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResultItem() *SearchResultItem {
	this := SearchResultItem{}
	return &this
}

// NewSearchResultItemWithDefaults instantiates a new SearchResultItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultItemWithDefaults() *SearchResultItem {
	this := SearchResultItem{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *SearchResultItem) GetScore() float32 {
	if o == nil || o.Score == nil {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItem) GetScoreOk() (*float32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *SearchResultItem) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *SearchResultItem) SetScore(v float32) {
	o.Score = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SearchResultItem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SearchResultItem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SearchResultItem) SetType(v string) {
	o.Type = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *SearchResultItem) GetItem() NamedDomainModel {
	if o == nil || o.Item == nil {
		var ret NamedDomainModel
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItem) GetItemOk() (*NamedDomainModel, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *SearchResultItem) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given NamedDomainModel and assigns it to the Item field.
func (o *SearchResultItem) SetItem(v NamedDomainModel) {
	o.Item = &v
}

func (o SearchResultItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableSearchResultItem struct {
	value *SearchResultItem
	isSet bool
}

func (v NullableSearchResultItem) Get() *SearchResultItem {
	return v.value
}

func (v *NullableSearchResultItem) Set(val *SearchResultItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResultItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResultItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResultItem(val *SearchResultItem) *NullableSearchResultItem {
	return &NullableSearchResultItem{value: val, isSet: true}
}

func (v NullableSearchResultItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResultItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


