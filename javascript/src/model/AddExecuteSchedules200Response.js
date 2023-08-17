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
import ExecuteSchedule from './ExecuteSchedule';
import Model200Success from './Model200Success';

/**
 * The AddExecuteSchedules200Response model module.
 * @module model/AddExecuteSchedules200Response
 * @version 6.1.1
 */
class AddExecuteSchedules200Response {
    /**
     * Constructs a new <code>AddExecuteSchedules200Response</code>.
     * @alias module:model/AddExecuteSchedules200Response
     * @implements module:model/Model200Success
     */
    constructor() { 
        Model200Success.initialize(this);
        AddExecuteSchedules200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AddExecuteSchedules200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AddExecuteSchedules200Response} obj Optional instance to populate.
     * @return {module:model/AddExecuteSchedules200Response} The populated <code>AddExecuteSchedules200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AddExecuteSchedules200Response();
            Model200Success.constructFromObject(data, obj);

            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
            if (data.hasOwnProperty('schedule')) {
                obj['schedule'] = ExecuteSchedule.constructFromObject(data['schedule']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AddExecuteSchedules200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AddExecuteSchedules200Response</code>.
     */
    static validateJSON(data) {
        // validate the optional field `schedule`
        if (data['schedule']) { // data not null
          ExecuteSchedule.validateJSON(data['schedule']);
        }

        return true;
    }


}



/**
 * @member {Boolean} success
 */
AddExecuteSchedules200Response.prototype['success'] = undefined;

/**
 * @member {module:model/ExecuteSchedule} schedule
 */
AddExecuteSchedules200Response.prototype['schedule'] = undefined;


// Implement Model200Success interface:
/**
 * @member {Boolean} success
 */
Model200Success.prototype['success'] = undefined;




export default AddExecuteSchedules200Response;

