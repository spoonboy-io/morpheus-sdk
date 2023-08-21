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
 * The ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool model module.
 * @module model/ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool
 * @version 6.2.1
 */
class ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool {
    /**
     * Constructs a new <code>ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool</code>.
     * @alias module:model/ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool
     */
    constructor() { 
        
        ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool} obj Optional instance to populate.
     * @return {module:model/ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool} The populated <code>ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('vipBalance')) {
                obj['vipBalance'] = ApiClient.convertToType(data['vipBalance'], 'String');
            }
            if (data.hasOwnProperty('minActive')) {
                obj['minActive'] = ApiClient.convertToType(data['minActive'], 'Number');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool.prototype['name'] = undefined;

/**
 * Description
 * @member {String} description
 */
ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool.prototype['description'] = undefined;

/**
 * Balance Algorithm
 * @member {String} vipBalance
 */
ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool.prototype['vipBalance'] = undefined;

/**
 * Min Active Members
 * @member {Number} minActive
 */
ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool.prototype['minActive'] = undefined;

/**
 * Configuration object with parameters that vary by type.
 * @member {Object} config
 */
ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool.prototype['config'] = undefined;






export default ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool;

