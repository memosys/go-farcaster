/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ReactorPfp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactorPfp{}

// ReactorPfp struct for ReactorPfp
type ReactorPfp struct {
	// The URL of the reactor's profile picture.
	Url string `json:"url"`
}

type _ReactorPfp ReactorPfp

// NewReactorPfp instantiates a new ReactorPfp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactorPfp(url string) *ReactorPfp {
	this := ReactorPfp{}
	this.Url = url
	return &this
}

// NewReactorPfpWithDefaults instantiates a new ReactorPfp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactorPfpWithDefaults() *ReactorPfp {
	this := ReactorPfp{}
	return &this
}

// GetUrl returns the Url field value
func (o *ReactorPfp) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ReactorPfp) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ReactorPfp) SetUrl(v string) {
	o.Url = v
}

func (o ReactorPfp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactorPfp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *ReactorPfp) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReactorPfp := _ReactorPfp{}

	err = json.Unmarshal(bytes, &varReactorPfp)

	if err != nil {
		return err
	}

	*o = ReactorPfp(varReactorPfp)

	return err
}

type NullableReactorPfp struct {
	value *ReactorPfp
	isSet bool
}

func (v NullableReactorPfp) Get() *ReactorPfp {
	return v.value
}

func (v *NullableReactorPfp) Set(val *ReactorPfp) {
	v.value = val
	v.isSet = true
}

func (v NullableReactorPfp) IsSet() bool {
	return v.isSet
}

func (v *NullableReactorPfp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactorPfp(val *ReactorPfp) *NullableReactorPfp {
	return &NullableReactorPfp{value: val, isSet: true}
}

func (v NullableReactorPfp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactorPfp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}