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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	// User identifier (unsigned integer)
	Fid int32 `json:"fid"`
	// The username of the user.
	Username string `json:"username"`
	// Custody Address of the user.
	CustodyAddress string `json:"custodyAddress"`
	// The display of the reactor.
	DisplayName string `json:"displayName"`
	Pfp UserPfp `json:"pfp"`
	Profile UserProfile `json:"profile"`
	// The number of followers the user has.
	FollowerCount int32 `json:"followerCount"`
	// The number of users the user is following.
	FollowingCount int32 `json:"followingCount"`
	Verifications []string `json:"verifications"`
	ActiveStatus ActiveStatus `json:"activeStatus"`
	ViewerContext *ViewerContext `json:"viewerContext,omitempty"`
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(fid int32, username string, custodyAddress string, displayName string, pfp UserPfp, profile UserProfile, followerCount int32, followingCount int32, verifications []string, activeStatus ActiveStatus) *User {
	this := User{}
	this.Fid = fid
	this.Username = username
	this.CustodyAddress = custodyAddress
	this.DisplayName = displayName
	this.Pfp = pfp
	this.Profile = profile
	this.FollowerCount = followerCount
	this.FollowingCount = followingCount
	this.Verifications = verifications
	this.ActiveStatus = activeStatus
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	var fid int32 = 3
	this.Fid = fid
	return &this
}

// GetFid returns the Fid field value
func (o *User) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *User) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *User) SetFid(v int32) {
	o.Fid = v
}

// GetUsername returns the Username field value
func (o *User) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *User) SetUsername(v string) {
	o.Username = v
}

// GetCustodyAddress returns the CustodyAddress field value
func (o *User) GetCustodyAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustodyAddress
}

// GetCustodyAddressOk returns a tuple with the CustodyAddress field value
// and a boolean to check if the value has been set.
func (o *User) GetCustodyAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustodyAddress, true
}

// SetCustodyAddress sets field value
func (o *User) SetCustodyAddress(v string) {
	o.CustodyAddress = v
}

// GetDisplayName returns the DisplayName field value
func (o *User) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *User) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *User) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPfp returns the Pfp field value
func (o *User) GetPfp() UserPfp {
	if o == nil {
		var ret UserPfp
		return ret
	}

	return o.Pfp
}

// GetPfpOk returns a tuple with the Pfp field value
// and a boolean to check if the value has been set.
func (o *User) GetPfpOk() (*UserPfp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pfp, true
}

// SetPfp sets field value
func (o *User) SetPfp(v UserPfp) {
	o.Pfp = v
}

// GetProfile returns the Profile field value
func (o *User) GetProfile() UserProfile {
	if o == nil {
		var ret UserProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *User) GetProfileOk() (*UserProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *User) SetProfile(v UserProfile) {
	o.Profile = v
}

// GetFollowerCount returns the FollowerCount field value
func (o *User) GetFollowerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowerCount
}

// GetFollowerCountOk returns a tuple with the FollowerCount field value
// and a boolean to check if the value has been set.
func (o *User) GetFollowerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowerCount, true
}

// SetFollowerCount sets field value
func (o *User) SetFollowerCount(v int32) {
	o.FollowerCount = v
}

// GetFollowingCount returns the FollowingCount field value
func (o *User) GetFollowingCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowingCount
}

// GetFollowingCountOk returns a tuple with the FollowingCount field value
// and a boolean to check if the value has been set.
func (o *User) GetFollowingCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingCount, true
}

// SetFollowingCount sets field value
func (o *User) SetFollowingCount(v int32) {
	o.FollowingCount = v
}

// GetVerifications returns the Verifications field value
func (o *User) GetVerifications() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value
// and a boolean to check if the value has been set.
func (o *User) GetVerificationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verifications, true
}

// SetVerifications sets field value
func (o *User) SetVerifications(v []string) {
	o.Verifications = v
}

// GetActiveStatus returns the ActiveStatus field value
func (o *User) GetActiveStatus() ActiveStatus {
	if o == nil {
		var ret ActiveStatus
		return ret
	}

	return o.ActiveStatus
}

// GetActiveStatusOk returns a tuple with the ActiveStatus field value
// and a boolean to check if the value has been set.
func (o *User) GetActiveStatusOk() (*ActiveStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveStatus, true
}

// SetActiveStatus sets field value
func (o *User) SetActiveStatus(v ActiveStatus) {
	o.ActiveStatus = v
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *User) GetViewerContext() ViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret ViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetViewerContextOk() (*ViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *User) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given ViewerContext and assigns it to the ViewerContext field.
func (o *User) SetViewerContext(v ViewerContext) {
	o.ViewerContext = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["username"] = o.Username
	toSerialize["custodyAddress"] = o.CustodyAddress
	toSerialize["displayName"] = o.DisplayName
	toSerialize["pfp"] = o.Pfp
	toSerialize["profile"] = o.Profile
	toSerialize["followerCount"] = o.FollowerCount
	toSerialize["followingCount"] = o.FollowingCount
	toSerialize["verifications"] = o.Verifications
	toSerialize["activeStatus"] = o.ActiveStatus
	if !IsNil(o.ViewerContext) {
		toSerialize["viewerContext"] = o.ViewerContext
	}
	return toSerialize, nil
}

func (o *User) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"username",
		"custodyAddress",
		"displayName",
		"pfp",
		"profile",
		"followerCount",
		"followingCount",
		"verifications",
		"activeStatus",
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

	varUser := _User{}

	err = json.Unmarshal(bytes, &varUser)

	if err != nil {
		return err
	}

	*o = User(varUser)

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


