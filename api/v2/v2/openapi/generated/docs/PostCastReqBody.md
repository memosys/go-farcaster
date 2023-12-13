# PostCastReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer | 
**Text** | Pointer to **string** |  | [optional] 
**Embeds** | Pointer to [**[]EmbeddedCast**](EmbeddedCast.md) |  | [optional] 
**Parent** | Pointer to **string** | Parent URL or Cast Hash | [optional] 

## Methods

### NewPostCastReqBody

`func NewPostCastReqBody(signerUuid string, ) *PostCastReqBody`

NewPostCastReqBody instantiates a new PostCastReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCastReqBodyWithDefaults

`func NewPostCastReqBodyWithDefaults() *PostCastReqBody`

NewPostCastReqBodyWithDefaults instantiates a new PostCastReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *PostCastReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *PostCastReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *PostCastReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetText

`func (o *PostCastReqBody) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostCastReqBody) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostCastReqBody) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PostCastReqBody) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEmbeds

`func (o *PostCastReqBody) GetEmbeds() []EmbeddedCast`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *PostCastReqBody) GetEmbedsOk() (*[]EmbeddedCast, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *PostCastReqBody) SetEmbeds(v []EmbeddedCast)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *PostCastReqBody) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetParent

`func (o *PostCastReqBody) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PostCastReqBody) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PostCastReqBody) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PostCastReqBody) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


