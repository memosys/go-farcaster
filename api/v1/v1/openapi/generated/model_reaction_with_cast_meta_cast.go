/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ReactionWithCastMetaCast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactionWithCastMetaCast{}

// ReactionWithCastMetaCast struct for ReactionWithCastMetaCast
type ReactionWithCastMetaCast struct {
	// User identifier (unsigned integer)
	CastFid int32 `json:"cast_fid"`
	CastHash string `json:"cast_hash"`
	CastText string `json:"cast_text"`
	CastEmbeds []EmbedUrl `json:"cast_embeds"`
	CastTimestamp time.Time `json:"cast_timestamp"`
}

type _ReactionWithCastMetaCast ReactionWithCastMetaCast

// NewReactionWithCastMetaCast instantiates a new ReactionWithCastMetaCast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionWithCastMetaCast(castFid int32, castHash string, castText string, castEmbeds []EmbedUrl, castTimestamp time.Time) *ReactionWithCastMetaCast {
	this := ReactionWithCastMetaCast{}
	this.CastFid = castFid
	this.CastHash = castHash
	this.CastText = castText
	this.CastEmbeds = castEmbeds
	this.CastTimestamp = castTimestamp
	return &this
}

// NewReactionWithCastMetaCastWithDefaults instantiates a new ReactionWithCastMetaCast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionWithCastMetaCastWithDefaults() *ReactionWithCastMetaCast {
	this := ReactionWithCastMetaCast{}
	var castFid int32 = 3
	this.CastFid = castFid
	return &this
}

// GetCastFid returns the CastFid field value
func (o *ReactionWithCastMetaCast) GetCastFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CastFid
}

// GetCastFidOk returns a tuple with the CastFid field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastMetaCast) GetCastFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CastFid, true
}

// SetCastFid sets field value
func (o *ReactionWithCastMetaCast) SetCastFid(v int32) {
	o.CastFid = v
}

// GetCastHash returns the CastHash field value
func (o *ReactionWithCastMetaCast) GetCastHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CastHash
}

// GetCastHashOk returns a tuple with the CastHash field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastMetaCast) GetCastHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CastHash, true
}

// SetCastHash sets field value
func (o *ReactionWithCastMetaCast) SetCastHash(v string) {
	o.CastHash = v
}

// GetCastText returns the CastText field value
func (o *ReactionWithCastMetaCast) GetCastText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CastText
}

// GetCastTextOk returns a tuple with the CastText field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastMetaCast) GetCastTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CastText, true
}

// SetCastText sets field value
func (o *ReactionWithCastMetaCast) SetCastText(v string) {
	o.CastText = v
}

// GetCastEmbeds returns the CastEmbeds field value
func (o *ReactionWithCastMetaCast) GetCastEmbeds() []EmbedUrl {
	if o == nil {
		var ret []EmbedUrl
		return ret
	}

	return o.CastEmbeds
}

// GetCastEmbedsOk returns a tuple with the CastEmbeds field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastMetaCast) GetCastEmbedsOk() ([]EmbedUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.CastEmbeds, true
}

// SetCastEmbeds sets field value
func (o *ReactionWithCastMetaCast) SetCastEmbeds(v []EmbedUrl) {
	o.CastEmbeds = v
}

// GetCastTimestamp returns the CastTimestamp field value
func (o *ReactionWithCastMetaCast) GetCastTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CastTimestamp
}

// GetCastTimestampOk returns a tuple with the CastTimestamp field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastMetaCast) GetCastTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CastTimestamp, true
}

// SetCastTimestamp sets field value
func (o *ReactionWithCastMetaCast) SetCastTimestamp(v time.Time) {
	o.CastTimestamp = v
}

func (o ReactionWithCastMetaCast) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactionWithCastMetaCast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cast_fid"] = o.CastFid
	toSerialize["cast_hash"] = o.CastHash
	toSerialize["cast_text"] = o.CastText
	toSerialize["cast_embeds"] = o.CastEmbeds
	toSerialize["cast_timestamp"] = o.CastTimestamp
	return toSerialize, nil
}

func (o *ReactionWithCastMetaCast) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cast_fid",
		"cast_hash",
		"cast_text",
		"cast_embeds",
		"cast_timestamp",
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

	varReactionWithCastMetaCast := _ReactionWithCastMetaCast{}

	err = json.Unmarshal(bytes, &varReactionWithCastMetaCast)

	if err != nil {
		return err
	}

	*o = ReactionWithCastMetaCast(varReactionWithCastMetaCast)

	return err
}

type NullableReactionWithCastMetaCast struct {
	value *ReactionWithCastMetaCast
	isSet bool
}

func (v NullableReactionWithCastMetaCast) Get() *ReactionWithCastMetaCast {
	return v.value
}

func (v *NullableReactionWithCastMetaCast) Set(val *ReactionWithCastMetaCast) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionWithCastMetaCast) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionWithCastMetaCast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionWithCastMetaCast(val *ReactionWithCastMetaCast) *NullableReactionWithCastMetaCast {
	return &NullableReactionWithCastMetaCast{value: val, isSet: true}
}

func (v NullableReactionWithCastMetaCast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionWithCastMetaCast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


