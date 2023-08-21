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
import NetworkRouterFirewallRuleCreate from './NetworkRouterFirewallRuleCreate';

/**
 * The InlineObject153 model module.
 * @module model/InlineObject153
 * @version 6.2.1
 */
class InlineObject153 {
    /**
     * Constructs a new <code>InlineObject153</code>.
     * The parameters for updating a network router is type dependent. The following lists the common parameters. See get a specific type to list available options for that network router type. 
     * @alias module:model/InlineObject153
     */
    constructor() { 
        
        InlineObject153.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject153</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject153} obj Optional instance to populate.
     * @return {module:model/InlineObject153} The populated <code>InlineObject153</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject153();

            if (data.hasOwnProperty('rule')) {
                obj['rule'] = NetworkRouterFirewallRuleCreate.constructFromObject(data['rule']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/NetworkRouterFirewallRuleCreate} rule
 */
InlineObject153.prototype['rule'] = undefined;






export default InlineObject153;

