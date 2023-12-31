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

// checks if the CastReactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastReactionsResponse{}

// CastReactionsResponse struct for CastReactionsResponse
type CastReactionsResponse struct {
	Result CastReactionsResponseResult `json:"result"`
}

type _CastReactionsResponse CastReactionsResponse

// NewCastReactionsResponse instantiates a new CastReactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastReactionsResponse(result CastReactionsResponseResult) *CastReactionsResponse {
	this := CastReactionsResponse{}
	this.Result = result
	return &this
}

// NewCastReactionsResponseWithDefaults instantiates a new CastReactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastReactionsResponseWithDefaults() *CastReactionsResponse {
	this := CastReactionsResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *CastReactionsResponse) GetResult() CastReactionsResponseResult {
	if o == nil {
		var ret CastReactionsResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CastReactionsResponse) GetResultOk() (*CastReactionsResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CastReactionsResponse) SetResult(v CastReactionsResponseResult) {
	o.Result = v
}

func (o CastReactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastReactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *CastReactionsResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varCastReactionsResponse := _CastReactionsResponse{}

	err = json.Unmarshal(bytes, &varCastReactionsResponse)

	if err != nil {
		return err
	}

	*o = CastReactionsResponse(varCastReactionsResponse)

	return err
}

type NullableCastReactionsResponse struct {
	value *CastReactionsResponse
	isSet bool
}

func (v NullableCastReactionsResponse) Get() *CastReactionsResponse {
	return v.value
}

func (v *NullableCastReactionsResponse) Set(val *CastReactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCastReactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCastReactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastReactionsResponse(val *CastReactionsResponse) *NullableCastReactionsResponse {
	return &NullableCastReactionsResponse{value: val, isSet: true}
}

func (v NullableCastReactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastReactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


