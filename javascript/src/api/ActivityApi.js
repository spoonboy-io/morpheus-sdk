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

/**
* Activity service.
* @module api/ActivityApi
* @version 6.2.1
*/
export default class ActivityApi {

    /**
    * Constructs a new ActivityApi. 
    * @alias module:api/ActivityApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the listActivity operation.
     * @callback module:api/ActivityApi~listActivityCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves Activity
     * Retrieves activity. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.order Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {String} opts.name Filter by name, wildcard may be specified as %, eg. example-%
     * @param {Number} opts.userId Filter by User ID
     * @param {Number} opts.tenantId Filter by Tenant ID. Only available to the master account.
     * @param {String} opts.timeframe Filter by a timeframe (ex - today, yesterday, week, month, 3months) (default to 'month')
     * @param {Date} opts.start Filter by activity on or after a date(time). Default is 1 month prior
     * @param {Date} opts.end Filter by activity on or before a date(time). Default is current date
     * @param {module:api/ActivityApi~listActivityCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    listActivity(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'max': opts['max'],
        'offset': opts['offset'],
        'sort': opts['sort'],
        'order': opts['order'],
        'phrase': opts['phrase'],
        'name': opts['name'],
        'userId': opts['userId'],
        'tenantId': opts['tenantId'],
        'timeframe': opts['timeframe'],
        'start': opts['start'],
        'end': opts['end']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Object;
      return this.apiClient.callApi(
        '/api/activity', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
