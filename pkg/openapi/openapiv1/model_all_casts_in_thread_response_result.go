/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv1

import (
	"encoding/json"
	"fmt"
)

// checks if the AllCastsInThreadResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllCastsInThreadResponseResult{}

// AllCastsInThreadResponseResult struct for AllCastsInThreadResponseResult
type AllCastsInThreadResponseResult struct {
	Casts []CastWithInteractions `json:"casts"`
}

type _AllCastsInThreadResponseResult AllCastsInThreadResponseResult

// NewAllCastsInThreadResponseResult instantiates a new AllCastsInThreadResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllCastsInThreadResponseResult(casts []CastWithInteractions) *AllCastsInThreadResponseResult {
	this := AllCastsInThreadResponseResult{}
	this.Casts = casts
	return &this
}

// NewAllCastsInThreadResponseResultWithDefaults instantiates a new AllCastsInThreadResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllCastsInThreadResponseResultWithDefaults() *AllCastsInThreadResponseResult {
	this := AllCastsInThreadResponseResult{}
	return &this
}

// GetCasts returns the Casts field value
func (o *AllCastsInThreadResponseResult) GetCasts() []CastWithInteractions {
	if o == nil {
		var ret []CastWithInteractions
		return ret
	}

	return o.Casts
}

// GetCastsOk returns a tuple with the Casts field value
// and a boolean to check if the value has been set.
func (o *AllCastsInThreadResponseResult) GetCastsOk() ([]CastWithInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Casts, true
}

// SetCasts sets field value
func (o *AllCastsInThreadResponseResult) SetCasts(v []CastWithInteractions) {
	o.Casts = v
}

func (o AllCastsInThreadResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllCastsInThreadResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["casts"] = o.Casts
	return toSerialize, nil
}

func (o *AllCastsInThreadResponseResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"casts",
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

	varAllCastsInThreadResponseResult := _AllCastsInThreadResponseResult{}

	err = json.Unmarshal(bytes, &varAllCastsInThreadResponseResult)

	if err != nil {
		return err
	}

	*o = AllCastsInThreadResponseResult(varAllCastsInThreadResponseResult)

	return err
}

type NullableAllCastsInThreadResponseResult struct {
	value *AllCastsInThreadResponseResult
	isSet bool
}

func (v NullableAllCastsInThreadResponseResult) Get() *AllCastsInThreadResponseResult {
	return v.value
}

func (v *NullableAllCastsInThreadResponseResult) Set(val *AllCastsInThreadResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAllCastsInThreadResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAllCastsInThreadResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllCastsInThreadResponseResult(val *AllCastsInThreadResponseResult) *NullableAllCastsInThreadResponseResult {
	return &NullableAllCastsInThreadResponseResult{value: val, isSet: true}
}

func (v NullableAllCastsInThreadResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllCastsInThreadResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


