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

// checks if the RemoveVerificationReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveVerificationReqBody{}

// RemoveVerificationReqBody struct for RemoveVerificationReqBody
type RemoveVerificationReqBody struct {
	// UUID of the signer
	SignerUuid string `json:"signer_uuid"`
	// Ethereum address
	Address string `json:"address"`
}

type _RemoveVerificationReqBody RemoveVerificationReqBody

// NewRemoveVerificationReqBody instantiates a new RemoveVerificationReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveVerificationReqBody(signerUuid string, address string) *RemoveVerificationReqBody {
	this := RemoveVerificationReqBody{}
	this.SignerUuid = signerUuid
	this.Address = address
	return &this
}

// NewRemoveVerificationReqBodyWithDefaults instantiates a new RemoveVerificationReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveVerificationReqBodyWithDefaults() *RemoveVerificationReqBody {
	this := RemoveVerificationReqBody{}
	return &this
}

// GetSignerUuid returns the SignerUuid field value
func (o *RemoveVerificationReqBody) GetSignerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerUuid
}

// GetSignerUuidOk returns a tuple with the SignerUuid field value
// and a boolean to check if the value has been set.
func (o *RemoveVerificationReqBody) GetSignerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerUuid, true
}

// SetSignerUuid sets field value
func (o *RemoveVerificationReqBody) SetSignerUuid(v string) {
	o.SignerUuid = v
}

// GetAddress returns the Address field value
func (o *RemoveVerificationReqBody) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *RemoveVerificationReqBody) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *RemoveVerificationReqBody) SetAddress(v string) {
	o.Address = v
}

func (o RemoveVerificationReqBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveVerificationReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signer_uuid"] = o.SignerUuid
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

func (o *RemoveVerificationReqBody) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signer_uuid",
		"address",
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

	varRemoveVerificationReqBody := _RemoveVerificationReqBody{}

	err = json.Unmarshal(bytes, &varRemoveVerificationReqBody)

	if err != nil {
		return err
	}

	*o = RemoveVerificationReqBody(varRemoveVerificationReqBody)

	return err
}

type NullableRemoveVerificationReqBody struct {
	value *RemoveVerificationReqBody
	isSet bool
}

func (v NullableRemoveVerificationReqBody) Get() *RemoveVerificationReqBody {
	return v.value
}

func (v *NullableRemoveVerificationReqBody) Set(val *RemoveVerificationReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveVerificationReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveVerificationReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveVerificationReqBody(val *RemoveVerificationReqBody) *NullableRemoveVerificationReqBody {
	return &NullableRemoveVerificationReqBody{value: val, isSet: true}
}

func (v NullableRemoveVerificationReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveVerificationReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}