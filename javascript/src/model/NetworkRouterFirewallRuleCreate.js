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
 * The NetworkRouterFirewallRuleCreate model module.
 * @module model/NetworkRouterFirewallRuleCreate
 * @version 6.2.1
 */
class NetworkRouterFirewallRuleCreate {
    /**
     * Constructs a new <code>NetworkRouterFirewallRuleCreate</code>.
     * For a full list of available firewall rule options, see ruleOptionTypes in the specific Network Router Type detail
     * @alias module:model/NetworkRouterFirewallRuleCreate
     * @param name {String} Firewall rule name
     */
    constructor(name) { 
        
        NetworkRouterFirewallRuleCreate.initialize(this, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name) { 
        obj['name'] = name;
    }

    /**
     * Constructs a <code>NetworkRouterFirewallRuleCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkRouterFirewallRuleCreate} obj Optional instance to populate.
     * @return {module:model/NetworkRouterFirewallRuleCreate} The populated <code>NetworkRouterFirewallRuleCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkRouterFirewallRuleCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('priority')) {
                obj['priority'] = ApiClient.convertToType(data['priority'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * Firewall rule name
 * @member {String} name
 */
NetworkRouterFirewallRuleCreate.prototype['name'] = undefined;

/**
 * Can be used to enable / disable the rule (true, false). Default is on
 * @member {Boolean} enabled
 * @default true
 */
NetworkRouterFirewallRuleCreate.prototype['enabled'] = true;

/**
 * Priority for rule
 * @member {Number} priority
 */
NetworkRouterFirewallRuleCreate.prototype['priority'] = undefined;






export default NetworkRouterFirewallRuleCreate;

