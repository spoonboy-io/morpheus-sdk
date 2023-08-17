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
 * The ForgotPassword200Response model module.
 * @module model/ForgotPassword200Response
 * @version 6.1.1
 */
class ForgotPassword200Response {
    /**
     * Constructs a new <code>ForgotPassword200Response</code>.
     * @alias module:model/ForgotPassword200Response
     */
    constructor() { 
        
        ForgotPassword200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ForgotPassword200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ForgotPassword200Response} obj Optional instance to populate.
     * @return {module:model/ForgotPassword200Response} The populated <code>ForgotPassword200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ForgotPassword200Response();

            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
            if (data.hasOwnProperty('msg')) {
                obj['msg'] = ApiClient.convertToType(data['msg'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ForgotPassword200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ForgotPassword200Response</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['msg'] && !(typeof data['msg'] === 'string' || data['msg'] instanceof String)) {
            throw new Error("Expected the field `msg` to be a primitive type in the JSON string but got " + data['msg']);
        }

        return true;
    }


}



/**
 * @member {Boolean} success
 */
ForgotPassword200Response.prototype['success'] = undefined;

/**
 * @member {String} msg
 */
ForgotPassword200Response.prototype['msg'] = undefined;






export default ForgotPassword200Response;

