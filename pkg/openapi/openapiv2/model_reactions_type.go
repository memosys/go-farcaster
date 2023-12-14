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

// ReactionsType the model 'ReactionsType'
type ReactionsType string

// List of ReactionsType
const (
	REACTIONSTYPE_LIKES ReactionsType = "likes"
	REACTIONSTYPE_RECASTS ReactionsType = "recasts"
	REACTIONSTYPE_ALL ReactionsType = "all"
)

// All allowed values of ReactionsType enum
var AllowedReactionsTypeEnumValues = []ReactionsType{
	"likes",
	"recasts",
	"all",
}

func (v *ReactionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReactionsType(value)
	for _, existing := range AllowedReactionsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReactionsType", value)
}

// NewReactionsTypeFromValue returns a pointer to a valid ReactionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReactionsTypeFromValue(v string) (*ReactionsType, error) {
	ev := ReactionsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReactionsType: valid values are %v", v, AllowedReactionsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReactionsType) IsValid() bool {
	for _, existing := range AllowedReactionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReactionsType value
func (v ReactionsType) Ptr() *ReactionsType {
	return &v
}

type NullableReactionsType struct {
	value *ReactionsType
	isSet bool
}

func (v NullableReactionsType) Get() *ReactionsType {
	return v.value
}

func (v *NullableReactionsType) Set(val *ReactionsType) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionsType) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionsType(val *ReactionsType) *NullableReactionsType {
	return &NullableReactionsType{value: val, isSet: true}
}

func (v NullableReactionsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
