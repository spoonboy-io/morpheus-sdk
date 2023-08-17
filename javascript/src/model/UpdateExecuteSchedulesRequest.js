/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import UpdateExecuteSchedulesRequestSchedule from './UpdateExecuteSchedulesRequestSchedule';

/**
 * The UpdateExecuteSchedulesRequest model module.
 * @module model/UpdateExecuteSchedulesRequest
 * @version 6.1.1
 */
class UpdateExecuteSchedulesRequest {
    /**
     * Constructs a new <code>UpdateExecuteSchedulesRequest</code>.
     * @alias module:model/UpdateExecuteSchedulesRequest
     * @param schedule {module:model/UpdateExecuteSchedulesRequestSchedule} 
     */
    constructor(schedule) { 
        
        UpdateExecuteSchedulesRequest.initialize(this, schedule);
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
     * Constructs a <code>UpdateExecuteSchedulesRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UpdateExecuteSchedulesRequest} obj Optional instance to populate.
     * @return {module:model/UpdateExecuteSchedulesRequest} The populated <code>UpdateExecuteSchedulesRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UpdateExecuteSchedulesRequest();

            if (data.hasOwnProperty('schedule')) {
                obj['schedule'] = UpdateExecuteSchedulesRequestSchedule.constructFromObject(data['schedule']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>UpdateExecuteSchedulesRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>UpdateExecuteSchedulesRequest</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of UpdateExecuteSchedulesRequest.RequiredProperties) {
            if (!data[property]) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // validate the optional field `schedule`
        if (data['schedule']) { // data not null
          UpdateExecuteSchedulesRequestSchedule.validateJSON(data['schedule']);
        }

        return true;
    }


}

UpdateExecuteSchedulesRequest.RequiredProperties = ["schedule"];

/**
 * @member {module:model/UpdateExecuteSchedulesRequestSchedule} schedule
 */
UpdateExecuteSchedulesRequest.prototype['schedule'] = undefined;






export default UpdateExecuteSchedulesRequest;
