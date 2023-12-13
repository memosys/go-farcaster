# ReactionReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer | 
**ReactionType** | [**ReactionType**](ReactionType.md) |  | 
**Target** | **string** |  | 

## Methods

### NewReactionReqBody

`func NewReactionReqBody(signerUuid string, reactionType ReactionType, target string, ) *ReactionReqBody`

NewReactionReqBody instantiates a new ReactionReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionReqBodyWithDefaults

`func NewReactionReqBodyWithDefaults() *ReactionReqBody`

NewReactionReqBodyWithDefaults instantiates a new ReactionReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *ReactionReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *ReactionReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *ReactionReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetReactionType

`func (o *ReactionReqBody) GetReactionType() ReactionType`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *ReactionReqBody) GetReactionTypeOk() (*ReactionType, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *ReactionReqBody) SetReactionType(v ReactionType)`

SetReactionType sets ReactionType field to given value.


### GetTarget

`func (o *ReactionReqBody) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ReactionReqBody) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ReactionReqBody) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


