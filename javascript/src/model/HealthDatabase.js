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
import HealthDatabaseInnodbStats from './HealthDatabaseInnodbStats';
import HealthDatabaseScans from './HealthDatabaseScans';
import HealthDatabaseSlowQueries from './HealthDatabaseSlowQueries';
import HealthDatabaseStats from './HealthDatabaseStats';

/**
 * The HealthDatabase model module.
 * @module model/HealthDatabase
 * @version 6.2.1
 */
class HealthDatabase {
    /**
     * Constructs a new <code>HealthDatabase</code>.
     * @alias module:model/HealthDatabase
     */
    constructor() { 
        
        HealthDatabase.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>HealthDatabase</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/HealthDatabase} obj Optional instance to populate.
     * @return {module:model/HealthDatabase} The populated <code>HealthDatabase</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new HealthDatabase();

            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
            if (data.hasOwnProperty('connectionList')) {
                obj['connectionList'] = ApiClient.convertToType(data['connectionList'], [Object]);
            }
            if (data.hasOwnProperty('busyConnections')) {
                obj['busyConnections'] = ApiClient.convertToType(data['busyConnections'], ['String']);
            }
            if (data.hasOwnProperty('maxConnections')) {
                obj['maxConnections'] = ApiClient.convertToType(data['maxConnections'], 'Number');
            }
            if (data.hasOwnProperty('maxUsedConnections')) {
                obj['maxUsedConnections'] = ApiClient.convertToType(data['maxUsedConnections'], 'Number');
            }
            if (data.hasOwnProperty('usedConnections')) {
                obj['usedConnections'] = ApiClient.convertToType(data['usedConnections'], 'Number');
            }
            if (data.hasOwnProperty('abortedConnections')) {
                obj['abortedConnections'] = ApiClient.convertToType(data['abortedConnections'], 'Number');
            }
            if (data.hasOwnProperty('innodbStatus')) {
                obj['innodbStatus'] = ApiClient.convertToType(data['innodbStatus'], 'String');
            }
            if (data.hasOwnProperty('stats')) {
                obj['stats'] = HealthDatabaseStats.constructFromObject(data['stats']);
            }
            if (data.hasOwnProperty('scans')) {
                obj['scans'] = HealthDatabaseScans.constructFromObject(data['scans']);
            }
            if (data.hasOwnProperty('slowQueries')) {
                obj['slowQueries'] = ApiClient.convertToType(data['slowQueries'], [HealthDatabaseSlowQueries]);
            }
            if (data.hasOwnProperty('innodbStats')) {
                obj['innodbStats'] = HealthDatabaseInnodbStats.constructFromObject(data['innodbStats']);
            }
            if (data.hasOwnProperty('scanPercent')) {
                obj['scanPercent'] = ApiClient.convertToType(data['scanPercent'], 'Number');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} success
 */
HealthDatabase.prototype['success'] = undefined;

/**
 * @member {Array.<Object>} connectionList
 */
HealthDatabase.prototype['connectionList'] = undefined;

/**
 * @member {Array.<String>} busyConnections
 */
HealthDatabase.prototype['busyConnections'] = undefined;

/**
 * @member {Number} maxConnections
 */
HealthDatabase.prototype['maxConnections'] = undefined;

/**
 * @member {Number} maxUsedConnections
 */
HealthDatabase.prototype['maxUsedConnections'] = undefined;

/**
 * @member {Number} usedConnections
 */
HealthDatabase.prototype['usedConnections'] = undefined;

/**
 * @member {Number} abortedConnections
 */
HealthDatabase.prototype['abortedConnections'] = undefined;

/**
 * @member {String} innodbStatus
 */
HealthDatabase.prototype['innodbStatus'] = undefined;

/**
 * @member {module:model/HealthDatabaseStats} stats
 */
HealthDatabase.prototype['stats'] = undefined;

/**
 * @member {module:model/HealthDatabaseScans} scans
 */
HealthDatabase.prototype['scans'] = undefined;

/**
 * @member {Array.<module:model/HealthDatabaseSlowQueries>} slowQueries
 */
HealthDatabase.prototype['slowQueries'] = undefined;

/**
 * @member {module:model/HealthDatabaseInnodbStats} innodbStats
 */
HealthDatabase.prototype['innodbStats'] = undefined;

/**
 * @member {Number} scanPercent
 */
HealthDatabase.prototype['scanPercent'] = undefined;

/**
 * @member {String} status
 */
HealthDatabase.prototype['status'] = undefined;






export default HealthDatabase;
