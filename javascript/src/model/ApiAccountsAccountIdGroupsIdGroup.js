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
 * The ApiAccountsAccountIdGroupsIdGroup model module.
 * @module model/ApiAccountsAccountIdGroupsIdGroup
 * @version 6.2.1
 */
class ApiAccountsAccountIdGroupsIdGroup {
    /**
     * Constructs a new <code>ApiAccountsAccountIdGroupsIdGroup</code>.
     * @alias module:model/ApiAccountsAccountIdGroupsIdGroup
     */
    constructor() { 
        
        ApiAccountsAccountIdGroupsIdGroup.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiAccountsAccountIdGroupsIdGroup</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiAccountsAccountIdGroupsIdGroup} obj Optional instance to populate.
     * @return {module:model/ApiAccountsAccountIdGroupsIdGroup} The populated <code>ApiAccountsAccountIdGroupsIdGroup</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiAccountsAccountIdGroupsIdGroup();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('location')) {
                obj['location'] = ApiClient.convertToType(data['location'], 'String');
            }
        }
        return obj;
    }


}

/**
 * A unique name scoped to the subtenant for the group
 * @member {String} name
 */
ApiAccountsAccountIdGroupsIdGroup.prototype['name'] = undefined;

/**
 * Optional description field if you want to put more info there
 * @member {String} description
 */
ApiAccountsAccountIdGroupsIdGroup.prototype['description'] = undefined;

/**
 * Optional code for use with policies
 * @member {String} code
 */
ApiAccountsAccountIdGroupsIdGroup.prototype['code'] = undefined;

/**
 * location
 * @member {String} location
 */
ApiAccountsAccountIdGroupsIdGroup.prototype['location'] = undefined;






export default ApiAccountsAccountIdGroupsIdGroup;
