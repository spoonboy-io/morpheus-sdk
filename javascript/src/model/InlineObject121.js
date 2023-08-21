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
import OptionTypeCreate from './OptionTypeCreate';

/**
 * The InlineObject121 model module.
 * @module model/InlineObject121
 * @version 6.2.1
 */
class InlineObject121 {
    /**
     * Constructs a new <code>InlineObject121</code>.
     * @alias module:model/InlineObject121
     */
    constructor() { 
        
        InlineObject121.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject121</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject121} obj Optional instance to populate.
     * @return {module:model/InlineObject121} The populated <code>InlineObject121</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject121();

            if (data.hasOwnProperty('optionType')) {
                obj['optionType'] = OptionTypeCreate.constructFromObject(data['optionType']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/OptionTypeCreate} optionType
 */
InlineObject121.prototype['optionType'] = undefined;






export default InlineObject121;
