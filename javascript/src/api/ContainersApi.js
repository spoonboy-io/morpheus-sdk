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
import InlineObject62 from '../model/InlineObject62';
import InlineObject63 from '../model/InlineObject63';
import InlineResponse20034 from '../model/InlineResponse20034';
import InlineResponse20035 from '../model/InlineResponse20035';
import InstancesCloneImage from '../model/InstancesCloneImage';
import Model200Success from '../model/Model200Success';
import SuccessMessage from '../model/SuccessMessage';

/**
* Containers service.
* @module api/ContainersApi
* @version 6.2.1
*/
export default class ContainersApi {

    /**
    * Constructs a new ContainersApi. 
    * @alias module:api/ContainersApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the cloneImageContainerAction operation.
     * @callback module:api/ContainersApi~cloneImageContainerActionCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SuccessMessage} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Clone Specific Container to Image
     * This endpoint clones and converts a Container in its current state to image in the source Cloud and adds Virtual Image Record with metadata matching the source configuration.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {module:model/InstancesCloneImage} opts.instancesCloneImage 
     * @param {module:api/ContainersApi~cloneImageContainerActionCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SuccessMessage}
     */
    cloneImageContainerAction(id, opts, callback) {
      opts = opts || {};
      let postBody = opts['instancesCloneImage'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling cloneImageContainerAction");
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
      let returnType = SuccessMessage;
      return this.apiClient.callApi(
        '/api/containers/{id}/clone-image', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the containersAttachFloatingIp operation.
     * @callback module:api/ContainersApi~containersAttachFloatingIpCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Attach Floating IP to Container
     * This endpoint attaches a floating IP to a container (node/VM).
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {module:model/InlineObject63} opts.inlineObject63 
     * @param {module:api/ContainersApi~containersAttachFloatingIpCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    containersAttachFloatingIp(id, opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject63'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling containersAttachFloatingIp");
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
      let returnType = Model200Success;
      return this.apiClient.callApi(
        '/api/containers/{id}/attach-floating-ip', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the containersDetachFloatingIp operation.
     * @callback module:api/ContainersApi~containersDetachFloatingIpCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Detach Floating IP from Container
     * This endpoint detaches a floating IP from a container (node/VM).
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ContainersApi~containersDetachFloatingIpCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    containersDetachFloatingIp(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling containersDetachFloatingIp");
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
        '/api/containers/{id}/detach-floating-ip', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the ejectContainerAction operation.
     * @callback module:api/ContainersApi~ejectContainerActionCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SuccessMessage} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Eject a Specific Container
     * This endpoint ejects attached disk/iso of a specific container.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ContainersApi~ejectContainerActionCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SuccessMessage}
     */
    ejectContainerAction(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling ejectContainerAction");
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
      let returnType = SuccessMessage;
      return this.apiClient.callApi(
        '/api/containers/{id}/eject', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the executeContainerAction operation.
     * @callback module:api/ContainersApi~executeContainerActionCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SuccessMessage} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Execute Container Action
     * This endpoint executes a container action on specific container.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {String} code Action code to be executed on a specific container.
     * @param {module:api/ContainersApi~executeContainerActionCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SuccessMessage}
     */
    executeContainerAction(id, code, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling executeContainerAction");
      }
      // verify the required parameter 'code' is set
      if (code === undefined || code === null) {
        throw new Error("Missing the required parameter 'code' when calling executeContainerAction");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        'code': code
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = SuccessMessage;
      return this.apiClient.callApi(
        '/api/containers/{id}/action', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getContainer operation.
     * @callback module:api/ContainersApi~getContainerCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20034} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a Specific Container
     * This endpoint retrieves a specific container.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ContainersApi~getContainerCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20034}
     */
    getContainer(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getContainer");
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
      let returnType = InlineResponse20034;
      return this.apiClient.callApi(
        '/api/containers/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getContainerActions operation.
     * @callback module:api/ContainersApi~getContainerActionsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20035} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List Container Actions
     * This endpoint lists available actions for specific container.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ContainersApi~getContainerActionsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20035}
     */
    getContainerActions(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getContainerActions");
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
      let returnType = InlineResponse20035;
      return this.apiClient.callApi(
        '/api/containers/{id}/actions', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the importContainerAction operation.
     * @callback module:api/ContainersApi~importContainerActionCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SuccessMessage} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Import a Specific Container
     * This endpoint clones and exports a Container in its current state to target Storage provider and adds Virtual Image Record with metadata matching the source configuration.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {module:model/InlineObject62} opts.inlineObject62 
     * @param {module:api/ContainersApi~importContainerActionCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SuccessMessage}
     */
    importContainerAction(id, opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject62'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling importContainerAction");
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
      let returnType = SuccessMessage;
      return this.apiClient.callApi(
        '/api/containers/{id}/import', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the restartContainerAction operation.
     * @callback module:api/ContainersApi~restartContainerActionCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SuccessMessage} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Restart a Specific Container
     * This endpoint restarts a specific container.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ContainersApi~restartContainerActionCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SuccessMessage}
     */
    restartContainerAction(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling restartContainerAction");
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
      let returnType = SuccessMessage;
      return this.apiClient.callApi(
        '/api/containers/{id}/restart', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the startContainerAction operation.
     * @callback module:api/ContainersApi~startContainerActionCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SuccessMessage} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Start a Specific Container
     * This endpoint starts a specific container.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ContainersApi~startContainerActionCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SuccessMessage}
     */
    startContainerAction(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling startContainerAction");
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
      let returnType = SuccessMessage;
      return this.apiClient.callApi(
        '/api/containers/{id}/start', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the stopContainerAction operation.
     * @callback module:api/ContainersApi~stopContainerActionCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SuccessMessage} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Stop a Specific Container
     * This endpoint stops a specific container.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ContainersApi~stopContainerActionCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SuccessMessage}
     */
    stopContainerAction(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling stopContainerAction");
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
      let returnType = SuccessMessage;
      return this.apiClient.callApi(
        '/api/containers/{id}/stop', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the suspendContainerAction operation.
     * @callback module:api/ContainersApi~suspendContainerActionCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SuccessMessage} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Suspend a Specific Container
     * This endpoint suspends a specific container.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/ContainersApi~suspendContainerActionCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SuccessMessage}
     */
    suspendContainerAction(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling suspendContainerAction");
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
      let returnType = SuccessMessage;
      return this.apiClient.callApi(
        '/api/containers/{id}/suspend', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
