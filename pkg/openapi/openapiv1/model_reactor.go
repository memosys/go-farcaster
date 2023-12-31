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

// checks if the Reactor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reactor{}

// Reactor struct for Reactor
type Reactor struct {
	// The unique identifier of the reactor.
	Fid int32 `json:"fid"`
	// The username of the reactor.
	Username string `json:"username"`
	// The display name of the reactor.
	DisplayName string `json:"displayName"`
	Pfp ReactorPfp `json:"pfp"`
	// The number of followers the reactor has.
	FollowerCount int32 `json:"followerCount"`
	// The number of users the reactor is following.
	FollowingCount int32 `json:"followingCount"`
	ViewerContext *ReactorViewerContext `json:"viewerContext,omitempty"`
}

type _Reactor Reactor

// NewReactor instantiates a new Reactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactor(fid int32, username string, displayName string, pfp ReactorPfp, followerCount int32, followingCount int32) *Reactor {
	this := Reactor{}
	this.Fid = fid
	this.Username = username
	this.DisplayName = displayName
	this.Pfp = pfp
	this.FollowerCount = followerCount
	this.FollowingCount = followingCount
	return &this
}

// NewReactorWithDefaults instantiates a new Reactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactorWithDefaults() *Reactor {
	this := Reactor{}
	return &this
}

// GetFid returns the Fid field value
func (o *Reactor) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *Reactor) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *Reactor) SetFid(v int32) {
	o.Fid = v
}

// GetUsername returns the Username field value
func (o *Reactor) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Reactor) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Reactor) SetUsername(v string) {
	o.Username = v
}

// GetDisplayName returns the DisplayName field value
func (o *Reactor) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Reactor) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Reactor) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPfp returns the Pfp field value
func (o *Reactor) GetPfp() ReactorPfp {
	if o == nil {
		var ret ReactorPfp
		return ret
	}

	return o.Pfp
}

// GetPfpOk returns a tuple with the Pfp field value
// and a boolean to check if the value has been set.
func (o *Reactor) GetPfpOk() (*ReactorPfp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pfp, true
}

// SetPfp sets field value
func (o *Reactor) SetPfp(v ReactorPfp) {
	o.Pfp = v
}

// GetFollowerCount returns the FollowerCount field value
func (o *Reactor) GetFollowerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowerCount
}

// GetFollowerCountOk returns a tuple with the FollowerCount field value
// and a boolean to check if the value has been set.
func (o *Reactor) GetFollowerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowerCount, true
}

// SetFollowerCount sets field value
func (o *Reactor) SetFollowerCount(v int32) {
	o.FollowerCount = v
}

// GetFollowingCount returns the FollowingCount field value
func (o *Reactor) GetFollowingCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowingCount
}

// GetFollowingCountOk returns a tuple with the FollowingCount field value
// and a boolean to check if the value has been set.
func (o *Reactor) GetFollowingCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingCount, true
}

// SetFollowingCount sets field value
func (o *Reactor) SetFollowingCount(v int32) {
	o.FollowingCount = v
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *Reactor) GetViewerContext() ReactorViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret ReactorViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reactor) GetViewerContextOk() (*ReactorViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *Reactor) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given ReactorViewerContext and assigns it to the ViewerContext field.
func (o *Reactor) SetViewerContext(v ReactorViewerContext) {
	o.ViewerContext = &v
}

func (o Reactor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reactor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["username"] = o.Username
	toSerialize["displayName"] = o.DisplayName
	toSerialize["pfp"] = o.Pfp
	toSerialize["followerCount"] = o.FollowerCount
	toSerialize["followingCount"] = o.FollowingCount
	if !IsNil(o.ViewerContext) {
		toSerialize["viewerContext"] = o.ViewerContext
	}
	return toSerialize, nil
}

func (o *Reactor) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"username",
		"displayName",
		"pfp",
		"followerCount",
		"followingCount",
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

	varReactor := _Reactor{}

	err = json.Unmarshal(bytes, &varReactor)

	if err != nil {
		return err
	}

	*o = Reactor(varReactor)

	return err
}

type NullableReactor struct {
	value *Reactor
	isSet bool
}

func (v NullableReactor) Get() *Reactor {
	return v.value
}

func (v *NullableReactor) Set(val *Reactor) {
	v.value = val
	v.isSet = true
}

func (v NullableReactor) IsSet() bool {
	return v.isSet
}

func (v *NullableReactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactor(val *Reactor) *NullableReactor {
	return &NullableReactor{value: val, isSet: true}
}

func (v NullableReactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


