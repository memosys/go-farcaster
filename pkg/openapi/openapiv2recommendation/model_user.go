/*
Recommendation API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv2recommendation

import (
	"encoding/json"
	"fmt"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	Object string `json:"object"`
	// User identifier (unsigned integer)
	Fid int32 `json:"fid"`
	Username string `json:"username"`
	DisplayName string `json:"display_name"`
	// Ethereum address
	CustodyAddress *string `json:"custody_address,omitempty"`
	// The URL of the user's profile picture
	PfpUrl string `json:"pfp_url"`
	Profile UserProfile `json:"profile"`
	// The number of followers the user has.
	FollowerCount int32 `json:"follower_count"`
	// The number of users the user is following.
	FollowingCount int32 `json:"following_count"`
	Verifications []string `json:"verifications"`
	ActiveStatus ActiveStatus `json:"activeStatus"`
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(object string, fid int32, username string, displayName string, pfpUrl string, profile UserProfile, followerCount int32, followingCount int32, verifications []string, activeStatus ActiveStatus) *User {
	this := User{}
	this.Object = object
	this.Fid = fid
	this.Username = username
	this.DisplayName = displayName
	this.PfpUrl = pfpUrl
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

// GetObject returns the Object field value
func (o *User) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *User) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *User) SetObject(v string) {
	o.Object = v
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

// GetCustodyAddress returns the CustodyAddress field value if set, zero value otherwise.
func (o *User) GetCustodyAddress() string {
	if o == nil || IsNil(o.CustodyAddress) {
		var ret string
		return ret
	}
	return *o.CustodyAddress
}

// GetCustodyAddressOk returns a tuple with the CustodyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCustodyAddressOk() (*string, bool) {
	if o == nil || IsNil(o.CustodyAddress) {
		return nil, false
	}
	return o.CustodyAddress, true
}

// HasCustodyAddress returns a boolean if a field has been set.
func (o *User) HasCustodyAddress() bool {
	if o != nil && !IsNil(o.CustodyAddress) {
		return true
	}

	return false
}

// SetCustodyAddress gets a reference to the given string and assigns it to the CustodyAddress field.
func (o *User) SetCustodyAddress(v string) {
	o.CustodyAddress = &v
}

// GetPfpUrl returns the PfpUrl field value
func (o *User) GetPfpUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PfpUrl
}

// GetPfpUrlOk returns a tuple with the PfpUrl field value
// and a boolean to check if the value has been set.
func (o *User) GetPfpUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PfpUrl, true
}

// SetPfpUrl sets field value
func (o *User) SetPfpUrl(v string) {
	o.PfpUrl = v
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

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["fid"] = o.Fid
	toSerialize["username"] = o.Username
	toSerialize["display_name"] = o.DisplayName
	if !IsNil(o.CustodyAddress) {
		toSerialize["custody_address"] = o.CustodyAddress
	}
	toSerialize["pfp_url"] = o.PfpUrl
	toSerialize["profile"] = o.Profile
	toSerialize["follower_count"] = o.FollowerCount
	toSerialize["following_count"] = o.FollowingCount
	toSerialize["verifications"] = o.Verifications
	toSerialize["activeStatus"] = o.ActiveStatus
	return toSerialize, nil
}

func (o *User) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"fid",
		"username",
		"display_name",
		"pfp_url",
		"profile",
		"follower_count",
		"following_count",
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


