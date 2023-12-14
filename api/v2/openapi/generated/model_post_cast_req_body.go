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

// checks if the PostCastReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCastReqBody{}

// PostCastReqBody struct for PostCastReqBody
type PostCastReqBody struct {
	// UUID of the signer
	SignerUuid string         `json:"signer_uuid"`
	Text       *string        `json:"text,omitempty"`
	Embeds     []EmbeddedCast `json:"embeds,omitempty"`
	// Parent URL or Cast Hash
	Parent *string `json:"parent,omitempty"`
}

type _PostCastReqBody PostCastReqBody

// NewPostCastReqBody instantiates a new PostCastReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCastReqBody(signerUuid string) *PostCastReqBody {
	this := PostCastReqBody{}
	this.SignerUuid = signerUuid
	return &this
}

// NewPostCastReqBodyWithDefaults instantiates a new PostCastReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCastReqBodyWithDefaults() *PostCastReqBody {
	this := PostCastReqBody{}
	return &this
}

// GetSignerUuid returns the SignerUuid field value
func (o *PostCastReqBody) GetSignerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerUuid
}

// GetSignerUuidOk returns a tuple with the SignerUuid field value
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetSignerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerUuid, true
}

// SetSignerUuid sets field value
func (o *PostCastReqBody) SetSignerUuid(v string) {
	o.SignerUuid = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PostCastReqBody) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PostCastReqBody) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PostCastReqBody) SetText(v string) {
	o.Text = &v
}

// GetEmbeds returns the Embeds field value if set, zero value otherwise.
func (o *PostCastReqBody) GetEmbeds() []EmbeddedCast {
	if o == nil || IsNil(o.Embeds) {
		var ret []EmbeddedCast
		return ret
	}
	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetEmbedsOk() ([]EmbeddedCast, bool) {
	if o == nil || IsNil(o.Embeds) {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *PostCastReqBody) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given []EmbeddedCast and assigns it to the Embeds field.
func (o *PostCastReqBody) SetEmbeds(v []EmbeddedCast) {
	o.Embeds = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *PostCastReqBody) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *PostCastReqBody) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *PostCastReqBody) SetParent(v string) {
	o.Parent = &v
}

func (o PostCastReqBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCastReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signer_uuid"] = o.SignerUuid
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Embeds) {
		toSerialize["embeds"] = o.Embeds
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	return toSerialize, nil
}

func (o *PostCastReqBody) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signer_uuid",
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

	varPostCastReqBody := _PostCastReqBody{}

	err = json.Unmarshal(bytes, &varPostCastReqBody)

	if err != nil {
		return err
	}

	*o = PostCastReqBody(varPostCastReqBody)

	return err
}

type NullablePostCastReqBody struct {
	value *PostCastReqBody
	isSet bool
}

func (v NullablePostCastReqBody) Get() *PostCastReqBody {
	return v.value
}

func (v *NullablePostCastReqBody) Set(val *PostCastReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCastReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCastReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCastReqBody(val *PostCastReqBody) *NullablePostCastReqBody {
	return &NullablePostCastReqBody{value: val, isSet: true}
}

func (v NullablePostCastReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCastReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}