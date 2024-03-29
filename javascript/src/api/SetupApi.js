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
import AnyOfobjectobject from '../model/AnyOfobjectobject';
import DefaultError from '../model/DefaultError';
import Setup from '../model/Setup';
import UNKNOWN_BASE_TYPE from '../model/UNKNOWN_BASE_TYPE';

/**
* Setup service.
* @module api/SetupApi
* @version 6.2.1
*/
export default class SetupApi {

    /**
    * Constructs a new SetupApi. 
    * @alias module:api/SetupApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the setup operation.
     * @callback module:api/SetupApi~setupCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Setup} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Setup appliance
     * Initialize a freshly installed appliance to create the master tenant and System Admin user.  Authorization is not required.  This operation can only be executed successfully once. 
     * @param {Object} opts Optional parameters
     * @param {module:model/UNKNOWN_BASE_TYPE} opts.UNKNOWN_BASE_TYPE 
     * @param {module:api/SetupApi~setupCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Setup}
     */
    setup(opts, callback) {
      opts = opts || {};
      let postBody = opts['UNKNOWN_BASE_TYPE'];

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = Setup;
      return this.apiClient.callApi(
        '/api/setup', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
