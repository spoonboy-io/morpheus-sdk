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
import BackupSettings from './BackupSettings';

/**
 * The ListBackupSettings200Response model module.
 * @module model/ListBackupSettings200Response
 * @version 6.1.1
 */
class ListBackupSettings200Response {
    /**
     * Constructs a new <code>ListBackupSettings200Response</code>.
     * @alias module:model/ListBackupSettings200Response
     */
    constructor() { 
        
        ListBackupSettings200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ListBackupSettings200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ListBackupSettings200Response} obj Optional instance to populate.
     * @return {module:model/ListBackupSettings200Response} The populated <code>ListBackupSettings200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ListBackupSettings200Response();

            if (data.hasOwnProperty('backupSettings')) {
                obj['backupSettings'] = BackupSettings.constructFromObject(data['backupSettings']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ListBackupSettings200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ListBackupSettings200Response</code>.
     */
    static validateJSON(data) {
        // validate the optional field `backupSettings`
        if (data['backupSettings']) { // data not null
          BackupSettings.validateJSON(data['backupSettings']);
        }

        return true;
    }


}



/**
 * @member {module:model/BackupSettings} backupSettings
 */
ListBackupSettings200Response.prototype['backupSettings'] = undefined;






export default ListBackupSettings200Response;

