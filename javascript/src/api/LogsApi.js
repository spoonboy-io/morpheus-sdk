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
import Log from '../model/Log';

/**
* Logs service.
* @module api/LogsApi
* @version 6.2.1
*/
export default class LogsApi {

    /**
    * Constructs a new LogsApi. 
    * @alias module:api/LogsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the listLogs operation.
     * @callback module:api/LogsApi~listLogsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Log} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves Logs
     * Retrieves logs based on filters provided. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.order Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.query Alias for phrase
     * @param {String} opts.message Filter by message
     * @param {String} opts.sourceType Filter by source type
     * @param {String} opts.typeCode Filter by code type
     * @param {Number} opts.objectId Filter by objectId
     * @param {String} opts.token Filter by token
     * @param {module:model/String} opts.level Filter by log level. Multiple values can be passed pipe delimited.
     * @param {Number} opts.startMs Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified.
     * @param {Number} opts.endMs Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified.
     * @param {Date} opts.startDateTime Start Date timestamp (ISO 8601)
     * @param {Date} opts.endDateTime End Date timestamp (ISO 8601)
     * @param {Number} opts.containers The Container ID(s) for filtering. Accepts multiple values.
     * @param {Number} opts.servers The Server ID(s) for filtering. Accepts multiple values.
     * @param {Number} opts.clusterId The Cluster ID(s) for filtering. Accepts multiple values.
     * @param {module:api/LogsApi~listLogsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Log}
     */
    listLogs(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'max': opts['max'],
        'offset': opts['offset'],
        'sort': opts['sort'],
        'order': opts['order'],
        'query': opts['query'],
        'message': opts['message'],
        'sourceType': opts['sourceType'],
        'typeCode': opts['typeCode'],
        'objectId': opts['objectId'],
        'token': opts['token'],
        'level': opts['level'],
        'startMs': opts['startMs'],
        'endMs': opts['endMs'],
        'startDateTime': opts['startDateTime'],
        'endDateTime': opts['endDateTime'],
        'containers': opts['containers'],
        'servers': opts['servers'],
        'clusterId': opts['clusterId']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Log;
      return this.apiClient.callApi(
        '/api/logs', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
