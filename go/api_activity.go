/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)


// ActivityAPIService ActivityAPI service
type ActivityAPIService service

type ApiListActivityRequest struct {
	ctx context.Context
	ApiService *ActivityAPIService
	max *int64
	offset *int64
	sort *string
	order *string
	phrase *string
	name *string
	userId *int64
	tenantId *float32
	timeframe *string
	start *time.Time
	end *time.Time
}

// Maximum number of records to return
func (r ApiListActivityRequest) Max(max int64) ApiListActivityRequest {
	r.max = &max
	return r
}

// Offset records, the number of records to skip, for paginating requests
func (r ApiListActivityRequest) Offset(offset int64) ApiListActivityRequest {
	r.offset = &offset
	return r
}

// Sort order, the name of the property to sort by
func (r ApiListActivityRequest) Sort(sort string) ApiListActivityRequest {
	r.sort = &sort
	return r
}

// Sort direction, use &#39;desc&#39; to reverse sort
func (r ApiListActivityRequest) Order(order string) ApiListActivityRequest {
	r.order = &order
	return r
}

// Search phrase for partial matches on name or description
func (r ApiListActivityRequest) Phrase(phrase string) ApiListActivityRequest {
	r.phrase = &phrase
	return r
}

// Filter by name, wildcard may be specified as %, eg. example-%
func (r ApiListActivityRequest) Name(name string) ApiListActivityRequest {
	r.name = &name
	return r
}

// Filter by User ID
func (r ApiListActivityRequest) UserId(userId int64) ApiListActivityRequest {
	r.userId = &userId
	return r
}

// Filter by Tenant ID. Only available to the master account.
func (r ApiListActivityRequest) TenantId(tenantId float32) ApiListActivityRequest {
	r.tenantId = &tenantId
	return r
}

// Filter by a timeframe (ex - today, yesterday, week, month, 3months)
func (r ApiListActivityRequest) Timeframe(timeframe string) ApiListActivityRequest {
	r.timeframe = &timeframe
	return r
}

// Filter by activity on or after a date(time). Default is 1 month prior
func (r ApiListActivityRequest) Start(start time.Time) ApiListActivityRequest {
	r.start = &start
	return r
}

// Filter by activity on or before a date(time). Default is current date
func (r ApiListActivityRequest) End(end time.Time) ApiListActivityRequest {
	r.end = &end
	return r
}

func (r ApiListActivityRequest) Execute() (*ListActivity200Response, *http.Response, error) {
	return r.ApiService.ListActivityExecute(r)
}

/*
ListActivity Retrieves Activity

Retrieves activity.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListActivityRequest
*/
func (a *ActivityAPIService) ListActivity(ctx context.Context) ApiListActivityRequest {
	return ApiListActivityRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListActivity200Response
func (a *ActivityAPIService) ListActivityExecute(r ApiListActivityRequest) (*ListActivity200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListActivity200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivityAPIService.ListActivity")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/activity"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.max != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max", r.max, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.phrase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "phrase", r.phrase, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.tenantId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "")
	}
	if r.timeframe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeframe", r.timeframe, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v DefaultError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v DefaultError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
