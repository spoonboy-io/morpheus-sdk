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
 * The UserSettingsAccessTokens model module.
 * @module model/UserSettingsAccessTokens
 * @version 6.2.1
 */
class UserSettingsAccessTokens {
    /**
     * Constructs a new <code>UserSettingsAccessTokens</code>.
     * @alias module:model/UserSettingsAccessTokens
     */
    constructor() { 
        
        UserSettingsAccessTokens.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>UserSettingsAccessTokens</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UserSettingsAccessTokens} obj Optional instance to populate.
     * @return {module:model/UserSettingsAccessTokens} The populated <code>UserSettingsAccessTokens</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UserSettingsAccessTokens();

            if (data.hasOwnProperty('clientId')) {
                obj['clientId'] = ApiClient.convertToType(data['clientId'], 'String');
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('expiration')) {
                obj['expiration'] = ApiClient.convertToType(data['expiration'], 'Date');
            }
            if (data.hasOwnProperty('tokenType')) {
                obj['tokenType'] = ApiClient.convertToType(data['tokenType'], 'String');
            }
            if (data.hasOwnProperty('maskedAccessToken')) {
                obj['maskedAccessToken'] = ApiClient.convertToType(data['maskedAccessToken'], 'String');
            }
            if (data.hasOwnProperty('maskedRefreshToken')) {
                obj['maskedRefreshToken'] = ApiClient.convertToType(data['maskedRefreshToken'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} clientId
 */
UserSettingsAccessTokens.prototype['clientId'] = undefined;

/**
 * @member {String} username
 */
UserSettingsAccessTokens.prototype['username'] = undefined;

/**
 * @member {Date} expiration
 */
UserSettingsAccessTokens.prototype['expiration'] = undefined;

/**
 * @member {String} tokenType
 */
UserSettingsAccessTokens.prototype['tokenType'] = undefined;

/**
 * @member {String} maskedAccessToken
 */
UserSettingsAccessTokens.prototype['maskedAccessToken'] = undefined;

/**
 * @member {String} maskedRefreshToken
 */
UserSettingsAccessTokens.prototype['maskedRefreshToken'] = undefined;






export default UserSettingsAccessTokens;
