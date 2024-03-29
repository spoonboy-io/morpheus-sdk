/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import DefaultError from '../model/DefaultError';
import Search from '../model/Search';

/**
* Search service.
* @module api/SearchApi
* @version 6.2.1
*/
export default class SearchApi {

    /**
    * Constructs a new SearchApi. 
    * @alias module:api/SearchApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the search operation.
     * @callback module:api/SearchApi~searchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Search} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Provides global search for all types of objects
     * This endpoint provides global search for all types of objects, with the results sorted in order of relevance.  The `phrase` parameter can be specified as part of the URL or as a query parameter. If phrase is not specified, then 0 results (hits) will be returned. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {String} opts.query Alias for phrase
     * @param {module:api/SearchApi~searchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Search}
     */
    search(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'max': opts['max'],
        'offset': opts['offset'],
        'phrase': opts['phrase'],
        'query': opts['query']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Search;
      return this.apiClient.callApi(
        '/api/search', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
