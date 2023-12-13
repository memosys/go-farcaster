# \UserAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FarcasterUserVerificationDelete**](UserAPI.md#FarcasterUserVerificationDelete) | **Delete** /farcaster/user/verification | Removes verification for an eth address for the user
[**FarcasterUserVerificationPost**](UserAPI.md#FarcasterUserVerificationPost) | **Post** /farcaster/user/verification | Adds verification for an eth address for the user
[**FollowUser**](UserAPI.md#FollowUser) | **Post** /farcaster/user/follow | Follow a user
[**LookupUserByCustodyAddress**](UserAPI.md#LookupUserByCustodyAddress) | **Get** /farcaster/user/custody-address | Lookup a user by custody-address
[**UnfollowUser**](UserAPI.md#UnfollowUser) | **Delete** /farcaster/user/follow | Unfollow a user
[**UpdateUser**](UserAPI.md#UpdateUser) | **Patch** /farcaster/user | Update user profile
[**UserBulk**](UserAPI.md#UserBulk) | **Get** /farcaster/user/bulk | Fetches information about multiple users based on FIDs
[**UserSearch**](UserAPI.md#UserSearch) | **Get** /farcaster/user/search | Search for Usernames



## FarcasterUserVerificationDelete

> OperationResponse FarcasterUserVerificationDelete(ctx).ApiKey(apiKey).RemoveVerificationReqBody(removeVerificationReqBody).Execute()

Removes verification for an eth address for the user



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
    removeVerificationReqBody := *openapiclient.NewRemoveVerificationReqBody("SignerUuid_example", "Address_example") // RemoveVerificationReqBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.FarcasterUserVerificationDelete(context.Background()).ApiKey(apiKey).RemoveVerificationReqBody(removeVerificationReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FarcasterUserVerificationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FarcasterUserVerificationDelete`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.FarcasterUserVerificationDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFarcasterUserVerificationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **removeVerificationReqBody** | [**RemoveVerificationReqBody**](RemoveVerificationReqBody.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FarcasterUserVerificationPost

> OperationResponse FarcasterUserVerificationPost(ctx).ApiKey(apiKey).AddVerificationReqBody(addVerificationReqBody).Execute()

Adds verification for an eth address for the user



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
    addVerificationReqBody := *openapiclient.NewAddVerificationReqBody("SignerUuid_example", "Address_example", "BlockHash_example", "EthSignature_example") // AddVerificationReqBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.FarcasterUserVerificationPost(context.Background()).ApiKey(apiKey).AddVerificationReqBody(addVerificationReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FarcasterUserVerificationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FarcasterUserVerificationPost`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.FarcasterUserVerificationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFarcasterUserVerificationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **addVerificationReqBody** | [**AddVerificationReqBody**](AddVerificationReqBody.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowUser

> BulkFollowResponse FollowUser(ctx).ApiKey(apiKey).FollowReqBody(followReqBody).Execute()

Follow a user



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
    followReqBody := *openapiclient.NewFollowReqBody("SignerUuid_example", []int32{int32(123)}) // FollowReqBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.FollowUser(context.Background()).ApiKey(apiKey).FollowReqBody(followReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FollowUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FollowUser`: BulkFollowResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.FollowUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFollowUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **followReqBody** | [**FollowReqBody**](FollowReqBody.md) |  | 

### Return type

[**BulkFollowResponse**](BulkFollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupUserByCustodyAddress

> UserResponse LookupUserByCustodyAddress(ctx).ApiKey(apiKey).CustodyAddress(custodyAddress).Execute()

Lookup a user by custody-address



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
    custodyAddress := "0xd1b702203b1b3b641a699997746bd4a12d157909" // string | Custody Address associated with mnemonic

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.LookupUserByCustodyAddress(context.Background()).ApiKey(apiKey).CustodyAddress(custodyAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.LookupUserByCustodyAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupUserByCustodyAddress`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.LookupUserByCustodyAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupUserByCustodyAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **custodyAddress** | **string** | Custody Address associated with mnemonic | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfollowUser

> BulkFollowResponse UnfollowUser(ctx).ApiKey(apiKey).FollowReqBody(followReqBody).Execute()

Unfollow a user



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
    followReqBody := *openapiclient.NewFollowReqBody("SignerUuid_example", []int32{int32(123)}) // FollowReqBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UnfollowUser(context.Background()).ApiKey(apiKey).FollowReqBody(followReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UnfollowUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnfollowUser`: BulkFollowResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UnfollowUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **followReqBody** | [**FollowReqBody**](FollowReqBody.md) |  | 

### Return type

[**BulkFollowResponse**](BulkFollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> OperationResponse UpdateUser(ctx).ApiKey(apiKey).UpdateUserReqBody(updateUserReqBody).Execute()

Update user profile



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
    updateUserReqBody := *openapiclient.NewUpdateUserReqBody("SignerUuid_example") // UpdateUserReqBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UpdateUser(context.Background()).ApiKey(apiKey).UpdateUserReqBody(updateUserReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **updateUserReqBody** | [**UpdateUserReqBody**](UpdateUserReqBody.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserBulk

> BulkUsersResponse UserBulk(ctx).ApiKey(apiKey).Fids(fids).ViewerFid(viewerFid).Execute()

Fetches information about multiple users based on FIDs



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
    fids := "194, 191, 6131" // string | 
    viewerFid := int32(3) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UserBulk(context.Background()).ApiKey(apiKey).Fids(fids).ViewerFid(viewerFid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserBulk`: BulkUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **fids** | **string** |  | 
 **viewerFid** | **int32** |  | 

### Return type

[**BulkUsersResponse**](BulkUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSearch

> UserSearchResponse UserSearch(ctx).ApiKey(apiKey).Q(q).ViewerFid(viewerFid).Execute()

Search for Usernames



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
    q := "r" // string | 
    viewerFid := int32(3) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UserSearch(context.Background()).ApiKey(apiKey).Q(q).ViewerFid(viewerFid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSearch`: UserSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **q** | **string** |  | 
 **viewerFid** | **int32** |  | 

### Return type

[**UserSearchResponse**](UserSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

