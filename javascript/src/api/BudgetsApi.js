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
import InlineObject22 from '../model/InlineObject22';
import InlineObject23 from '../model/InlineObject23';
import Model200Success from '../model/Model200Success';

/**
* Budgets service.
* @module api/BudgetsApi
* @version 6.2.1
*/
export default class BudgetsApi {

    /**
    * Constructs a new BudgetsApi. 
    * @alias module:api/BudgetsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the addBudgets operation.
     * @callback module:api/BudgetsApi~addBudgetsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Creates a Budget
     * Creates a budget. 
     * @param {Object} opts Optional parameters
     * @param {module:model/InlineObject22} opts.inlineObject22 
     * @param {module:api/BudgetsApi~addBudgetsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    addBudgets(opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject22'];

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = Object;
      return this.apiClient.callApi(
        '/api/budgets', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getBudgets operation.
     * @callback module:api/BudgetsApi~getBudgetsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves a Specific Budget
     * Retrieves a specific budget. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/BudgetsApi~getBudgetsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    getBudgets(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getBudgets");
      }

      let pathParams = {
        'id': id
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
      let returnType = Object;
      return this.apiClient.callApi(
        '/api/budgets/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listBudgets operation.
     * @callback module:api/BudgetsApi~listBudgetsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves all Budgets
     * Retrieves all budgets. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.direction Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {String} opts.name Filter by name, wildcard may be specified as %, eg. example-%
     * @param {module:api/BudgetsApi~listBudgetsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    listBudgets(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'max': opts['max'],
        'offset': opts['offset'],
        'sort': opts['sort'],
        'direction': opts['direction'],
        'phrase': opts['phrase'],
        'name': opts['name']
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
        '/api/budgets', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the removeBudgets operation.
     * @callback module:api/BudgetsApi~removeBudgetsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Deletes a Budget
     * Deletes a specified Budget. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/BudgetsApi~removeBudgetsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    removeBudgets(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling removeBudgets");
      }

      let pathParams = {
        'id': id
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
      let returnType = Model200Success;
      return this.apiClient.callApi(
        '/api/budgets/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateBudgets operation.
     * @callback module:api/BudgetsApi~updateBudgetsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Updates a Budget
     * Updates a budget. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {module:model/InlineObject23} opts.inlineObject23 
     * @param {module:api/BudgetsApi~updateBudgetsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    updateBudgets(id, opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject23'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling updateBudgets");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = Object;
      return this.apiClient.callApi(
        '/api/budgets/{id}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
