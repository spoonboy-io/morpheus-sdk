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
 * The ApiGroupsIdUpdateZonesGroup model module.
 * @module model/ApiGroupsIdUpdateZonesGroup
 * @version 6.2.1
 */
class ApiGroupsIdUpdateZonesGroup {
    /**
     * Constructs a new <code>ApiGroupsIdUpdateZonesGroup</code>.
     * @alias module:model/ApiGroupsIdUpdateZonesGroup
     * @param zones {Array.<Object>} An array of all the zones assigned to this group.
     */
    constructor(zones) { 
        
        ApiGroupsIdUpdateZonesGroup.initialize(this, zones);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, zones) { 
        obj['zones'] = zones;
    }

    /**
     * Constructs a <code>ApiGroupsIdUpdateZonesGroup</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiGroupsIdUpdateZonesGroup} obj Optional instance to populate.
     * @return {module:model/ApiGroupsIdUpdateZonesGroup} The populated <code>ApiGroupsIdUpdateZonesGroup</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiGroupsIdUpdateZonesGroup();

            if (data.hasOwnProperty('zones')) {
                obj['zones'] = ApiClient.convertToType(data['zones'], [Object]);
            }
        }
        return obj;
    }


}

/**
 * An array of all the zones assigned to this group.
 * @member {Array.<Object>} zones
 */
ApiGroupsIdUpdateZonesGroup.prototype['zones'] = undefined;






export default ApiGroupsIdUpdateZonesGroup;

