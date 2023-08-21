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
import BackupBackupProvider from './BackupBackupProvider';
import BackupBackupRespository from './BackupBackupRespository';
import BackupBackupType from './BackupBackupType';
import BackupInstance from './BackupInstance';
import BackupJob from './BackupJob';
import BackupLastResult from './BackupLastResult';
import BackupSchedule from './BackupSchedule';
import BackupStats from './BackupStats';
import BackupStorageProvider from './BackupStorageProvider';

/**
 * The Backup model module.
 * @module model/Backup
 * @version 6.2.1
 */
class Backup {
    /**
     * Constructs a new <code>Backup</code>.
     * @alias module:model/Backup
     */
    constructor() { 
        
        Backup.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Backup</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Backup} obj Optional instance to populate.
     * @return {module:model/Backup} The populated <code>Backup</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Backup();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('locationType')) {
                obj['locationType'] = ApiClient.convertToType(data['locationType'], 'String');
            }
            if (data.hasOwnProperty('instance')) {
                obj['instance'] = BackupInstance.constructFromObject(data['instance']);
            }
            if (data.hasOwnProperty('containerId')) {
                obj['containerId'] = ApiClient.convertToType(data['containerId'], 'Number');
            }
            if (data.hasOwnProperty('job')) {
                obj['job'] = BackupJob.constructFromObject(data['job']);
            }
            if (data.hasOwnProperty('schedule')) {
                obj['schedule'] = BackupSchedule.constructFromObject(data['schedule']);
            }
            if (data.hasOwnProperty('retentionCount')) {
                obj['retentionCount'] = ApiClient.convertToType(data['retentionCount'], 'Number');
            }
            if (data.hasOwnProperty('backupType')) {
                obj['backupType'] = BackupBackupType.constructFromObject(data['backupType']);
            }
            if (data.hasOwnProperty('storageProvider')) {
                obj['storageProvider'] = BackupStorageProvider.constructFromObject(data['storageProvider']);
            }
            if (data.hasOwnProperty('backupProvider')) {
                obj['backupProvider'] = BackupBackupProvider.constructFromObject(data['backupProvider']);
            }
            if (data.hasOwnProperty('backupRespository')) {
                obj['backupRespository'] = BackupBackupRespository.constructFromObject(data['backupRespository']);
            }
            if (data.hasOwnProperty('cronExpression')) {
                obj['cronExpression'] = ApiClient.convertToType(data['cronExpression'], 'String');
            }
            if (data.hasOwnProperty('nextFire')) {
                obj['nextFire'] = ApiClient.convertToType(data['nextFire'], 'Date');
            }
            if (data.hasOwnProperty('lastStatus')) {
                obj['lastStatus'] = ApiClient.convertToType(data['lastStatus'], 'String');
            }
            if (data.hasOwnProperty('lastResult')) {
                obj['lastResult'] = BackupLastResult.constructFromObject(data['lastResult']);
            }
            if (data.hasOwnProperty('stats')) {
                obj['stats'] = BackupStats.constructFromObject(data['stats']);
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * Backup ID
 * @member {Number} id
 */
Backup.prototype['id'] = undefined;

/**
 * Name
 * @member {String} name
 */
Backup.prototype['name'] = undefined;

/**
 * Source Type (instance, server, storage)
 * @member {String} locationType
 */
Backup.prototype['locationType'] = undefined;

/**
 * @member {module:model/BackupInstance} instance
 */
Backup.prototype['instance'] = undefined;

/**
 * @member {Number} containerId
 */
Backup.prototype['containerId'] = undefined;

/**
 * @member {module:model/BackupJob} job
 */
Backup.prototype['job'] = undefined;

/**
 * @member {module:model/BackupSchedule} schedule
 */
Backup.prototype['schedule'] = undefined;

/**
 * @member {Number} retentionCount
 */
Backup.prototype['retentionCount'] = undefined;

/**
 * @member {module:model/BackupBackupType} backupType
 */
Backup.prototype['backupType'] = undefined;

/**
 * @member {module:model/BackupStorageProvider} storageProvider
 */
Backup.prototype['storageProvider'] = undefined;

/**
 * @member {module:model/BackupBackupProvider} backupProvider
 */
Backup.prototype['backupProvider'] = undefined;

/**
 * @member {module:model/BackupBackupRespository} backupRespository
 */
Backup.prototype['backupRespository'] = undefined;

/**
 * Cron Expression
 * @member {String} cronExpression
 */
Backup.prototype['cronExpression'] = undefined;

/**
 * Next Fire
 * @member {Date} nextFire
 */
Backup.prototype['nextFire'] = undefined;

/**
 * Last Status
 * @member {String} lastStatus
 */
Backup.prototype['lastStatus'] = undefined;

/**
 * @member {module:model/BackupLastResult} lastResult
 */
Backup.prototype['lastResult'] = undefined;

/**
 * @member {module:model/BackupStats} stats
 */
Backup.prototype['stats'] = undefined;

/**
 * Enabled
 * @member {Boolean} enabled
 */
Backup.prototype['enabled'] = undefined;

/**
 * Date Created
 * @member {Date} dateCreated
 */
Backup.prototype['dateCreated'] = undefined;

/**
 * Last Updated
 * @member {Date} lastUpdated
 */
Backup.prototype['lastUpdated'] = undefined;






export default Backup;

