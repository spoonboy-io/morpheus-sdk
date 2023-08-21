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
import ApiExecuteSchedulesIdSchedule from './ApiExecuteSchedulesIdSchedule';

/**
 * The InlineObject13 model module.
 * @module model/InlineObject13
 * @version 6.2.1
 */
class InlineObject13 {
    /**
     * Constructs a new <code>InlineObject13</code>.
     * @alias module:model/InlineObject13
     * @param schedule {module:model/ApiExecuteSchedulesIdSchedule} 
     */
    constructor(schedule) { 
        
        InlineObject13.initialize(this, schedule);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, schedule) { 
        obj['schedule'] = schedule;
    }

    /**
     * Constructs a <code>InlineObject13</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject13} obj Optional instance to populate.
     * @return {module:model/InlineObject13} The populated <code>InlineObject13</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject13();

            if (data.hasOwnProperty('schedule')) {
                obj['schedule'] = ApiExecuteSchedulesIdSchedule.constructFromObject(data['schedule']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiExecuteSchedulesIdSchedule} schedule
 */
InlineObject13.prototype['schedule'] = undefined;






export default InlineObject13;

