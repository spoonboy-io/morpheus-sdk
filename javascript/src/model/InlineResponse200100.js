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
import InlineResponse20095NetworkRouterFirewallRuleGroups from './InlineResponse20095NetworkRouterFirewallRuleGroups';

/**
 * The InlineResponse200100 model module.
 * @module model/InlineResponse200100
 * @version 6.2.1
 */
class InlineResponse200100 {
    /**
     * Constructs a new <code>InlineResponse200100</code>.
     * @alias module:model/InlineResponse200100
     */
    constructor() { 
        
        InlineResponse200100.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200100</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200100} obj Optional instance to populate.
     * @return {module:model/InlineResponse200100} The populated <code>InlineResponse200100</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200100();

            if (data.hasOwnProperty('ruleGroups')) {
                obj['ruleGroups'] = ApiClient.convertToType(data['ruleGroups'], [InlineResponse20095NetworkRouterFirewallRuleGroups]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/InlineResponse20095NetworkRouterFirewallRuleGroups>} ruleGroups
 */
InlineResponse200100.prototype['ruleGroups'] = undefined;






export default InlineResponse200100;

