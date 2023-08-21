/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiCallback;
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.ApiResponse;
import org.openapitools.client.Configuration;
import org.openapitools.client.Pair;
import org.openapitools.client.ProgressRequestBody;
import org.openapitools.client.ProgressResponseBody;

import com.google.gson.reflect.TypeToken;

import java.io.IOException;


import org.openapitools.client.model.DefaultError;
import org.openapitools.client.model.InlineObject200;
import org.openapitools.client.model.InlineObject201;
import org.openapitools.client.model.InlineResponse200124;
import org.openapitools.client.model.Model200Success;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class PricesApi {
    private ApiClient localVarApiClient;

    public PricesApi() {
        this(Configuration.getDefaultApiClient());
    }

    public PricesApi(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    public ApiClient getApiClient() {
        return localVarApiClient;
    }

    public void setApiClient(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    /**
     * Build call for addPrices
     * @param inlineObject200  (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call addPricesCall(InlineObject200 inlineObject200, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = inlineObject200;

        // create path and map variables
        String localVarPath = "/api/prices";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "bearerAuth" };
        return localVarApiClient.buildCall(localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call addPricesValidateBeforeCall(InlineObject200 inlineObject200, final ApiCallback _callback) throws ApiException {
        

        okhttp3.Call localVarCall = addPricesCall(inlineObject200, _callback);
        return localVarCall;

    }

    /**
     * Creates a Price
     * Creates a price. 
     * @param inlineObject200  (optional)
     * @return Object
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public Object addPrices(InlineObject200 inlineObject200) throws ApiException {
        ApiResponse<Object> localVarResp = addPricesWithHttpInfo(inlineObject200);
        return localVarResp.getData();
    }

    /**
     * Creates a Price
     * Creates a price. 
     * @param inlineObject200  (optional)
     * @return ApiResponse&lt;Object&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<Object> addPricesWithHttpInfo(InlineObject200 inlineObject200) throws ApiException {
        okhttp3.Call localVarCall = addPricesValidateBeforeCall(inlineObject200, null);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Creates a Price (asynchronously)
     * Creates a price. 
     * @param inlineObject200  (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call addPricesAsync(InlineObject200 inlineObject200, final ApiCallback<Object> _callback) throws ApiException {

        okhttp3.Call localVarCall = addPricesValidateBeforeCall(inlineObject200, _callback);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for deactivatePrices
     * @param id Morpheus ID of the Object being referenced (required)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call deactivatePricesCall(Long id, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/api/prices/{id}/deactivate"
            .replaceAll("\\{" + "id" + "\\}", localVarApiClient.escapeString(id.toString()));

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "bearerAuth" };
        return localVarApiClient.buildCall(localVarPath, "PUT", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call deactivatePricesValidateBeforeCall(Long id, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'id' is set
        if (id == null) {
            throw new ApiException("Missing the required parameter 'id' when calling deactivatePrices(Async)");
        }
        

        okhttp3.Call localVarCall = deactivatePricesCall(id, _callback);
        return localVarCall;

    }

    /**
     * Deactivates a Price
     * Deactivates a price. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 
     * @param id Morpheus ID of the Object being referenced (required)
     * @return Model200Success
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public Model200Success deactivatePrices(Long id) throws ApiException {
        ApiResponse<Model200Success> localVarResp = deactivatePricesWithHttpInfo(id);
        return localVarResp.getData();
    }

    /**
     * Deactivates a Price
     * Deactivates a price. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 
     * @param id Morpheus ID of the Object being referenced (required)
     * @return ApiResponse&lt;Model200Success&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<Model200Success> deactivatePricesWithHttpInfo(Long id) throws ApiException {
        okhttp3.Call localVarCall = deactivatePricesValidateBeforeCall(id, null);
        Type localVarReturnType = new TypeToken<Model200Success>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Deactivates a Price (asynchronously)
     * Deactivates a price. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 
     * @param id Morpheus ID of the Object being referenced (required)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call deactivatePricesAsync(Long id, final ApiCallback<Model200Success> _callback) throws ApiException {

        okhttp3.Call localVarCall = deactivatePricesValidateBeforeCall(id, _callback);
        Type localVarReturnType = new TypeToken<Model200Success>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for getPrices
     * @param id Morpheus ID of the Object being referenced (required)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getPricesCall(Long id, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/api/prices/{id}"
            .replaceAll("\\{" + "id" + "\\}", localVarApiClient.escapeString(id.toString()));

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "bearerAuth" };
        return localVarApiClient.buildCall(localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getPricesValidateBeforeCall(Long id, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'id' is set
        if (id == null) {
            throw new ApiException("Missing the required parameter 'id' when calling getPrices(Async)");
        }
        

        okhttp3.Call localVarCall = getPricesCall(id, _callback);
        return localVarCall;

    }

    /**
     * Retrieves a Specific Price
     * Retrieves a specific price. 
     * @param id Morpheus ID of the Object being referenced (required)
     * @return InlineResponse200124
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public InlineResponse200124 getPrices(Long id) throws ApiException {
        ApiResponse<InlineResponse200124> localVarResp = getPricesWithHttpInfo(id);
        return localVarResp.getData();
    }

    /**
     * Retrieves a Specific Price
     * Retrieves a specific price. 
     * @param id Morpheus ID of the Object being referenced (required)
     * @return ApiResponse&lt;InlineResponse200124&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<InlineResponse200124> getPricesWithHttpInfo(Long id) throws ApiException {
        okhttp3.Call localVarCall = getPricesValidateBeforeCall(id, null);
        Type localVarReturnType = new TypeToken<InlineResponse200124>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Retrieves a Specific Price (asynchronously)
     * Retrieves a specific price. 
     * @param id Morpheus ID of the Object being referenced (required)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getPricesAsync(Long id, final ApiCallback<InlineResponse200124> _callback) throws ApiException {

        okhttp3.Call localVarCall = getPricesValidateBeforeCall(id, _callback);
        Type localVarReturnType = new TypeToken<InlineResponse200124>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for listPrices
     * @param max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25l)
     * @param offset Offset records, the number of records to skip, for paginating requests (optional, default to 0l)
     * @param sort Sort order, the name of the property to sort by (optional, default to &quot;name&quot;)
     * @param direction Sort direction, use &#39;desc&#39; to reverse sort (optional, default to asc)
     * @param phrase Search phrase for partial matches on name or description (optional)
     * @param name Filter by name, wildcard may be specified as %, eg. example-% (optional)
     * @param includeInactive If true, include inactive prices in the results (optional)
     * @param priceType Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software  (optional)
     * @param platform Restricts query to only load only prices with specified platform (optional)
     * @param currency Restricts query to only load only prices with specified currency (optional)
     * @param type Filter by type code (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call listPricesCall(Long max, Long offset, String sort, String direction, String phrase, String name, Boolean includeInactive, String priceType, String platform, String currency, String type, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/api/prices";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        if (max != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("max", max));
        }

        if (offset != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("offset", offset));
        }

        if (sort != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("sort", sort));
        }

        if (direction != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("direction", direction));
        }

        if (phrase != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("phrase", phrase));
        }

        if (name != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("name", name));
        }

        if (includeInactive != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("includeInactive", includeInactive));
        }

        if (priceType != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("priceType", priceType));
        }

        if (platform != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("platform", platform));
        }

        if (currency != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("currency", currency));
        }

        if (type != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("type", type));
        }

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "bearerAuth" };
        return localVarApiClient.buildCall(localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call listPricesValidateBeforeCall(Long max, Long offset, String sort, String direction, String phrase, String name, Boolean includeInactive, String priceType, String platform, String currency, String type, final ApiCallback _callback) throws ApiException {
        

        okhttp3.Call localVarCall = listPricesCall(max, offset, sort, direction, phrase, name, includeInactive, priceType, platform, currency, type, _callback);
        return localVarCall;

    }

    /**
     * Retrieves all Prices
     * Retrieves all prices. 
     * @param max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25l)
     * @param offset Offset records, the number of records to skip, for paginating requests (optional, default to 0l)
     * @param sort Sort order, the name of the property to sort by (optional, default to &quot;name&quot;)
     * @param direction Sort direction, use &#39;desc&#39; to reverse sort (optional, default to asc)
     * @param phrase Search phrase for partial matches on name or description (optional)
     * @param name Filter by name, wildcard may be specified as %, eg. example-% (optional)
     * @param includeInactive If true, include inactive prices in the results (optional)
     * @param priceType Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software  (optional)
     * @param platform Restricts query to only load only prices with specified platform (optional)
     * @param currency Restricts query to only load only prices with specified currency (optional)
     * @param type Filter by type code (optional)
     * @return Object
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public Object listPrices(Long max, Long offset, String sort, String direction, String phrase, String name, Boolean includeInactive, String priceType, String platform, String currency, String type) throws ApiException {
        ApiResponse<Object> localVarResp = listPricesWithHttpInfo(max, offset, sort, direction, phrase, name, includeInactive, priceType, platform, currency, type);
        return localVarResp.getData();
    }

    /**
     * Retrieves all Prices
     * Retrieves all prices. 
     * @param max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25l)
     * @param offset Offset records, the number of records to skip, for paginating requests (optional, default to 0l)
     * @param sort Sort order, the name of the property to sort by (optional, default to &quot;name&quot;)
     * @param direction Sort direction, use &#39;desc&#39; to reverse sort (optional, default to asc)
     * @param phrase Search phrase for partial matches on name or description (optional)
     * @param name Filter by name, wildcard may be specified as %, eg. example-% (optional)
     * @param includeInactive If true, include inactive prices in the results (optional)
     * @param priceType Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software  (optional)
     * @param platform Restricts query to only load only prices with specified platform (optional)
     * @param currency Restricts query to only load only prices with specified currency (optional)
     * @param type Filter by type code (optional)
     * @return ApiResponse&lt;Object&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<Object> listPricesWithHttpInfo(Long max, Long offset, String sort, String direction, String phrase, String name, Boolean includeInactive, String priceType, String platform, String currency, String type) throws ApiException {
        okhttp3.Call localVarCall = listPricesValidateBeforeCall(max, offset, sort, direction, phrase, name, includeInactive, priceType, platform, currency, type, null);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Retrieves all Prices (asynchronously)
     * Retrieves all prices. 
     * @param max Maximum number of records to return, -1 can be used to fetch all records (optional, default to 25l)
     * @param offset Offset records, the number of records to skip, for paginating requests (optional, default to 0l)
     * @param sort Sort order, the name of the property to sort by (optional, default to &quot;name&quot;)
     * @param direction Sort direction, use &#39;desc&#39; to reverse sort (optional, default to asc)
     * @param phrase Search phrase for partial matches on name or description (optional)
     * @param name Filter by name, wildcard may be specified as %, eg. example-% (optional)
     * @param includeInactive If true, include inactive prices in the results (optional)
     * @param priceType Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software  (optional)
     * @param platform Restricts query to only load only prices with specified platform (optional)
     * @param currency Restricts query to only load only prices with specified currency (optional)
     * @param type Filter by type code (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call listPricesAsync(Long max, Long offset, String sort, String direction, String phrase, String name, Boolean includeInactive, String priceType, String platform, String currency, String type, final ApiCallback<Object> _callback) throws ApiException {

        okhttp3.Call localVarCall = listPricesValidateBeforeCall(max, offset, sort, direction, phrase, name, includeInactive, priceType, platform, currency, type, _callback);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for updatePrices
     * @param id Morpheus ID of the Object being referenced (required)
     * @param inlineObject201  (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call updatePricesCall(Long id, InlineObject201 inlineObject201, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = inlineObject201;

        // create path and map variables
        String localVarPath = "/api/prices/{id}"
            .replaceAll("\\{" + "id" + "\\}", localVarApiClient.escapeString(id.toString()));

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "bearerAuth" };
        return localVarApiClient.buildCall(localVarPath, "PUT", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call updatePricesValidateBeforeCall(Long id, InlineObject201 inlineObject201, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'id' is set
        if (id == null) {
            throw new ApiException("Missing the required parameter 'id' when calling updatePrices(Async)");
        }
        

        okhttp3.Call localVarCall = updatePricesCall(id, inlineObject201, _callback);
        return localVarCall;

    }

    /**
     * Updates a Price
     * Updates a price. 
     * @param id Morpheus ID of the Object being referenced (required)
     * @param inlineObject201  (optional)
     * @return Object
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public Object updatePrices(Long id, InlineObject201 inlineObject201) throws ApiException {
        ApiResponse<Object> localVarResp = updatePricesWithHttpInfo(id, inlineObject201);
        return localVarResp.getData();
    }

    /**
     * Updates a Price
     * Updates a price. 
     * @param id Morpheus ID of the Object being referenced (required)
     * @param inlineObject201  (optional)
     * @return ApiResponse&lt;Object&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<Object> updatePricesWithHttpInfo(Long id, InlineObject201 inlineObject201) throws ApiException {
        okhttp3.Call localVarCall = updatePricesValidateBeforeCall(id, inlineObject201, null);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Updates a Price (asynchronously)
     * Updates a price. 
     * @param id Morpheus ID of the Object being referenced (required)
     * @param inlineObject201  (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Request </td><td>  -  </td></tr>
        <tr><td> 4XX </td><td> Error Codes </td><td>  -  </td></tr>
        <tr><td> 5XX </td><td> Error Codes </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call updatePricesAsync(Long id, InlineObject201 inlineObject201, final ApiCallback<Object> _callback) throws ApiException {

        okhttp3.Call localVarCall = updatePricesValidateBeforeCall(id, inlineObject201, _callback);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
}