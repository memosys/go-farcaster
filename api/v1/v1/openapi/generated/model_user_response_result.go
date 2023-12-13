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

// checks if the UserResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponseResult{}

// UserResponseResult struct for UserResponseResult
type UserResponseResult struct {
	User User `json:"user"`
}

type _UserResponseResult UserResponseResult

// NewUserResponseResult instantiates a new UserResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponseResult(user User) *UserResponseResult {
	this := UserResponseResult{}
	this.User = user
	return &this
}

// NewUserResponseResultWithDefaults instantiates a new UserResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseResultWithDefaults() *UserResponseResult {
	this := UserResponseResult{}
	return &this
}

// GetUser returns the User field value
func (o *UserResponseResult) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserResponseResult) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserResponseResult) SetUser(v User) {
	o.User = v
}

func (o UserResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *UserResponseResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
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

	varUserResponseResult := _UserResponseResult{}

	err = json.Unmarshal(bytes, &varUserResponseResult)

	if err != nil {
		return err
	}

	*o = UserResponseResult(varUserResponseResult)

	return err
}

type NullableUserResponseResult struct {
	value *UserResponseResult
	isSet bool
}

func (v NullableUserResponseResult) Get() *UserResponseResult {
	return v.value
}

func (v *NullableUserResponseResult) Set(val *UserResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseResult(val *UserResponseResult) *NullableUserResponseResult {
	return &NullableUserResponseResult{value: val, isSet: true}
}

func (v NullableUserResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


