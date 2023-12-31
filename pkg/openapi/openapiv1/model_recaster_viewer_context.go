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

// checks if the RecasterViewerContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecasterViewerContext{}

// RecasterViewerContext struct for RecasterViewerContext
type RecasterViewerContext struct {
	// Indicates if the viewer is following the recaster.
	Following bool `json:"following"`
	// Indicates if the recaster is followed by the viewer.
	FollowedBy bool `json:"followedBy"`
}

type _RecasterViewerContext RecasterViewerContext

// NewRecasterViewerContext instantiates a new RecasterViewerContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecasterViewerContext(following bool, followedBy bool) *RecasterViewerContext {
	this := RecasterViewerContext{}
	this.Following = following
	this.FollowedBy = followedBy
	return &this
}

// NewRecasterViewerContextWithDefaults instantiates a new RecasterViewerContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecasterViewerContextWithDefaults() *RecasterViewerContext {
	this := RecasterViewerContext{}
	return &this
}

// GetFollowing returns the Following field value
func (o *RecasterViewerContext) GetFollowing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Following
}

// GetFollowingOk returns a tuple with the Following field value
// and a boolean to check if the value has been set.
func (o *RecasterViewerContext) GetFollowingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Following, true
}

// SetFollowing sets field value
func (o *RecasterViewerContext) SetFollowing(v bool) {
	o.Following = v
}

// GetFollowedBy returns the FollowedBy field value
func (o *RecasterViewerContext) GetFollowedBy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FollowedBy
}

// GetFollowedByOk returns a tuple with the FollowedBy field value
// and a boolean to check if the value has been set.
func (o *RecasterViewerContext) GetFollowedByOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowedBy, true
}

// SetFollowedBy sets field value
func (o *RecasterViewerContext) SetFollowedBy(v bool) {
	o.FollowedBy = v
}

func (o RecasterViewerContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecasterViewerContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["following"] = o.Following
	toSerialize["followedBy"] = o.FollowedBy
	return toSerialize, nil
}

func (o *RecasterViewerContext) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"following",
		"followedBy",
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

	varRecasterViewerContext := _RecasterViewerContext{}

	err = json.Unmarshal(bytes, &varRecasterViewerContext)

	if err != nil {
		return err
	}

	*o = RecasterViewerContext(varRecasterViewerContext)

	return err
}

type NullableRecasterViewerContext struct {
	value *RecasterViewerContext
	isSet bool
}

func (v NullableRecasterViewerContext) Get() *RecasterViewerContext {
	return v.value
}

func (v *NullableRecasterViewerContext) Set(val *RecasterViewerContext) {
	v.value = val
	v.isSet = true
}

func (v NullableRecasterViewerContext) IsSet() bool {
	return v.isSet
}

func (v *NullableRecasterViewerContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecasterViewerContext(val *RecasterViewerContext) *NullableRecasterViewerContext {
	return &NullableRecasterViewerContext{value: val, isSet: true}
}

func (v NullableRecasterViewerContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecasterViewerContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


