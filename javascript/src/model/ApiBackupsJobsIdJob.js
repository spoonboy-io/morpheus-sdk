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
 * The ApiBackupsJobsIdJob model module.
 * @module model/ApiBackupsJobsIdJob
 * @version 6.2.1
 */
class ApiBackupsJobsIdJob {
    /**
     * Constructs a new <code>ApiBackupsJobsIdJob</code>.
     * @alias module:model/ApiBackupsJobsIdJob
     */
    constructor() { 
        
        ApiBackupsJobsIdJob.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiBackupsJobsIdJob</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiBackupsJobsIdJob} obj Optional instance to populate.
     * @return {module:model/ApiBackupsJobsIdJob} The populated <code>ApiBackupsJobsIdJob</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiBackupsJobsIdJob();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('scheduleId')) {
                obj['scheduleId'] = ApiClient.convertToType(data['scheduleId'], 'Number');
            }
            if (data.hasOwnProperty('retentionCount')) {
                obj['retentionCount'] = ApiClient.convertToType(data['retentionCount'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * A name for the backup job
 * @member {String} name
 */
ApiBackupsJobsIdJob.prototype['name'] = undefined;

/**
 * A code for the backup job
 * @member {String} code
 */
ApiBackupsJobsIdJob.prototype['code'] = undefined;

/**
 * Execute Schedule ID to use for the backup job
 * @member {Number} scheduleId
 */
ApiBackupsJobsIdJob.prototype['scheduleId'] = undefined;

/**
 * Retention Count
 * @member {Number} retentionCount
 */
ApiBackupsJobsIdJob.prototype['retentionCount'] = undefined;






export default ApiBackupsJobsIdJob;
