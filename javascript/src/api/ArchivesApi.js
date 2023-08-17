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
import AddArchiveBucket200Response from '../model/AddArchiveBucket200Response';
import AddArchiveBucketRequest from '../model/AddArchiveBucketRequest';
import AddArchiveFile200Response from '../model/AddArchiveFile200Response';
import AddArchiveFileLink200Response from '../model/AddArchiveFileLink200Response';
import DefaultError from '../model/DefaultError';
import GetArchiveBucket200Response from '../model/GetArchiveBucket200Response';
import GetArchiveFileDetail200Response from '../model/GetArchiveFileDetail200Response';
import GetArchiveFileLinks200Response from '../model/GetArchiveFileLinks200Response';
import ListArchiveBuckets200Response from '../model/ListArchiveBuckets200Response';
import ListArchiveFiles200Response from '../model/ListArchiveFiles200Response';
import Model200Success from '../model/Model200Success';
import UpdateArchiveBucketRequest from '../model/UpdateArchiveBucketRequest';

/**
* Archives service.
* @module api/ArchivesApi
* @version 6.1.1
*/
export default class ArchivesApi {

    /**
    * Constructs a new ArchivesApi. 
    * @alias module:api/ArchivesApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the addArchiveBucket operation.
     * @callback module:api/ArchivesApi~addArchiveBucketCallback
     * @param {String} error Error message, if any.
     * @param {module:model/AddArchiveBucket200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create an Archive Bucket
     * Create an Archive Bucket
     * @param {Object} opts Optional parameters
     * @param {module:model/AddArchiveBucketRequest} [addArchiveBucketRequest] 
     * @param {module:api/ArchivesApi~addArchiveBucketCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/AddArchiveBucket200Response}
     */
    addArchiveBucket(opts, callback) {
      opts = opts || {};
      let postBody = opts['addArchiveBucketRequest'];

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
      let returnType = AddArchiveBucket200Response;
      return this.apiClient.callApi(
        '/api/archives/buckets', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the addArchiveFile operation.
     * @callback module:api/ArchivesApi~addArchiveFileCallback
     * @param {String} error Error message, if any.
     * @param {module:model/AddArchiveFile200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Upload Archive File
     * Upload a file to the specified archive bucket and file path.  This will overwrite the file if it already exists. 
     * @param {String} bucket Bucket name
     * @param {String} filepath The path to to search for files under. Default is the root directory /.
     * @param {Object} opts Optional parameters
     * @param {String} [filename] Specify a filename for archive file. The base filename of the uploaded file is the default.
     * @param {File} [file] 
     * @param {module:api/ArchivesApi~addArchiveFileCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/AddArchiveFile200Response}
     */
    addArchiveFile(bucket, filepath, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'bucket' is set
      if (bucket === undefined || bucket === null) {
        throw new Error("Missing the required parameter 'bucket' when calling addArchiveFile");
      }
      // verify the required parameter 'filepath' is set
      if (filepath === undefined || filepath === null) {
        throw new Error("Missing the required parameter 'filepath' when calling addArchiveFile");
      }

      let pathParams = {
        'bucket': bucket,
        'filepath': filepath
      };
      let queryParams = {
        'filename': opts['filename']
      };
      let headerParams = {
      };
      let formParams = {
        'file': opts['file']
      };

      let authNames = ['bearerAuth'];
      let contentTypes = ['multipart/form-data'];
      let accepts = ['application/json'];
      let returnType = AddArchiveFile200Response;
      return this.apiClient.callApi(
        '/api/archives/buckets/{bucket}/files/{filepath}', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the addArchiveFileLink operation.
     * @callback module:api/ArchivesApi~addArchiveFileLinkCallback
     * @param {String} error Error message, if any.
     * @param {module:model/AddArchiveFileLink200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create an Archive File Link
     * This returns a secret token that can be used to download the file via a public url, without any other authentication or authorization. File links can be set to expire after a certain amount of time.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {Number} [expireSeconds = 0)] Time to live in seconds. 0 means do not expire.
     * @param {module:api/ArchivesApi~addArchiveFileLinkCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/AddArchiveFileLink200Response}
     */
    addArchiveFileLink(id, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling addArchiveFileLink");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        'expireSeconds': opts['expireSeconds']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = AddArchiveFileLink200Response;
      return this.apiClient.callApi(
        '/api/archives/files/{id}/links', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteArchiveBucket operation.
     * @callback module:api/ArchivesApi~deleteArchiveBucketCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete an Archive Bucket
     * Will delete an archive bucket from the system and make it no longer usable.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ArchivesApi~deleteArchiveBucketCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    deleteArchiveBucket(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling deleteArchiveBucket");
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
        '/api/archives/buckets/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteArchiveFile operation.
     * @callback module:api/ArchivesApi~deleteArchiveFileCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete Archive File
     * Permanently delete a file or directory.  Deleting a directory will also delete all the files under it. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ArchivesApi~deleteArchiveFileCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    deleteArchiveFile(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling deleteArchiveFile");
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
        '/api/archives/files/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteArchiveFileLink operation.
     * @callback module:api/ArchivesApi~deleteArchiveFileLinkCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete an Archive File Link
     * This will delete the link from the system, so it can no longer be used.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Number} linkId The ID of the archive file link.
     * @param {module:api/ArchivesApi~deleteArchiveFileLinkCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    deleteArchiveFileLink(id, linkId, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling deleteArchiveFileLink");
      }
      // verify the required parameter 'linkId' is set
      if (linkId === undefined || linkId === null) {
        throw new Error("Missing the required parameter 'linkId' when calling deleteArchiveFileLink");
      }

      let pathParams = {
        'id': id,
        'linkId': linkId
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
        '/api/archives/files/{id}/links/{linkId}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getArchiveBucket operation.
     * @callback module:api/ArchivesApi~getArchiveBucketCallback
     * @param {String} error Error message, if any.
     * @param {module:model/GetArchiveBucket200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a Specific Archive Bucket
     * This endpoint retrieves a specific archive bucket.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ArchivesApi~getArchiveBucketCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/GetArchiveBucket200Response}
     */
    getArchiveBucket(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getArchiveBucket");
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
      let returnType = GetArchiveBucket200Response;
      return this.apiClient.callApi(
        '/api/archives/buckets/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getArchiveFile operation.
     * @callback module:api/ArchivesApi~getArchiveFileCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Download an Archive File
     * Download the file as an authorized user with access to the bucket.  Downloading a directory will return a .zip file containing all files under it. 
     * @param {String} bucket Bucket name
     * @param {String} filepath The path to to search for files under. Default is the root directory /.
     * @param {module:api/ArchivesApi~getArchiveFileCallback} callback The callback function, accepting three arguments: error, data, response
     */
    getArchiveFile(bucket, filepath, callback) {
      let postBody = null;
      // verify the required parameter 'bucket' is set
      if (bucket === undefined || bucket === null) {
        throw new Error("Missing the required parameter 'bucket' when calling getArchiveFile");
      }
      // verify the required parameter 'filepath' is set
      if (filepath === undefined || filepath === null) {
        throw new Error("Missing the required parameter 'filepath' when calling getArchiveFile");
      }

      let pathParams = {
        'bucket': bucket,
        'filepath': filepath
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
      let returnType = null;
      return this.apiClient.callApi(
        '/api/archives/download/{bucket}/{filepath}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getArchiveFileDetail operation.
     * @callback module:api/ArchivesApi~getArchiveFileDetailCallback
     * @param {String} error Error message, if any.
     * @param {module:model/GetArchiveFileDetail200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get Archive File Details
     * Get details about a specific archive file.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ArchivesApi~getArchiveFileDetailCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/GetArchiveFileDetail200Response}
     */
    getArchiveFileDetail(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getArchiveFileDetail");
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
      let returnType = GetArchiveFileDetail200Response;
      return this.apiClient.callApi(
        '/api/archives/files/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getArchiveFileLinks operation.
     * @callback module:api/ArchivesApi~getArchiveFileLinksCallback
     * @param {String} error Error message, if any.
     * @param {module:model/GetArchiveFileLinks200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get Archive File Links
     * This endpoint retrieves the links that have been created for the specified file.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ArchivesApi~getArchiveFileLinksCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/GetArchiveFileLinks200Response}
     */
    getArchiveFileLinks(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getArchiveFileLinks");
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
      let returnType = GetArchiveFileLinks200Response;
      return this.apiClient.callApi(
        '/api/archives/files/{id}/links', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listArchiveBuckets operation.
     * @callback module:api/ArchivesApi~listArchiveBucketsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ListArchiveBuckets200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get All Archive Buckets
     * This endpoint retrieves all archive buckets associated with the account.
     * @param {Object} opts Optional parameters
     * @param {String} [name] Filter by name, wildcard may be specified as %, eg. example-%
     * @param {String} [phrase] Search phrase for partial matches on name or description
     * @param {module:api/ArchivesApi~listArchiveBucketsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ListArchiveBuckets200Response}
     */
    listArchiveBuckets(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'name': opts['name'],
        'phrase': opts['phrase']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ListArchiveBuckets200Response;
      return this.apiClient.callApi(
        '/api/archives/buckets', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listArchiveFiles operation.
     * @callback module:api/ArchivesApi~listArchiveFilesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ListArchiveFiles200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get All Archive Files
     * This endpoint retrieves all files in an archive bucket under the specified `filePath`.
     * @param {String} bucket Bucket name
     * @param {String} filepath The path to to search for files under. Default is the root directory /.
     * @param {Object} opts Optional parameters
     * @param {String} [name] Filter by name, wildcard may be specified as %, eg. example-%
     * @param {String} [phrase] Search phrase for partial matches on name or description
     * @param {Boolean} [fullTree = false)] Include files under sub directories too. This is always true when searching with phrase.
     * @param {module:api/ArchivesApi~listArchiveFilesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ListArchiveFiles200Response}
     */
    listArchiveFiles(bucket, filepath, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'bucket' is set
      if (bucket === undefined || bucket === null) {
        throw new Error("Missing the required parameter 'bucket' when calling listArchiveFiles");
      }
      // verify the required parameter 'filepath' is set
      if (filepath === undefined || filepath === null) {
        throw new Error("Missing the required parameter 'filepath' when calling listArchiveFiles");
      }

      let pathParams = {
        'bucket': bucket,
        'filepath': filepath
      };
      let queryParams = {
        'name': opts['name'],
        'phrase': opts['phrase'],
        'fullTree': opts['fullTree']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ListArchiveFiles200Response;
      return this.apiClient.callApi(
        '/api/archives/buckets/{bucket}/files/{filepath}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateArchiveBucket operation.
     * @callback module:api/ArchivesApi~updateArchiveBucketCallback
     * @param {String} error Error message, if any.
     * @param {module:model/AddArchiveBucket200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update an Archive Bucket
     * Update an Archive Bucket
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {module:model/UpdateArchiveBucketRequest} [updateArchiveBucketRequest] 
     * @param {module:api/ArchivesApi~updateArchiveBucketCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/AddArchiveBucket200Response}
     */
    updateArchiveBucket(id, opts, callback) {
      opts = opts || {};
      let postBody = opts['updateArchiveBucketRequest'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling updateArchiveBucket");
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
      let returnType = AddArchiveBucket200Response;
      return this.apiClient.callApi(
        '/api/archives/buckets/{id}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}