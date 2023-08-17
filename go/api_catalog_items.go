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
	"strings"
	"os"
)


// CatalogItemsAPIService CatalogItemsAPI service
type CatalogItemsAPIService service

type ApiAddCatalogItemTypeRequest struct {
	ctx context.Context
	ApiService *CatalogItemsAPIService
	addCatalogItemTypeRequest *AddCatalogItemTypeRequest
}

func (r ApiAddCatalogItemTypeRequest) AddCatalogItemTypeRequest(addCatalogItemTypeRequest AddCatalogItemTypeRequest) ApiAddCatalogItemTypeRequest {
	r.addCatalogItemTypeRequest = &addCatalogItemTypeRequest
	return r
}

func (r ApiAddCatalogItemTypeRequest) Execute() (*AddCatalogItemType200Response, *http.Response, error) {
	return r.ApiService.AddCatalogItemTypeExecute(r)
}

/*
AddCatalogItemType Create a Catalog Item Type

Use this command to create a catalog item type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddCatalogItemTypeRequest
*/
func (a *CatalogItemsAPIService) AddCatalogItemType(ctx context.Context) ApiAddCatalogItemTypeRequest {
	return ApiAddCatalogItemTypeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AddCatalogItemType200Response
func (a *CatalogItemsAPIService) AddCatalogItemTypeExecute(r ApiAddCatalogItemTypeRequest) (*AddCatalogItemType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddCatalogItemType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogItemsAPIService.AddCatalogItemType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/catalog-item-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "multipart/form-data"}

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
	localVarPostBody = r.addCatalogItemTypeRequest
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

type ApiGetCatalogItemTypeRequest struct {
	ctx context.Context
	ApiService *CatalogItemsAPIService
	id int64
}

func (r ApiGetCatalogItemTypeRequest) Execute() (*GetCatalogItemType200Response, *http.Response, error) {
	return r.ApiService.GetCatalogItemTypeExecute(r)
}

/*
GetCatalogItemType Get a Specific Catalog Item Type

This endpoint retrieves a specific category item type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Morpheus ID of the Object being referenced
 @return ApiGetCatalogItemTypeRequest
*/
func (a *CatalogItemsAPIService) GetCatalogItemType(ctx context.Context, id int64) ApiGetCatalogItemTypeRequest {
	return ApiGetCatalogItemTypeRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetCatalogItemType200Response
func (a *CatalogItemsAPIService) GetCatalogItemTypeExecute(r ApiGetCatalogItemTypeRequest) (*GetCatalogItemType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCatalogItemType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogItemsAPIService.GetCatalogItemType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/catalog-item-types/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListCatalogItemTypesRequest struct {
	ctx context.Context
	ApiService *CatalogItemsAPIService
	max *int64
	offset *int64
	sort *string
	direction *string
	phrase *string
	name *string
	description *string
	enabled *bool
	featured *bool
	labels *string
	allLabels *string
	code *string
}

// Maximum number of records to return
func (r ApiListCatalogItemTypesRequest) Max(max int64) ApiListCatalogItemTypesRequest {
	r.max = &max
	return r
}

// Offset records, the number of records to skip, for paginating requests
func (r ApiListCatalogItemTypesRequest) Offset(offset int64) ApiListCatalogItemTypesRequest {
	r.offset = &offset
	return r
}

// Sort order, the name of the property to sort by
func (r ApiListCatalogItemTypesRequest) Sort(sort string) ApiListCatalogItemTypesRequest {
	r.sort = &sort
	return r
}

// Sort direction, use &#39;desc&#39; to reverse sort
func (r ApiListCatalogItemTypesRequest) Direction(direction string) ApiListCatalogItemTypesRequest {
	r.direction = &direction
	return r
}

// Search phrase for partial matches on name or description
func (r ApiListCatalogItemTypesRequest) Phrase(phrase string) ApiListCatalogItemTypesRequest {
	r.phrase = &phrase
	return r
}

// Filter by name, wildcard may be specified as %, eg. example-%
func (r ApiListCatalogItemTypesRequest) Name(name string) ApiListCatalogItemTypesRequest {
	r.name = &name
	return r
}

// Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60;
func (r ApiListCatalogItemTypesRequest) Description(description string) ApiListCatalogItemTypesRequest {
	r.description = &description
	return r
}

// Filter by enabled
func (r ApiListCatalogItemTypesRequest) Enabled(enabled bool) ApiListCatalogItemTypesRequest {
	r.enabled = &enabled
	return r
}

// Filter by featured
func (r ApiListCatalogItemTypesRequest) Featured(featured bool) ApiListCatalogItemTypesRequest {
	r.featured = &featured
	return r
}

// Filter by label(s), matches records that contain any of the specified labels
func (r ApiListCatalogItemTypesRequest) Labels(labels string) ApiListCatalogItemTypesRequest {
	r.labels = &labels
	return r
}

// Filter by label(s), matches records that contain all of the specified labels
func (r ApiListCatalogItemTypesRequest) AllLabels(allLabels string) ApiListCatalogItemTypesRequest {
	r.allLabels = &allLabels
	return r
}

// If specified will return an exact match on code
func (r ApiListCatalogItemTypesRequest) Code(code string) ApiListCatalogItemTypesRequest {
	r.code = &code
	return r
}

func (r ApiListCatalogItemTypesRequest) Execute() (*ListCatalogItemTypes200Response, *http.Response, error) {
	return r.ApiService.ListCatalogItemTypesExecute(r)
}

/*
ListCatalogItemTypes Get All Catalog Item Types

This endpoint retrieves all catalog item types.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCatalogItemTypesRequest
*/
func (a *CatalogItemsAPIService) ListCatalogItemTypes(ctx context.Context) ApiListCatalogItemTypesRequest {
	return ApiListCatalogItemTypesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListCatalogItemTypes200Response
func (a *CatalogItemsAPIService) ListCatalogItemTypesExecute(r ApiListCatalogItemTypesRequest) (*ListCatalogItemTypes200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCatalogItemTypes200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogItemsAPIService.ListCatalogItemTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/catalog-item-types"

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
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.phrase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "phrase", r.phrase, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "description", r.description, "")
	}
	if r.enabled != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enabled", r.enabled, "")
	}
	if r.featured != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "featured", r.featured, "")
	}
	if r.labels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "labels", r.labels, "")
	}
	if r.allLabels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "allLabels", r.allLabels, "")
	}
	if r.code != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "code", r.code, "")
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

