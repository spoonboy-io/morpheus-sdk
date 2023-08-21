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
import ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions from './ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions';
import ApiZonesZoneIdFoldersIdFolderTenantPermissions from './ApiZonesZoneIdFoldersIdFolderTenantPermissions';

/**
 * The ApiZonesZoneIdResourcePoolsIdResourcePool model module.
 * @module model/ApiZonesZoneIdResourcePoolsIdResourcePool
 * @version 6.2.1
 */
class ApiZonesZoneIdResourcePoolsIdResourcePool {
    /**
     * Constructs a new <code>ApiZonesZoneIdResourcePoolsIdResourcePool</code>.
     * @alias module:model/ApiZonesZoneIdResourcePoolsIdResourcePool
     */
    constructor() { 
        
        ApiZonesZoneIdResourcePoolsIdResourcePool.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiZonesZoneIdResourcePoolsIdResourcePool</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiZonesZoneIdResourcePoolsIdResourcePool} obj Optional instance to populate.
     * @return {module:model/ApiZonesZoneIdResourcePoolsIdResourcePool} The populated <code>ApiZonesZoneIdResourcePoolsIdResourcePool</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiZonesZoneIdResourcePoolsIdResourcePool();

            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('displayName')) {
                obj['displayName'] = ApiClient.convertToType(data['displayName'], 'String');
            }
            if (data.hasOwnProperty('inventory')) {
                obj['inventory'] = ApiClient.convertToType(data['inventory'], 'Boolean');
            }
            if (data.hasOwnProperty('tenantPermissions')) {
                obj['tenantPermissions'] = ApiClient.convertToType(data['tenantPermissions'], [ApiZonesZoneIdFoldersIdFolderTenantPermissions]);
            }
            if (data.hasOwnProperty('resourcePermissions')) {
                obj['resourcePermissions'] = ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.constructFromObject(data['resourcePermissions']);
            }
        }
        return obj;
    }


}

/**
 * Activate `true` or disable `false` the datastore
 * @member {Boolean} active
 */
ApiZonesZoneIdResourcePoolsIdResourcePool.prototype['active'] = undefined;

/**
 * Setting `private` or `public`
 * @member {module:model/ApiZonesZoneIdResourcePoolsIdResourcePool.VisibilityEnum} visibility
 * @default 'private'
 */
ApiZonesZoneIdResourcePoolsIdResourcePool.prototype['visibility'] = 'private';

/**
 * Optional Display Name (VMware only)
 * @member {String} displayName
 */
ApiZonesZoneIdResourcePoolsIdResourcePool.prototype['displayName'] = undefined;

/**
 * Enable `True` or disable `False` inventory sync for resource pool during cloud refresh
 * @member {Boolean} inventory
 */
ApiZonesZoneIdResourcePoolsIdResourcePool.prototype['inventory'] = undefined;

/**
 * @member {Array.<module:model/ApiZonesZoneIdFoldersIdFolderTenantPermissions>} tenantPermissions
 */
ApiZonesZoneIdResourcePoolsIdResourcePool.prototype['tenantPermissions'] = undefined;

/**
 * @member {module:model/ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions} resourcePermissions
 */
ApiZonesZoneIdResourcePoolsIdResourcePool.prototype['resourcePermissions'] = undefined;





/**
 * Allowed values for the <code>visibility</code> property.
 * @enum {String}
 * @readonly
 */
ApiZonesZoneIdResourcePoolsIdResourcePool['VisibilityEnum'] = {

    /**
     * value: "public"
     * @const
     */
    "public": "public",

    /**
     * value: "private"
     * @const
     */
    "private": "private"
};



export default ApiZonesZoneIdResourcePoolsIdResourcePool;
