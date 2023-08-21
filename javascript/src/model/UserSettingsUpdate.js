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
import UserSettingsUpdateDefaultCloud from './UserSettingsUpdateDefaultCloud';
import UserSettingsUpdateDefaultGroup from './UserSettingsUpdateDefaultGroup';
import UserSettingsUpdateDefaultPersona from './UserSettingsUpdateDefaultPersona';

/**
 * The UserSettingsUpdate model module.
 * @module model/UserSettingsUpdate
 * @version 6.2.1
 */
class UserSettingsUpdate {
    /**
     * Constructs a new <code>UserSettingsUpdate</code>.
     * @alias module:model/UserSettingsUpdate
     */
    constructor() { 
        
        UserSettingsUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>UserSettingsUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UserSettingsUpdate} obj Optional instance to populate.
     * @return {module:model/UserSettingsUpdate} The populated <code>UserSettingsUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UserSettingsUpdate();

            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('email')) {
                obj['email'] = ApiClient.convertToType(data['email'], 'String');
            }
            if (data.hasOwnProperty('firstName')) {
                obj['firstName'] = ApiClient.convertToType(data['firstName'], 'String');
            }
            if (data.hasOwnProperty('lastName')) {
                obj['lastName'] = ApiClient.convertToType(data['lastName'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
            if (data.hasOwnProperty('linuxUsername')) {
                obj['linuxUsername'] = ApiClient.convertToType(data['linuxUsername'], 'String');
            }
            if (data.hasOwnProperty('linuxPassword')) {
                obj['linuxPassword'] = ApiClient.convertToType(data['linuxPassword'], 'String');
            }
            if (data.hasOwnProperty('linuxKeyPairId')) {
                obj['linuxKeyPairId'] = ApiClient.convertToType(data['linuxKeyPairId'], 'Number');
            }
            if (data.hasOwnProperty('windowsUsername')) {
                obj['windowsUsername'] = ApiClient.convertToType(data['windowsUsername'], 'String');
            }
            if (data.hasOwnProperty('windowsPassword')) {
                obj['windowsPassword'] = ApiClient.convertToType(data['windowsPassword'], 'String');
            }
            if (data.hasOwnProperty('receiveNotifications')) {
                obj['receiveNotifications'] = ApiClient.convertToType(data['receiveNotifications'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultGroup')) {
                obj['defaultGroup'] = UserSettingsUpdateDefaultGroup.constructFromObject(data['defaultGroup']);
            }
            if (data.hasOwnProperty('defaultCloud')) {
                obj['defaultCloud'] = UserSettingsUpdateDefaultCloud.constructFromObject(data['defaultCloud']);
            }
            if (data.hasOwnProperty('defaultPersona')) {
                obj['defaultPersona'] = UserSettingsUpdateDefaultPersona.constructFromObject(data['defaultPersona']);
            }
        }
        return obj;
    }


}

/**
 * Username
 * @member {String} username
 */
UserSettingsUpdate.prototype['username'] = undefined;

/**
 * Email
 * @member {String} email
 */
UserSettingsUpdate.prototype['email'] = undefined;

/**
 * First Name
 * @member {String} firstName
 */
UserSettingsUpdate.prototype['firstName'] = undefined;

/**
 * Last Name
 * @member {String} lastName
 */
UserSettingsUpdate.prototype['lastName'] = undefined;

/**
 * Change your password
 * @member {String} password
 */
UserSettingsUpdate.prototype['password'] = undefined;

/**
 * Linux Username
 * @member {String} linuxUsername
 */
UserSettingsUpdate.prototype['linuxUsername'] = undefined;

/**
 * Linux Password
 * @member {String} linuxPassword
 */
UserSettingsUpdate.prototype['linuxPassword'] = undefined;

/**
 * Linux Key Pair ID
 * @member {Number} linuxKeyPairId
 */
UserSettingsUpdate.prototype['linuxKeyPairId'] = undefined;

/**
 * Windows Username
 * @member {String} windowsUsername
 */
UserSettingsUpdate.prototype['windowsUsername'] = undefined;

/**
 * Windows Password
 * @member {String} windowsPassword
 */
UserSettingsUpdate.prototype['windowsPassword'] = undefined;

/**
 * Receive Notifications (true or false)
 * @member {Boolean} receiveNotifications
 */
UserSettingsUpdate.prototype['receiveNotifications'] = undefined;

/**
 * @member {module:model/UserSettingsUpdateDefaultGroup} defaultGroup
 */
UserSettingsUpdate.prototype['defaultGroup'] = undefined;

/**
 * @member {module:model/UserSettingsUpdateDefaultCloud} defaultCloud
 */
UserSettingsUpdate.prototype['defaultCloud'] = undefined;

/**
 * @member {module:model/UserSettingsUpdateDefaultPersona} defaultPersona
 */
UserSettingsUpdate.prototype['defaultPersona'] = undefined;






export default UserSettingsUpdate;

