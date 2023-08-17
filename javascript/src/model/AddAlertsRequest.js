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
import AddAlertsRequestAlert from './AddAlertsRequestAlert';

/**
 * The AddAlertsRequest model module.
 * @module model/AddAlertsRequest
 * @version 6.1.1
 */
class AddAlertsRequest {
    /**
     * Constructs a new <code>AddAlertsRequest</code>.
     * @alias module:model/AddAlertsRequest
     * @param alert {module:model/AddAlertsRequestAlert} 
     */
    constructor(alert) { 
        
        AddAlertsRequest.initialize(this, alert);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, alert) { 
        obj['alert'] = alert;
    }

    /**
     * Constructs a <code>AddAlertsRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AddAlertsRequest} obj Optional instance to populate.
     * @return {module:model/AddAlertsRequest} The populated <code>AddAlertsRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AddAlertsRequest();

            if (data.hasOwnProperty('alert')) {
                obj['alert'] = AddAlertsRequestAlert.constructFromObject(data['alert']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AddAlertsRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AddAlertsRequest</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of AddAlertsRequest.RequiredProperties) {
            if (!data[property]) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // validate the optional field `alert`
        if (data['alert']) { // data not null
          AddAlertsRequestAlert.validateJSON(data['alert']);
        }

        return true;
    }


}

AddAlertsRequest.RequiredProperties = ["alert"];

/**
 * @member {module:model/AddAlertsRequestAlert} alert
 */
AddAlertsRequest.prototype['alert'] = undefined;






export default AddAlertsRequest;
