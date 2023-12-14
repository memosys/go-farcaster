# Reactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Cast** | [**ReactionsCast**](ReactionsCast.md) |  | 
**User** | [**User**](User.md) |  | 

## Methods

### NewReactions

`func NewReactions(object string, cast ReactionsCast, user User, ) *Reactions`

NewReactions instantiates a new Reactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionsWithDefaults

`func NewReactionsWithDefaults() *Reactions`

NewReactionsWithDefaults instantiates a new Reactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Reactions) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Reactions) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Reactions) SetObject(v string)`

SetObject sets Object field to given value.


### GetCast

`func (o *Reactions) GetCast() ReactionsCast`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *Reactions) GetCastOk() (*ReactionsCast, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *Reactions) SetCast(v ReactionsCast)`

SetCast sets Cast field to given value.


### GetUser

`func (o *Reactions) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Reactions) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Reactions) SetUser(v User)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


