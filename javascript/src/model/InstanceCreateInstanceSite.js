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
 * The InstanceCreateInstanceSite model module.
 * @module model/InstanceCreateInstanceSite
 * @version 6.2.1
 */
class InstanceCreateInstanceSite {
    /**
     * Constructs a new <code>InstanceCreateInstanceSite</code>.
     * @alias module:model/InstanceCreateInstanceSite
     * @param id {Number} The Group ID to provision the instance into.
     */
    constructor(id) { 
        
        InstanceCreateInstanceSite.initialize(this, id);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id) { 
        obj['id'] = id;
    }

    /**
     * Constructs a <code>InstanceCreateInstanceSite</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceCreateInstanceSite} obj Optional instance to populate.
     * @return {module:model/InstanceCreateInstanceSite} The populated <code>InstanceCreateInstanceSite</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceCreateInstanceSite();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * The Group ID to provision the instance into.
 * @member {Number} id
 */
InstanceCreateInstanceSite.prototype['id'] = undefined;






export default InstanceCreateInstanceSite;

