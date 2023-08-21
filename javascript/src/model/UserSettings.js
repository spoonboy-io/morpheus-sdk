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
import UserSettingsAccessTokens from './UserSettingsAccessTokens';
import UserSettingsUser from './UserSettingsUser';

/**
 * The UserSettings model module.
 * @module model/UserSettings
 * @version 6.2.1
 */
class UserSettings {
    /**
     * Constructs a new <code>UserSettings</code>.
     * @alias module:model/UserSettings
     */
    constructor() { 
        
        UserSettings.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>UserSettings</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UserSettings} obj Optional instance to populate.
     * @return {module:model/UserSettings} The populated <code>UserSettings</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UserSettings();

            if (data.hasOwnProperty('user')) {
                obj['user'] = UserSettingsUser.constructFromObject(data['user']);
            }
            if (data.hasOwnProperty('accessTokens')) {
                obj['accessTokens'] = ApiClient.convertToType(data['accessTokens'], [UserSettingsAccessTokens]);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/UserSettingsUser} user
 */
UserSettings.prototype['user'] = undefined;

/**
 * @member {Array.<module:model/UserSettingsAccessTokens>} accessTokens
 */
UserSettings.prototype['accessTokens'] = undefined;






export default UserSettings;

