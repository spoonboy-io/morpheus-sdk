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
import CheckType from './CheckType';

/**
 * The InlineResponse20018 model module.
 * @module model/InlineResponse20018
 * @version 6.2.1
 */
class InlineResponse20018 {
    /**
     * Constructs a new <code>InlineResponse20018</code>.
     * @alias module:model/InlineResponse20018
     */
    constructor() { 
        
        InlineResponse20018.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20018</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20018} obj Optional instance to populate.
     * @return {module:model/InlineResponse20018} The populated <code>InlineResponse20018</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20018();

            if (data.hasOwnProperty('checkType')) {
                obj['checkType'] = CheckType.constructFromObject(data['checkType']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/CheckType} checkType
 */
InlineResponse20018.prototype['checkType'] = undefined;






export default InlineResponse20018;
