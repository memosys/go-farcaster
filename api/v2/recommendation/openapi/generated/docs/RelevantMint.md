# RelevantMint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | **string** | Ethereum address | 
**TokenId** | **string** |  | 
**BlockNumber** | **int32** |  | 
**TxHash** | **string** |  | 
**Minter** | [**User**](User.md) |  | 

## Methods

### NewRelevantMint

`func NewRelevantMint(contractAddress string, tokenId string, blockNumber int32, txHash string, minter User, ) *RelevantMint`

NewRelevantMint instantiates a new RelevantMint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelevantMintWithDefaults

`func NewRelevantMintWithDefaults() *RelevantMint`

NewRelevantMintWithDefaults instantiates a new RelevantMint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *RelevantMint) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *RelevantMint) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *RelevantMint) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetTokenId

`func (o *RelevantMint) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *RelevantMint) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *RelevantMint) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetBlockNumber

`func (o *RelevantMint) GetBlockNumber() int32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *RelevantMint) GetBlockNumberOk() (*int32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *RelevantMint) SetBlockNumber(v int32)`

SetBlockNumber sets BlockNumber field to given value.


### GetTxHash

`func (o *RelevantMint) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *RelevantMint) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *RelevantMint) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetMinter

`func (o *RelevantMint) GetMinter() User`

GetMinter returns the Minter field if non-nil, zero value otherwise.

### GetMinterOk

`func (o *RelevantMint) GetMinterOk() (*User, bool)`

GetMinterOk returns a tuple with the Minter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinter

`func (o *RelevantMint) SetMinter(v User)`

SetMinter sets Minter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


