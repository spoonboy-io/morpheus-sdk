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
 * The InlineObject6 model module.
 * @module model/InlineObject6
 * @version 6.2.1
 */
class InlineObject6 {
    /**
     * Constructs a new <code>InlineObject6</code>.
     * @alias module:model/InlineObject6
     */
    constructor() { 
        
        InlineObject6.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject6</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject6} obj Optional instance to populate.
     * @return {module:model/InlineObject6} The populated <code>InlineObject6</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject6();

            if (data.hasOwnProperty('templateParameter')) {
                obj['templateParameter'] = ApiClient.convertToType(data['templateParameter'], Object);
            }
        }
        return obj;
    }


}

/**
 * Template Parameter object. A map of key-value pairs that correspond to the template variables i.e. tfvars
 * @member {Object} templateParameter
 */
InlineObject6.prototype['templateParameter'] = undefined;






export default InlineObject6;

