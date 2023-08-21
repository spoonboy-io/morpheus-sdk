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
 * The InlineResponse200131ResourcePermissionSites model module.
 * @module model/InlineResponse200131ResourcePermissionSites
 * @version 6.2.1
 */
class InlineResponse200131ResourcePermissionSites {
    /**
     * Constructs a new <code>InlineResponse200131ResourcePermissionSites</code>.
     * @alias module:model/InlineResponse200131ResourcePermissionSites
     */
    constructor() { 
        
        InlineResponse200131ResourcePermissionSites.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200131ResourcePermissionSites</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200131ResourcePermissionSites} obj Optional instance to populate.
     * @return {module:model/InlineResponse200131ResourcePermissionSites} The populated <code>InlineResponse200131ResourcePermissionSites</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200131ResourcePermissionSites();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('default')) {
                obj['default'] = ApiClient.convertToType(data['default'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse200131ResourcePermissionSites.prototype['id'] = undefined;

/**
 * @member {Boolean} default
 */
InlineResponse200131ResourcePermissionSites.prototype['default'] = undefined;






export default InlineResponse200131ResourcePermissionSites;

