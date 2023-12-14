# \SignerAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSigner**](SignerAPI.md#CreateSigner) | **Post** /farcaster/signer | Creates a signer and returns the signer status
[**RegisterSignedKey**](SignerAPI.md#RegisterSignedKey) | **Post** /farcaster/signer/signed_key | Register Signed Key
[**Signer**](SignerAPI.md#Signer) | **Get** /farcaster/signer | Fetches the status of a signer



## CreateSigner

> Signer CreateSigner(ctx).ApiKey(apiKey).Execute()

Creates a signer and returns the signer status



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignerAPI.CreateSigner(context.Background()).ApiKey(apiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.CreateSigner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSigner`: Signer
    fmt.Fprintf(os.Stdout, "Response from `SignerAPI.CreateSigner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]

### Return type

[**Signer**](Signer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSignedKey

> Signer RegisterSignedKey(ctx).ApiKey(apiKey).RegisterSignerKeyReqBody(registerSignerKeyReqBody).Execute()

Register Signed Key



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
    registerSignerKeyReqBody := *openapiclient.NewRegisterSignerKeyReqBody("SignerUuid_example", "Signature_example", int32(123), int32(123)) // RegisterSignerKeyReqBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignerAPI.RegisterSignedKey(context.Background()).ApiKey(apiKey).RegisterSignerKeyReqBody(registerSignerKeyReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.RegisterSignedKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterSignedKey`: Signer
    fmt.Fprintf(os.Stdout, "Response from `SignerAPI.RegisterSignedKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSignedKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **registerSignerKeyReqBody** | [**RegisterSignerKeyReqBody**](RegisterSignerKeyReqBody.md) |  | 

### Return type

[**Signer**](Signer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Signer

> Signer Signer(ctx).ApiKey(apiKey).SignerUuid(signerUuid).Execute()

Fetches the status of a signer



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
    signerUuid := "19d0c5fd-9b33-4a48-a0e2-bc7b0555baec" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignerAPI.Signer(context.Background()).ApiKey(apiKey).SignerUuid(signerUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.Signer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Signer`: Signer
    fmt.Fprintf(os.Stdout, "Response from `SignerAPI.Signer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **signerUuid** | **string** |  | 

### Return type

[**Signer**](Signer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

