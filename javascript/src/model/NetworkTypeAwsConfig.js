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
import NetworkTypeAwsConfigZonePool from './NetworkTypeAwsConfigZonePool';

/**
 * The NetworkTypeAwsConfig model module.
 * @module model/NetworkTypeAwsConfig
 * @version 6.2.1
 */
class NetworkTypeAwsConfig {
    /**
     * Constructs a new <code>NetworkTypeAwsConfig</code>.
     * @alias module:model/NetworkTypeAwsConfig
     * @param availabilityZone {String} Availability Zone Name
     * @param cidr {String} Network CIDR
     * @param assignPublicIp {Boolean} Assign public IPs by default.
     * @param zonePool {module:model/NetworkTypeAwsConfigZonePool} 
     */
    constructor(availabilityZone, cidr, assignPublicIp, zonePool) { 
        
        NetworkTypeAwsConfig.initialize(this, availabilityZone, cidr, assignPublicIp, zonePool);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, availabilityZone, cidr, assignPublicIp, zonePool) { 
        obj['availabilityZone'] = availabilityZone;
        obj['cidr'] = cidr;
        obj['assignPublicIp'] = assignPublicIp;
        obj['zonePool'] = zonePool;
    }

    /**
     * Constructs a <code>NetworkTypeAwsConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkTypeAwsConfig} obj Optional instance to populate.
     * @return {module:model/NetworkTypeAwsConfig} The populated <code>NetworkTypeAwsConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkTypeAwsConfig();

            if (data.hasOwnProperty('availabilityZone')) {
                obj['availabilityZone'] = ApiClient.convertToType(data['availabilityZone'], 'String');
            }
            if (data.hasOwnProperty('cidr')) {
                obj['cidr'] = ApiClient.convertToType(data['cidr'], 'String');
            }
            if (data.hasOwnProperty('assignPublicIp')) {
                obj['assignPublicIp'] = ApiClient.convertToType(data['assignPublicIp'], 'Boolean');
            }
            if (data.hasOwnProperty('zonePool')) {
                obj['zonePool'] = NetworkTypeAwsConfigZonePool.constructFromObject(data['zonePool']);
            }
        }
        return obj;
    }


}

/**
 * Availability Zone Name
 * @member {String} availabilityZone
 */
NetworkTypeAwsConfig.prototype['availabilityZone'] = undefined;

/**
 * Network CIDR
 * @member {String} cidr
 */
NetworkTypeAwsConfig.prototype['cidr'] = undefined;

/**
 * Assign public IPs by default.
 * @member {Boolean} assignPublicIp
 */
NetworkTypeAwsConfig.prototype['assignPublicIp'] = undefined;

/**
 * @member {module:model/NetworkTypeAwsConfigZonePool} zonePool
 */
NetworkTypeAwsConfig.prototype['zonePool'] = undefined;






export default NetworkTypeAwsConfig;

