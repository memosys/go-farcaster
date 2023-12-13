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

// checks if the RecasterProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecasterProfile{}

// RecasterProfile struct for RecasterProfile
type RecasterProfile struct {
	Bio RecasterProfileBio `json:"bio"`
}

type _RecasterProfile RecasterProfile

// NewRecasterProfile instantiates a new RecasterProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecasterProfile(bio RecasterProfileBio) *RecasterProfile {
	this := RecasterProfile{}
	this.Bio = bio
	return &this
}

// NewRecasterProfileWithDefaults instantiates a new RecasterProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecasterProfileWithDefaults() *RecasterProfile {
	this := RecasterProfile{}
	return &this
}

// GetBio returns the Bio field value
func (o *RecasterProfile) GetBio() RecasterProfileBio {
	if o == nil {
		var ret RecasterProfileBio
		return ret
	}

	return o.Bio
}

// GetBioOk returns a tuple with the Bio field value
// and a boolean to check if the value has been set.
func (o *RecasterProfile) GetBioOk() (*RecasterProfileBio, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bio, true
}

// SetBio sets field value
func (o *RecasterProfile) SetBio(v RecasterProfileBio) {
	o.Bio = v
}

func (o RecasterProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecasterProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bio"] = o.Bio
	return toSerialize, nil
}

func (o *RecasterProfile) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bio",
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

	varRecasterProfile := _RecasterProfile{}

	err = json.Unmarshal(bytes, &varRecasterProfile)

	if err != nil {
		return err
	}

	*o = RecasterProfile(varRecasterProfile)

	return err
}

type NullableRecasterProfile struct {
	value *RecasterProfile
	isSet bool
}

func (v NullableRecasterProfile) Get() *RecasterProfile {
	return v.value
}

func (v *NullableRecasterProfile) Set(val *RecasterProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableRecasterProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableRecasterProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecasterProfile(val *RecasterProfile) *NullableRecasterProfile {
	return &NullableRecasterProfile{value: val, isSet: true}
}

func (v NullableRecasterProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecasterProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


