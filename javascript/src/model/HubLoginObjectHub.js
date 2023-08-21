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

/**
 * The HubLoginObjectHub model module.
 * @module model/HubLoginObjectHub
 * @version 6.2.1
 */
class HubLoginObjectHub {
    /**
     * Constructs a new <code>HubLoginObjectHub</code>.
     * @alias module:model/HubLoginObjectHub
     * @param email {String} Email for existing Morpheus Hub user
     * @param password {String} Password for existing Morpheus Hub user
     */
    constructor(email, password) { 
        
        HubLoginObjectHub.initialize(this, email, password);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, email, password) { 
        obj['email'] = email;
        obj['password'] = password;
    }

    /**
     * Constructs a <code>HubLoginObjectHub</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/HubLoginObjectHub} obj Optional instance to populate.
     * @return {module:model/HubLoginObjectHub} The populated <code>HubLoginObjectHub</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new HubLoginObjectHub();

            if (data.hasOwnProperty('email')) {
                obj['email'] = ApiClient.convertToType(data['email'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Email for existing Morpheus Hub user
 * @member {String} email
 */
HubLoginObjectHub.prototype['email'] = undefined;

/**
 * Password for existing Morpheus Hub user
 * @member {String} password
 */
HubLoginObjectHub.prototype['password'] = undefined;






export default HubLoginObjectHub;

