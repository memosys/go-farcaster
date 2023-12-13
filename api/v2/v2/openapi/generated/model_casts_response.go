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

// checks if the CastsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastsResponse{}

// CastsResponse struct for CastsResponse
type CastsResponse struct {
	Result CastsResponseResult `json:"result"`
}

type _CastsResponse CastsResponse

// NewCastsResponse instantiates a new CastsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastsResponse(result CastsResponseResult) *CastsResponse {
	this := CastsResponse{}
	this.Result = result
	return &this
}

// NewCastsResponseWithDefaults instantiates a new CastsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastsResponseWithDefaults() *CastsResponse {
	this := CastsResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *CastsResponse) GetResult() CastsResponseResult {
	if o == nil {
		var ret CastsResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CastsResponse) GetResultOk() (*CastsResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CastsResponse) SetResult(v CastsResponseResult) {
	o.Result = v
}

func (o CastsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *CastsResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varCastsResponse := _CastsResponse{}

	err = json.Unmarshal(bytes, &varCastsResponse)

	if err != nil {
		return err
	}

	*o = CastsResponse(varCastsResponse)

	return err
}

type NullableCastsResponse struct {
	value *CastsResponse
	isSet bool
}

func (v NullableCastsResponse) Get() *CastsResponse {
	return v.value
}

func (v *NullableCastsResponse) Set(val *CastsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCastsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCastsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastsResponse(val *CastsResponse) *NullableCastsResponse {
	return &NullableCastsResponse{value: val, isSet: true}
}

func (v NullableCastsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


