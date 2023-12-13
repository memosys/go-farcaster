/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ReactionReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactionReqBody{}

// ReactionReqBody struct for ReactionReqBody
type ReactionReqBody struct {
	// UUID of the signer
	SignerUuid string `json:"signer_uuid"`
	ReactionType ReactionType `json:"reaction_type"`
	Target string `json:"target"`
}

type _ReactionReqBody ReactionReqBody

// NewReactionReqBody instantiates a new ReactionReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionReqBody(signerUuid string, reactionType ReactionType, target string) *ReactionReqBody {
	this := ReactionReqBody{}
	this.SignerUuid = signerUuid
	this.ReactionType = reactionType
	this.Target = target
	return &this
}

// NewReactionReqBodyWithDefaults instantiates a new ReactionReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionReqBodyWithDefaults() *ReactionReqBody {
	this := ReactionReqBody{}
	return &this
}

// GetSignerUuid returns the SignerUuid field value
func (o *ReactionReqBody) GetSignerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerUuid
}

// GetSignerUuidOk returns a tuple with the SignerUuid field value
// and a boolean to check if the value has been set.
func (o *ReactionReqBody) GetSignerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerUuid, true
}

// SetSignerUuid sets field value
func (o *ReactionReqBody) SetSignerUuid(v string) {
	o.SignerUuid = v
}

// GetReactionType returns the ReactionType field value
func (o *ReactionReqBody) GetReactionType() ReactionType {
	if o == nil {
		var ret ReactionType
		return ret
	}

	return o.ReactionType
}

// GetReactionTypeOk returns a tuple with the ReactionType field value
// and a boolean to check if the value has been set.
func (o *ReactionReqBody) GetReactionTypeOk() (*ReactionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReactionType, true
}

// SetReactionType sets field value
func (o *ReactionReqBody) SetReactionType(v ReactionType) {
	o.ReactionType = v
}

// GetTarget returns the Target field value
func (o *ReactionReqBody) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ReactionReqBody) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *ReactionReqBody) SetTarget(v string) {
	o.Target = v
}

func (o ReactionReqBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactionReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signer_uuid"] = o.SignerUuid
	toSerialize["reaction_type"] = o.ReactionType
	toSerialize["target"] = o.Target
	return toSerialize, nil
}

func (o *ReactionReqBody) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signer_uuid",
		"reaction_type",
		"target",
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

	varReactionReqBody := _ReactionReqBody{}

	err = json.Unmarshal(bytes, &varReactionReqBody)

	if err != nil {
		return err
	}

	*o = ReactionReqBody(varReactionReqBody)

	return err
}

type NullableReactionReqBody struct {
	value *ReactionReqBody
	isSet bool
}

func (v NullableReactionReqBody) Get() *ReactionReqBody {
	return v.value
}

func (v *NullableReactionReqBody) Set(val *ReactionReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionReqBody(val *ReactionReqBody) *NullableReactionReqBody {
	return &NullableReactionReqBody{value: val, isSet: true}
}

func (v NullableReactionReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


