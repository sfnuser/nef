# \AsSessionWithQoSAPISCSASLevelGETOperationApi

All URIs are relative to *https://example.com/3gpp-as-session-with-qos/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScsAsIdSubscriptionsGet**](AsSessionWithQoSAPISCSASLevelGETOperationApi.md#ScsAsIdSubscriptionsGet) | **Get** /{scsAsId}/subscriptions | read all of the active subscriptions for the SCS/AS



## ScsAsIdSubscriptionsGet

> []AsSessionWithQoSSubscription ScsAsIdSubscriptionsGet(ctx, scsAsId).Execute()

read all of the active subscriptions for the SCS/AS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsSessionWithQoSAPISCSASLevelGETOperationApi.ScsAsIdSubscriptionsGet(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsSessionWithQoSAPISCSASLevelGETOperationApi.ScsAsIdSubscriptionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScsAsIdSubscriptionsGet`: []AsSessionWithQoSSubscription
    fmt.Fprintf(os.Stdout, "Response from `AsSessionWithQoSAPISCSASLevelGETOperationApi.ScsAsIdSubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiScsAsIdSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

