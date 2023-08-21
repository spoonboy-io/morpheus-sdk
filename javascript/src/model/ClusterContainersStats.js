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
 * The ClusterContainersStats model module.
 * @module model/ClusterContainersStats
 * @version 6.2.1
 */
class ClusterContainersStats {
    /**
     * Constructs a new <code>ClusterContainersStats</code>.
     * @alias module:model/ClusterContainersStats
     */
    constructor() { 
        
        ClusterContainersStats.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterContainersStats</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterContainersStats} obj Optional instance to populate.
     * @return {module:model/ClusterContainersStats} The populated <code>ClusterContainersStats</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterContainersStats();

            if (data.hasOwnProperty('ts')) {
                obj['ts'] = ApiClient.convertToType(data['ts'], 'Date');
            }
            if (data.hasOwnProperty('running')) {
                obj['running'] = ApiClient.convertToType(data['running'], 'Boolean');
            }
            if (data.hasOwnProperty('userCpuUsage')) {
                obj['userCpuUsage'] = ApiClient.convertToType(data['userCpuUsage'], 'Number');
            }
            if (data.hasOwnProperty('systemCpuUsage')) {
                obj['systemCpuUsage'] = ApiClient.convertToType(data['systemCpuUsage'], 'Number');
            }
            if (data.hasOwnProperty('usedMemory')) {
                obj['usedMemory'] = ApiClient.convertToType(data['usedMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
            if (data.hasOwnProperty('cacheMemory')) {
                obj['cacheMemory'] = ApiClient.convertToType(data['cacheMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxStorage')) {
                obj['maxStorage'] = ApiClient.convertToType(data['maxStorage'], 'Number');
            }
            if (data.hasOwnProperty('usedStorage')) {
                obj['usedStorage'] = ApiClient.convertToType(data['usedStorage'], 'Number');
            }
            if (data.hasOwnProperty('readIOPS')) {
                obj['readIOPS'] = ApiClient.convertToType(data['readIOPS'], 'Number');
            }
            if (data.hasOwnProperty('writeIOPS')) {
                obj['writeIOPS'] = ApiClient.convertToType(data['writeIOPS'], 'Number');
            }
            if (data.hasOwnProperty('totalIOPS')) {
                obj['totalIOPS'] = ApiClient.convertToType(data['totalIOPS'], 'Number');
            }
            if (data.hasOwnProperty('netTxUsage')) {
                obj['netTxUsage'] = ApiClient.convertToType(data['netTxUsage'], 'Number');
            }
            if (data.hasOwnProperty('netRxUsage')) {
                obj['netRxUsage'] = ApiClient.convertToType(data['netRxUsage'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Date} ts
 */
ClusterContainersStats.prototype['ts'] = undefined;

/**
 * @member {Boolean} running
 */
ClusterContainersStats.prototype['running'] = undefined;

/**
 * @member {Number} userCpuUsage
 */
ClusterContainersStats.prototype['userCpuUsage'] = undefined;

/**
 * @member {Number} systemCpuUsage
 */
ClusterContainersStats.prototype['systemCpuUsage'] = undefined;

/**
 * @member {Number} usedMemory
 */
ClusterContainersStats.prototype['usedMemory'] = undefined;

/**
 * @member {Number} maxMemory
 */
ClusterContainersStats.prototype['maxMemory'] = undefined;

/**
 * @member {Number} cacheMemory
 */
ClusterContainersStats.prototype['cacheMemory'] = undefined;

/**
 * @member {Number} maxStorage
 */
ClusterContainersStats.prototype['maxStorage'] = undefined;

/**
 * @member {Number} usedStorage
 */
ClusterContainersStats.prototype['usedStorage'] = undefined;

/**
 * @member {Number} readIOPS
 */
ClusterContainersStats.prototype['readIOPS'] = undefined;

/**
 * @member {Number} writeIOPS
 */
ClusterContainersStats.prototype['writeIOPS'] = undefined;

/**
 * @member {Number} totalIOPS
 */
ClusterContainersStats.prototype['totalIOPS'] = undefined;

/**
 * @member {Number} netTxUsage
 */
ClusterContainersStats.prototype['netTxUsage'] = undefined;

/**
 * @member {Number} netRxUsage
 */
ClusterContainersStats.prototype['netRxUsage'] = undefined;






export default ClusterContainersStats;
