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

// checks if the UserSearchResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSearchResponseResult{}

// UserSearchResponseResult struct for UserSearchResponseResult
type UserSearchResponseResult struct {
	Users []SearchedUser `json:"users"`
}

type _UserSearchResponseResult UserSearchResponseResult

// NewUserSearchResponseResult instantiates a new UserSearchResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSearchResponseResult(users []SearchedUser) *UserSearchResponseResult {
	this := UserSearchResponseResult{}
	this.Users = users
	return &this
}

// NewUserSearchResponseResultWithDefaults instantiates a new UserSearchResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSearchResponseResultWithDefaults() *UserSearchResponseResult {
	this := UserSearchResponseResult{}
	return &this
}

// GetUsers returns the Users field value
func (o *UserSearchResponseResult) GetUsers() []SearchedUser {
	if o == nil {
		var ret []SearchedUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *UserSearchResponseResult) GetUsersOk() ([]SearchedUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UserSearchResponseResult) SetUsers(v []SearchedUser) {
	o.Users = v
}

func (o UserSearchResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSearchResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *UserSearchResponseResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
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

	varUserSearchResponseResult := _UserSearchResponseResult{}

	err = json.Unmarshal(bytes, &varUserSearchResponseResult)

	if err != nil {
		return err
	}

	*o = UserSearchResponseResult(varUserSearchResponseResult)

	return err
}

type NullableUserSearchResponseResult struct {
	value *UserSearchResponseResult
	isSet bool
}

func (v NullableUserSearchResponseResult) Get() *UserSearchResponseResult {
	return v.value
}

func (v *NullableUserSearchResponseResult) Set(val *UserSearchResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSearchResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSearchResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSearchResponseResult(val *UserSearchResponseResult) *NullableUserSearchResponseResult {
	return &NullableUserSearchResponseResult{value: val, isSet: true}
}

func (v NullableUserSearchResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSearchResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


