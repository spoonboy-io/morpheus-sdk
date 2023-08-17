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
import BackupSettingsUpdateDefaultSchedule from './BackupSettingsUpdateDefaultSchedule';
import BackupSettingsUpdateDefaultStorageBucket from './BackupSettingsUpdateDefaultStorageBucket';

/**
 * The BackupSettingsUpdate model module.
 * @module model/BackupSettingsUpdate
 * @version 6.1.1
 */
class BackupSettingsUpdate {
    /**
     * Constructs a new <code>BackupSettingsUpdate</code>.
     * @alias module:model/BackupSettingsUpdate
     */
    constructor() { 
        
        BackupSettingsUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BackupSettingsUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BackupSettingsUpdate} obj Optional instance to populate.
     * @return {module:model/BackupSettingsUpdate} The populated <code>BackupSettingsUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BackupSettingsUpdate();

            if (data.hasOwnProperty('backupsEnabled')) {
                obj['backupsEnabled'] = ApiClient.convertToType(data['backupsEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('retentionCount')) {
                obj['retentionCount'] = ApiClient.convertToType(data['retentionCount'], 'Number');
            }
            if (data.hasOwnProperty('createBackups')) {
                obj['createBackups'] = ApiClient.convertToType(data['createBackups'], 'Boolean');
            }
            if (data.hasOwnProperty('backupAppliance')) {
                obj['backupAppliance'] = ApiClient.convertToType(data['backupAppliance'], 'Boolean');
            }
            if (data.hasOwnProperty('updateExisting')) {
                obj['updateExisting'] = ApiClient.convertToType(data['updateExisting'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultSchedule')) {
                obj['defaultSchedule'] = BackupSettingsUpdateDefaultSchedule.constructFromObject(data['defaultSchedule']);
            }
            if (data.hasOwnProperty('clearDefaultSchedule')) {
                obj['clearDefaultSchedule'] = ApiClient.convertToType(data['clearDefaultSchedule'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultStorageBucket')) {
                obj['defaultStorageBucket'] = BackupSettingsUpdateDefaultStorageBucket.constructFromObject(data['defaultStorageBucket']);
            }
            if (data.hasOwnProperty('clearDefaultStorageBucket')) {
                obj['clearDefaultStorageBucket'] = ApiClient.convertToType(data['clearDefaultStorageBucket'], 'Boolean');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BackupSettingsUpdate</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BackupSettingsUpdate</code>.
     */
    static validateJSON(data) {
        // validate the optional field `defaultSchedule`
        if (data['defaultSchedule']) { // data not null
          BackupSettingsUpdateDefaultSchedule.validateJSON(data['defaultSchedule']);
        }
        // validate the optional field `defaultStorageBucket`
        if (data['defaultStorageBucket']) { // data not null
          BackupSettingsUpdateDefaultStorageBucket.validateJSON(data['defaultStorageBucket']);
        }

        return true;
    }


}



/**
 * Use this to enable / disable scheduled backups
 * @member {Boolean} backupsEnabled
 */
BackupSettingsUpdate.prototype['backupsEnabled'] = undefined;

/**
 * Maximum number of successful backups to retain
 * @member {Number} retentionCount
 */
BackupSettingsUpdate.prototype['retentionCount'] = undefined;

/**
 * Use this to enable / disable create backups
 * @member {Boolean} createBackups
 */
BackupSettingsUpdate.prototype['createBackups'] = undefined;

/**
 * When enabled, a Backup will be created to backup the Morpheus appliance database
 * @member {Boolean} backupAppliance
 */
BackupSettingsUpdate.prototype['backupAppliance'] = undefined;

/**
 * Use this to update existing backups with new settings
 * @member {Boolean} updateExisting
 */
BackupSettingsUpdate.prototype['updateExisting'] = undefined;

/**
 * @member {module:model/BackupSettingsUpdateDefaultSchedule} defaultSchedule
 */
BackupSettingsUpdate.prototype['defaultSchedule'] = undefined;

/**
 * Use this to clear existing default backup schedule
 * @member {Boolean} clearDefaultSchedule
 */
BackupSettingsUpdate.prototype['clearDefaultSchedule'] = undefined;

/**
 * @member {module:model/BackupSettingsUpdateDefaultStorageBucket} defaultStorageBucket
 */
BackupSettingsUpdate.prototype['defaultStorageBucket'] = undefined;

/**
 * Use this to clear default store bucket
 * @member {Boolean} clearDefaultStorageBucket
 */
BackupSettingsUpdate.prototype['clearDefaultStorageBucket'] = undefined;






export default BackupSettingsUpdate;
