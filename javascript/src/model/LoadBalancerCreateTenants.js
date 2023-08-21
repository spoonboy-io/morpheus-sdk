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
 * The LoadBalancerCreateTenants model module.
 * @module model/LoadBalancerCreateTenants
 * @version 6.2.1
 */
class LoadBalancerCreateTenants {
    /**
     * Constructs a new <code>LoadBalancerCreateTenants</code>.
     * @alias module:model/LoadBalancerCreateTenants
     */
    constructor() { 
        
        LoadBalancerCreateTenants.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>LoadBalancerCreateTenants</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/LoadBalancerCreateTenants} obj Optional instance to populate.
     * @return {module:model/LoadBalancerCreateTenants} The populated <code>LoadBalancerCreateTenants</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new LoadBalancerCreateTenants();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
LoadBalancerCreateTenants.prototype['id'] = undefined;






export default LoadBalancerCreateTenants;

