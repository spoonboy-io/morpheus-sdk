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
 * The BillingLoadBalancers model module.
 * @module model/BillingLoadBalancers
 * @version 6.2.1
 */
class BillingLoadBalancers {
    /**
     * Constructs a new <code>BillingLoadBalancers</code>.
     * @alias module:model/BillingLoadBalancers
     */
    constructor() { 
        
        BillingLoadBalancers.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingLoadBalancers</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingLoadBalancers} obj Optional instance to populate.
     * @return {module:model/BillingLoadBalancers} The populated <code>BillingLoadBalancers</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingLoadBalancers();

            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
            if (data.hasOwnProperty('loadBalancers')) {
                obj['loadBalancers'] = ApiClient.convertToType(data['loadBalancers'], [Object]);
            }
            if (data.hasOwnProperty('count')) {
                obj['count'] = ApiClient.convertToType(data['count'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} price
 */
BillingLoadBalancers.prototype['price'] = undefined;

/**
 * @member {Number} cost
 */
BillingLoadBalancers.prototype['cost'] = undefined;

/**
 * @member {Array.<Object>} loadBalancers
 */
BillingLoadBalancers.prototype['loadBalancers'] = undefined;

/**
 * @member {Number} count
 */
BillingLoadBalancers.prototype['count'] = undefined;






export default BillingLoadBalancers;

