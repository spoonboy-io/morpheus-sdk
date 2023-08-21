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
 * The InlineResponse200101 model module.
 * @module model/InlineResponse200101
 * @version 6.2.1
 */
class InlineResponse200101 {
    /**
     * Constructs a new <code>InlineResponse200101</code>.
     * @alias module:model/InlineResponse200101
     */
    constructor() { 
        
        InlineResponse200101.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200101</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200101} obj Optional instance to populate.
     * @return {module:model/InlineResponse200101} The populated <code>InlineResponse200101</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200101();

            if (data.hasOwnProperty('ruleGroup')) {
                obj['ruleGroup'] = ApiClient.convertToType(data['ruleGroup'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {Object} ruleGroup
 */
InlineResponse200101.prototype['ruleGroup'] = undefined;






export default InlineResponse200101;
