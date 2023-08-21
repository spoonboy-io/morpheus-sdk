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

import ApiClient from '../ApiClient';
import ApiBlueprintsIdUpdatePermissionsResourcePermissionSites from './ApiBlueprintsIdUpdatePermissionsResourcePermissionSites';

/**
 * The ApiStorageServersStorageServer model module.
 * @module model/ApiStorageServersStorageServer
 * @version 6.2.1
 */
class ApiStorageServersStorageServer {
    /**
     * Constructs a new <code>ApiStorageServersStorageServer</code>.
     * @alias module:model/ApiStorageServersStorageServer
     * @param name {String} Name
     * @param type {String} The `Storage Type` Code or ID
     * @param config {Object} Configuration object with parameters that vary by `type`
     */
    constructor(name, type, config) { 
        
        ApiStorageServersStorageServer.initialize(this, name, type, config);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, type, config) { 
        obj['name'] = name;
        obj['type'] = type;
        obj['config'] = config;
    }

    /**
     * Constructs a <code>ApiStorageServersStorageServer</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiStorageServersStorageServer} obj Optional instance to populate.
     * @return {module:model/ApiStorageServersStorageServer} The populated <code>ApiStorageServersStorageServer</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiStorageServersStorageServer();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], [ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]);
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
ApiStorageServersStorageServer.prototype['name'] = undefined;

/**
 * The `Storage Type` Code or ID
 * @member {String} type
 */
ApiStorageServersStorageServer.prototype['type'] = undefined;

/**
 * description
 * @member {String} description
 */
ApiStorageServersStorageServer.prototype['description'] = undefined;

/**
 * The enabled flag
 * @member {Boolean} enabled
 * @default true
 */
ApiStorageServersStorageServer.prototype['enabled'] = true;

/**
 * Configuration object with parameters that vary by `type`
 * @member {Object} config
 */
ApiStorageServersStorageServer.prototype['config'] = undefined;

/**
 * private or public
 * @member {module:model/ApiStorageServersStorageServer.VisibilityEnum} visibility
 * @default 'private'
 */
ApiStorageServersStorageServer.prototype['visibility'] = 'private';

/**
 * Array of tenant account ids that are allowed access
 * @member {Array.<module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>} tenants
 */
ApiStorageServersStorageServer.prototype['tenants'] = undefined;





/**
 * Allowed values for the <code>visibility</code> property.
 * @enum {String}
 * @readonly
 */
ApiStorageServersStorageServer['VisibilityEnum'] = {

    /**
     * value: "private"
     * @const
     */
    "private": "private",

    /**
     * value: "public"
     * @const
     */
    "public": "public"
};



export default ApiStorageServersStorageServer;
