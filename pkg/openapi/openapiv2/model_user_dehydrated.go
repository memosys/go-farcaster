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

// checks if the UserDehydrated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDehydrated{}

// UserDehydrated struct for UserDehydrated
type UserDehydrated struct {
	Object *string `json:"object,omitempty"`
	// User identifier (unsigned integer)
	Fid *int32 `json:"fid,omitempty"`
}

// NewUserDehydrated instantiates a new UserDehydrated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDehydrated() *UserDehydrated {
	this := UserDehydrated{}
	return &this
}

// NewUserDehydratedWithDefaults instantiates a new UserDehydrated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDehydratedWithDefaults() *UserDehydrated {
	this := UserDehydrated{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *UserDehydrated) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDehydrated) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *UserDehydrated) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *UserDehydrated) SetObject(v string) {
	o.Object = &v
}

// GetFid returns the Fid field value if set, zero value otherwise.
func (o *UserDehydrated) GetFid() int32 {
	if o == nil || IsNil(o.Fid) {
		var ret int32
		return ret
	}
	return *o.Fid
}

// GetFidOk returns a tuple with the Fid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDehydrated) GetFidOk() (*int32, bool) {
	if o == nil || IsNil(o.Fid) {
		return nil, false
	}
	return o.Fid, true
}

// HasFid returns a boolean if a field has been set.
func (o *UserDehydrated) HasFid() bool {
	if o != nil && !IsNil(o.Fid) {
		return true
	}

	return false
}

// SetFid gets a reference to the given int32 and assigns it to the Fid field.
func (o *UserDehydrated) SetFid(v int32) {
	o.Fid = &v
}

func (o UserDehydrated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDehydrated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Fid) {
		toSerialize["fid"] = o.Fid
	}
	return toSerialize, nil
}

type NullableUserDehydrated struct {
	value *UserDehydrated
	isSet bool
}

func (v NullableUserDehydrated) Get() *UserDehydrated {
	return v.value
}

func (v *NullableUserDehydrated) Set(val *UserDehydrated) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDehydrated) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDehydrated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDehydrated(val *UserDehydrated) *NullableUserDehydrated {
	return &NullableUserDehydrated{value: val, isSet: true}
}

func (v NullableUserDehydrated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDehydrated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


