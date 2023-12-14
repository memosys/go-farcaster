/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv1

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Recaster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recaster{}

// Recaster struct for Recaster
type Recaster struct {
	// The unique identifier of the recaster.
	Fid int32 `json:"fid"`
	// The username of the recaster.
	Username string `json:"username"`
	// The display name of the recaster.
	DisplayName string `json:"displayName"`
	Pfp RecasterPfp `json:"pfp"`
	Profile RecasterProfile `json:"profile"`
	// The number of followers the recaster has.
	FollowerCount int32 `json:"followerCount"`
	// The number of users the recaster is following.
	FollowingCount int32 `json:"followingCount"`
	Timestamp time.Time `json:"timestamp"`
	ViewerContext *RecasterViewerContext `json:"viewerContext,omitempty"`
}

type _Recaster Recaster

// NewRecaster instantiates a new Recaster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecaster(fid int32, username string, displayName string, pfp RecasterPfp, profile RecasterProfile, followerCount int32, followingCount int32, timestamp time.Time) *Recaster {
	this := Recaster{}
	this.Fid = fid
	this.Username = username
	this.DisplayName = displayName
	this.Pfp = pfp
	this.Profile = profile
	this.FollowerCount = followerCount
	this.FollowingCount = followingCount
	this.Timestamp = timestamp
	return &this
}

// NewRecasterWithDefaults instantiates a new Recaster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecasterWithDefaults() *Recaster {
	this := Recaster{}
	return &this
}

// GetFid returns the Fid field value
func (o *Recaster) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *Recaster) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *Recaster) SetFid(v int32) {
	o.Fid = v
}

// GetUsername returns the Username field value
func (o *Recaster) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Recaster) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Recaster) SetUsername(v string) {
	o.Username = v
}

// GetDisplayName returns the DisplayName field value
func (o *Recaster) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Recaster) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Recaster) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPfp returns the Pfp field value
func (o *Recaster) GetPfp() RecasterPfp {
	if o == nil {
		var ret RecasterPfp
		return ret
	}

	return o.Pfp
}

// GetPfpOk returns a tuple with the Pfp field value
// and a boolean to check if the value has been set.
func (o *Recaster) GetPfpOk() (*RecasterPfp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pfp, true
}

// SetPfp sets field value
func (o *Recaster) SetPfp(v RecasterPfp) {
	o.Pfp = v
}

// GetProfile returns the Profile field value
func (o *Recaster) GetProfile() RecasterProfile {
	if o == nil {
		var ret RecasterProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *Recaster) GetProfileOk() (*RecasterProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *Recaster) SetProfile(v RecasterProfile) {
	o.Profile = v
}

// GetFollowerCount returns the FollowerCount field value
func (o *Recaster) GetFollowerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowerCount
}

// GetFollowerCountOk returns a tuple with the FollowerCount field value
// and a boolean to check if the value has been set.
func (o *Recaster) GetFollowerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowerCount, true
}

// SetFollowerCount sets field value
func (o *Recaster) SetFollowerCount(v int32) {
	o.FollowerCount = v
}

// GetFollowingCount returns the FollowingCount field value
func (o *Recaster) GetFollowingCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowingCount
}

// GetFollowingCountOk returns a tuple with the FollowingCount field value
// and a boolean to check if the value has been set.
func (o *Recaster) GetFollowingCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingCount, true
}

// SetFollowingCount sets field value
func (o *Recaster) SetFollowingCount(v int32) {
	o.FollowingCount = v
}

// GetTimestamp returns the Timestamp field value
func (o *Recaster) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Recaster) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Recaster) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *Recaster) GetViewerContext() RecasterViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret RecasterViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recaster) GetViewerContextOk() (*RecasterViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *Recaster) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given RecasterViewerContext and assigns it to the ViewerContext field.
func (o *Recaster) SetViewerContext(v RecasterViewerContext) {
	o.ViewerContext = &v
}

func (o Recaster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Recaster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["username"] = o.Username
	toSerialize["displayName"] = o.DisplayName
	toSerialize["pfp"] = o.Pfp
	toSerialize["profile"] = o.Profile
	toSerialize["followerCount"] = o.FollowerCount
	toSerialize["followingCount"] = o.FollowingCount
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.ViewerContext) {
		toSerialize["viewerContext"] = o.ViewerContext
	}
	return toSerialize, nil
}

func (o *Recaster) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"username",
		"displayName",
		"pfp",
		"profile",
		"followerCount",
		"followingCount",
		"timestamp",
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

	varRecaster := _Recaster{}

	err = json.Unmarshal(bytes, &varRecaster)

	if err != nil {
		return err
	}

	*o = Recaster(varRecaster)

	return err
}

type NullableRecaster struct {
	value *Recaster
	isSet bool
}

func (v NullableRecaster) Get() *Recaster {
	return v.value
}

func (v *NullableRecaster) Set(val *Recaster) {
	v.value = val
	v.isSet = true
}

func (v NullableRecaster) IsSet() bool {
	return v.isSet
}

func (v *NullableRecaster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecaster(val *Recaster) *NullableRecaster {
	return &NullableRecaster{value: val, isSet: true}
}

func (v NullableRecaster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecaster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


