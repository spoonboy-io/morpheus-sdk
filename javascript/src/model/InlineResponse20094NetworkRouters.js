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
import InlineResponse20079LoadBalancerMonitorLoadBalancerType from './InlineResponse20079LoadBalancerMonitorLoadBalancerType';
import InlineResponse20094Interfaces from './InlineResponse20094Interfaces';
import InlineResponse20094Type from './InlineResponse20094Type';

/**
 * The InlineResponse20094NetworkRouters model module.
 * @module model/InlineResponse20094NetworkRouters
 * @version 6.2.1
 */
class InlineResponse20094NetworkRouters {
    /**
     * Constructs a new <code>InlineResponse20094NetworkRouters</code>.
     * @alias module:model/InlineResponse20094NetworkRouters
     */
    constructor() { 
        
        InlineResponse20094NetworkRouters.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20094NetworkRouters</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20094NetworkRouters} obj Optional instance to populate.
     * @return {module:model/InlineResponse20094NetworkRouters} The populated <code>InlineResponse20094NetworkRouters</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20094NetworkRouters();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('routerType')) {
                obj['routerType'] = ApiClient.convertToType(data['routerType'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('externalIp')) {
                obj['externalIp'] = ApiClient.convertToType(data['externalIp'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('providerId')) {
                obj['providerId'] = ApiClient.convertToType(data['providerId'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = InlineResponse20094Type.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('networkServer')) {
                obj['networkServer'] = ApiClient.convertToType(data['networkServer'], 'String');
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = InlineResponse20079LoadBalancerMonitorLoadBalancerType.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('instance')) {
                obj['instance'] = ApiClient.convertToType(data['instance'], 'String');
            }
            if (data.hasOwnProperty('externalNetwork')) {
                obj['externalNetwork'] = InlineResponse20079LoadBalancerMonitorLoadBalancerType.constructFromObject(data['externalNetwork']);
            }
            if (data.hasOwnProperty('site')) {
                obj['site'] = ApiClient.convertToType(data['site'], 'String');
            }
            if (data.hasOwnProperty('interfaces')) {
                obj['interfaces'] = ApiClient.convertToType(data['interfaces'], [InlineResponse20094Interfaces]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse20094NetworkRouters.prototype['id'] = undefined;

/**
 * @member {String} code
 */
InlineResponse20094NetworkRouters.prototype['code'] = undefined;

/**
 * @member {String} name
 */
InlineResponse20094NetworkRouters.prototype['name'] = undefined;

/**
 * @member {String} description
 */
InlineResponse20094NetworkRouters.prototype['description'] = undefined;

/**
 * @member {String} category
 */
InlineResponse20094NetworkRouters.prototype['category'] = undefined;

/**
 * @member {Date} dateCreated
 */
InlineResponse20094NetworkRouters.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
InlineResponse20094NetworkRouters.prototype['lastUpdated'] = undefined;

/**
 * @member {String} routerType
 */
InlineResponse20094NetworkRouters.prototype['routerType'] = undefined;

/**
 * @member {String} status
 */
InlineResponse20094NetworkRouters.prototype['status'] = undefined;

/**
 * @member {Boolean} enabled
 */
InlineResponse20094NetworkRouters.prototype['enabled'] = undefined;

/**
 * @member {String} externalIp
 */
InlineResponse20094NetworkRouters.prototype['externalIp'] = undefined;

/**
 * @member {String} externalId
 */
InlineResponse20094NetworkRouters.prototype['externalId'] = undefined;

/**
 * @member {String} providerId
 */
InlineResponse20094NetworkRouters.prototype['providerId'] = undefined;

/**
 * @member {module:model/InlineResponse20094Type} type
 */
InlineResponse20094NetworkRouters.prototype['type'] = undefined;

/**
 * @member {String} networkServer
 */
InlineResponse20094NetworkRouters.prototype['networkServer'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} zone
 */
InlineResponse20094NetworkRouters.prototype['zone'] = undefined;

/**
 * @member {String} instance
 */
InlineResponse20094NetworkRouters.prototype['instance'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} externalNetwork
 */
InlineResponse20094NetworkRouters.prototype['externalNetwork'] = undefined;

/**
 * @member {String} site
 */
InlineResponse20094NetworkRouters.prototype['site'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20094Interfaces>} interfaces
 */
InlineResponse20094NetworkRouters.prototype['interfaces'] = undefined;






export default InlineResponse20094NetworkRouters;

