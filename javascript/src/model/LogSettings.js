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
 * The LogSettings model module.
 * @module model/LogSettings
 * @version 6.2.1
 */
class LogSettings {
    /**
     * Constructs a new <code>LogSettings</code>.
     * @alias module:model/LogSettings
     */
    constructor() { 
        
        LogSettings.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>LogSettings</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/LogSettings} obj Optional instance to populate.
     * @return {module:model/LogSettings} The populated <code>LogSettings</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new LogSettings();

            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('retentionDays')) {
                obj['retentionDays'] = ApiClient.convertToType(data['retentionDays'], 'String');
            }
            if (data.hasOwnProperty('syslogRules')) {
                obj['syslogRules'] = ApiClient.convertToType(data['syslogRules'], [Object]);
            }
            if (data.hasOwnProperty('integrations')) {
                obj['integrations'] = ApiClient.convertToType(data['integrations'], [Object]);
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} enabled
 */
LogSettings.prototype['enabled'] = undefined;

/**
 * @member {String} retentionDays
 */
LogSettings.prototype['retentionDays'] = undefined;

/**
 * @member {Array.<Object>} syslogRules
 */
LogSettings.prototype['syslogRules'] = undefined;

/**
 * @member {Array.<Object>} integrations
 */
LogSettings.prototype['integrations'] = undefined;






export default LogSettings;

