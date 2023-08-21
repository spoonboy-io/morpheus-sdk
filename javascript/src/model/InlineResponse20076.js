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
import SpecTemplate from './SpecTemplate';

/**
 * The InlineResponse20076 model module.
 * @module model/InlineResponse20076
 * @version 6.2.1
 */
class InlineResponse20076 {
    /**
     * Constructs a new <code>InlineResponse20076</code>.
     * @alias module:model/InlineResponse20076
     */
    constructor() { 
        
        InlineResponse20076.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20076</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20076} obj Optional instance to populate.
     * @return {module:model/InlineResponse20076} The populated <code>InlineResponse20076</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20076();

            if (data.hasOwnProperty('specTemplate')) {
                obj['specTemplate'] = SpecTemplate.constructFromObject(data['specTemplate']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/SpecTemplate} specTemplate
 */
InlineResponse20076.prototype['specTemplate'] = undefined;






export default InlineResponse20076;
