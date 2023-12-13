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

// checks if the CastWithInteractionsReplies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastWithInteractionsReplies{}

// CastWithInteractionsReplies struct for CastWithInteractionsReplies
type CastWithInteractionsReplies struct {
	Count int32 `json:"count"`
}

type _CastWithInteractionsReplies CastWithInteractionsReplies

// NewCastWithInteractionsReplies instantiates a new CastWithInteractionsReplies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastWithInteractionsReplies(count int32) *CastWithInteractionsReplies {
	this := CastWithInteractionsReplies{}
	this.Count = count
	return &this
}

// NewCastWithInteractionsRepliesWithDefaults instantiates a new CastWithInteractionsReplies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastWithInteractionsRepliesWithDefaults() *CastWithInteractionsReplies {
	this := CastWithInteractionsReplies{}
	return &this
}

// GetCount returns the Count field value
func (o *CastWithInteractionsReplies) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsReplies) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *CastWithInteractionsReplies) SetCount(v int32) {
	o.Count = v
}

func (o CastWithInteractionsReplies) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastWithInteractionsReplies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

func (o *CastWithInteractionsReplies) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
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

	varCastWithInteractionsReplies := _CastWithInteractionsReplies{}

	err = json.Unmarshal(bytes, &varCastWithInteractionsReplies)

	if err != nil {
		return err
	}

	*o = CastWithInteractionsReplies(varCastWithInteractionsReplies)

	return err
}

type NullableCastWithInteractionsReplies struct {
	value *CastWithInteractionsReplies
	isSet bool
}

func (v NullableCastWithInteractionsReplies) Get() *CastWithInteractionsReplies {
	return v.value
}

func (v *NullableCastWithInteractionsReplies) Set(val *CastWithInteractionsReplies) {
	v.value = val
	v.isSet = true
}

func (v NullableCastWithInteractionsReplies) IsSet() bool {
	return v.isSet
}

func (v *NullableCastWithInteractionsReplies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastWithInteractionsReplies(val *CastWithInteractionsReplies) *NullableCastWithInteractionsReplies {
	return &NullableCastWithInteractionsReplies{value: val, isSet: true}
}

func (v NullableCastWithInteractionsReplies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastWithInteractionsReplies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


