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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// SecurityPackagesApiService SecurityPackagesApi service
type SecurityPackagesApiService service

type ApiAddSecurityPackagesRequest struct {
	ctx _context.Context
	ApiService *SecurityPackagesApiService
	inlineObject218 *InlineObject218
}

func (r ApiAddSecurityPackagesRequest) InlineObject218(inlineObject218 InlineObject218) ApiAddSecurityPackagesRequest {
	r.inlineObject218 = &inlineObject218
	return r
}

func (r ApiAddSecurityPackagesRequest) Execute() (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.AddSecurityPackagesExecute(r)
}

/*
 * AddSecurityPackages Creates a Security Package
 * Creates a security package.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAddSecurityPackagesRequest
 */
func (a *SecurityPackagesApiService) AddSecurityPackages(ctx _context.Context) ApiAddSecurityPackagesRequest {
	return ApiAddSecurityPackagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *SecurityPackagesApiService) AddSecurityPackagesExecute(r ApiAddSecurityPackagesRequest) (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityPackagesApiService.AddSecurityPackages")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/security-packages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = r.inlineObject218
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

type ApiGetSecurityPackagesRequest struct {
	ctx _context.Context
	ApiService *SecurityPackagesApiService
	id int64
}


func (r ApiGetSecurityPackagesRequest) Execute() (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.GetSecurityPackagesExecute(r)
}

/*
 * GetSecurityPackages Retrieves a Specific Security Package
 * Retrieves a specific security package.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Morpheus ID of the Object being referenced
 * @return ApiGetSecurityPackagesRequest
 */
func (a *SecurityPackagesApiService) GetSecurityPackages(ctx _context.Context, id int64) ApiGetSecurityPackagesRequest {
	return ApiGetSecurityPackagesRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *SecurityPackagesApiService) GetSecurityPackagesExecute(r ApiGetSecurityPackagesRequest) (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityPackagesApiService.GetSecurityPackages")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/security-packages/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiListSecurityPackagesRequest struct {
	ctx _context.Context
	ApiService *SecurityPackagesApiService
	max *int64
	offset *int64
	sort *string
	direction *string
	phrase *string
	name *string
	labels *string
	allLabels *string
}

func (r ApiListSecurityPackagesRequest) Max(max int64) ApiListSecurityPackagesRequest {
	r.max = &max
	return r
}
func (r ApiListSecurityPackagesRequest) Offset(offset int64) ApiListSecurityPackagesRequest {
	r.offset = &offset
	return r
}
func (r ApiListSecurityPackagesRequest) Sort(sort string) ApiListSecurityPackagesRequest {
	r.sort = &sort
	return r
}
func (r ApiListSecurityPackagesRequest) Direction(direction string) ApiListSecurityPackagesRequest {
	r.direction = &direction
	return r
}
func (r ApiListSecurityPackagesRequest) Phrase(phrase string) ApiListSecurityPackagesRequest {
	r.phrase = &phrase
	return r
}
func (r ApiListSecurityPackagesRequest) Name(name string) ApiListSecurityPackagesRequest {
	r.name = &name
	return r
}
func (r ApiListSecurityPackagesRequest) Labels(labels string) ApiListSecurityPackagesRequest {
	r.labels = &labels
	return r
}
func (r ApiListSecurityPackagesRequest) AllLabels(allLabels string) ApiListSecurityPackagesRequest {
	r.allLabels = &allLabels
	return r
}

func (r ApiListSecurityPackagesRequest) Execute() (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.ListSecurityPackagesExecute(r)
}

/*
 * ListSecurityPackages Retrieves all Security Packages
 * Retrieves all security packages.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListSecurityPackagesRequest
 */
func (a *SecurityPackagesApiService) ListSecurityPackages(ctx _context.Context) ApiListSecurityPackagesRequest {
	return ApiListSecurityPackagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *SecurityPackagesApiService) ListSecurityPackagesExecute(r ApiListSecurityPackagesRequest) (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityPackagesApiService.ListSecurityPackages")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/security-packages"

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
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.labels != nil {
		localVarQueryParams.Add("labels", parameterToString(*r.labels, ""))
	}
	if r.allLabels != nil {
		localVarQueryParams.Add("allLabels", parameterToString(*r.allLabels, ""))
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

type ApiRemoveSecurityPackagesRequest struct {
	ctx _context.Context
	ApiService *SecurityPackagesApiService
	id int64
}


func (r ApiRemoveSecurityPackagesRequest) Execute() (Model200Success, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.RemoveSecurityPackagesExecute(r)
}

/*
 * RemoveSecurityPackages Deletes a Security Package
 * Deletes a specified security package.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Morpheus ID of the Object being referenced
 * @return ApiRemoveSecurityPackagesRequest
 */
func (a *SecurityPackagesApiService) RemoveSecurityPackages(ctx _context.Context, id int64) ApiRemoveSecurityPackagesRequest {
	return ApiRemoveSecurityPackagesRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Model200Success
 */
func (a *SecurityPackagesApiService) RemoveSecurityPackagesExecute(r ApiRemoveSecurityPackagesRequest) (Model200Success, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  Model200Success
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityPackagesApiService.RemoveSecurityPackages")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/security-packages/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiUpdateSecurityPackagesRequest struct {
	ctx _context.Context
	ApiService *SecurityPackagesApiService
	id int64
	inlineObject219 *InlineObject219
}

func (r ApiUpdateSecurityPackagesRequest) InlineObject219(inlineObject219 InlineObject219) ApiUpdateSecurityPackagesRequest {
	r.inlineObject219 = &inlineObject219
	return r
}

func (r ApiUpdateSecurityPackagesRequest) Execute() (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.UpdateSecurityPackagesExecute(r)
}

/*
 * UpdateSecurityPackages Updates a Security Package
 * Updates a security package.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Morpheus ID of the Object being referenced
 * @return ApiUpdateSecurityPackagesRequest
 */
func (a *SecurityPackagesApiService) UpdateSecurityPackages(ctx _context.Context, id int64) ApiUpdateSecurityPackagesRequest {
	return ApiUpdateSecurityPackagesRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *SecurityPackagesApiService) UpdateSecurityPackagesExecute(r ApiUpdateSecurityPackagesRequest) (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityPackagesApiService.UpdateSecurityPackages")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/security-packages/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = r.inlineObject219
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
