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
 * The InlineResponse20079LoadBalancerMonitorLoadBalancerType model module.
 * @module model/InlineResponse20079LoadBalancerMonitorLoadBalancerType
 * @version 6.2.1
 */
class InlineResponse20079LoadBalancerMonitorLoadBalancerType {
    /**
     * Constructs a new <code>InlineResponse20079LoadBalancerMonitorLoadBalancerType</code>.
     * @alias module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType
     */
    constructor() { 
        
        InlineResponse20079LoadBalancerMonitorLoadBalancerType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20079LoadBalancerMonitorLoadBalancerType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} obj Optional instance to populate.
     * @return {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} The populated <code>InlineResponse20079LoadBalancerMonitorLoadBalancerType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20079LoadBalancerMonitorLoadBalancerType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse20079LoadBalancerMonitorLoadBalancerType.prototype['id'] = undefined;

/**
 * @member {String} code
 */
InlineResponse20079LoadBalancerMonitorLoadBalancerType.prototype['code'] = undefined;

/**
 * @member {String} name
 */
InlineResponse20079LoadBalancerMonitorLoadBalancerType.prototype['name'] = undefined;






export default InlineResponse20079LoadBalancerMonitorLoadBalancerType;

