# \FollowsAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RelevantFollowers**](FollowsAPI.md#RelevantFollowers) | **Get** /farcaster/followers/relevant | Retrieve relevant followers for a given user



## RelevantFollowers

> RelevantFollowersResponse RelevantFollowers(ctx).ApiKey(apiKey).TargetFid(targetFid).ViewerFid(viewerFid).Execute()

Retrieve relevant followers for a given user



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
    targetFid := int32(56) // int32 | User who's profile you are looking at
    viewerFid := int32(56) // int32 | Viewer who's looking at the profile

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FollowsAPI.RelevantFollowers(context.Background()).ApiKey(apiKey).TargetFid(targetFid).ViewerFid(viewerFid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.RelevantFollowers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RelevantFollowers`: RelevantFollowersResponse
    fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.RelevantFollowers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelevantFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **targetFid** | **int32** | User who&#39;s profile you are looking at | 
 **viewerFid** | **int32** | Viewer who&#39;s looking at the profile | 

### Return type

[**RelevantFollowersResponse**](RelevantFollowersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

