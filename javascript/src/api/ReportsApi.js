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
import InlineObject205 from '../model/InlineObject205';
import InlineResponse200129 from '../model/InlineResponse200129';
import InlineResponse200130 from '../model/InlineResponse200130';
import Model200Success from '../model/Model200Success';

/**
* Reports service.
* @module api/ReportsApi
* @version 6.2.1
*/
export default class ReportsApi {

    /**
    * Constructs a new ReportsApi. 
    * @alias module:api/ReportsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the downloadReports operation.
     * @callback module:api/ReportsApi~downloadReportsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Downloads a specific report result as a file attachment
     * This endpoint downloads a specific report result as a file attachment. The default file format is `.json`. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:model/String} format Format of the rendered report file, `.json` or `.csv`.
     * @param {module:api/ReportsApi~downloadReportsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    downloadReports(id, format, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling downloadReports");
      }
      // verify the required parameter 'format' is set
      if (format === undefined || format === null) {
        throw new Error("Missing the required parameter 'format' when calling downloadReports");
      }

      let pathParams = {
        'id': id,
        'format': format
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
        '/api/reports/download/{id}{format}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getReportTypes operation.
     * @callback module:api/ReportsApi~getReportTypesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse200129} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * This endpoint retrieves all available report types
     * This endpoint retrieves all available report types. A report type has optionTypes that define the parameters available when executing a report of that type. The sample response has been abbreviated. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.direction Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {String} opts.name Filter by name, wildcard may be specified as %, eg. example-%
     * @param {String} opts.code If specified will return an exact match on code
     * @param {String} opts.category If specified will return an exact match on category
     * @param {module:api/ReportsApi~getReportTypesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse200129}
     */
    getReportTypes(opts, callback) {
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
        'name': opts['name'],
        'code': opts['code'],
        'category': opts['category']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = InlineResponse200129;
      return this.apiClient.callApi(
        '/api/report-types', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getReports operation.
     * @callback module:api/ReportsApi~getReportsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse200130} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * This endpoint returns a specific report result
     * This endpoint retrieves a specific report result. The response includes the result data as rows which can be used to render the report. Each report type will have sections for data and headers that vary by type, use Download a Specific Report to get the results organized by section. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ReportsApi~getReportsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse200130}
     */
    getReports(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getReports");
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
      let returnType = InlineResponse200130;
      return this.apiClient.callApi(
        '/api/reports/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listReports operation.
     * @callback module:api/ReportsApi~listReportsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Returns all reports
     * This endpoint returns all reports. This is results of reports that have been executed in the past. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.direction Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {String} opts.name Filter by name, wildcard may be specified as %, eg. example-%
     * @param {String} opts.reportType If specified will return an exact match on report type code, accepts multiple values
     * @param {String} opts.category If specified will return an exact match on category
     * @param {module:api/ReportsApi~listReportsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    listReports(opts, callback) {
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
        'name': opts['name'],
        'reportType': opts['reportType'],
        'category': opts['category']
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
        '/api/reports', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the removeReports operation.
     * @callback module:api/ReportsApi~removeReportsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * This endpoint will delete a report result
     * This endpoint will delete a report result. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ReportsApi~removeReportsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    removeReports(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling removeReports");
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
        '/api/reports/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the runReports operation.
     * @callback module:api/ReportsApi~runReportsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Run specified report
     * This endpoint execute the specified report type and create a new report result.  The available parameters vary by report type. Refer to the defined `inputs` for each report. 
     * @param {Object} opts Optional parameters
     * @param {module:model/InlineObject205} opts.inlineObject205 
     * @param {module:api/ReportsApi~runReportsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    runReports(opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject205'];

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
        '/api/reports', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}