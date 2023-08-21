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
import AnyOfobjectmeta from '../model/AnyOfobjectmeta';
import DefaultError from '../model/DefaultError';
import InlineObject100 from '../model/InlineObject100';
import InlineObject101 from '../model/InlineObject101';
import InlineResponse20061 from '../model/InlineResponse20061';
import InlineResponse20062 from '../model/InlineResponse20062';
import InlineResponse20063 from '../model/InlineResponse20063';
import InlineResponse20064 from '../model/InlineResponse20064';
import InlineResponse20065 from '../model/InlineResponse20065';
import Model200Success from '../model/Model200Success';
import OneOfintegrationConfigintegrationAnsibleConfigintegrationSNOWConfigintegrationSaltMasterConfigintegrationDockerRepoConfigintegrationGitRepoConfigintegrationGitHubConfig from '../model/OneOfintegrationConfigintegrationAnsibleConfigintegrationSNOWConfigintegrationSaltMasterConfigintegrationDockerRepoConfigintegrationGitRepoConfigintegrationGitHubConfig';
import UNKNOWN_BASE_TYPE from '../model/UNKNOWN_BASE_TYPE';

/**
* Integrations service.
* @module api/IntegrationsApi
* @version 6.2.1
*/
export default class IntegrationsApi {

    /**
    * Constructs a new IntegrationsApi. 
    * @alias module:api/IntegrationsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the addIntegrationSnowObjects operation.
     * @callback module:api/IntegrationsApi~addIntegrationSnowObjectsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Creates an Exposed ServiceNow Catalog Item
     * This endpoint creates an Exposed Catalog Item. This is an integration object of type `catalog` that references a `Catalog Item Type.` 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {module:model/InlineObject100} opts.inlineObject100 
     * @param {module:api/IntegrationsApi~addIntegrationSnowObjectsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    addIntegrationSnowObjects(id, opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject100'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling addIntegrationSnowObjects");
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
        '/api/integrations/{id}/objects', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the addIntegrations operation.
     * @callback module:api/IntegrationsApi~addIntegrationsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Creates an Integration
     * Creates an integration. 
     * @param {Object} opts Optional parameters
     * @param {module:model/UNKNOWN_BASE_TYPE} opts.UNKNOWN_BASE_TYPE 
     * @param {module:api/IntegrationsApi~addIntegrationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    addIntegrations(opts, callback) {
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

      let authNames = ['bearerAuth'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = Object;
      return this.apiClient.callApi(
        '/api/integrations', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getIntegrationInventory operation.
     * @callback module:api/IntegrationsApi~getIntegrationInventoryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20065} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a Specific Integration Inventory
     * This endpoint retrieves a specific integration inventory. Only certain types of integrations support this operation, such as Ansible Tower. 
     * @param {Number} id Morpheus ID of the Integration
     * @param {Number} inventoryId Morpheus ID of the Integration Inventory
     * @param {module:api/IntegrationsApi~getIntegrationInventoryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20065}
     */
    getIntegrationInventory(id, inventoryId, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getIntegrationInventory");
      }
      // verify the required parameter 'inventoryId' is set
      if (inventoryId === undefined || inventoryId === null) {
        throw new Error("Missing the required parameter 'inventoryId' when calling getIntegrationInventory");
      }

      let pathParams = {
        'id': id,
        'inventoryId': inventoryId
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
      let returnType = InlineResponse20065;
      return this.apiClient.callApi(
        '/api/integrations/{id}/inventory/{inventoryId}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getIntegrationObjects operation.
     * @callback module:api/IntegrationsApi~getIntegrationObjectsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20064} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a Specific ServiceNow Integration Object
     * This endpoint retrieves a specific ServiceNow integration object.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Number} objectId Morpheus ID of the Object being created or referenced
     * @param {module:api/IntegrationsApi~getIntegrationObjectsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20064}
     */
    getIntegrationObjects(id, objectId, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getIntegrationObjects");
      }
      // verify the required parameter 'objectId' is set
      if (objectId === undefined || objectId === null) {
        throw new Error("Missing the required parameter 'objectId' when calling getIntegrationObjects");
      }

