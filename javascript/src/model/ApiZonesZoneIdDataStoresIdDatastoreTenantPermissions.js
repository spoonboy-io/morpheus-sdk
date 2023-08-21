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

/**
 * The ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions model module.
 * @module model/ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions
 * @version 6.2.1
 */
class ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions {
    /**
     * Constructs a new <code>ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions</code>.
     * @alias module:model/ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions
     */
    constructor() { 
        
        ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions} obj Optional instance to populate.
     * @return {module:model/ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions} The populated <code>ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions();

            if (data.hasOwnProperty('accounts')) {
                obj['accounts'] = ApiClient.convertToType(data['accounts'], ['Number']);
            }
            if (data.hasOwnProperty('defaultTarget')) {
                obj['defaultTarget'] = ApiClient.convertToType(data['defaultTarget'], ['Number']);
            }
            if (data.hasOwnProperty('defaultStore')) {
                obj['defaultStore'] = ApiClient.convertToType(data['defaultStore'], ['Number']);
            }
        }
        return obj;
    }


}

/**
 * Array of tenant account ids that are allowed access
 * @member {Array.<Number>} accounts
 */
ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions.prototype['accounts'] = undefined;

/**
 * Array of tenant account ids which should use the data store as the Default
 * @member {Array.<Number>} defaultTarget
 */
ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions.prototype['defaultTarget'] = undefined;

/**
 * Array of tenant account ids which should use the data store as the Image Target
 * @member {Array.<Number>} defaultStore
 */
ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions.prototype['defaultStore'] = undefined;






export default ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions;
