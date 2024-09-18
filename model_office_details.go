/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the OfficeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfficeDetails{}

// OfficeDetails struct for OfficeDetails
type OfficeDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	Office *Office `json:"office,omitempty"`
	Country *Country `json:"country,omitempty"`
}

// NewOfficeDetails instantiates a new OfficeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfficeDetails() *OfficeDetails {
	this := OfficeDetails{}
	return &this
}

// NewOfficeDetailsWithDefaults instantiates a new OfficeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfficeDetailsWithDefaults() *OfficeDetails {
	this := OfficeDetails{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *OfficeDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *OfficeDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *OfficeDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetOffice returns the Office field value if set, zero value otherwise.
func (o *OfficeDetails) GetOffice() Office {
	if o == nil || IsNil(o.Office) {
		var ret Office
		return ret
	}
	return *o.Office
}

// GetOfficeOk returns a tuple with the Office field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeDetails) GetOfficeOk() (*Office, bool) {
	if o == nil || IsNil(o.Office) {
		return nil, false
	}
	return o.Office, true
}

// HasOffice returns a boolean if a field has been set.
func (o *OfficeDetails) HasOffice() bool {
	if o != nil && !IsNil(o.Office) {
		return true
	}

	return false
}

// SetOffice gets a reference to the given Office and assigns it to the Office field.
func (o *OfficeDetails) SetOffice(v Office) {
	o.Office = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *OfficeDetails) GetCountry() Country {
	if o == nil || IsNil(o.Country) {
		var ret Country
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeDetails) GetCountryOk() (*Country, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *OfficeDetails) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given Country and assigns it to the Country field.
func (o *OfficeDetails) SetCountry(v Country) {
	o.Country = &v
}

func (o OfficeDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfficeDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.Office) {
		toSerialize["office"] = o.Office
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	return toSerialize, nil
}

type NullableOfficeDetails struct {
	value *OfficeDetails
	isSet bool
}

func (v NullableOfficeDetails) Get() *OfficeDetails {
	return v.value
}

func (v *NullableOfficeDetails) Set(val *OfficeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOfficeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOfficeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfficeDetails(val *OfficeDetails) *NullableOfficeDetails {
	return &NullableOfficeDetails{value: val, isSet: true}
}

func (v NullableOfficeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfficeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


