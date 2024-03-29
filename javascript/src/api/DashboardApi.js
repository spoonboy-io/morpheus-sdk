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
import Dashboard from '../model/Dashboard';
import DefaultError from '../model/DefaultError';

/**
* Dashboard service.
* @module api/DashboardApi
* @version 6.2.1
*/
export default class DashboardApi {

    /**
    * Constructs a new DashboardApi. 
    * @alias module:api/DashboardApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the dashboard operation.
     * @callback module:api/DashboardApi~dashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Overview of Dashboard information
     * This endpoint can be used to view dashboard information about the remote Morpheus appliance. This is an overview and summary of data available to the user that can be used to render a dashboard. 
     * @param {module:api/DashboardApi~dashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Dashboard}
     */
    dashboard(callback) {
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
      let returnType = Dashboard;
      return this.apiClient.callApi(
        '/api/dashboard', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
