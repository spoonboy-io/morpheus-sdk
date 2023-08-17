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

/**
 * The AppPrepareApplyDataTerraform model module.
 * @module model/AppPrepareApplyDataTerraform
 * @version 6.1.1
 */
class AppPrepareApplyDataTerraform {
    /**
     * Constructs a new <code>AppPrepareApplyDataTerraform</code>.
     * @alias module:model/AppPrepareApplyDataTerraform
     */
    constructor() { 
        
        AppPrepareApplyDataTerraform.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AppPrepareApplyDataTerraform</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AppPrepareApplyDataTerraform} obj Optional instance to populate.
     * @return {module:model/AppPrepareApplyDataTerraform} The populated <code>AppPrepareApplyDataTerraform</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AppPrepareApplyDataTerraform();

            if (data.hasOwnProperty('refreshMode')) {
                obj['refreshMode'] = ApiClient.convertToType(data['refreshMode'], 'String');
            }
            if (data.hasOwnProperty('backendType')) {
                obj['backendType'] = ApiClient.convertToType(data['backendType'], 'String');
            }
            if (data.hasOwnProperty('timeoutMode')) {
                obj['timeoutMode'] = ApiClient.convertToType(data['timeoutMode'], 'String');
            }
            if (data.hasOwnProperty('configType')) {
                obj['configType'] = ApiClient.convertToType(data['configType'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AppPrepareApplyDataTerraform</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AppPrepareApplyDataTerraform</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['refreshMode'] && !(typeof data['refreshMode'] === 'string' || data['refreshMode'] instanceof String)) {
            throw new Error("Expected the field `refreshMode` to be a primitive type in the JSON string but got " + data['refreshMode']);
        }
        // ensure the json data is a string
        if (data['backendType'] && !(typeof data['backendType'] === 'string' || data['backendType'] instanceof String)) {
            throw new Error("Expected the field `backendType` to be a primitive type in the JSON string but got " + data['backendType']);
        }
        // ensure the json data is a string
        if (data['timeoutMode'] && !(typeof data['timeoutMode'] === 'string' || data['timeoutMode'] instanceof String)) {
            throw new Error("Expected the field `timeoutMode` to be a primitive type in the JSON string but got " + data['timeoutMode']);
        }
        // ensure the json data is a string
        if (data['configType'] && !(typeof data['configType'] === 'string' || data['configType'] instanceof String)) {
            throw new Error("Expected the field `configType` to be a primitive type in the JSON string but got " + data['configType']);
        }

        return true;
    }


}



/**
 * @member {String} refreshMode
 */
AppPrepareApplyDataTerraform.prototype['refreshMode'] = undefined;

/**
 * @member {String} backendType
 */
AppPrepareApplyDataTerraform.prototype['backendType'] = undefined;

/**
 * @member {String} timeoutMode
 */
AppPrepareApplyDataTerraform.prototype['timeoutMode'] = undefined;

/**
 * @member {String} configType
 */
AppPrepareApplyDataTerraform.prototype['configType'] = undefined;






export default AppPrepareApplyDataTerraform;

