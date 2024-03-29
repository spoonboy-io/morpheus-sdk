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
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The InlineResponse20095NetworkRouterFirewallRules model module.
 * @module model/InlineResponse20095NetworkRouterFirewallRules
 * @version 6.2.1
 */
class InlineResponse20095NetworkRouterFirewallRules {
    /**
     * Constructs a new <code>InlineResponse20095NetworkRouterFirewallRules</code>.
     * @alias module:model/InlineResponse20095NetworkRouterFirewallRules
     */
    constructor() { 
        
        InlineResponse20095NetworkRouterFirewallRules.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20095NetworkRouterFirewallRules</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20095NetworkRouterFirewallRules} obj Optional instance to populate.
     * @return {module:model/InlineResponse20095NetworkRouterFirewallRules} The populated <code>InlineResponse20095NetworkRouterFirewallRules</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20095NetworkRouterFirewallRules();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('groupName')) {
                obj['groupName'] = ApiClient.convertToType(data['groupName'], 'String');
            }
            if (data.hasOwnProperty('direction')) {
                obj['direction'] = ApiClient.convertToType(data['direction'], 'String');
            }
            if (data.hasOwnProperty('ruleType')) {
                obj['ruleType'] = ApiClient.convertToType(data['ruleType'], 'String');
            }
            if (data.hasOwnProperty('policy')) {
                obj['policy'] = ApiClient.convertToType(data['policy'], 'String');
            }
            if (data.hasOwnProperty('source')) {
                obj['source'] = ApiClient.convertToType(data['source'], ['String']);
            }
            if (data.hasOwnProperty('sourceType')) {
                obj['sourceType'] = ApiClient.convertToType(data['sourceType'], 'String');
            }
            if (data.hasOwnProperty('destination')) {
                obj['destination'] = ApiClient.convertToType(data['destination'], ['String']);
            }
            if (data.hasOwnProperty('destinationType')) {
                obj['destinationType'] = ApiClient.convertToType(data['destinationType'], 'String');
            }
            if (data.hasOwnProperty('profiles')) {
                obj['profiles'] = ApiClient.convertToType(data['profiles'], ['String']);
            }
            if (data.hasOwnProperty('protocol')) {
                obj['protocol'] = ApiClient.convertToType(data['protocol'], 'String');
            }
            if (data.hasOwnProperty('application')) {
                obj['application'] = ApiClient.convertToType(data['application'], 'String');
            }
            if (data.hasOwnProperty('applicationType')) {
                obj['applicationType'] = ApiClient.convertToType(data['applicationType'], 'String');
            }
            if (data.hasOwnProperty('portRange')) {
                obj['portRange'] = ApiClient.convertToType(data['portRange'], 'String');
            }
            if (data.hasOwnProperty('sourcePortRange')) {
                obj['sourcePortRange'] = ApiClient.convertToType(data['sourcePortRange'], 'String');
            }
            if (data.hasOwnProperty('sourceGroup')) {
                obj['sourceGroup'] = ApiClient.convertToType(data['sourceGroup'], 'String');
            }
            if (data.hasOwnProperty('sourceTier')) {
                obj['sourceTier'] = ApiClient.convertToType(data['sourceTier'], 'String');
            }
            if (data.hasOwnProperty('applications')) {
                obj['applications'] = ApiClient.convertToType(data['applications'], [InlineResponse20040AppDeployInstance]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['id'] = undefined;

/**
 * @member {String} name
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['name'] = undefined;

/**
 * @member {String} code
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['code'] = undefined;

/**
 * @member {Boolean} enabled
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['enabled'] = undefined;

/**
 * @member {String} groupName
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['groupName'] = undefined;

/**
 * @member {String} direction
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['direction'] = undefined;

/**
 * @member {String} ruleType
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['ruleType'] = undefined;

/**
 * @member {String} policy
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['policy'] = undefined;

/**
 * @member {Array.<String>} source
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['source'] = undefined;

/**
 * @member {String} sourceType
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['sourceType'] = undefined;

/**
 * @member {Array.<String>} destination
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['destination'] = undefined;

/**
 * @member {String} destinationType
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['destinationType'] = undefined;

/**
 * @member {Array.<String>} profiles
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['profiles'] = undefined;

/**
 * @member {String} protocol
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['protocol'] = undefined;

/**
 * @member {String} application
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['application'] = undefined;

/**
 * @member {String} applicationType
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['applicationType'] = undefined;

/**
 * @member {String} portRange
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['portRange'] = undefined;

/**
 * @member {String} sourcePortRange
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['sourcePortRange'] = undefined;

/**
 * @member {String} sourceGroup
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['sourceGroup'] = undefined;

/**
 * @member {String} sourceTier
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['sourceTier'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20040AppDeployInstance>} applications
 */
InlineResponse20095NetworkRouterFirewallRules.prototype['applications'] = undefined;






export default InlineResponse20095NetworkRouterFirewallRules;

