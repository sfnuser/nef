/*
3gpp-as-session-with-qos

API for setting us an AS session with required QoS. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package assessionwithqos

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type AsSessionWithQoSAPISubscriptionLevelPOSTOperationApi interface {

	/*
	ScsAsIdSubscriptionsPost Creates a new subscription resource

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param scsAsId Identifier of the SCS/AS
	 @return ApiScsAsIdSubscriptionsPostRequest
	*/
	ScsAsIdSubscriptionsPost(ctx _context.Context, scsAsId string) ApiScsAsIdSubscriptionsPostRequest

	// ScsAsIdSubscriptionsPostExecute executes the request
	//  @return AsSessionWithQoSSubscription
	ScsAsIdSubscriptionsPostExecute(r ApiScsAsIdSubscriptionsPostRequest) (AsSessionWithQoSSubscription, *_nethttp.Response, error)
}

// AsSessionWithQoSAPISubscriptionLevelPOSTOperationApiService AsSessionWithQoSAPISubscriptionLevelPOSTOperationApi service
type AsSessionWithQoSAPISubscriptionLevelPOSTOperationApiService service

type ApiScsAsIdSubscriptionsPostRequest struct {
	ctx _context.Context
	ApiService AsSessionWithQoSAPISubscriptionLevelPOSTOperationApi
	scsAsId string
	asSessionWithQoSSubscription *AsSessionWithQoSSubscription
}

// Request to create a new subscription resource
func (r ApiScsAsIdSubscriptionsPostRequest) AsSessionWithQoSSubscription(asSessionWithQoSSubscription AsSessionWithQoSSubscription) ApiScsAsIdSubscriptionsPostRequest {
	r.asSessionWithQoSSubscription = &asSessionWithQoSSubscription
	return r
}

func (r ApiScsAsIdSubscriptionsPostRequest) Execute() (AsSessionWithQoSSubscription, *_nethttp.Response, error) {
	return r.ApiService.ScsAsIdSubscriptionsPostExecute(r)
}

/*
ScsAsIdSubscriptionsPost Creates a new subscription resource

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param scsAsId Identifier of the SCS/AS
 @return ApiScsAsIdSubscriptionsPostRequest
*/
func (a *AsSessionWithQoSAPISubscriptionLevelPOSTOperationApiService) ScsAsIdSubscriptionsPost(ctx _context.Context, scsAsId string) ApiScsAsIdSubscriptionsPostRequest {
	return ApiScsAsIdSubscriptionsPostRequest{
		ApiService: a,
		ctx: ctx,
		scsAsId: scsAsId,
	}
}

// Execute executes the request
//  @return AsSessionWithQoSSubscription
func (a *AsSessionWithQoSAPISubscriptionLevelPOSTOperationApiService) ScsAsIdSubscriptionsPostExecute(r ApiScsAsIdSubscriptionsPostRequest) (AsSessionWithQoSSubscription, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AsSessionWithQoSSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AsSessionWithQoSAPISubscriptionLevelPOSTOperationApiService.ScsAsIdSubscriptionsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{scsAsId}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"scsAsId"+"}", _neturl.PathEscape(parameterToString(r.scsAsId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.asSessionWithQoSSubscription == nil {
		return localVarReturnValue, nil, reportError("asSessionWithQoSSubscription is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.asSessionWithQoSSubscription
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 411 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
