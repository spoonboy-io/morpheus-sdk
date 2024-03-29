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

/**
 * The InlineResponse20095NetworkRouterNetworkServerIntegration model module.
 * @module model/InlineResponse20095NetworkRouterNetworkServerIntegration
 * @version 6.2.1
 */
class InlineResponse20095NetworkRouterNetworkServerIntegration {
    /**
     * Constructs a new <code>InlineResponse20095NetworkRouterNetworkServerIntegration</code>.
     * @alias module:model/InlineResponse20095NetworkRouterNetworkServerIntegration
     */
    constructor() { 
        
        InlineResponse20095NetworkRouterNetworkServerIntegration.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20095NetworkRouterNetworkServerIntegration</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20095NetworkRouterNetworkServerIntegration} obj Optional instance to populate.
     * @return {module:model/InlineResponse20095NetworkRouterNetworkServerIntegration} The populated <code>InlineResponse20095NetworkRouterNetworkServerIntegration</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20095NetworkRouterNetworkServerIntegration();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('integrationType')) {
                obj['integrationType'] = InlineResponse20079LoadBalancerMonitorLoadBalancerType.constructFromObject(data['integrationType']);
            }
            if (data.hasOwnProperty('url')) {
                obj['url'] = ApiClient.convertToType(data['url'], 'String');
            }
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'String');
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'String');
            }
            if (data.hasOwnProperty('isPlugin')) {
                obj['isPlugin'] = ApiClient.convertToType(data['isPlugin'], 'Boolean');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('statusDate')) {
                obj['statusDate'] = ApiClient.convertToType(data['statusDate'], 'Date');
            }
            if (data.hasOwnProperty('statusMessage')) {
                obj['statusMessage'] = ApiClient.convertToType(data['statusMessage'], 'String');
            }
            if (data.hasOwnProperty('lastSync')) {
                obj['lastSync'] = ApiClient.convertToType(data['lastSync'], 'Date');
            }
            if (data.hasOwnProperty('lastSyncDuration')) {
                obj['lastSyncDuration'] = ApiClient.convertToType(data['lastSyncDuration'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['id'] = undefined;

/**
 * @member {String} name
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['name'] = undefined;

/**
 * @member {Boolean} enabled
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['enabled'] = undefined;

/**
 * @member {String} type
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['type'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} integrationType
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['integrationType'] = undefined;

/**
 * @member {String} url
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['url'] = undefined;

/**
 * @member {String} port
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['port'] = undefined;

/**
 * @member {String} username
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['username'] = undefined;

/**
 * @member {String} password
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['password'] = undefined;

/**
 * @member {String} refType
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['refType'] = undefined;

/**
 * @member {String} refId
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['refId'] = undefined;

/**
 * @member {Boolean} isPlugin
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['isPlugin'] = undefined;

/**
 * @member {Object} config
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['config'] = undefined;

/**
 * @member {String} status
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['status'] = undefined;

/**
 * @member {Date} statusDate
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['statusDate'] = undefined;

/**
 * @member {String} statusMessage
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['statusMessage'] = undefined;

/**
 * @member {Date} lastSync
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['lastSync'] = undefined;

/**
 * @member {Number} lastSyncDuration
 */
InlineResponse20095NetworkRouterNetworkServerIntegration.prototype['lastSyncDuration'] = undefined;






export default InlineResponse20095NetworkRouterNetworkServerIntegration;

