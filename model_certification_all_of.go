/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CertificationAllOf struct for CertificationAllOf
type CertificationAllOf struct {
	Validity *string `json:"validity,omitempty"`
}

// NewCertificationAllOf instantiates a new CertificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificationAllOf() *CertificationAllOf {
	this := CertificationAllOf{}
	return &this
}

// NewCertificationAllOfWithDefaults instantiates a new CertificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationAllOfWithDefaults() *CertificationAllOf {
	this := CertificationAllOf{}
	return &this
}

// GetValidity returns the Validity field value if set, zero value otherwise.
func (o *CertificationAllOf) GetValidity() string {
	if o == nil || o.Validity == nil {
		var ret string
		return ret
	}
	return *o.Validity
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationAllOf) GetValidityOk() (*string, bool) {
	if o == nil || o.Validity == nil {
		return nil, false
	}
	return o.Validity, true
}

// HasValidity returns a boolean if a field has been set.
func (o *CertificationAllOf) HasValidity() bool {
	if o != nil && o.Validity != nil {
		return true
	}

	return false
}

// SetValidity gets a reference to the given string and assigns it to the Validity field.
func (o *CertificationAllOf) SetValidity(v string) {
	o.Validity = &v
}

func (o CertificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Validity != nil {
		toSerialize["validity"] = o.Validity
	}
	return json.Marshal(toSerialize)
}

type NullableCertificationAllOf struct {
	value *CertificationAllOf
	isSet bool
}

func (v NullableCertificationAllOf) Get() *CertificationAllOf {
	return v.value
}

func (v *NullableCertificationAllOf) Set(val *CertificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificationAllOf(val *CertificationAllOf) *NullableCertificationAllOf {
	return &NullableCertificationAllOf{value: val, isSet: true}
}

func (v NullableCertificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


