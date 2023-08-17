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
 * The Client model module.
 * @module model/Client
 * @version 6.1.1
 */
class Client {
    /**
     * Constructs a new <code>Client</code>.
     * @alias module:model/Client
     */
    constructor() { 
        
        Client.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Client</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Client} obj Optional instance to populate.
     * @return {module:model/Client} The populated <code>Client</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Client();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('clientId')) {
                obj['clientId'] = ApiClient.convertToType(data['clientId'], 'String');
            }
            if (data.hasOwnProperty('accessTokenValiditySeconds')) {
                obj['accessTokenValiditySeconds'] = ApiClient.convertToType(data['accessTokenValiditySeconds'], 'Number');
            }
            if (data.hasOwnProperty('refreshTokenValiditySeconds')) {
                obj['refreshTokenValiditySeconds'] = ApiClient.convertToType(data['refreshTokenValiditySeconds'], 'Number');
            }
            if (data.hasOwnProperty('authorities')) {
                obj['authorities'] = ApiClient.convertToType(data['authorities'], ['String']);
            }
            if (data.hasOwnProperty('authorizedGrantTypes')) {
                obj['authorizedGrantTypes'] = ApiClient.convertToType(data['authorizedGrantTypes'], ['String']);
            }
            if (data.hasOwnProperty('scopes')) {
                obj['scopes'] = ApiClient.convertToType(data['scopes'], ['String']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>Client</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>Client</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['clientId'] && !(typeof data['clientId'] === 'string' || data['clientId'] instanceof String)) {
            throw new Error("Expected the field `clientId` to be a primitive type in the JSON string but got " + data['clientId']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['authorities'])) {
            throw new Error("Expected the field `authorities` to be an array in the JSON data but got " + data['authorities']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['authorizedGrantTypes'])) {
            throw new Error("Expected the field `authorizedGrantTypes` to be an array in the JSON data but got " + data['authorizedGrantTypes']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['scopes'])) {
            throw new Error("Expected the field `scopes` to be an array in the JSON data but got " + data['scopes']);
        }

        return true;
    }


}



/**
 * @member {Number} id
 */
Client.prototype['id'] = undefined;

/**
 * @member {String} clientId
 */
Client.prototype['clientId'] = undefined;

/**
 * @member {Number} accessTokenValiditySeconds
 */
Client.prototype['accessTokenValiditySeconds'] = undefined;

/**
 * @member {Number} refreshTokenValiditySeconds
 */
Client.prototype['refreshTokenValiditySeconds'] = undefined;

/**
 * @member {Array.<String>} authorities
 */
Client.prototype['authorities'] = undefined;

/**
 * @member {Array.<String>} authorizedGrantTypes
 */
Client.prototype['authorizedGrantTypes'] = undefined;

/**
 * @member {Array.<String>} scopes
 */
Client.prototype['scopes'] = undefined;






export default Client;
