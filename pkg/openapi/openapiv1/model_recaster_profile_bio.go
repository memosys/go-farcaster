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

// checks if the RecasterProfileBio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecasterProfileBio{}

// RecasterProfileBio struct for RecasterProfileBio
type RecasterProfileBio struct {
	Text string `json:"text"`
	Mentions []string `json:"mentions"`
}

type _RecasterProfileBio RecasterProfileBio

// NewRecasterProfileBio instantiates a new RecasterProfileBio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecasterProfileBio(text string, mentions []string) *RecasterProfileBio {
	this := RecasterProfileBio{}
	this.Text = text
	this.Mentions = mentions
	return &this
}

// NewRecasterProfileBioWithDefaults instantiates a new RecasterProfileBio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecasterProfileBioWithDefaults() *RecasterProfileBio {
	this := RecasterProfileBio{}
	return &this
}

// GetText returns the Text field value
func (o *RecasterProfileBio) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *RecasterProfileBio) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *RecasterProfileBio) SetText(v string) {
	o.Text = v
}

// GetMentions returns the Mentions field value
func (o *RecasterProfileBio) GetMentions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value
// and a boolean to check if the value has been set.
func (o *RecasterProfileBio) GetMentionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mentions, true
}

// SetMentions sets field value
func (o *RecasterProfileBio) SetMentions(v []string) {
	o.Mentions = v
}

func (o RecasterProfileBio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecasterProfileBio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	toSerialize["mentions"] = o.Mentions
	return toSerialize, nil
}

func (o *RecasterProfileBio) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
		"mentions",
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

	varRecasterProfileBio := _RecasterProfileBio{}

	err = json.Unmarshal(bytes, &varRecasterProfileBio)

	if err != nil {
		return err
	}

	*o = RecasterProfileBio(varRecasterProfileBio)

	return err
}

type NullableRecasterProfileBio struct {
	value *RecasterProfileBio
	isSet bool
}

func (v NullableRecasterProfileBio) Get() *RecasterProfileBio {
	return v.value
}

func (v *NullableRecasterProfileBio) Set(val *RecasterProfileBio) {
	v.value = val
	v.isSet = true
}

func (v NullableRecasterProfileBio) IsSet() bool {
	return v.isSet
}

func (v *NullableRecasterProfileBio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecasterProfileBio(val *RecasterProfileBio) *NullableRecasterProfileBio {
	return &NullableRecasterProfileBio{value: val, isSet: true}
}

func (v NullableRecasterProfileBio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecasterProfileBio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


