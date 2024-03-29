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
import ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissionsResourcePermissions from './ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissionsResourcePermissions';

/**
 * The ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions model module.
 * @module model/ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions
 * @version 6.2.1
 */
class ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions {
    /**
     * Constructs a new <code>ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions</code>.
     * Permissions object for upgrading group access
     * @alias module:model/ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions
     */
    constructor() { 
        
        ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions} obj Optional instance to populate.
     * @return {module:model/ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions} The populated <code>ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions();

            if (data.hasOwnProperty('resourcePermissions')) {
                obj['resourcePermissions'] = ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissionsResourcePermissions.constructFromObject(data['resourcePermissions']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissionsResourcePermissions} resourcePermissions
 */
ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions.prototype['resourcePermissions'] = undefined;






export default ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions;

