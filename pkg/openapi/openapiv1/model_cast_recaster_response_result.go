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

// checks if the CastRecasterResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastRecasterResponseResult{}

// CastRecasterResponseResult struct for CastRecasterResponseResult
type CastRecasterResponseResult struct {
	Users []Recaster `json:"users"`
	Next NextCursor `json:"next"`
}

type _CastRecasterResponseResult CastRecasterResponseResult

// NewCastRecasterResponseResult instantiates a new CastRecasterResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastRecasterResponseResult(users []Recaster, next NextCursor) *CastRecasterResponseResult {
	this := CastRecasterResponseResult{}
	this.Users = users
	this.Next = next
	return &this
}

// NewCastRecasterResponseResultWithDefaults instantiates a new CastRecasterResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastRecasterResponseResultWithDefaults() *CastRecasterResponseResult {
	this := CastRecasterResponseResult{}
	return &this
}

// GetUsers returns the Users field value
func (o *CastRecasterResponseResult) GetUsers() []Recaster {
	if o == nil {
		var ret []Recaster
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *CastRecasterResponseResult) GetUsersOk() ([]Recaster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *CastRecasterResponseResult) SetUsers(v []Recaster) {
	o.Users = v
}

// GetNext returns the Next field value
func (o *CastRecasterResponseResult) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *CastRecasterResponseResult) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *CastRecasterResponseResult) SetNext(v NextCursor) {
	o.Next = v
}

func (o CastRecasterResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastRecasterResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *CastRecasterResponseResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
		"next",
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

	varCastRecasterResponseResult := _CastRecasterResponseResult{}

	err = json.Unmarshal(bytes, &varCastRecasterResponseResult)

	if err != nil {
		return err
	}

	*o = CastRecasterResponseResult(varCastRecasterResponseResult)

	return err
}

type NullableCastRecasterResponseResult struct {
	value *CastRecasterResponseResult
	isSet bool
}

func (v NullableCastRecasterResponseResult) Get() *CastRecasterResponseResult {
	return v.value
}

func (v *NullableCastRecasterResponseResult) Set(val *CastRecasterResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCastRecasterResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCastRecasterResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastRecasterResponseResult(val *CastRecasterResponseResult) *NullableCastRecasterResponseResult {
	return &NullableCastRecasterResponseResult{value: val, isSet: true}
}

func (v NullableCastRecasterResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastRecasterResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


