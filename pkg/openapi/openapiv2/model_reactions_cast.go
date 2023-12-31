/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv2

import (
	"encoding/json"
	"fmt"
)

// checks if the ReactionsCast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactionsCast{}

// ReactionsCast struct for ReactionsCast
type ReactionsCast struct {
	Hash string `json:"hash"`
	Object string `json:"object"`
}

type _ReactionsCast ReactionsCast

// NewReactionsCast instantiates a new ReactionsCast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionsCast(hash string, object string) *ReactionsCast {
	this := ReactionsCast{}
	this.Hash = hash
	this.Object = object
	return &this
}

// NewReactionsCastWithDefaults instantiates a new ReactionsCast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionsCastWithDefaults() *ReactionsCast {
	this := ReactionsCast{}
	return &this
}

// GetHash returns the Hash field value
func (o *ReactionsCast) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *ReactionsCast) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *ReactionsCast) SetHash(v string) {
	o.Hash = v
}

// GetObject returns the Object field value
func (o *ReactionsCast) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ReactionsCast) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ReactionsCast) SetObject(v string) {
	o.Object = v
}

func (o ReactionsCast) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactionsCast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	toSerialize["object"] = o.Object
	return toSerialize, nil
}

func (o *ReactionsCast) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hash",
		"object",
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

	varReactionsCast := _ReactionsCast{}

	err = json.Unmarshal(bytes, &varReactionsCast)

	if err != nil {
		return err
	}

	*o = ReactionsCast(varReactionsCast)

	return err
}

type NullableReactionsCast struct {
	value *ReactionsCast
	isSet bool
}

func (v NullableReactionsCast) Get() *ReactionsCast {
	return v.value
}

func (v *NullableReactionsCast) Set(val *ReactionsCast) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionsCast) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionsCast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionsCast(val *ReactionsCast) *NullableReactionsCast {
	return &NullableReactionsCast{value: val, isSet: true}
}

func (v NullableReactionsCast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionsCast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


