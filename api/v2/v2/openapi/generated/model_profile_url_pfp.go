/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ProfileUrlPfp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileUrlPfp{}

// ProfileUrlPfp struct for ProfileUrlPfp
type ProfileUrlPfp struct {
	Url string `json:"url"`
}

type _ProfileUrlPfp ProfileUrlPfp

// NewProfileUrlPfp instantiates a new ProfileUrlPfp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileUrlPfp(url string) *ProfileUrlPfp {
	this := ProfileUrlPfp{}
	this.Url = url
	return &this
}

// NewProfileUrlPfpWithDefaults instantiates a new ProfileUrlPfp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileUrlPfpWithDefaults() *ProfileUrlPfp {
	this := ProfileUrlPfp{}
	return &this
}

// GetUrl returns the Url field value
func (o *ProfileUrlPfp) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ProfileUrlPfp) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ProfileUrlPfp) SetUrl(v string) {
	o.Url = v
}

func (o ProfileUrlPfp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileUrlPfp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *ProfileUrlPfp) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProfileUrlPfp := _ProfileUrlPfp{}

	err = json.Unmarshal(bytes, &varProfileUrlPfp)

	if err != nil {
		return err
	}

	*o = ProfileUrlPfp(varProfileUrlPfp)

	return err
}

type NullableProfileUrlPfp struct {
	value *ProfileUrlPfp
	isSet bool
}

func (v NullableProfileUrlPfp) Get() *ProfileUrlPfp {
	return v.value
}

func (v *NullableProfileUrlPfp) Set(val *ProfileUrlPfp) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileUrlPfp) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileUrlPfp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileUrlPfp(val *ProfileUrlPfp) *NullableProfileUrlPfp {
	return &NullableProfileUrlPfp{value: val, isSet: true}
}

func (v NullableProfileUrlPfp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileUrlPfp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


