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
import PreseedScript from './PreseedScript';

/**
 * The InlineResponse200122 model module.
 * @module model/InlineResponse200122
 * @version 6.2.1
 */
class InlineResponse200122 {
    /**
     * Constructs a new <code>InlineResponse200122</code>.
     * @alias module:model/InlineResponse200122
     */
    constructor() { 
        
        InlineResponse200122.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200122</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200122} obj Optional instance to populate.
     * @return {module:model/InlineResponse200122} The populated <code>InlineResponse200122</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200122();

            if (data.hasOwnProperty('preseedScript')) {
                obj['preseedScript'] = PreseedScript.constructFromObject(data['preseedScript']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/PreseedScript} preseedScript
 */
InlineResponse200122.prototype['preseedScript'] = undefined;






export default InlineResponse200122;

