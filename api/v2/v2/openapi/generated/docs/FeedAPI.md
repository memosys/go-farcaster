# \FeedAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Feed**](FeedAPI.md#Feed) | **Get** /farcaster/feed | Retrieve casts based on filters



## Feed

> FeedResponse Feed(ctx).ApiKey(apiKey).FeedType(feedType).FilterType(filterType).Fid(fid).Fids(fids).ParentUrl(parentUrl).EmbedUrl(embedUrl).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).Execute()

Retrieve casts based on filters

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
    feedType := openapiclient.FeedType("following") // FeedType | Defaults to following (requires fid or address). If set to filter (requires filter_type)
    filterType := openapiclient.FilterType("fids") // FilterType | Used when feed_type=filter. Can be set to fids (requires fids) or parent_url (requires parent_url) (optional)
    fid := int32(56) // int32 | (Optional) fid of user whose feed you want to create. By default, the API expects this field, except if you pass a filter_type (optional)
    fids := "3,2,194" // string | Used when filter_type=fids . Create a feed based on a list of fids. Max array size is 250. Requires feed_type and filter_type. (optional)
    parentUrl := "parentUrl_example" // string | Used when filter_type=parent_url can be used to fetch content under any parent url e.g. FIP-2 channels on Warpcast. Requires feed_type and filter_type (optional)
    embedUrl := "embedUrl_example" // string | Used when filter_type=embed_url can be used to fetch all casts with an embed url that contains embed_url. Requires feed_type and filter_type (optional)
    withRecasts := true // bool | Include recasts in the response, true by default (optional) (default to true)
    limit := int32(56) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
    cursor := "cursor_example" // string | Pagination cursor. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedAPI.Feed(context.Background()).ApiKey(apiKey).FeedType(feedType).FilterType(filterType).Fid(fid).Fids(fids).ParentUrl(parentUrl).EmbedUrl(embedUrl).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.Feed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Feed`: FeedResponse
    fmt.Fprintf(os.Stdout, "Response from `FeedAPI.Feed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]
 **feedType** | [**FeedType**](FeedType.md) | Defaults to following (requires fid or address). If set to filter (requires filter_type) | 
 **filterType** | [**FilterType**](FilterType.md) | Used when feed_type&#x3D;filter. Can be set to fids (requires fids) or parent_url (requires parent_url) | 
 **fid** | **int32** | (Optional) fid of user whose feed you want to create. By default, the API expects this field, except if you pass a filter_type | 
 **fids** | **string** | Used when filter_type&#x3D;fids . Create a feed based on a list of fids. Max array size is 250. Requires feed_type and filter_type. | 
 **parentUrl** | **string** | Used when filter_type&#x3D;parent_url can be used to fetch content under any parent url e.g. FIP-2 channels on Warpcast. Requires feed_type and filter_type | 
 **embedUrl** | **string** | Used when filter_type&#x3D;embed_url can be used to fetch all casts with an embed url that contains embed_url. Requires feed_type and filter_type | 
 **withRecasts** | **bool** | Include recasts in the response, true by default | [default to true]
 **limit** | **int32** | Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

