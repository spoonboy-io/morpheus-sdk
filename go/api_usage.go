/*
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * API version: 6.2.1
 * Contact: dev@morpheusdata.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// UsageApiService UsageApi service
type UsageApiService service

type ApiListUsagesRequest struct {
	ctx _context.Context
	ApiService *UsageApiService
	max *int64
	offset *int64
	sort *string
	direction *string
	phrase *string
	startDate *string
	endDate *string
}

func (r ApiListUsagesRequest) Max(max int64) ApiListUsagesRequest {
	r.max = &max
	return r
}
func (r ApiListUsagesRequest) Offset(offset int64) ApiListUsagesRequest {
	r.offset = &offset
	return r
}
func (r ApiListUsagesRequest) Sort(sort string) ApiListUsagesRequest {
	r.sort = &sort
	return r
}
func (r ApiListUsagesRequest) Direction(direction string) ApiListUsagesRequest {
	r.direction = &direction
	return r
}
func (r ApiListUsagesRequest) Phrase(phrase string) ApiListUsagesRequest {
	r.phrase = &phrase
	return r
}
func (r ApiListUsagesRequest) StartDate(startDate string) ApiListUsagesRequest {
	r.startDate = &startDate
	return r
}
func (r ApiListUsagesRequest) EndDate(endDate string) ApiListUsagesRequest {
	r.endDate = &endDate
	return r
}

func (r ApiListUsagesRequest) Execute() (Usages, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.ListUsagesExecute(r)
}

/*
 * ListUsages Retrieves Usage Records
 * Retrieves a paginated list of usage records. The usages are scoped to only include resources you have access to.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListUsagesRequest
 */
func (a *UsageApiService) ListUsages(ctx _context.Context) ApiListUsagesRequest {
	return ApiListUsagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Usages
 */
func (a *UsageApiService) ListUsagesExecute(r ApiListUsagesRequest) (Usages, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  Usages
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageApiService.ListUsages")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.max != nil {
		localVarQueryParams.Add("max", parameterToString(*r.max, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.direction != nil {
		localVarQueryParams.Add("direction", parameterToString(*r.direction, ""))
	}
	if r.phrase != nil {
		localVarQueryParams.Add("phrase", parameterToString(*r.phrase, ""))
	}
	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 4XX {
			var v DefaultError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 5XX {
			var v DefaultError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

	return localVarReturnValue, localVarHTTPResponse, executionError
}