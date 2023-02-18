# \AsSessionWithQoSAPISubscriptionLevelPOSTOperationApi

All URIs are relative to *https://example.com/3gpp-as-session-with-qos/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScsAsIdSubscriptionsPost**](AsSessionWithQoSAPISubscriptionLevelPOSTOperationApi.md#ScsAsIdSubscriptionsPost) | **Post** /{scsAsId}/subscriptions | Creates a new subscription resource



## ScsAsIdSubscriptionsPost

> AsSessionWithQoSSubscription ScsAsIdSubscriptionsPost(ctx, scsAsId).AsSessionWithQoSSubscription(asSessionWithQoSSubscription).Execute()

Creates a new subscription resource

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
    asSessionWithQoSSubscription := *openapiclient.NewAsSessionWithQoSSubscription("NotificationDestination_example") // AsSessionWithQoSSubscription | Request to create a new subscription resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsSessionWithQoSAPISubscriptionLevelPOSTOperationApi.ScsAsIdSubscriptionsPost(context.Background(), scsAsId).AsSessionWithQoSSubscription(asSessionWithQoSSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsSessionWithQoSAPISubscriptionLevelPOSTOperationApi.ScsAsIdSubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScsAsIdSubscriptionsPost`: AsSessionWithQoSSubscription
    fmt.Fprintf(os.Stdout, "Response from `AsSessionWithQoSAPISubscriptionLevelPOSTOperationApi.ScsAsIdSubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiScsAsIdSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asSessionWithQoSSubscription** | [**AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md) | Request to create a new subscription resource | 

### Return type

[**AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

