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

// checks if the CastWithInteractionsReactionsOrRecasts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastWithInteractionsReactionsOrRecasts{}

// CastWithInteractionsReactionsOrRecasts struct for CastWithInteractionsReactionsOrRecasts
type CastWithInteractionsReactionsOrRecasts struct {
	Count  int32    `json:"count"`
	Fids   []int32  `json:"fids"`
	Fnames []string `json:"fnames"`
}

type _CastWithInteractionsReactionsOrRecasts CastWithInteractionsReactionsOrRecasts

// NewCastWithInteractionsReactionsOrRecasts instantiates a new CastWithInteractionsReactionsOrRecasts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastWithInteractionsReactionsOrRecasts(count int32, fids []int32, fnames []string) *CastWithInteractionsReactionsOrRecasts {
	this := CastWithInteractionsReactionsOrRecasts{}
	this.Count = count
	this.Fids = fids
	this.Fnames = fnames
	return &this
}

// NewCastWithInteractionsReactionsOrRecastsWithDefaults instantiates a new CastWithInteractionsReactionsOrRecasts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastWithInteractionsReactionsOrRecastsWithDefaults() *CastWithInteractionsReactionsOrRecasts {
	this := CastWithInteractionsReactionsOrRecasts{}
	return &this
}

// GetCount returns the Count field value
func (o *CastWithInteractionsReactionsOrRecasts) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsReactionsOrRecasts) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *CastWithInteractionsReactionsOrRecasts) SetCount(v int32) {
	o.Count = v
}

// GetFids returns the Fids field value
func (o *CastWithInteractionsReactionsOrRecasts) GetFids() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Fids
}

// GetFidsOk returns a tuple with the Fids field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsReactionsOrRecasts) GetFidsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fids, true
}

// SetFids sets field value
func (o *CastWithInteractionsReactionsOrRecasts) SetFids(v []int32) {
	o.Fids = v
}

// GetFnames returns the Fnames field value
func (o *CastWithInteractionsReactionsOrRecasts) GetFnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fnames
}

// GetFnamesOk returns a tuple with the Fnames field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsReactionsOrRecasts) GetFnamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fnames, true
}

// SetFnames sets field value
func (o *CastWithInteractionsReactionsOrRecasts) SetFnames(v []string) {
	o.Fnames = v
}

func (o CastWithInteractionsReactionsOrRecasts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastWithInteractionsReactionsOrRecasts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["fids"] = o.Fids
	toSerialize["fnames"] = o.Fnames
	return toSerialize, nil
}

func (o *CastWithInteractionsReactionsOrRecasts) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"fids",
		"fnames",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCastWithInteractionsReactionsOrRecasts := _CastWithInteractionsReactionsOrRecasts{}

	err = json.Unmarshal(bytes, &varCastWithInteractionsReactionsOrRecasts)

	if err != nil {
		return err
	}

	*o = CastWithInteractionsReactionsOrRecasts(varCastWithInteractionsReactionsOrRecasts)

	return err
}

type NullableCastWithInteractionsReactionsOrRecasts struct {
	value *CastWithInteractionsReactionsOrRecasts
	isSet bool
}

func (v NullableCastWithInteractionsReactionsOrRecasts) Get() *CastWithInteractionsReactionsOrRecasts {
	return v.value
}

func (v *NullableCastWithInteractionsReactionsOrRecasts) Set(val *CastWithInteractionsReactionsOrRecasts) {
	v.value = val
	v.isSet = true
}

func (v NullableCastWithInteractionsReactionsOrRecasts) IsSet() bool {
	return v.isSet
}

func (v *NullableCastWithInteractionsReactionsOrRecasts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastWithInteractionsReactionsOrRecasts(val *CastWithInteractionsReactionsOrRecasts) *NullableCastWithInteractionsReactionsOrRecasts {
	return &NullableCastWithInteractionsReactionsOrRecasts{value: val, isSet: true}
}

func (v NullableCastWithInteractionsReactionsOrRecasts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastWithInteractionsReactionsOrRecasts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}