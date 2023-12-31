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

// checks if the FollowResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowResponseResult{}

// FollowResponseResult struct for FollowResponseResult
type FollowResponseResult struct {
	Users []FollowResponseUser `json:"users"`
	Next NextCursor `json:"next"`
}

type _FollowResponseResult FollowResponseResult

// NewFollowResponseResult instantiates a new FollowResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowResponseResult(users []FollowResponseUser, next NextCursor) *FollowResponseResult {
	this := FollowResponseResult{}
	this.Users = users
	this.Next = next
	return &this
}

// NewFollowResponseResultWithDefaults instantiates a new FollowResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowResponseResultWithDefaults() *FollowResponseResult {
	this := FollowResponseResult{}
	return &this
}

// GetUsers returns the Users field value
func (o *FollowResponseResult) GetUsers() []FollowResponseUser {
	if o == nil {
		var ret []FollowResponseUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *FollowResponseResult) GetUsersOk() ([]FollowResponseUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *FollowResponseResult) SetUsers(v []FollowResponseUser) {
	o.Users = v
}

// GetNext returns the Next field value
func (o *FollowResponseResult) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *FollowResponseResult) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *FollowResponseResult) SetNext(v NextCursor) {
	o.Next = v
}

func (o FollowResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *FollowResponseResult) UnmarshalJSON(bytes []byte) (err error) {
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

	varFollowResponseResult := _FollowResponseResult{}

	err = json.Unmarshal(bytes, &varFollowResponseResult)

	if err != nil {
		return err
	}

	*o = FollowResponseResult(varFollowResponseResult)

	return err
}

type NullableFollowResponseResult struct {
	value *FollowResponseResult
	isSet bool
}

func (v NullableFollowResponseResult) Get() *FollowResponseResult {
	return v.value
}

func (v *NullableFollowResponseResult) Set(val *FollowResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowResponseResult(val *FollowResponseResult) *NullableFollowResponseResult {
	return &NullableFollowResponseResult{value: val, isSet: true}
}

func (v NullableFollowResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


