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
import InlineResponse20044 from '../model/InlineResponse20044';
import InlineResponse20045 from '../model/InlineResponse20045';
import InlineResponse20046 from '../model/InlineResponse20046';
import Model200Success from '../model/Model200Success';

/**
* Guidance service.
* @module api/GuidanceApi
* @version 6.2.1
*/
export default class GuidanceApi {

    /**
    * Constructs a new GuidanceApi. 
    * @alias module:api/GuidanceApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the executeGuidances operation.
     * @callback module:api/GuidanceApi~executeGuidancesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Executes a Specific Guidance Recommendation
     * Executes a specific guidance recommendation. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/GuidanceApi~executeGuidancesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    executeGuidances(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling executeGuidances");
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
        '/api/guidance/{id}/execute', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getGuidanceStats operation.
     * @callback module:api/GuidanceApi~getGuidanceStatsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20045} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves Guidance Stats
     * This endpoint retrieves a summary of actionable guidance. 
     * @param {module:api/GuidanceApi~getGuidanceStatsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20045}
     */
    getGuidanceStats(callback) {
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
      let returnType = InlineResponse20045;
      return this.apiClient.callApi(
        '/api/guidance/stats', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getGuidanceTypes operation.
     * @callback module:api/GuidanceApi~getGuidanceTypesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20046} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves Guidance Types
     * This endpoint retrieves all guidance types. 
     * @param {module:api/GuidanceApi~getGuidanceTypesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20046}
     */
    getGuidanceTypes(callback) {
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
      let returnType = InlineResponse20046;
      return this.apiClient.callApi(
        '/api/guidance/types', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getGuidances operation.
     * @callback module:api/GuidanceApi~getGuidancesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20044} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves a Specific Guidance Recommendation
     * Retrieves a specific guidance recommendation. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/GuidanceApi~getGuidancesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20044}
     */
    getGuidances(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getGuidances");
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
      let returnType = InlineResponse20044;
      return this.apiClient.callApi(
        '/api/guidance/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the ignoreGuidances operation.
     * @callback module:api/GuidanceApi~ignoreGuidancesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Ignores a Specific Guidance Recommendation
     * Ignores a specific guidance recommendation. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/GuidanceApi~ignoreGuidancesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    ignoreGuidances(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling ignoreGuidances");
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
        '/api/guidance/{id}/ignore', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listGuidances operation.
     * @callback module:api/GuidanceApi~listGuidancesCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves all Guidance Recommendations
     * Retrieves all guidance recommendations. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.direction Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {module:model/String} opts.severity Filter by severity
     * @param {module:model/String} opts.type Filter by Guidance Type.  See `Retrieves Guidance Types` for most up to date list of types.
     * @param {module:model/String} opts.state Filter by state, restricts query to only load discoveries by state
     * @param {module:api/GuidanceApi~listGuidancesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    listGuidances(opts, callback) {
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
        'severity': opts['severity'],
        'type': opts['type'],
        'state': opts['state']
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
        '/api/guidance', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}