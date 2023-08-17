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
import ApplianceSettings from './ApplianceSettings';

/**
 * The ListApplianceSettings200Response model module.
 * @module model/ListApplianceSettings200Response
 * @version 6.1.1
 */
class ListApplianceSettings200Response {
    /**
     * Constructs a new <code>ListApplianceSettings200Response</code>.
     * @alias module:model/ListApplianceSettings200Response
     */
    constructor() { 
        
        ListApplianceSettings200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ListApplianceSettings200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ListApplianceSettings200Response} obj Optional instance to populate.
     * @return {module:model/ListApplianceSettings200Response} The populated <code>ListApplianceSettings200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ListApplianceSettings200Response();

            if (data.hasOwnProperty('applianceSettings')) {
                obj['applianceSettings'] = ApplianceSettings.constructFromObject(data['applianceSettings']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ListApplianceSettings200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ListApplianceSettings200Response</code>.
     */
    static validateJSON(data) {
        // validate the optional field `applianceSettings`
        if (data['applianceSettings']) { // data not null
          ApplianceSettings.validateJSON(data['applianceSettings']);
        }

        return true;
    }


}



/**
 * @member {module:model/ApplianceSettings} applianceSettings
 */
ListApplianceSettings200Response.prototype['applianceSettings'] = undefined;






export default ListApplianceSettings200Response;
