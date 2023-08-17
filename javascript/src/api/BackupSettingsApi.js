/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import DefaultError from '../model/DefaultError';
import ListBackupSettings200Response from '../model/ListBackupSettings200Response';
import Model200Success from '../model/Model200Success';
import UpdateBackupSettingsRequest from '../model/UpdateBackupSettingsRequest';

/**
* BackupSettings service.
* @module api/BackupSettingsApi
* @version 6.1.1
*/
export default class BackupSettingsApi {

    /**
    * Constructs a new BackupSettingsApi. 
    * @alias module:api/BackupSettingsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the listBackupSettings operation.
     * @callback module:api/BackupSettingsApi~listBackupSettingsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ListBackupSettings200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get Backup Settings
     * This endpoint retrieves backup settings.
     * @param {module:api/BackupSettingsApi~listBackupSettingsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ListBackupSettings200Response}
     */
    listBackupSettings(callback) {
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
      let returnType = ListBackupSettings200Response;
      return this.apiClient.callApi(
        '/api/backup-settings', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateBackupSettings operation.
     * @callback module:api/BackupSettingsApi~updateBackupSettingsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update Backup Settings
     * Update Backup Settings
     * @param {Object} opts Optional parameters
     * @param {module:model/UpdateBackupSettingsRequest} [updateBackupSettingsRequest] 
     * @param {module:api/BackupSettingsApi~updateBackupSettingsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    updateBackupSettings(opts, callback) {
      opts = opts || {};
      let postBody = opts['updateBackupSettingsRequest'];

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
      let returnType = Model200Success;
      return this.apiClient.callApi(
        '/api/backup-settings', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}