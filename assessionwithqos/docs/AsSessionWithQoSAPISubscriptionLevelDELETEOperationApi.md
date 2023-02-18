# \AsSessionWithQoSAPISubscriptionLevelDELETEOperationApi

All URIs are relative to *https://example.com/3gpp-as-session-with-qos/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScsAsIdSubscriptionsSubscriptionIdDelete**](AsSessionWithQoSAPISubscriptionLevelDELETEOperationApi.md#ScsAsIdSubscriptionsSubscriptionIdDelete) | **Delete** /{scsAsId}/subscriptions/{subscriptionId} | Deletes an already existing subscription



## ScsAsIdSubscriptionsSubscriptionIdDelete

> UserPlaneNotificationData ScsAsIdSubscriptionsSubscriptionIdDelete(ctx, scsAsId, subscriptionId).Execute()

Deletes an already existing subscription

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsSessionWithQoSAPISubscriptionLevelDELETEOperationApi.ScsAsIdSubscriptionsSubscriptionIdDelete(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsSessionWithQoSAPISubscriptionLevelDELETEOperationApi.ScsAsIdSubscriptionsSubscriptionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScsAsIdSubscriptionsSubscriptionIdDelete`: UserPlaneNotificationData
    fmt.Fprintf(os.Stdout, "Response from `AsSessionWithQoSAPISubscriptionLevelDELETEOperationApi.ScsAsIdSubscriptionsSubscriptionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiScsAsIdSubscriptionsSubscriptionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserPlaneNotificationData**](UserPlaneNotificationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

