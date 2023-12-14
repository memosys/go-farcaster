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

// checks if the MentionsAndRepliesResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MentionsAndRepliesResponseResult{}

// MentionsAndRepliesResponseResult struct for MentionsAndRepliesResponseResult
type MentionsAndRepliesResponseResult struct {
	Notifications []CastWithInteractions `json:"notifications"`
	Next NextCursor `json:"next"`
}

type _MentionsAndRepliesResponseResult MentionsAndRepliesResponseResult

// NewMentionsAndRepliesResponseResult instantiates a new MentionsAndRepliesResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionsAndRepliesResponseResult(notifications []CastWithInteractions, next NextCursor) *MentionsAndRepliesResponseResult {
	this := MentionsAndRepliesResponseResult{}
	this.Notifications = notifications
	this.Next = next
	return &this
}

// NewMentionsAndRepliesResponseResultWithDefaults instantiates a new MentionsAndRepliesResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionsAndRepliesResponseResultWithDefaults() *MentionsAndRepliesResponseResult {
	this := MentionsAndRepliesResponseResult{}
	return &this
}

// GetNotifications returns the Notifications field value
func (o *MentionsAndRepliesResponseResult) GetNotifications() []CastWithInteractions {
	if o == nil {
		var ret []CastWithInteractions
		return ret
	}

	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value
// and a boolean to check if the value has been set.
func (o *MentionsAndRepliesResponseResult) GetNotificationsOk() ([]CastWithInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notifications, true
}

// SetNotifications sets field value
func (o *MentionsAndRepliesResponseResult) SetNotifications(v []CastWithInteractions) {
	o.Notifications = v
}

// GetNext returns the Next field value
func (o *MentionsAndRepliesResponseResult) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *MentionsAndRepliesResponseResult) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *MentionsAndRepliesResponseResult) SetNext(v NextCursor) {
	o.Next = v
}

func (o MentionsAndRepliesResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MentionsAndRepliesResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifications"] = o.Notifications
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *MentionsAndRepliesResponseResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notifications",
		"next",
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

	varMentionsAndRepliesResponseResult := _MentionsAndRepliesResponseResult{}

	err = json.Unmarshal(bytes, &varMentionsAndRepliesResponseResult)

	if err != nil {
		return err
	}

	*o = MentionsAndRepliesResponseResult(varMentionsAndRepliesResponseResult)

	return err
}

type NullableMentionsAndRepliesResponseResult struct {
	value *MentionsAndRepliesResponseResult
	isSet bool
}

func (v NullableMentionsAndRepliesResponseResult) Get() *MentionsAndRepliesResponseResult {
	return v.value
}

func (v *NullableMentionsAndRepliesResponseResult) Set(val *MentionsAndRepliesResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionsAndRepliesResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionsAndRepliesResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionsAndRepliesResponseResult(val *MentionsAndRepliesResponseResult) *NullableMentionsAndRepliesResponseResult {
	return &NullableMentionsAndRepliesResponseResult{value: val, isSet: true}
}

func (v NullableMentionsAndRepliesResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionsAndRepliesResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


