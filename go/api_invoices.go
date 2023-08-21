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

// InvoicesApiService InvoicesApi service
type InvoicesApiService service

type ApiGetInvoiceLineItemsRequest struct {
	ctx _context.Context
	ApiService *InvoicesApiService
	id int64
}


func (r ApiGetInvoiceLineItemsRequest) Execute() (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.GetInvoiceLineItemsExecute(r)
}

/*
 * GetInvoiceLineItems Get a Specific Invoice Line Item
 * Get details about a specific invoice line item.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Morpheus ID of the Object being referenced
 * @return ApiGetInvoiceLineItemsRequest
 */
func (a *InvoicesApiService) GetInvoiceLineItems(ctx _context.Context, id int64) ApiGetInvoiceLineItemsRequest {
	return ApiGetInvoiceLineItemsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *InvoicesApiService) GetInvoiceLineItemsExecute(r ApiGetInvoiceLineItemsRequest) (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.GetInvoiceLineItems")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/invoice-line-items/{id}"
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

type ApiGetInvoicesRequest struct {
	ctx _context.Context
	ApiService *InvoicesApiService
	id int64
}


func (r ApiGetInvoicesRequest) Execute() (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.GetInvoicesExecute(r)
}

/*
 * GetInvoices Get a Specific Invoice
 * Get details about a specific invoice.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Morpheus ID of the Object being referenced
 * @return ApiGetInvoicesRequest
 */
func (a *InvoicesApiService) GetInvoices(ctx _context.Context, id int64) ApiGetInvoicesRequest {
	return ApiGetInvoicesRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *InvoicesApiService) GetInvoicesExecute(r ApiGetInvoicesRequest) (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.GetInvoices")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/invoices/{id}"
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

type ApiListInvoiceLineItemsRequest struct {
	ctx _context.Context
	ApiService *InvoicesApiService
	max *int64
	offset *int64
	sort *string
	direction *string
	phrase *string
	name *string
	startDate *string
	endDate *string
	period *string
	refType *string
	refId *int64
	zoneId *int64
	siteId *int64
	instanceId *int64
	containerId *int64
	serverId *int64
	userId *int64
	projectId *int64
	active *bool
	accountId *int64
	includeTotals *bool
}

func (r ApiListInvoiceLineItemsRequest) Max(max int64) ApiListInvoiceLineItemsRequest {
	r.max = &max
	return r
}
func (r ApiListInvoiceLineItemsRequest) Offset(offset int64) ApiListInvoiceLineItemsRequest {
	r.offset = &offset
	return r
}
func (r ApiListInvoiceLineItemsRequest) Sort(sort string) ApiListInvoiceLineItemsRequest {
	r.sort = &sort
	return r
}
func (r ApiListInvoiceLineItemsRequest) Direction(direction string) ApiListInvoiceLineItemsRequest {
	r.direction = &direction
	return r
}
func (r ApiListInvoiceLineItemsRequest) Phrase(phrase string) ApiListInvoiceLineItemsRequest {
	r.phrase = &phrase
	return r
}
func (r ApiListInvoiceLineItemsRequest) Name(name string) ApiListInvoiceLineItemsRequest {
	r.name = &name
	return r
}
func (r ApiListInvoiceLineItemsRequest) StartDate(startDate string) ApiListInvoiceLineItemsRequest {
	r.startDate = &startDate
	return r
}
func (r ApiListInvoiceLineItemsRequest) EndDate(endDate string) ApiListInvoiceLineItemsRequest {
	r.endDate = &endDate
	return r
}
func (r ApiListInvoiceLineItemsRequest) Period(period string) ApiListInvoiceLineItemsRequest {
	r.period = &period
	return r
}
func (r ApiListInvoiceLineItemsRequest) RefType(refType string) ApiListInvoiceLineItemsRequest {
	r.refType = &refType
	return r
}
func (r ApiListInvoiceLineItemsRequest) RefId(refId int64) ApiListInvoiceLineItemsRequest {
	r.refId = &refId
	return r
}
func (r ApiListInvoiceLineItemsRequest) ZoneId(zoneId int64) ApiListInvoiceLineItemsRequest {
	r.zoneId = &zoneId
	return r
}
func (r ApiListInvoiceLineItemsRequest) SiteId(siteId int64) ApiListInvoiceLineItemsRequest {
	r.siteId = &siteId
	return r
}
func (r ApiListInvoiceLineItemsRequest) InstanceId(instanceId int64) ApiListInvoiceLineItemsRequest {
	r.instanceId = &instanceId
	return r
}
func (r ApiListInvoiceLineItemsRequest) ContainerId(containerId int64) ApiListInvoiceLineItemsRequest {
	r.containerId = &containerId
	return r
}
func (r ApiListInvoiceLineItemsRequest) ServerId(serverId int64) ApiListInvoiceLineItemsRequest {
	r.serverId = &serverId
	return r
}
func (r ApiListInvoiceLineItemsRequest) UserId(userId int64) ApiListInvoiceLineItemsRequest {
	r.userId = &userId
	return r
}
func (r ApiListInvoiceLineItemsRequest) ProjectId(projectId int64) ApiListInvoiceLineItemsRequest {
	r.projectId = &projectId
	return r
}
func (r ApiListInvoiceLineItemsRequest) Active(active bool) ApiListInvoiceLineItemsRequest {
	r.active = &active
	return r
}
func (r ApiListInvoiceLineItemsRequest) AccountId(accountId int64) ApiListInvoiceLineItemsRequest {
	r.accountId = &accountId
	return r
}
func (r ApiListInvoiceLineItemsRequest) IncludeTotals(includeTotals bool) ApiListInvoiceLineItemsRequest {
	r.includeTotals = &includeTotals
	return r
}

func (r ApiListInvoiceLineItemsRequest) Execute() (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.ListInvoiceLineItemsExecute(r)
}

/*
 * ListInvoiceLineItems List All Invoice Line Items
 * This endpoint retrieves a list of all invoice line items for the specified parameters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListInvoiceLineItemsRequest
 */
func (a *InvoicesApiService) ListInvoiceLineItems(ctx _context.Context) ApiListInvoiceLineItemsRequest {
	return ApiListInvoiceLineItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *InvoicesApiService) ListInvoiceLineItemsExecute(r ApiListInvoiceLineItemsRequest) (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.ListInvoiceLineItems")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/invoice-line-items"

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
	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	}
	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
	}
	if r.refType != nil {
		localVarQueryParams.Add("refType", parameterToString(*r.refType, ""))
	}
	if r.refId != nil {
		localVarQueryParams.Add("refId", parameterToString(*r.refId, ""))
	}
	if r.zoneId != nil {
		localVarQueryParams.Add("zoneId", parameterToString(*r.zoneId, ""))
	}
	if r.siteId != nil {
		localVarQueryParams.Add("siteId", parameterToString(*r.siteId, ""))
	}
	if r.instanceId != nil {
		localVarQueryParams.Add("instanceId", parameterToString(*r.instanceId, ""))
	}
	if r.containerId != nil {
		localVarQueryParams.Add("containerId", parameterToString(*r.containerId, ""))
	}
	if r.serverId != nil {
		localVarQueryParams.Add("serverId", parameterToString(*r.serverId, ""))
	}
	if r.userId != nil {
		localVarQueryParams.Add("userId", parameterToString(*r.userId, ""))
	}
	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.active != nil {
		localVarQueryParams.Add("active", parameterToString(*r.active, ""))
	}
	if r.accountId != nil {
		localVarQueryParams.Add("accountId", parameterToString(*r.accountId, ""))
	}
	if r.includeTotals != nil {
		localVarQueryParams.Add("includeTotals", parameterToString(*r.includeTotals, ""))
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

type ApiListInvoicesRequest struct {
	ctx _context.Context
	ApiService *InvoicesApiService
	max *int64
	offset *int64
	sort *string
	direction *string
	phrase *string
	name *string
	startDate *string
	endDate *string
	period *string
	refType *string
	refId *int64
	refStatus *string
	zoneId *int64
	siteId *int64
	instanceId *int64
	containerId *int64
	serverId *int64
	userId *int64
	projectId *int64
	active *bool
	accountId *int64
	includeLineItems *bool
	includeTotals *bool
	tags *string
}

func (r ApiListInvoicesRequest) Max(max int64) ApiListInvoicesRequest {
	r.max = &max
	return r
}
func (r ApiListInvoicesRequest) Offset(offset int64) ApiListInvoicesRequest {
	r.offset = &offset
	return r
}
func (r ApiListInvoicesRequest) Sort(sort string) ApiListInvoicesRequest {
	r.sort = &sort
	return r
}
func (r ApiListInvoicesRequest) Direction(direction string) ApiListInvoicesRequest {
	r.direction = &direction
	return r
}
func (r ApiListInvoicesRequest) Phrase(phrase string) ApiListInvoicesRequest {
	r.phrase = &phrase
	return r
}
func (r ApiListInvoicesRequest) Name(name string) ApiListInvoicesRequest {
	r.name = &name
	return r
}
func (r ApiListInvoicesRequest) StartDate(startDate string) ApiListInvoicesRequest {
	r.startDate = &startDate
	return r
}
func (r ApiListInvoicesRequest) EndDate(endDate string) ApiListInvoicesRequest {
	r.endDate = &endDate
	return r
}
func (r ApiListInvoicesRequest) Period(period string) ApiListInvoicesRequest {
	r.period = &period
	return r
}
func (r ApiListInvoicesRequest) RefType(refType string) ApiListInvoicesRequest {
	r.refType = &refType
	return r
}
func (r ApiListInvoicesRequest) RefId(refId int64) ApiListInvoicesRequest {
	r.refId = &refId
	return r
}
func (r ApiListInvoicesRequest) RefStatus(refStatus string) ApiListInvoicesRequest {
	r.refStatus = &refStatus
	return r
}
func (r ApiListInvoicesRequest) ZoneId(zoneId int64) ApiListInvoicesRequest {
	r.zoneId = &zoneId
	return r
}
func (r ApiListInvoicesRequest) SiteId(siteId int64) ApiListInvoicesRequest {
	r.siteId = &siteId
	return r
}
func (r ApiListInvoicesRequest) InstanceId(instanceId int64) ApiListInvoicesRequest {
	r.instanceId = &instanceId
	return r
}
func (r ApiListInvoicesRequest) ContainerId(containerId int64) ApiListInvoicesRequest {
	r.containerId = &containerId
	return r
}
func (r ApiListInvoicesRequest) ServerId(serverId int64) ApiListInvoicesRequest {
	r.serverId = &serverId
	return r
}
func (r ApiListInvoicesRequest) UserId(userId int64) ApiListInvoicesRequest {
	r.userId = &userId
	return r
}
func (r ApiListInvoicesRequest) ProjectId(projectId int64) ApiListInvoicesRequest {
	r.projectId = &projectId
	return r
}
func (r ApiListInvoicesRequest) Active(active bool) ApiListInvoicesRequest {
	r.active = &active
	return r
}
func (r ApiListInvoicesRequest) AccountId(accountId int64) ApiListInvoicesRequest {
	r.accountId = &accountId
	return r
}
func (r ApiListInvoicesRequest) IncludeLineItems(includeLineItems bool) ApiListInvoicesRequest {
	r.includeLineItems = &includeLineItems
	return r
}
func (r ApiListInvoicesRequest) IncludeTotals(includeTotals bool) ApiListInvoicesRequest {
	r.includeTotals = &includeTotals
	return r
}
func (r ApiListInvoicesRequest) Tags(tags string) ApiListInvoicesRequest {
	r.tags = &tags
	return r
}

func (r ApiListInvoicesRequest) Execute() (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.ListInvoicesExecute(r)
}

/*
 * ListInvoices List All Invoices
 * This endpoint retrieves a list of invoices for the specified parameters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListInvoicesRequest
 */
func (a *InvoicesApiService) ListInvoices(ctx _context.Context) ApiListInvoicesRequest {
	return ApiListInvoicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *InvoicesApiService) ListInvoicesExecute(r ApiListInvoicesRequest) (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.ListInvoices")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/invoices"

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
	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	}
	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
	}
	if r.refType != nil {
		localVarQueryParams.Add("refType", parameterToString(*r.refType, ""))
	}
	if r.refId != nil {
		localVarQueryParams.Add("refId", parameterToString(*r.refId, ""))
	}
	if r.refStatus != nil {
		localVarQueryParams.Add("refStatus", parameterToString(*r.refStatus, ""))
	}
	if r.zoneId != nil {
		localVarQueryParams.Add("zoneId", parameterToString(*r.zoneId, ""))
	}
	if r.siteId != nil {
		localVarQueryParams.Add("siteId", parameterToString(*r.siteId, ""))
	}
	if r.instanceId != nil {
		localVarQueryParams.Add("instanceId", parameterToString(*r.instanceId, ""))
	}
	if r.containerId != nil {
		localVarQueryParams.Add("containerId", parameterToString(*r.containerId, ""))
	}
	if r.serverId != nil {
		localVarQueryParams.Add("serverId", parameterToString(*r.serverId, ""))
	}
	if r.userId != nil {
		localVarQueryParams.Add("userId", parameterToString(*r.userId, ""))
	}
	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.active != nil {
		localVarQueryParams.Add("active", parameterToString(*r.active, ""))
	}
	if r.accountId != nil {
		localVarQueryParams.Add("accountId", parameterToString(*r.accountId, ""))
	}
	if r.includeLineItems != nil {
		localVarQueryParams.Add("includeLineItems", parameterToString(*r.includeLineItems, ""))
	}
	if r.includeTotals != nil {
		localVarQueryParams.Add("includeTotals", parameterToString(*r.includeTotals, ""))
	}
	if r.tags != nil {
		localVarQueryParams.Add("tags", parameterToString(*r.tags, ""))
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

type ApiUpdateInvoicesRequest struct {
	ctx _context.Context
	ApiService *InvoicesApiService
	id int64
	inlineObject102 *InlineObject102
}

func (r ApiUpdateInvoicesRequest) InlineObject102(inlineObject102 InlineObject102) ApiUpdateInvoicesRequest {
	r.inlineObject102 = &inlineObject102
	return r
}

func (r ApiUpdateInvoicesRequest) Execute() (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.UpdateInvoicesExecute(r)
}

/*
 * UpdateInvoices Update Invoice Tags
 * Update, Add, or Remove invoice tag(s).
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Morpheus ID of the Object being referenced
 * @return ApiUpdateInvoicesRequest
 */
func (a *InvoicesApiService) UpdateInvoices(ctx _context.Context, id int64) ApiUpdateInvoicesRequest {
	return ApiUpdateInvoicesRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *InvoicesApiService) UpdateInvoicesExecute(r ApiUpdateInvoicesRequest) (map[string]interface{}, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.UpdateInvoices")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/api/invoices/{id}"
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
	localVarPostBody = r.inlineObject102
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