type ApiRemoveCatalogItemTypeRequest struct {
	ctx context.Context
	ApiService *CatalogItemsAPIService
	id int64
}

func (r ApiRemoveCatalogItemTypeRequest) Execute() (*Model200Success, *http.Response, error) {
	return r.ApiService.RemoveCatalogItemTypeExecute(r)
}

/*
RemoveCatalogItemType Delete a Catalog Item Type

Will delete a catalog item type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Morpheus ID of the Object being referenced
 @return ApiRemoveCatalogItemTypeRequest
*/
func (a *CatalogItemsAPIService) RemoveCatalogItemType(ctx context.Context, id int64) ApiRemoveCatalogItemTypeRequest {
	return ApiRemoveCatalogItemTypeRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Model200Success
func (a *CatalogItemsAPIService) RemoveCatalogItemTypeExecute(r ApiRemoveCatalogItemTypeRequest) (*Model200Success, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Model200Success
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogItemsAPIService.RemoveCatalogItemType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/catalog-item-types/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiUpdateCatalogItemTypeRequest struct {
	ctx context.Context
	ApiService *CatalogItemsAPIService
	id int64
	updateCatalogItemTypeRequest *UpdateCatalogItemTypeRequest
}

func (r ApiUpdateCatalogItemTypeRequest) UpdateCatalogItemTypeRequest(updateCatalogItemTypeRequest UpdateCatalogItemTypeRequest) ApiUpdateCatalogItemTypeRequest {
	r.updateCatalogItemTypeRequest = &updateCatalogItemTypeRequest
	return r
}

func (r ApiUpdateCatalogItemTypeRequest) Execute() (*UpdateCatalogItemType200Response, *http.Response, error) {
	return r.ApiService.UpdateCatalogItemTypeExecute(r)
}

/*
UpdateCatalogItemType Update a Catalog Item Type

Use this command to update an existing catalog item type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Morpheus ID of the Object being referenced
 @return ApiUpdateCatalogItemTypeRequest
*/
func (a *CatalogItemsAPIService) UpdateCatalogItemType(ctx context.Context, id int64) ApiUpdateCatalogItemTypeRequest {
	return ApiUpdateCatalogItemTypeRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UpdateCatalogItemType200Response
func (a *CatalogItemsAPIService) UpdateCatalogItemTypeExecute(r ApiUpdateCatalogItemTypeRequest) (*UpdateCatalogItemType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateCatalogItemType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogItemsAPIService.UpdateCatalogItemType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/catalog-item-types/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "multipart/form-data"}

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
	localVarPostBody = r.updateCatalogItemTypeRequest
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

type ApiUpdateCatalogItemTypeLogoRequest struct {
	ctx context.Context
	ApiService *CatalogItemsAPIService
	id int64
	catalogItemTypeLogo *os.File
	catalogItemTypeDarkLogo *os.File
}

// Logo File png,jpg,svg
func (r ApiUpdateCatalogItemTypeLogoRequest) CatalogItemTypeLogo(catalogItemTypeLogo *os.File) ApiUpdateCatalogItemTypeLogoRequest {
	r.catalogItemTypeLogo = catalogItemTypeLogo
	return r
}

// Dark Logo File png,jpg,svg
func (r ApiUpdateCatalogItemTypeLogoRequest) CatalogItemTypeDarkLogo(catalogItemTypeDarkLogo *os.File) ApiUpdateCatalogItemTypeLogoRequest {
	r.catalogItemTypeDarkLogo = catalogItemTypeDarkLogo
	return r
}

func (r ApiUpdateCatalogItemTypeLogoRequest) Execute() (*UpdateCatalogItemType200Response, *http.Response, error) {
	return r.ApiService.UpdateCatalogItemTypeLogoExecute(r)
}

/*
UpdateCatalogItemTypeLogo Update Logo For Catalog Item Type

Use this command to update the logo and dark logo images for an existing catalog item type. This endpoint expects multipart form data as the request format, not JSON.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Morpheus ID of the Object being referenced
 @return ApiUpdateCatalogItemTypeLogoRequest
*/
func (a *CatalogItemsAPIService) UpdateCatalogItemTypeLogo(ctx context.Context, id int64) ApiUpdateCatalogItemTypeLogoRequest {
	return ApiUpdateCatalogItemTypeLogoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UpdateCatalogItemType200Response
func (a *CatalogItemsAPIService) UpdateCatalogItemTypeLogoExecute(r ApiUpdateCatalogItemTypeLogoRequest) (*UpdateCatalogItemType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateCatalogItemType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogItemsAPIService.UpdateCatalogItemTypeLogo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/catalog-item-types/{id}/update-logo"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var catalogItemTypeLogoLocalVarFormFileName string
	var catalogItemTypeLogoLocalVarFileName     string
	var catalogItemTypeLogoLocalVarFileBytes    []byte

	catalogItemTypeLogoLocalVarFormFileName = "catalogItemType.logo"


	catalogItemTypeLogoLocalVarFile := r.catalogItemTypeLogo

	if catalogItemTypeLogoLocalVarFile != nil {
		fbs, _ := io.ReadAll(catalogItemTypeLogoLocalVarFile)

		catalogItemTypeLogoLocalVarFileBytes = fbs
		catalogItemTypeLogoLocalVarFileName = catalogItemTypeLogoLocalVarFile.Name()
		catalogItemTypeLogoLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: catalogItemTypeLogoLocalVarFileBytes, fileName: catalogItemTypeLogoLocalVarFileName, formFileName: catalogItemTypeLogoLocalVarFormFileName})
	}
	var catalogItemTypeDarkLogoLocalVarFormFileName string
	var catalogItemTypeDarkLogoLocalVarFileName     string
	var catalogItemTypeDarkLogoLocalVarFileBytes    []byte

	catalogItemTypeDarkLogoLocalVarFormFileName = "catalogItemType.darkLogo"


	catalogItemTypeDarkLogoLocalVarFile := r.catalogItemTypeDarkLogo

	if catalogItemTypeDarkLogoLocalVarFile != nil {
		fbs, _ := io.ReadAll(catalogItemTypeDarkLogoLocalVarFile)

		catalogItemTypeDarkLogoLocalVarFileBytes = fbs
		catalogItemTypeDarkLogoLocalVarFileName = catalogItemTypeDarkLogoLocalVarFile.Name()
		catalogItemTypeDarkLogoLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: catalogItemTypeDarkLogoLocalVarFileBytes, fileName: catalogItemTypeDarkLogoLocalVarFileName, formFileName: catalogItemTypeDarkLogoLocalVarFormFileName})
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
