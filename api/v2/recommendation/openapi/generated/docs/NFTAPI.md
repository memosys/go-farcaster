# \NFTAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRelevantMints**](NFTAPI.md#FetchRelevantMints) | **Get** /nft/relevant_mints | Relevant Mints for a User



## FetchRelevantMints

> RelevantMints FetchRelevantMints(ctx).ApiKey(apiKey).Address(address).ContractAddress(contractAddress).TokenId(tokenId).Execute()

Relevant Mints for a User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
    address := "0x5a927ac639636e534b678e81768ca19e2c6280b7" // string | 
    contractAddress := "0xe8e0e543a3dd32d366cb756fa4d112f30172bcb1" // string | 
    tokenId := "1" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFTAPI.FetchRelevantMints(context.Background()).ApiKey(apiKey).Address(address).ContractAddress(contractAddress).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTAPI.FetchRelevantMints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRelevantMints`: RelevantMints
    fmt.Fprintf(os.Stdout, "Response from `NFTAPI.FetchRelevantMints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchRelevantMintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **address** | **string** |  | 
 **contractAddress** | **string** |  | 
 **tokenId** | **string** |  | 

### Return type

[**RelevantMints**](RelevantMints.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

