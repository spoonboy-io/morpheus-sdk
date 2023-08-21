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
import Ping from '../model/Ping';

/**
* Ping service.
* @module api/PingApi
* @version 6.2.1
*/
export default class PingApi {

    /**
    * Constructs a new PingApi. 
    * @alias module:api/PingApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the ping operation.
     * @callback module:api/PingApi~pingCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Ping} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Basic information about current Morpheus Installation
     * This endpoint can be used to check the remote appliance build version and some other basic information.  This is an unsecured endpoint and does not require authorization. However, build version will not be returned unless you are authenticated with a valid access token. 
     * @param {module:api/PingApi~pingCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Ping}
     */
    ping(callback) {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Ping;
      return this.apiClient.callApi(
        '/api/ping', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}