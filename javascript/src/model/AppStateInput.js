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
import AppStateInputDataInner from './AppStateInputDataInner';
import AppStateInputProvidersInner from './AppStateInputProvidersInner';
import AppStateInputVariablesInner from './AppStateInputVariablesInner';

/**
 * The AppStateInput model module.
 * @module model/AppStateInput
 * @version 6.1.1
 */
class AppStateInput {
    /**
     * Constructs a new <code>AppStateInput</code>.
     * @alias module:model/AppStateInput
     */
    constructor() { 
        
        AppStateInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AppStateInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AppStateInput} obj Optional instance to populate.
     * @return {module:model/AppStateInput} The populated <code>AppStateInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AppStateInput();

            if (data.hasOwnProperty('variables')) {
                obj['variables'] = ApiClient.convertToType(data['variables'], [AppStateInputVariablesInner]);
            }
            if (data.hasOwnProperty('providers')) {
                obj['providers'] = ApiClient.convertToType(data['providers'], [AppStateInputProvidersInner]);
            }
            if (data.hasOwnProperty('data')) {
                obj['data'] = ApiClient.convertToType(data['data'], [AppStateInputDataInner]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AppStateInput</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AppStateInput</code>.
     */
    static validateJSON(data) {
        if (data['variables']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['variables'])) {
                throw new Error("Expected the field `variables` to be an array in the JSON data but got " + data['variables']);
            }
            // validate the optional field `variables` (array)
            for (const item of data['variables']) {
                AppStateInputVariablesInner.validateJSON(item);
            };
        }
        if (data['providers']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['providers'])) {
                throw new Error("Expected the field `providers` to be an array in the JSON data but got " + data['providers']);
            }
            // validate the optional field `providers` (array)
            for (const item of data['providers']) {
                AppStateInputProvidersInner.validateJSON(item);
            };
        }
        if (data['data']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['data'])) {
                throw new Error("Expected the field `data` to be an array in the JSON data but got " + data['data']);
            }
            // validate the optional field `data` (array)
            for (const item of data['data']) {
                AppStateInputDataInner.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {Array.<module:model/AppStateInputVariablesInner>} variables
 */
AppStateInput.prototype['variables'] = undefined;

/**
 * @member {Array.<module:model/AppStateInputProvidersInner>} providers
 */
AppStateInput.prototype['providers'] = undefined;

/**
 * @member {Array.<module:model/AppStateInputDataInner>} data
 */
AppStateInput.prototype['data'] = undefined;






export default AppStateInput;

