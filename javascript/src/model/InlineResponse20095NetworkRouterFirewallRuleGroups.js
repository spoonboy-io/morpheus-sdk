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
import InlineResponse20095NetworkRouterFirewallRules from './InlineResponse20095NetworkRouterFirewallRules';

/**
 * The InlineResponse20095NetworkRouterFirewallRuleGroups model module.
 * @module model/InlineResponse20095NetworkRouterFirewallRuleGroups
 * @version 6.2.1
 */
class InlineResponse20095NetworkRouterFirewallRuleGroups {
    /**
     * Constructs a new <code>InlineResponse20095NetworkRouterFirewallRuleGroups</code>.
     * @alias module:model/InlineResponse20095NetworkRouterFirewallRuleGroups
     */
    constructor() { 
        
        InlineResponse20095NetworkRouterFirewallRuleGroups.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20095NetworkRouterFirewallRuleGroups</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20095NetworkRouterFirewallRuleGroups} obj Optional instance to populate.
     * @return {module:model/InlineResponse20095NetworkRouterFirewallRuleGroups} The populated <code>InlineResponse20095NetworkRouterFirewallRuleGroups</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20095NetworkRouterFirewallRuleGroups();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('iacId')) {
                obj['iacId'] = ApiClient.convertToType(data['iacId'], 'String');
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = ApiClient.convertToType(data['zone'], 'String');
            }
            if (data.hasOwnProperty('zonePool')) {
                obj['zonePool'] = ApiClient.convertToType(data['zonePool'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('priority')) {
                obj['priority'] = ApiClient.convertToType(data['priority'], 'Number');
            }
            if (data.hasOwnProperty('groupLayer')) {
                obj['groupLayer'] = ApiClient.convertToType(data['groupLayer'], 'String');
            }
            if (data.hasOwnProperty('rules')) {
                obj['rules'] = ApiClient.convertToType(data['rules'], [InlineResponse20095NetworkRouterFirewallRules]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['id'] = undefined;

/**
 * @member {String} name
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['name'] = undefined;

/**
 * @member {String} description
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['description'] = undefined;

/**
 * @member {String} externalId
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['externalId'] = undefined;

/**
 * @member {String} iacId
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['iacId'] = undefined;

/**
 * @member {String} zone
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['zone'] = undefined;

/**
 * @member {String} zonePool
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['zonePool'] = undefined;

/**
 * @member {String} status
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['status'] = undefined;

/**
 * @member {Number} priority
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['priority'] = undefined;

/**
 * @member {String} groupLayer
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['groupLayer'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20095NetworkRouterFirewallRules>} rules
 */
InlineResponse20095NetworkRouterFirewallRuleGroups.prototype['rules'] = undefined;






export default InlineResponse20095NetworkRouterFirewallRuleGroups;
