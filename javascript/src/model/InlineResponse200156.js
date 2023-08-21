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
import TaskType from './TaskType';

/**
 * The InlineResponse200156 model module.
 * @module model/InlineResponse200156
 * @version 6.2.1
 */
class InlineResponse200156 {
    /**
     * Constructs a new <code>InlineResponse200156</code>.
     * @alias module:model/InlineResponse200156
     */
    constructor() { 
        
        InlineResponse200156.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200156</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200156} obj Optional instance to populate.
     * @return {module:model/InlineResponse200156} The populated <code>InlineResponse200156</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200156();

            if (data.hasOwnProperty('taskType')) {
                obj['taskType'] = TaskType.constructFromObject(data['taskType']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/TaskType} taskType
 */
InlineResponse200156.prototype['taskType'] = undefined;






export default InlineResponse200156;

