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
 * The ResetPasswordRequest model module.
 * @module model/ResetPasswordRequest
 * @version 6.1.1
 */
class ResetPasswordRequest {
    /**
     * Constructs a new <code>ResetPasswordRequest</code>.
     * @alias module:model/ResetPasswordRequest
     * @param token {String} The secret Reset Password token that was included in the **Forgot Password Email**.
     * @param password {String} User new password. This is the new password for your user.
     */
    constructor(token, password) { 
        
        ResetPasswordRequest.initialize(this, token, password);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, token, password) { 
        obj['token'] = token;
        obj['password'] = password;
    }

    /**
     * Constructs a <code>ResetPasswordRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ResetPasswordRequest} obj Optional instance to populate.
     * @return {module:model/ResetPasswordRequest} The populated <code>ResetPasswordRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ResetPasswordRequest();

            if (data.hasOwnProperty('token')) {
                obj['token'] = ApiClient.convertToType(data['token'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ResetPasswordRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ResetPasswordRequest</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of ResetPasswordRequest.RequiredProperties) {
            if (!data[property]) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['token'] && !(typeof data['token'] === 'string' || data['token'] instanceof String)) {
            throw new Error("Expected the field `token` to be a primitive type in the JSON string but got " + data['token']);
        }
        // ensure the json data is a string
        if (data['password'] && !(typeof data['password'] === 'string' || data['password'] instanceof String)) {
            throw new Error("Expected the field `password` to be a primitive type in the JSON string but got " + data['password']);
        }

        return true;
    }


}

ResetPasswordRequest.RequiredProperties = ["token", "password"];

/**
 * The secret Reset Password token that was included in the **Forgot Password Email**.
 * @member {String} token
 */
ResetPasswordRequest.prototype['token'] = undefined;

/**
 * User new password. This is the new password for your user.
 * @member {String} password
 */
ResetPasswordRequest.prototype['password'] = undefined;






export default ResetPasswordRequest;
