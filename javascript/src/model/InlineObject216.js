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
import ApiSecurityGroupsIdRulesRule from './ApiSecurityGroupsIdRulesRule';

/**
 * The InlineObject216 model module.
 * @module model/InlineObject216
 * @version 6.2.1
 */
class InlineObject216 {
    /**
     * Constructs a new <code>InlineObject216</code>.
     * @alias module:model/InlineObject216
     * @param rule {module:model/ApiSecurityGroupsIdRulesRule} 
     */
    constructor(rule) { 
        
        InlineObject216.initialize(this, rule);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, rule) { 
        obj['rule'] = rule;
    }

    /**
     * Constructs a <code>InlineObject216</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject216} obj Optional instance to populate.
     * @return {module:model/InlineObject216} The populated <code>InlineObject216</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject216();

            if (data.hasOwnProperty('rule')) {
                obj['rule'] = ApiSecurityGroupsIdRulesRule.constructFromObject(data['rule']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiSecurityGroupsIdRulesRule} rule
 */
InlineObject216.prototype['rule'] = undefined;






export default InlineObject216;

