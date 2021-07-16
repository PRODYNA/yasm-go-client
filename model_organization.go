/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.7.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Organization struct for Organization
type Organization struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Suggestion bool `json:"suggestion"`
	Synonyms *[]string `json:"synonyms,omitempty"`
	Partner *bool `json:"partner,omitempty"`
	Customer *bool `json:"customer,omitempty"`
	Geolocation *Geolocation `json:"geolocation,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(id string, name string, suggestion bool) *Organization {
	this := Organization{}
	this.Id = id
	this.Name = name
	this.Suggestion = suggestion
	var partner bool = false
	this.Partner = &partner
	var customer bool = false
	this.Customer = &customer
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	var suggestion bool = false
	this.Suggestion = suggestion
	var partner bool = false
	this.Partner = &partner
	var customer bool = false
	this.Customer = &customer
	return &this
}

// GetId returns the Id field value
func (o *Organization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Organization) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Organization) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Organization) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Organization) SetName(v string) {
	o.Name = v
}

// GetSuggestion returns the Suggestion field value
func (o *Organization) GetSuggestion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value
// and a boolean to check if the value has been set.
func (o *Organization) GetSuggestionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Suggestion, true
}

// SetSuggestion sets field value
func (o *Organization) SetSuggestion(v bool) {
	o.Suggestion = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Organization) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return *o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetSynonymsOk() (*[]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Organization) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Organization) SetSynonyms(v []string) {
	o.Synonyms = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *Organization) GetPartner() bool {
	if o == nil || o.Partner == nil {
		var ret bool
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPartnerOk() (*bool, bool) {
	if o == nil || o.Partner == nil {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *Organization) HasPartner() bool {
	if o != nil && o.Partner != nil {
		return true
	}

	return false
}

// SetPartner gets a reference to the given bool and assigns it to the Partner field.
func (o *Organization) SetPartner(v bool) {
	o.Partner = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *Organization) GetCustomer() bool {
	if o == nil || o.Customer == nil {
		var ret bool
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCustomerOk() (*bool, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *Organization) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given bool and assigns it to the Customer field.
func (o *Organization) SetCustomer(v bool) {
	o.Customer = &v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *Organization) GetGeolocation() Geolocation {
	if o == nil || o.Geolocation == nil {
		var ret Geolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetGeolocationOk() (*Geolocation, bool) {
	if o == nil || o.Geolocation == nil {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *Organization) HasGeolocation() bool {
	if o != nil && o.Geolocation != nil {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given Geolocation and assigns it to the Geolocation field.
func (o *Organization) SetGeolocation(v Geolocation) {
	o.Geolocation = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["suggestion"] = o.Suggestion
	}
	if o.Synonyms != nil {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.Partner != nil {
		toSerialize["partner"] = o.Partner
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Geolocation != nil {
		toSerialize["geolocation"] = o.Geolocation
	}
	return json.Marshal(toSerialize)
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


