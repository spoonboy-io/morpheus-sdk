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
import InlineObject252 from '../model/InlineObject252';
import InlineObject255 from '../model/InlineObject255';
import InlineObject256 from '../model/InlineObject256';
import InlineResponse200157 from '../model/InlineResponse200157';
import InlineResponse200158 from '../model/InlineResponse200158';
import InlineResponse200159 from '../model/InlineResponse200159';
import InlineResponse200167 from '../model/InlineResponse200167';
import Model200Success from '../model/Model200Success';
import UserSettings from '../model/UserSettings';
import UserSettingsRegenerateAccessToken from '../model/UserSettingsRegenerateAccessToken';
import UsersAvailableRoles from '../model/UsersAvailableRoles';

/**
* Users service.
* @module api/UsersApi
* @version 6.2.1
*/
export default class UsersApi {

    /**
    * Constructs a new UsersApi. 
    * @alias module:api/UsersApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the addUser operation.
     * @callback module:api/UsersApi~addUserCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create a New User
     * Create a new user.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.accountId Tenant ID, create user in a sub tenant account instead of your own.
     * @param {module:model/InlineObject255} opts.inlineObject255 
     * @param {module:api/UsersApi~addUserCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    addUser(opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject255'];

      let pathParams = {
      };
      let queryParams = {
        'accountId': opts['accountId']
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
        '/api/users', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteUser operation.
     * @callback module:api/UsersApi~deleteUserCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete a User
     * Delete an existing user.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/UsersApi~deleteUserCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    deleteUser(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling deleteUser");
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
        '/api/users/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteUserSettingsAccessToken operation.
     * @callback module:api/UsersApi~deleteUserSettingsAccessTokenCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Revoke API Access Token
     * This endpoint revokes your API access token for the specified client.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.userId ID of User (Only available to the master tenant.)  Default is the current user
     * @param {String} opts.clientId Client ID.  See `Get Available API Clients` for a list of valid `clientId` values.
     * @param {module:api/UsersApi~deleteUserSettingsAccessTokenCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    deleteUserSettingsAccessToken(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'userId': opts['userId'],
        'clientId': opts['clientId']
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
        '/api/user-settings/clear-access-token', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteUserSettingsAvatar operation.
     * @callback module:api/UsersApi~deleteUserSettingsAvatarCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete Avatar
     * Delete your avatar image.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.userId ID of User (Only available to the master tenant.)  Default is the current user
     * @param {module:api/UsersApi~deleteUserSettingsAvatarCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    deleteUserSettingsAvatar(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'userId': opts['userId']
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
        '/api/user-settings/avatar', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteUserSettingsDesktopBackground operation.
     * @callback module:api/UsersApi~deleteUserSettingsDesktopBackgroundCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete Desktop Background
     * Delete your desktop background image.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.userId ID of User (Only available to the master tenant.)  Default is the current user
     * @param {module:api/UsersApi~deleteUserSettingsDesktopBackgroundCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    deleteUserSettingsDesktopBackground(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'userId': opts['userId']
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
        '/api/user-settings/desktop-background', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getUser operation.
     * @callback module:api/UsersApi~getUserCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse200158} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a Specific User
     * This endpoint will retrieve a specific user by id if the user has permission to access the user.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {Boolean} opts.includeAccess Include `access` information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s).
     * @param {module:api/UsersApi~getUserCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse200158}
     */
    getUser(id, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getUser");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
        'includeAccess': opts['includeAccess']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = InlineResponse200158;
      return this.apiClient.callApi(
        '/api/users/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getUserPermissions operation.
     * @callback module:api/UsersApi~getUserPermissionsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse200159} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a Specific User Permissions
     * This will list all the permissions for a specific user.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {module:api/UsersApi~getUserPermissionsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse200159}
     */
    getUserPermissions(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getUserPermissions");
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
      let returnType = InlineResponse200159;
      return this.apiClient.callApi(
        '/api/users/{id}/permissions', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getUserSettingsApiClients operation.
     * @callback module:api/UsersApi~getUserSettingsApiClientsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse200157} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get Available API Clients
     * This endpoint retrieves a list of available API clients.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.userId ID of User (Only available to the master tenant.)  Default is the current user
     * @param {module:api/UsersApi~getUserSettingsApiClientsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse200157}
     */
    getUserSettingsApiClients(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'userId': opts['userId']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = InlineResponse200157;
      return this.apiClient.callApi(
        '/api/user-settings/api-clients', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listUserSettings operation.
     * @callback module:api/UsersApi~listUserSettingsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/UserSettings} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * User Settings
     * Provides API for managing your own user settings and api access tokens.  This endpoint retrieves your user settings and API access token information. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.userId ID of User (Only available to the master tenant.)  Default is the current user
     * @param {module:api/UsersApi~listUserSettingsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/UserSettings}
     */
    listUserSettings(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'userId': opts['userId']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = UserSettings;
      return this.apiClient.callApi(
        '/api/user-settings', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listUsers operation.
     * @callback module:api/UsersApi~listUsersCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List All Users
     * This endpoint retrieves all users in the current user's tenant account. Master tenant users with permission to manage subtenants can use `global=true` to find users across all tenants. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.max Maximum number of records to return, -1 can be used to fetch all records (default to 25)
     * @param {Number} opts.offset Offset records, the number of records to skip, for paginating requests (default to 0)
     * @param {String} opts.sort Sort order, the name of the property to sort by (default to 'name')
     * @param {module:model/String} opts.direction Sort direction, use 'desc' to reverse sort (default to 'asc')
     * @param {String} opts.phrase Search phrase for partial matches on username or email
     * @param {String} opts.username Username
     * @param {String} opts.email E-Mail Address
     * @param {Number} opts.roleId Filter by Role ID (supports multiple values)
     * @param {Date} opts.lastUpdated Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
     * @param {Number} opts.accountId Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
     * @param {Boolean} opts.global Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. (default to false)
     * @param {module:api/UsersApi~listUsersCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    listUsers(opts, callback) {
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
        'username': opts['username'],
        'email': opts['email'],
        'roleId': opts['roleId'],
        'lastUpdated': opts['lastUpdated'],
        'accountId': opts['accountId'],
        'global': opts['global']
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
        '/api/users', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listUsersAvailableRoles operation.
     * @callback module:api/UsersApi~listUsersAvailableRolesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/UsersAvailableRoles} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List available roles for a user
     * Get a list of roles that can be assigned to a user.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.accountId Tenant ID, find roles available to users of a sub tenant account.
     * @param {module:api/UsersApi~listUsersAvailableRolesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/UsersAvailableRoles}
     */
    listUsersAvailableRoles(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'accountId': opts['accountId']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = UsersAvailableRoles;
      return this.apiClient.callApi(
        '/api/users/available-roles', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateUserSettings operation.
     * @callback module:api/UsersApi~updateUserSettingsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update User Settings
     * This endpoint allows updating user settings. 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.userId ID of User (Only available to the master tenant.)  Default is the current user
     * @param {module:model/InlineObject252} opts.inlineObject252 
     * @param {module:api/UsersApi~updateUserSettingsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    updateUserSettings(opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject252'];

      let pathParams = {
      };
      let queryParams = {
        'userId': opts['userId']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = ['application/json', 'multipart/form-data'];
      let accepts = ['application/json'];
      let returnType = Model200Success;
      return this.apiClient.callApi(
        '/api/user-settings', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateUserSettingsAccessToken operation.
     * @callback module:api/UsersApi~updateUserSettingsAccessTokenCallback
     * @param {String} error Error message, if any.
     * @param {module:model/UserSettingsRegenerateAccessToken} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Regenerate API Access Token
     * This endpoint regenerates your API access token for the specified client. If a current token exists, it is revoked and a new token is returned.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.userId ID of User (Only available to the master tenant.)  Default is the current user
     * @param {String} opts.clientId Client ID.  See `Get Available API Clients` for a list of valid `clientId` values.
     * @param {module:api/UsersApi~updateUserSettingsAccessTokenCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/UserSettingsRegenerateAccessToken}
     */
    updateUserSettingsAccessToken(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'userId': opts['userId'],
        'clientId': opts['clientId']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['bearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = UserSettingsRegenerateAccessToken;
      return this.apiClient.callApi(
        '/api/user-settings/regenerate-access-token', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateUserSettingsAvatar operation.
     * @callback module:api/UsersApi~updateUserSettingsAvatarCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update Avatar
     * This endpoint updates your avatar image. Expects multipart form data as the request format, not JSON.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.userId ID of User (Only available to the master tenant.)  Default is the current user
     * @param {String} opts.userAvatar 
     * @param {module:api/UsersApi~updateUserSettingsAvatarCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    updateUserSettingsAvatar(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'userId': opts['userId']
      };
      let headerParams = {
      };
      let formParams = {
        'user.avatar': opts['userAvatar']
      };

      let authNames = ['bearerAuth'];
      let contentTypes = ['multipart/form-data'];
      let accepts = ['application/json'];
      let returnType = Model200Success;
      return this.apiClient.callApi(
        '/api/user-settings/avatar', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateUserSettingsDesktopBackground operation.
     * @callback module:api/UsersApi~updateUserSettingsDesktopBackgroundCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Model200Success} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update Desktop Background
     * This endpoint updates your desktop background image that is used in the Virtual Desktop persona. Expects multipart form data as the request format, not JSON.
     * @param {Object} opts Optional parameters
     * @param {Number} opts.userId ID of User (Only available to the master tenant.)  Default is the current user
     * @param {String} opts.userDesktopBackground 
     * @param {module:api/UsersApi~updateUserSettingsDesktopBackgroundCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Model200Success}
     */
    updateUserSettingsDesktopBackground(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'userId': opts['userId']
      };
      let headerParams = {
      };
      let formParams = {
        'user.desktopBackground': opts['userDesktopBackground']
      };

      let authNames = ['bearerAuth'];
      let contentTypes = ['multipart/form-data'];
      let accepts = ['application/json'];
      let returnType = Model200Success;
      return this.apiClient.callApi(
        '/api/user-settings/desktop-background', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateUsers operation.
     * @callback module:api/UsersApi~updateUsersCallback
     * @param {String} error Error message, if any.
     * @param {Object} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update user
     * Update an existing user.
     * @param {Number} id Morpheus ID of the Object being referenced
     * @param {Object} opts Optional parameters
     * @param {module:model/InlineObject256} opts.inlineObject256 
     * @param {module:api/UsersApi~updateUsersCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object}
     */
    updateUsers(id, opts, callback) {
      opts = opts || {};
      let postBody = opts['inlineObject256'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling updateUsers");
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
        '/api/users/{id}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the whoami operation.
     * @callback module:api/UsersApi~whoamiCallback
     * @param {String} error Error message, if any.
     * @param {module:model/InlineResponse200167} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieves information about current user roles and permissions
     * Provides API to retrieve information about yourself, including your roles and permissions.  The appliance build version is also returned. 
     * @param {module:api/UsersApi~whoamiCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/InlineResponse200167}
     */
    whoami(callback) {
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
      let returnType = InlineResponse200167;
      return this.apiClient.callApi(
        '/api/whoami', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
