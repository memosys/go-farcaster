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

// checks if the FeedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedResponse{}

// FeedResponse struct for FeedResponse
type FeedResponse struct {
	Casts []CastWithInteractions `json:"casts"`
	Next NextCursor `json:"next"`
}

type _FeedResponse FeedResponse

// NewFeedResponse instantiates a new FeedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedResponse(casts []CastWithInteractions, next NextCursor) *FeedResponse {
	this := FeedResponse{}
	this.Casts = casts
	this.Next = next
	return &this
}

// NewFeedResponseWithDefaults instantiates a new FeedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedResponseWithDefaults() *FeedResponse {
	this := FeedResponse{}
	return &this
}

// GetCasts returns the Casts field value
func (o *FeedResponse) GetCasts() []CastWithInteractions {
	if o == nil {
		var ret []CastWithInteractions
		return ret
	}

	return o.Casts
}

// GetCastsOk returns a tuple with the Casts field value
// and a boolean to check if the value has been set.
func (o *FeedResponse) GetCastsOk() ([]CastWithInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Casts, true
}

// SetCasts sets field value
func (o *FeedResponse) SetCasts(v []CastWithInteractions) {
	o.Casts = v
}

// GetNext returns the Next field value
func (o *FeedResponse) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *FeedResponse) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *FeedResponse) SetNext(v NextCursor) {
	o.Next = v
}

func (o FeedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["casts"] = o.Casts
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *FeedResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"casts",
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

	varFeedResponse := _FeedResponse{}

	err = json.Unmarshal(bytes, &varFeedResponse)

	if err != nil {
		return err
	}

	*o = FeedResponse(varFeedResponse)

	return err
}

type NullableFeedResponse struct {
	value *FeedResponse
	isSet bool
}

func (v NullableFeedResponse) Get() *FeedResponse {
	return v.value
}

func (v *NullableFeedResponse) Set(val *FeedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedResponse(val *FeedResponse) *NullableFeedResponse {
	return &NullableFeedResponse{value: val, isSet: true}
}

func (v NullableFeedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


