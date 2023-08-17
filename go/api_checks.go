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


// ChecksAPIService ChecksAPI service
type ChecksAPIService service

type ApiAddCheckAppsRequest struct {
	ctx context.Context
	ApiService *ChecksAPIService
	addCheckAppsRequest *AddCheckAppsRequest
}

func (r ApiAddCheckAppsRequest) AddCheckAppsRequest(addCheckAppsRequest AddCheckAppsRequest) ApiAddCheckAppsRequest {
	r.addCheckAppsRequest = &addCheckAppsRequest
	return r
}

func (r ApiAddCheckAppsRequest) Execute() (*AddCheckApps200Response, *http.Response, error) {
	return r.ApiService.AddCheckAppsExecute(r)
}

/*
AddCheckApps Create a New Check App

Create a new check app.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddCheckAppsRequest
*/
func (a *ChecksAPIService) AddCheckApps(ctx context.Context) ApiAddCheckAppsRequest {
	return ApiAddCheckAppsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AddCheckApps200Response
func (a *ChecksAPIService) AddCheckAppsExecute(r ApiAddCheckAppsRequest) (*AddCheckApps200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddCheckApps200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChecksAPIService.AddCheckApps")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/monitoring/apps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.addCheckAppsRequest
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

type ApiListCheckAppsRequest struct {
	ctx context.Context
	ApiService *ChecksAPIService
	max *int64
	offset *int64
	sort *string
	name *string
	phrase *string
	status *string
	lastUpdated *time.Time
	deleted *bool
}

// Maximum number of records to return
func (r ApiListCheckAppsRequest) Max(max int64) ApiListCheckAppsRequest {
	r.max = &max
	return r
}

// Offset records, the number of records to skip, for paginating requests
func (r ApiListCheckAppsRequest) Offset(offset int64) ApiListCheckAppsRequest {
	r.offset = &offset
	return r
}

// Sort order, the name of the property to sort by
func (r ApiListCheckAppsRequest) Sort(sort string) ApiListCheckAppsRequest {
	r.sort = &sort
	return r
}

// Filter by name, wildcard may be specified as %, eg. example-%
func (r ApiListCheckAppsRequest) Name(name string) ApiListCheckAppsRequest {
	r.name = &name
	return r
}

// Search phrase for partial matches on name or description
func (r ApiListCheckAppsRequest) Phrase(phrase string) ApiListCheckAppsRequest {
	r.phrase = &phrase
	return r
}

// The instance status for filtering.
func (r ApiListCheckAppsRequest) Status(status string) ApiListCheckAppsRequest {
	r.status = &status
	return r
}

// Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
func (r ApiListCheckAppsRequest) LastUpdated(lastUpdated time.Time) ApiListCheckAppsRequest {
	r.lastUpdated = &lastUpdated
	return r
}

// If true, only deleted resources or instances in pending removal status are returned.
func (r ApiListCheckAppsRequest) Deleted(deleted bool) ApiListCheckAppsRequest {
	r.deleted = &deleted
	return r
}

func (r ApiListCheckAppsRequest) Execute() (*ListCheckApps200Response, *http.Response, error) {
	return r.ApiService.ListCheckAppsExecute(r)
}

/*
ListCheckApps List All Check Apps

Get a list of check apps.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCheckAppsRequest
*/
func (a *ChecksAPIService) ListCheckApps(ctx context.Context) ApiListCheckAppsRequest {
	return ApiListCheckAppsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListCheckApps200Response
func (a *ChecksAPIService) ListCheckAppsExecute(r ApiListCheckAppsRequest) (*ListCheckApps200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCheckApps200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChecksAPIService.ListCheckApps")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/monitoring/apps"

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
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.phrase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "phrase", r.phrase, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.lastUpdated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastUpdated", r.lastUpdated, "")
	}
	if r.deleted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deleted", r.deleted, "")
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