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
import ApplianceSettingsEnabledZoneTypesInner from './ApplianceSettingsEnabledZoneTypesInner';

/**
 * The ApplianceSettings model module.
 * @module model/ApplianceSettings
 * @version 6.1.1
 */
class ApplianceSettings {
    /**
     * Constructs a new <code>ApplianceSettings</code>.
     * @alias module:model/ApplianceSettings
     */
    constructor() { 
        
        ApplianceSettings.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApplianceSettings</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApplianceSettings} obj Optional instance to populate.
     * @return {module:model/ApplianceSettings} The populated <code>ApplianceSettings</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApplianceSettings();

            if (data.hasOwnProperty('applianceUrl')) {
                obj['applianceUrl'] = ApiClient.convertToType(data['applianceUrl'], 'String');
            }
            if (data.hasOwnProperty('internalApplianceUrl')) {
                obj['internalApplianceUrl'] = ApiClient.convertToType(data['internalApplianceUrl'], 'String');
            }
            if (data.hasOwnProperty('corsAllowed')) {
                obj['corsAllowed'] = ApiClient.convertToType(data['corsAllowed'], 'String');
            }
            if (data.hasOwnProperty('registrationEnabled')) {
                obj['registrationEnabled'] = ApiClient.convertToType(data['registrationEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultRoleId')) {
                obj['defaultRoleId'] = ApiClient.convertToType(data['defaultRoleId'], 'String');
            }
            if (data.hasOwnProperty('defaultUserRoleId')) {
                obj['defaultUserRoleId'] = ApiClient.convertToType(data['defaultUserRoleId'], 'String');
            }
            if (data.hasOwnProperty('dockerPrivilegedMode')) {
                obj['dockerPrivilegedMode'] = ApiClient.convertToType(data['dockerPrivilegedMode'], 'Boolean');
            }
            if (data.hasOwnProperty('expirePwdDays')) {
                obj['expirePwdDays'] = ApiClient.convertToType(data['expirePwdDays'], 'String');
            }
            if (data.hasOwnProperty('disableAfterAttempts')) {
                obj['disableAfterAttempts'] = ApiClient.convertToType(data['disableAfterAttempts'], 'String');
            }
            if (data.hasOwnProperty('disableAfterDaysInactive')) {
                obj['disableAfterDaysInactive'] = ApiClient.convertToType(data['disableAfterDaysInactive'], 'String');
            }
            if (data.hasOwnProperty('warnUserDaysBefore')) {
                obj['warnUserDaysBefore'] = ApiClient.convertToType(data['warnUserDaysBefore'], 'String');
            }
            if (data.hasOwnProperty('smtpMailFrom')) {
                obj['smtpMailFrom'] = ApiClient.convertToType(data['smtpMailFrom'], 'String');
            }
            if (data.hasOwnProperty('smtpServer')) {
                obj['smtpServer'] = ApiClient.convertToType(data['smtpServer'], 'String');
            }
            if (data.hasOwnProperty('smtpPort')) {
                obj['smtpPort'] = ApiClient.convertToType(data['smtpPort'], 'String');
            }
            if (data.hasOwnProperty('smtpSSL')) {
                obj['smtpSSL'] = ApiClient.convertToType(data['smtpSSL'], 'Boolean');
            }
            if (data.hasOwnProperty('smtpTLS')) {
                obj['smtpTLS'] = ApiClient.convertToType(data['smtpTLS'], 'Boolean');
            }
            if (data.hasOwnProperty('smtpUser')) {
                obj['smtpUser'] = ApiClient.convertToType(data['smtpUser'], 'String');
            }
            if (data.hasOwnProperty('smtpPassword')) {
                obj['smtpPassword'] = ApiClient.convertToType(data['smtpPassword'], 'String');
            }
            if (data.hasOwnProperty('smtpPasswordHash')) {
                obj['smtpPasswordHash'] = ApiClient.convertToType(data['smtpPasswordHash'], 'String');
            }
            if (data.hasOwnProperty('proxyHost')) {
                obj['proxyHost'] = ApiClient.convertToType(data['proxyHost'], 'String');
            }
            if (data.hasOwnProperty('proxyPort')) {
                obj['proxyPort'] = ApiClient.convertToType(data['proxyPort'], 'String');
            }
            if (data.hasOwnProperty('proxyUser')) {
                obj['proxyUser'] = ApiClient.convertToType(data['proxyUser'], 'String');
            }
            if (data.hasOwnProperty('proxyPassword')) {
                obj['proxyPassword'] = ApiClient.convertToType(data['proxyPassword'], 'String');
            }
            if (data.hasOwnProperty('proxyPasswordHash')) {
                obj['proxyPasswordHash'] = ApiClient.convertToType(data['proxyPasswordHash'], 'String');
            }
            if (data.hasOwnProperty('proxyDomain')) {
                obj['proxyDomain'] = ApiClient.convertToType(data['proxyDomain'], 'String');
            }
            if (data.hasOwnProperty('proxyWorkstation')) {
                obj['proxyWorkstation'] = ApiClient.convertToType(data['proxyWorkstation'], 'String');
            }
            if (data.hasOwnProperty('currencyProvider')) {
                obj['currencyProvider'] = ApiClient.convertToType(data['currencyProvider'], 'String');
            }
            if (data.hasOwnProperty('currencyKey')) {
                obj['currencyKey'] = ApiClient.convertToType(data['currencyKey'], 'String');
            }
            if (data.hasOwnProperty('enabledZoneTypes')) {
                obj['enabledZoneTypes'] = ApiClient.convertToType(data['enabledZoneTypes'], [ApplianceSettingsEnabledZoneTypesInner]);
            }
            if (data.hasOwnProperty('statsRetainmentPeriod')) {
                obj['statsRetainmentPeriod'] = ApiClient.convertToType(data['statsRetainmentPeriod'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ApplianceSettings</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ApplianceSettings</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['applianceUrl'] && !(typeof data['applianceUrl'] === 'string' || data['applianceUrl'] instanceof String)) {
            throw new Error("Expected the field `applianceUrl` to be a primitive type in the JSON string but got " + data['applianceUrl']);
        }
        // ensure the json data is a string
        if (data['internalApplianceUrl'] && !(typeof data['internalApplianceUrl'] === 'string' || data['internalApplianceUrl'] instanceof String)) {
            throw new Error("Expected the field `internalApplianceUrl` to be a primitive type in the JSON string but got " + data['internalApplianceUrl']);
        }
        // ensure the json data is a string
        if (data['corsAllowed'] && !(typeof data['corsAllowed'] === 'string' || data['corsAllowed'] instanceof String)) {
            throw new Error("Expected the field `corsAllowed` to be a primitive type in the JSON string but got " + data['corsAllowed']);
        }
        // ensure the json data is a string
        if (data['defaultRoleId'] && !(typeof data['defaultRoleId'] === 'string' || data['defaultRoleId'] instanceof String)) {
            throw new Error("Expected the field `defaultRoleId` to be a primitive type in the JSON string but got " + data['defaultRoleId']);
        }
        // ensure the json data is a string
        if (data['defaultUserRoleId'] && !(typeof data['defaultUserRoleId'] === 'string' || data['defaultUserRoleId'] instanceof String)) {
            throw new Error("Expected the field `defaultUserRoleId` to be a primitive type in the JSON string but got " + data['defaultUserRoleId']);
        }
        // ensure the json data is a string
        if (data['expirePwdDays'] && !(typeof data['expirePwdDays'] === 'string' || data['expirePwdDays'] instanceof String)) {
            throw new Error("Expected the field `expirePwdDays` to be a primitive type in the JSON string but got " + data['expirePwdDays']);
        }
        // ensure the json data is a string
        if (data['disableAfterAttempts'] && !(typeof data['disableAfterAttempts'] === 'string' || data['disableAfterAttempts'] instanceof String)) {
            throw new Error("Expected the field `disableAfterAttempts` to be a primitive type in the JSON string but got " + data['disableAfterAttempts']);
        }
        // ensure the json data is a string
        if (data['disableAfterDaysInactive'] && !(typeof data['disableAfterDaysInactive'] === 'string' || data['disableAfterDaysInactive'] instanceof String)) {
            throw new Error("Expected the field `disableAfterDaysInactive` to be a primitive type in the JSON string but got " + data['disableAfterDaysInactive']);
        }
        // ensure the json data is a string
        if (data['warnUserDaysBefore'] && !(typeof data['warnUserDaysBefore'] === 'string' || data['warnUserDaysBefore'] instanceof String)) {
            throw new Error("Expected the field `warnUserDaysBefore` to be a primitive type in the JSON string but got " + data['warnUserDaysBefore']);
        }
        // ensure the json data is a string
        if (data['smtpMailFrom'] && !(typeof data['smtpMailFrom'] === 'string' || data['smtpMailFrom'] instanceof String)) {
            throw new Error("Expected the field `smtpMailFrom` to be a primitive type in the JSON string but got " + data['smtpMailFrom']);
        }
        // ensure the json data is a string
        if (data['smtpServer'] && !(typeof data['smtpServer'] === 'string' || data['smtpServer'] instanceof String)) {
            throw new Error("Expected the field `smtpServer` to be a primitive type in the JSON string but got " + data['smtpServer']);
        }
        // ensure the json data is a string
        if (data['smtpPort'] && !(typeof data['smtpPort'] === 'string' || data['smtpPort'] instanceof String)) {
            throw new Error("Expected the field `smtpPort` to be a primitive type in the JSON string but got " + data['smtpPort']);
        }
        // ensure the json data is a string
        if (data['smtpUser'] && !(typeof data['smtpUser'] === 'string' || data['smtpUser'] instanceof String)) {
            throw new Error("Expected the field `smtpUser` to be a primitive type in the JSON string but got " + data['smtpUser']);
        }
        // ensure the json data is a string
        if (data['smtpPassword'] && !(typeof data['smtpPassword'] === 'string' || data['smtpPassword'] instanceof String)) {
            throw new Error("Expected the field `smtpPassword` to be a primitive type in the JSON string but got " + data['smtpPassword']);
        }
        // ensure the json data is a string
        if (data['smtpPasswordHash'] && !(typeof data['smtpPasswordHash'] === 'string' || data['smtpPasswordHash'] instanceof String)) {
            throw new Error("Expected the field `smtpPasswordHash` to be a primitive type in the JSON string but got " + data['smtpPasswordHash']);
        }
        // ensure the json data is a string
        if (data['proxyHost'] && !(typeof data['proxyHost'] === 'string' || data['proxyHost'] instanceof String)) {
            throw new Error("Expected the field `proxyHost` to be a primitive type in the JSON string but got " + data['proxyHost']);
        }
        // ensure the json data is a string
        if (data['proxyPort'] && !(typeof data['proxyPort'] === 'string' || data['proxyPort'] instanceof String)) {
            throw new Error("Expected the field `proxyPort` to be a primitive type in the JSON string but got " + data['proxyPort']);
        }
        // ensure the json data is a string
        if (data['proxyUser'] && !(typeof data['proxyUser'] === 'string' || data['proxyUser'] instanceof String)) {
            throw new Error("Expected the field `proxyUser` to be a primitive type in the JSON string but got " + data['proxyUser']);
        }
        // ensure the json data is a string
        if (data['proxyPassword'] && !(typeof data['proxyPassword'] === 'string' || data['proxyPassword'] instanceof String)) {
            throw new Error("Expected the field `proxyPassword` to be a primitive type in the JSON string but got " + data['proxyPassword']);
        }
        // ensure the json data is a string
        if (data['proxyPasswordHash'] && !(typeof data['proxyPasswordHash'] === 'string' || data['proxyPasswordHash'] instanceof String)) {
            throw new Error("Expected the field `proxyPasswordHash` to be a primitive type in the JSON string but got " + data['proxyPasswordHash']);
        }
        // ensure the json data is a string
        if (data['proxyDomain'] && !(typeof data['proxyDomain'] === 'string' || data['proxyDomain'] instanceof String)) {
            throw new Error("Expected the field `proxyDomain` to be a primitive type in the JSON string but got " + data['proxyDomain']);
        }
        // ensure the json data is a string
        if (data['proxyWorkstation'] && !(typeof data['proxyWorkstation'] === 'string' || data['proxyWorkstation'] instanceof String)) {
            throw new Error("Expected the field `proxyWorkstation` to be a primitive type in the JSON string but got " + data['proxyWorkstation']);
        }
        // ensure the json data is a string
        if (data['currencyProvider'] && !(typeof data['currencyProvider'] === 'string' || data['currencyProvider'] instanceof String)) {
            throw new Error("Expected the field `currencyProvider` to be a primitive type in the JSON string but got " + data['currencyProvider']);
        }
        // ensure the json data is a string
        if (data['currencyKey'] && !(typeof data['currencyKey'] === 'string' || data['currencyKey'] instanceof String)) {
            throw new Error("Expected the field `currencyKey` to be a primitive type in the JSON string but got " + data['currencyKey']);
        }
        if (data['enabledZoneTypes']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['enabledZoneTypes'])) {
                throw new Error("Expected the field `enabledZoneTypes` to be an array in the JSON data but got " + data['enabledZoneTypes']);
            }
            // validate the optional field `enabledZoneTypes` (array)
            for (const item of data['enabledZoneTypes']) {
                ApplianceSettingsEnabledZoneTypesInner.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {String} applianceUrl
 */
ApplianceSettings.prototype['applianceUrl'] = undefined;

/**
 * @member {String} internalApplianceUrl
 */
ApplianceSettings.prototype['internalApplianceUrl'] = undefined;

/**
 * @member {String} corsAllowed
 */
ApplianceSettings.prototype['corsAllowed'] = undefined;

/**
 * @member {Boolean} registrationEnabled
 */
ApplianceSettings.prototype['registrationEnabled'] = undefined;

/**
 * @member {String} defaultRoleId
 */
ApplianceSettings.prototype['defaultRoleId'] = undefined;

/**
 * @member {String} defaultUserRoleId
 */
ApplianceSettings.prototype['defaultUserRoleId'] = undefined;

/**
 * @member {Boolean} dockerPrivilegedMode
 */
ApplianceSettings.prototype['dockerPrivilegedMode'] = undefined;

/**
 * @member {String} expirePwdDays
 */
ApplianceSettings.prototype['expirePwdDays'] = undefined;

/**
 * @member {String} disableAfterAttempts
 */
ApplianceSettings.prototype['disableAfterAttempts'] = undefined;

/**
 * @member {String} disableAfterDaysInactive
 */
ApplianceSettings.prototype['disableAfterDaysInactive'] = undefined;

/**
 * @member {String} warnUserDaysBefore
 */
ApplianceSettings.prototype['warnUserDaysBefore'] = undefined;

/**
 * @member {String} smtpMailFrom
 */
ApplianceSettings.prototype['smtpMailFrom'] = undefined;

/**
 * @member {String} smtpServer
 */
ApplianceSettings.prototype['smtpServer'] = undefined;

/**
 * @member {String} smtpPort
 */
ApplianceSettings.prototype['smtpPort'] = undefined;

/**
 * @member {Boolean} smtpSSL
 */
ApplianceSettings.prototype['smtpSSL'] = undefined;

/**
 * @member {Boolean} smtpTLS
 */
ApplianceSettings.prototype['smtpTLS'] = undefined;

/**
 * @member {String} smtpUser
 */
ApplianceSettings.prototype['smtpUser'] = undefined;

/**
 * @member {String} smtpPassword
 */
ApplianceSettings.prototype['smtpPassword'] = undefined;

/**
 * @member {String} smtpPasswordHash
 */
ApplianceSettings.prototype['smtpPasswordHash'] = undefined;

/**
 * @member {String} proxyHost
 */
ApplianceSettings.prototype['proxyHost'] = undefined;

/**
 * @member {String} proxyPort
 */
ApplianceSettings.prototype['proxyPort'] = undefined;

/**
 * @member {String} proxyUser
 */
ApplianceSettings.prototype['proxyUser'] = undefined;

/**
 * @member {String} proxyPassword
 */
ApplianceSettings.prototype['proxyPassword'] = undefined;

/**
 * @member {String} proxyPasswordHash
 */
ApplianceSettings.prototype['proxyPasswordHash'] = undefined;

/**
 * @member {String} proxyDomain
 */
ApplianceSettings.prototype['proxyDomain'] = undefined;

/**
 * @member {String} proxyWorkstation
 */
ApplianceSettings.prototype['proxyWorkstation'] = undefined;

/**
 * @member {String} currencyProvider
 */
ApplianceSettings.prototype['currencyProvider'] = undefined;

/**
 * @member {String} currencyKey
 */
ApplianceSettings.prototype['currencyKey'] = undefined;

/**
 * @member {Array.<module:model/ApplianceSettingsEnabledZoneTypesInner>} enabledZoneTypes
 */
ApplianceSettings.prototype['enabledZoneTypes'] = undefined;

/**
 * @member {Number} statsRetainmentPeriod
 */
ApplianceSettings.prototype['statsRetainmentPeriod'] = undefined;






export default ApplianceSettings;
