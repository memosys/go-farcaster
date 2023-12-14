/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv2

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Notification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Notification{}

// Notification struct for Notification
type Notification struct {
	Object string `json:"object"`
	MostRecentTimestamp time.Time `json:"most_recent_timestamp"`
	Type string `json:"type"`
	Follows []Follow `json:"follows,omitempty"`
	Cast *CastWithInteractions `json:"cast,omitempty"`
	Reactions []Reactions `json:"reactions,omitempty"`
}

type _Notification Notification

// NewNotification instantiates a new Notification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotification(object string, mostRecentTimestamp time.Time, type_ string) *Notification {
	this := Notification{}
	this.Object = object
	this.MostRecentTimestamp = mostRecentTimestamp
	this.Type = type_
	return &this
}

// NewNotificationWithDefaults instantiates a new Notification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithDefaults() *Notification {
	this := Notification{}
	return &this
}

// GetObject returns the Object field value
func (o *Notification) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Notification) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Notification) SetObject(v string) {
	o.Object = v
}

// GetMostRecentTimestamp returns the MostRecentTimestamp field value
func (o *Notification) GetMostRecentTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.MostRecentTimestamp
}

// GetMostRecentTimestampOk returns a tuple with the MostRecentTimestamp field value
// and a boolean to check if the value has been set.
func (o *Notification) GetMostRecentTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MostRecentTimestamp, true
}

// SetMostRecentTimestamp sets field value
func (o *Notification) SetMostRecentTimestamp(v time.Time) {
	o.MostRecentTimestamp = v
}

// GetType returns the Type field value
func (o *Notification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Notification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Notification) SetType(v string) {
	o.Type = v
}

// GetFollows returns the Follows field value if set, zero value otherwise.
func (o *Notification) GetFollows() []Follow {
	if o == nil || IsNil(o.Follows) {
		var ret []Follow
		return ret
	}
	return o.Follows
}

// GetFollowsOk returns a tuple with the Follows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetFollowsOk() ([]Follow, bool) {
	if o == nil || IsNil(o.Follows) {
		return nil, false
	}
	return o.Follows, true
}

// HasFollows returns a boolean if a field has been set.
func (o *Notification) HasFollows() bool {
	if o != nil && !IsNil(o.Follows) {
		return true
	}

	return false
}

// SetFollows gets a reference to the given []Follow and assigns it to the Follows field.
func (o *Notification) SetFollows(v []Follow) {
	o.Follows = v
}

// GetCast returns the Cast field value if set, zero value otherwise.
func (o *Notification) GetCast() CastWithInteractions {
	if o == nil || IsNil(o.Cast) {
		var ret CastWithInteractions
		return ret
	}
	return *o.Cast
}

// GetCastOk returns a tuple with the Cast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetCastOk() (*CastWithInteractions, bool) {
	if o == nil || IsNil(o.Cast) {
		return nil, false
	}
	return o.Cast, true
}

// HasCast returns a boolean if a field has been set.
func (o *Notification) HasCast() bool {
	if o != nil && !IsNil(o.Cast) {
		return true
	}

	return false
}

// SetCast gets a reference to the given CastWithInteractions and assigns it to the Cast field.
func (o *Notification) SetCast(v CastWithInteractions) {
	o.Cast = &v
}

// GetReactions returns the Reactions field value if set, zero value otherwise.
func (o *Notification) GetReactions() []Reactions {
	if o == nil || IsNil(o.Reactions) {
		var ret []Reactions
		return ret
	}
	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetReactionsOk() ([]Reactions, bool) {
	if o == nil || IsNil(o.Reactions) {
		return nil, false
	}
	return o.Reactions, true
}

// HasReactions returns a boolean if a field has been set.
func (o *Notification) HasReactions() bool {
	if o != nil && !IsNil(o.Reactions) {
		return true
	}

	return false
}

// SetReactions gets a reference to the given []Reactions and assigns it to the Reactions field.
func (o *Notification) SetReactions(v []Reactions) {
	o.Reactions = v
}

func (o Notification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Notification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["most_recent_timestamp"] = o.MostRecentTimestamp
	toSerialize["type"] = o.Type
	if !IsNil(o.Follows) {
		toSerialize["follows"] = o.Follows
	}
	if !IsNil(o.Cast) {
		toSerialize["cast"] = o.Cast
	}
	if !IsNil(o.Reactions) {
		toSerialize["reactions"] = o.Reactions
	}
	return toSerialize, nil
}

func (o *Notification) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"most_recent_timestamp",
		"type",
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

	varNotification := _Notification{}

	err = json.Unmarshal(bytes, &varNotification)

	if err != nil {
		return err
	}

	*o = Notification(varNotification)

	return err
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


