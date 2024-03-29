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
 * The ApiIntegrationsIdInventoryInventoryIdInventory model module.
 * @module model/ApiIntegrationsIdInventoryInventoryIdInventory
 * @version 6.2.1
 */
class ApiIntegrationsIdInventoryInventoryIdInventory {
    /**
     * Constructs a new <code>ApiIntegrationsIdInventoryInventoryIdInventory</code>.
     * @alias module:model/ApiIntegrationsIdInventoryInventoryIdInventory
     */
    constructor() { 
        
        ApiIntegrationsIdInventoryInventoryIdInventory.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiIntegrationsIdInventoryInventoryIdInventory</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiIntegrationsIdInventoryInventoryIdInventory} obj Optional instance to populate.
     * @return {module:model/ApiIntegrationsIdInventoryInventoryIdInventory} The populated <code>ApiIntegrationsIdInventoryInventoryIdInventory</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiIntegrationsIdInventoryInventoryIdInventory();

            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], [ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]);
            }
        }
        return obj;
    }


}

/**
 * Array of tenant accounts that will use this inventory as Default. Used by jobs set to 'Use Tenant Default'
 * @member {Array.<module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>} tenants
 */
ApiIntegrationsIdInventoryInventoryIdInventory.prototype['tenants'] = undefined;






export default ApiIntegrationsIdInventoryInventoryIdInventory;

