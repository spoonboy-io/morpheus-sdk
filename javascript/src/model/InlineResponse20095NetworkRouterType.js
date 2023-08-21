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
import OptionType from './OptionType';

/**
 * The InlineResponse20095NetworkRouterType model module.
 * @module model/InlineResponse20095NetworkRouterType
 * @version 6.2.1
 */
class InlineResponse20095NetworkRouterType {
    /**
     * Constructs a new <code>InlineResponse20095NetworkRouterType</code>.
     * @alias module:model/InlineResponse20095NetworkRouterType
     */
    constructor() { 
        
        InlineResponse20095NetworkRouterType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20095NetworkRouterType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20095NetworkRouterType} obj Optional instance to populate.
     * @return {module:model/InlineResponse20095NetworkRouterType} The populated <code>InlineResponse20095NetworkRouterType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20095NetworkRouterType();

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
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('creatable')) {
                obj['creatable'] = ApiClient.convertToType(data['creatable'], 'Boolean');
            }
            if (data.hasOwnProperty('selectable')) {
                obj['selectable'] = ApiClient.convertToType(data['selectable'], 'Boolean');
            }
            if (data.hasOwnProperty('hasFirewall')) {
                obj['hasFirewall'] = ApiClient.convertToType(data['hasFirewall'], 'Boolean');
            }
            if (data.hasOwnProperty('hasDhcp')) {
                obj['hasDhcp'] = ApiClient.convertToType(data['hasDhcp'], 'Boolean');
            }
            if (data.hasOwnProperty('hasRouting')) {
                obj['hasRouting'] = ApiClient.convertToType(data['hasRouting'], 'Boolean');
            }
            if (data.hasOwnProperty('hasNat')) {
                obj['hasNat'] = ApiClient.convertToType(data['hasNat'], 'Boolean');
            }
            if (data.hasOwnProperty('hasNetworkServer')) {
                obj['hasNetworkServer'] = ApiClient.convertToType(data['hasNetworkServer'], 'Boolean');
            }
            if (data.hasOwnProperty('hasFirewallGroups')) {
                obj['hasFirewallGroups'] = ApiClient.convertToType(data['hasFirewallGroups'], 'Boolean');
            }
            if (data.hasOwnProperty('hasSecurityGroupPriority')) {
                obj['hasSecurityGroupPriority'] = ApiClient.convertToType(data['hasSecurityGroupPriority'], 'Boolean');
            }
            if (data.hasOwnProperty('optionTypes')) {
                obj['optionTypes'] = ApiClient.convertToType(data['optionTypes'], [OptionType]);
            }
            if (data.hasOwnProperty('ruleOptionTypes')) {
                obj['ruleOptionTypes'] = ApiClient.convertToType(data['ruleOptionTypes'], [OptionType]);
            }
            if (data.hasOwnProperty('firewallGroupOptionTypes')) {
                obj['firewallGroupOptionTypes'] = ApiClient.convertToType(data['firewallGroupOptionTypes'], [OptionType]);
            }
            if (data.hasOwnProperty('natOptionTypes')) {
                obj['natOptionTypes'] = ApiClient.convertToType(data['natOptionTypes'], [OptionType]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse20095NetworkRouterType.prototype['id'] = undefined;

/**
 * @member {String} code
 */
InlineResponse20095NetworkRouterType.prototype['code'] = undefined;

/**
 * @member {String} name
 */
InlineResponse20095NetworkRouterType.prototype['name'] = undefined;

/**
 * @member {String} description
 */
InlineResponse20095NetworkRouterType.prototype['description'] = undefined;

/**
 * @member {Boolean} enabled
 */
InlineResponse20095NetworkRouterType.prototype['enabled'] = undefined;

/**
 * @member {Boolean} creatable
 */
InlineResponse20095NetworkRouterType.prototype['creatable'] = undefined;

/**
 * @member {Boolean} selectable
 */
InlineResponse20095NetworkRouterType.prototype['selectable'] = undefined;

/**
 * @member {Boolean} hasFirewall
 */
InlineResponse20095NetworkRouterType.prototype['hasFirewall'] = undefined;

/**
 * @member {Boolean} hasDhcp
 */
InlineResponse20095NetworkRouterType.prototype['hasDhcp'] = undefined;

/**
 * @member {Boolean} hasRouting
 */
InlineResponse20095NetworkRouterType.prototype['hasRouting'] = undefined;

/**
 * @member {Boolean} hasNat
 */
InlineResponse20095NetworkRouterType.prototype['hasNat'] = undefined;

/**
 * @member {Boolean} hasNetworkServer
 */
InlineResponse20095NetworkRouterType.prototype['hasNetworkServer'] = undefined;

/**
 * @member {Boolean} hasFirewallGroups
 */
InlineResponse20095NetworkRouterType.prototype['hasFirewallGroups'] = undefined;

/**
 * @member {Boolean} hasSecurityGroupPriority
 */
InlineResponse20095NetworkRouterType.prototype['hasSecurityGroupPriority'] = undefined;

/**
 * @member {Array.<module:model/OptionType>} optionTypes
 */
InlineResponse20095NetworkRouterType.prototype['optionTypes'] = undefined;

/**
 * @member {Array.<module:model/OptionType>} ruleOptionTypes
 */
InlineResponse20095NetworkRouterType.prototype['ruleOptionTypes'] = undefined;

/**
 * @member {Array.<module:model/OptionType>} firewallGroupOptionTypes
 */
InlineResponse20095NetworkRouterType.prototype['firewallGroupOptionTypes'] = undefined;

/**
 * @member {Array.<module:model/OptionType>} natOptionTypes
 */
InlineResponse20095NetworkRouterType.prototype['natOptionTypes'] = undefined;






export default InlineResponse20095NetworkRouterType;

