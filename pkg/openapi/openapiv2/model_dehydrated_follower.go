/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv2

import (
	"encoding/json"
)

// checks if the DehydratedFollower type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DehydratedFollower{}

// DehydratedFollower struct for DehydratedFollower
type DehydratedFollower struct {
	Object *string `json:"object,omitempty"`
	User *UserDehydrated `json:"user,omitempty"`
}

// NewDehydratedFollower instantiates a new DehydratedFollower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDehydratedFollower() *DehydratedFollower {
	this := DehydratedFollower{}
	return &this
}

// NewDehydratedFollowerWithDefaults instantiates a new DehydratedFollower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDehydratedFollowerWithDefaults() *DehydratedFollower {
	this := DehydratedFollower{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DehydratedFollower) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DehydratedFollower) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DehydratedFollower) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *DehydratedFollower) SetObject(v string) {
	o.Object = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DehydratedFollower) GetUser() UserDehydrated {
	if o == nil || IsNil(o.User) {
		var ret UserDehydrated
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DehydratedFollower) GetUserOk() (*UserDehydrated, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DehydratedFollower) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserDehydrated and assigns it to the User field.
func (o *DehydratedFollower) SetUser(v UserDehydrated) {
	o.User = &v
}

func (o DehydratedFollower) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DehydratedFollower) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableDehydratedFollower struct {
	value *DehydratedFollower
	isSet bool
}

func (v NullableDehydratedFollower) Get() *DehydratedFollower {
	return v.value
}

func (v *NullableDehydratedFollower) Set(val *DehydratedFollower) {
	v.value = val
	v.isSet = true
}

func (v NullableDehydratedFollower) IsSet() bool {
	return v.isSet
}

func (v *NullableDehydratedFollower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDehydratedFollower(val *DehydratedFollower) *NullableDehydratedFollower {
	return &NullableDehydratedFollower{value: val, isSet: true}
}

func (v NullableDehydratedFollower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDehydratedFollower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


