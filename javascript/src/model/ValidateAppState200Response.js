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
import Model200Success from './Model200Success';

/**
 * The ValidateAppState200Response model module.
 * @module model/ValidateAppState200Response
 * @version 6.1.1
 */
class ValidateAppState200Response {
    /**
     * Constructs a new <code>ValidateAppState200Response</code>.
     * @alias module:model/ValidateAppState200Response
     * @implements module:model/Model200Success
     */
    constructor() { 
        Model200Success.initialize(this);
        ValidateAppState200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ValidateAppState200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ValidateAppState200Response} obj Optional instance to populate.
     * @return {module:model/ValidateAppState200Response} The populated <code>ValidateAppState200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ValidateAppState200Response();
            Model200Success.constructFromObject(data, obj);

            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
            if (data.hasOwnProperty('executionId')) {
                obj['executionId'] = ApiClient.convertToType(data['executionId'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ValidateAppState200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ValidateAppState200Response</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['executionId'] && !(typeof data['executionId'] === 'string' || data['executionId'] instanceof String)) {
            throw new Error("Expected the field `executionId` to be a primitive type in the JSON string but got " + data['executionId']);
        }

        return true;
    }


}



/**
 * @member {Boolean} success
 */
ValidateAppState200Response.prototype['success'] = undefined;

/**
 * @member {String} executionId
 */
ValidateAppState200Response.prototype['executionId'] = undefined;


// Implement Model200Success interface:
/**
 * @member {Boolean} success
 */
Model200Success.prototype['success'] = undefined;




export default ValidateAppState200Response;