      let pathParams = {
        'id': id,
        'objectId': objectId
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
      let returnType = InlineResponse20064;
      return this.apiClient.callApi(
        '/api/integrations/{id}/objects/{objectId}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getIntegrationTypeOptionTypes operation.
     * @callback module:api/IntegrationsApi~getIntegrationTypeOptionTypesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20062} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves a Option Types for a Specific Integration Type
     * This endpoint will retrieve the list of option types for a specific integration type. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/IntegrationsApi~getIntegrationTypeOptionTypesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20062}
     */
    getIntegrationTypeOptionTypes(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getIntegrationTypeOptionTypes");
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
      let returnType = InlineResponse20062;
      return this.apiClient.callApi(
        '/api/integration-types/{id}/option-types', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getIntegrationTypes operation.
     * @callback module:api/IntegrationsApi~getIntegrationTypesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20061} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves a Specific Integration Type
     * Retrieves a specific integration type. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {Boolean} opts.optiontypes Pass `true` to include optionTypes in the response for each integration type (default to false)
     * @param {module:api/IntegrationsApi~getIntegrationTypesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20061}
     */
    getIntegrationTypes(id, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getIntegrationTypes");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        'optiontypes': opts['optiontypes']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = InlineResponse20061;
      return this.apiClient.callApi(
        '/api/integration-types/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getIntegrations operation.
     * @callback module:api/IntegrationsApi~getIntegrationsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves a Specific Integration
     * Retrieves a specific integration. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/IntegrationsApi~getIntegrationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    getIntegrations(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getIntegrations");
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
        '/api/integrations/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listIntegrationInventory operation.
     * @callback module:api/IntegrationsApi~listIntegrationInventoryCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get All Integration Inventory
     * This endpoint retrieves a list of inventory for a specific integration. Only certain types of integrations support this operation, such as Ansible Tower. 
     * @param {Number} id Morpheus ID of the Integration
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.direction Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {String} opts.name Filter by name, wildcard may be specified as %, eg. example-%
     * @param {module:api/IntegrationsApi~listIntegrationInventoryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    listIntegrationInventory(id, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling listIntegrationInventory");
      }

      let pathParams = {
        'id': id
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
        '/api/integrations/{id}/inventory', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listIntegrationObjects operation.
     * @callback module:api/IntegrationsApi~listIntegrationObjectsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse20063} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get ServiceNow Integration Objects
     * This endpoint retrieves a list of exposed objects for a specific ServiceNow integration. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.direction Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {String} opts.name Filter by name, wildcard may be specified as %, eg. example-%
     * @param {module:model/String} opts.value The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing `type=string`. This means the data value returned by the API will be a string instead of an object. 
     * @param {Number} opts.refId If specified will return an exact match on refId
     * @param {module:api/IntegrationsApi~listIntegrationObjectsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse20063}
     */
    listIntegrationObjects(id, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling listIntegrationObjects");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        'max': opts['max'],
        'offset': opts['offset'],
        'sort': opts['sort'],
        'direction': opts['direction'],
        'phrase': opts['phrase'],
        'name': opts['name'],
        'value': opts['value'],
        'refId': opts['refId']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = InlineResponse20063;
      return this.apiClient.callApi(
        '/api/integrations/{id}/objects', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listIntegrationTypes operation.
     * @callback module:api/IntegrationsApi~listIntegrationTypesCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves all Integration Types
     * An Integration Type is specific third party software that the appliance can be configured to integrate with, such as Ansible, Chef, Git, ServiceNOW, etc. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.direction Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {String} opts.name Filter by name, wildcard may be specified as %, eg. example-%
     * @param {String} opts.code If specified will return an exact match on code
     * @param {Boolean} opts.optiontypes Pass `true` to include optionTypes in the response for each integration type (default to false)
     * @param {String} opts.description Filter by description, wildcard may be specified as %. eg. `example-%`
     * @param {String} opts.category If specified will return an exact match on category
     * @param {Boolean} opts.creatable Filter by creatable
     * @param {Boolean} opts.enabled Filter by enabled
     * @param {Boolean} opts.hasCMDB Filter by hasCMDB
     * @param {Boolean} opts.hasCMDBDiscovery Filter by hasCMDBDiscovery
     * @param {Boolean} opts.hasCM Filter by hasCM
     * @param {Boolean} opts.hasDNS Filter by hasDNS
     * @param {Boolean} opts.hasApprovals Filter by hasApprovals
     * @param {Boolean} opts.isPlugin Filter by isPlugin
     * @param {module:api/IntegrationsApi~listIntegrationTypesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    listIntegrationTypes(opts, callback) {
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
        'optiontypes': opts['optiontypes'],
        'description': opts['description'],
        'category': opts['category'],
        'creatable': opts['creatable'],
        'enabled': opts['enabled'],
        'hasCMDB': opts['hasCMDB'],
        'hasCMDBDiscovery': opts['hasCMDBDiscovery'],
        'hasCM': opts['hasCM'],
        'hasDNS': opts['hasDNS'],
        'hasApprovals': opts['hasApprovals'],
        'isPlugin': opts['isPlugin']
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
        '/api/integration-types', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listIntegrations operation.
     * @callback module:api/IntegrationsApi~listIntegrationsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/AnyOfobjectmeta} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves all Integrations
     * Retrieves all integrations. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.direction Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on name or description
     * @param {String} opts.name Filter by name, wildcard may be specified as %, eg. example-%
     * @param {Number} opts.id Morpheus ID of the Object being created or referenced
     * @param {String} opts.url Download the file from a remote url. This can be used instead of uploading a local file.
     * @param {String} opts.host Filter by integration Host
     * @param {String} opts.port Filter by integration Port
     * @param {String} opts.username Username
     * @param {Number} opts.version Filter by version number (userVersion)
     * @param {String} opts.windowsVersion Filter by integration Windows Version
     * @param {String} opts.status The instance status for filtering.
     * @param {Boolean} opts.objects Include `objects=true` to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2.  (default to false)
     * @param {module:api/IntegrationsApi~listIntegrationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/AnyOfobjectmeta}
     */
    listIntegrations(opts, callback) {
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
        'id': opts['id'],
        'url': opts['url'],
        'host': opts['host'],
        'port': opts['port'],
        'username': opts['username'],
        'version': opts['version'],
        'windowsVersion': opts['windowsVersion'],
        'status': opts['status'],
        'objects': opts['objects']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = AnyOfobjectmeta;
      return this.apiClient.callApi(
        '/api/integrations', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the refreshIntegrations operation.
     * @callback module:api/IntegrationsApi~refreshIntegrationsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Refresh an Integration
     * This endpoint will refresh an existing Integration. Only some types support this and will actually perform an action as a result. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/IntegrationsApi~refreshIntegrationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    refreshIntegrations(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling refreshIntegrations");
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
        '/api/integrations/{id}/refresh', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the removeIntegrationObjects operation.
     * @callback module:api/IntegrationsApi~removeIntegrationObjectsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Deletes a ServiceNow Integration object
     * Deletes a specified ServiceNow integration object. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Number} objectId Morpheus ID of the Object being created or referenced
     * @param {module:api/IntegrationsApi~removeIntegrationObjectsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    removeIntegrationObjects(id, objectId, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling removeIntegrationObjects");
      }
      // verify the required parameter 'objectId' is set
      if (objectId === undefined || objectId === null) {
        throw new Error("Missing the required parameter 'objectId' when calling removeIntegrationObjects");
      }

      let pathParams = {
        'id': id,
        'objectId': objectId
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
        '/api/integrations/{id}/objects/{objectId}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the removeIntegrations operation.
     * @callback module:api/IntegrationsApi~removeIntegrationsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Deletes an Integration
     * Deletes a specified integration. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/IntegrationsApi~removeIntegrationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    removeIntegrations(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling removeIntegrations");
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
        '/api/integrations/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateIntegrationInventory operation.
     * @callback module:api/IntegrationsApi~updateIntegrationInventoryCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Updating an Integration Inventory
     * This endpoint provides updating of integration inventory
     * @param {Number} id Morpheus ID of the Integration
     * @param {Number} inventoryId Morpheus ID of the Integration Inventory
     * @param {Object} opts Optional parameters
     * @param {module:model/InlineObject101} opts.inlineObject101 
     * @param {module:api/IntegrationsApi~updateIntegrationInventoryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    updateIntegrationInventory(id, inventoryId, opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject101'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling updateIntegrationInventory");
      }
      // verify the required parameter 'inventoryId' is set
      if (inventoryId === undefined || inventoryId === null) {
        throw new Error("Missing the required parameter 'inventoryId' when calling updateIntegrationInventory");
      }

      let pathParams = {
        'id': id,
        'inventoryId': inventoryId
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
        '/api/integrations/{id}/inventory/{inventoryId}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateIntegrations operation.
     * @callback module:api/IntegrationsApi~updateIntegrationsCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Updates an Integration
     * Updates an integration. 
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {module:model/UNKNOWN_BASE_TYPE} opts.UNKNOWN_BASE_TYPE 
     * @param {module:api/IntegrationsApi~updateIntegrationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    updateIntegrations(id, opts, callback) {
      opts = opts || {};
      let postBody = opts['UNKNOWN_BASE_TYPE'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling updateIntegrations");
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
        '/api/integrations/{id}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}