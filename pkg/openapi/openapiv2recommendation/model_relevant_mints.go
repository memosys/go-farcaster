/*
Recommendation API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv2recommendation

import (
	"encoding/json"
	"fmt"
)

// checks if the RelevantMints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelevantMints{}

// RelevantMints A list of mint objects relevant to the user
type RelevantMints struct {
	Mints []RelevantMint `json:"mints"`
}

type _RelevantMints RelevantMints

// NewRelevantMints instantiates a new RelevantMints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelevantMints(mints []RelevantMint) *RelevantMints {
	this := RelevantMints{}
	this.Mints = mints
	return &this
}

// NewRelevantMintsWithDefaults instantiates a new RelevantMints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelevantMintsWithDefaults() *RelevantMints {
	this := RelevantMints{}
	return &this
}

// GetMints returns the Mints field value
func (o *RelevantMints) GetMints() []RelevantMint {
	if o == nil {
		var ret []RelevantMint
		return ret
	}

	return o.Mints
}

// GetMintsOk returns a tuple with the Mints field value
// and a boolean to check if the value has been set.
func (o *RelevantMints) GetMintsOk() ([]RelevantMint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mints, true
}

// SetMints sets field value
func (o *RelevantMints) SetMints(v []RelevantMint) {
	o.Mints = v
}

func (o RelevantMints) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelevantMints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mints"] = o.Mints
	return toSerialize, nil
}

func (o *RelevantMints) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mints",
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

	varRelevantMints := _RelevantMints{}

	err = json.Unmarshal(bytes, &varRelevantMints)

	if err != nil {
		return err
	}

	*o = RelevantMints(varRelevantMints)

	return err
}

type NullableRelevantMints struct {
	value *RelevantMints
	isSet bool
}

func (v NullableRelevantMints) Get() *RelevantMints {
	return v.value
}

func (v *NullableRelevantMints) Set(val *RelevantMints) {
	v.value = val
	v.isSet = true
}

func (v NullableRelevantMints) IsSet() bool {
	return v.isSet
}

func (v *NullableRelevantMints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelevantMints(val *RelevantMints) *NullableRelevantMints {
	return &NullableRelevantMints{value: val, isSet: true}
}

func (v NullableRelevantMints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelevantMints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


