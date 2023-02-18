# \AsSessionWithQoSAPISubscriptionLevelPATCHOperationApi

All URIs are relative to *https://example.com/3gpp-as-session-with-qos/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScsAsIdSubscriptionsSubscriptionIdPatch**](AsSessionWithQoSAPISubscriptionLevelPATCHOperationApi.md#ScsAsIdSubscriptionsSubscriptionIdPatch) | **Patch** /{scsAsId}/subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource



## ScsAsIdSubscriptionsSubscriptionIdPatch

> AsSessionWithQoSSubscription ScsAsIdSubscriptionsSubscriptionIdPatch(ctx, scsAsId, subscriptionId).AsSessionWithQoSSubscriptionPatch(asSessionWithQoSSubscriptionPatch).Execute()

Updates/replaces an existing subscription resource

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
    asSessionWithQoSSubscriptionPatch := *openapiclient.NewAsSessionWithQoSSubscriptionPatch() // AsSessionWithQoSSubscriptionPatch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsSessionWithQoSAPISubscriptionLevelPATCHOperationApi.ScsAsIdSubscriptionsSubscriptionIdPatch(context.Background(), scsAsId, subscriptionId).AsSessionWithQoSSubscriptionPatch(asSessionWithQoSSubscriptionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsSessionWithQoSAPISubscriptionLevelPATCHOperationApi.ScsAsIdSubscriptionsSubscriptionIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScsAsIdSubscriptionsSubscriptionIdPatch`: AsSessionWithQoSSubscription
    fmt.Fprintf(os.Stdout, "Response from `AsSessionWithQoSAPISubscriptionLevelPATCHOperationApi.ScsAsIdSubscriptionsSubscriptionIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiScsAsIdSubscriptionsSubscriptionIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asSessionWithQoSSubscriptionPatch** | [**AsSessionWithQoSSubscriptionPatch**](AsSessionWithQoSSubscriptionPatch.md) |  | 

### Return type

[**AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

