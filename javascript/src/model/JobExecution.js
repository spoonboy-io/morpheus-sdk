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
import JobExecutionJob from './JobExecutionJob';

/**
 * The JobExecution model module.
 * @module model/JobExecution
 * @version 6.2.1
 */
class JobExecution {
    /**
     * Constructs a new <code>JobExecution</code>.
     * @alias module:model/JobExecution
     */
    constructor() { 
        
        JobExecution.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>JobExecution</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/JobExecution} obj Optional instance to populate.
     * @return {module:model/JobExecution} The populated <code>JobExecution</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobExecution();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('process')) {
                obj['process'] = ApiClient.convertToType(data['process'], 'String');
            }
            if (data.hasOwnProperty('job')) {
                obj['job'] = JobExecutionJob.constructFromObject(data['job']);
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('duration')) {
                obj['duration'] = ApiClient.convertToType(data['duration'], 'Number');
            }
            if (data.hasOwnProperty('resultData')) {
                obj['resultData'] = ApiClient.convertToType(data['resultData'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('statusMessage')) {
                obj['statusMessage'] = ApiClient.convertToType(data['statusMessage'], 'String');
            }
            if (data.hasOwnProperty('createdBy')) {
                obj['createdBy'] = ApiClient.convertToType(data['createdBy'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
JobExecution.prototype['id'] = undefined;

/**
 * @member {String} name
 */
JobExecution.prototype['name'] = undefined;

/**
 * @member {String} process
 */
JobExecution.prototype['process'] = undefined;

/**
 * @member {module:model/JobExecutionJob} job
 */
JobExecution.prototype['job'] = undefined;

/**
 * @member {String} description
 */
JobExecution.prototype['description'] = undefined;

/**
 * @member {Date} dateCreated
 */
JobExecution.prototype['dateCreated'] = undefined;

/**
 * @member {Date} startDate
 */
JobExecution.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
JobExecution.prototype['endDate'] = undefined;

/**
 * @member {Number} duration
 */
JobExecution.prototype['duration'] = undefined;

/**
 * @member {String} resultData
 */
JobExecution.prototype['resultData'] = undefined;

/**
 * @member {String} status
 */
JobExecution.prototype['status'] = undefined;

/**
 * @member {String} statusMessage
 */
JobExecution.prototype['statusMessage'] = undefined;

/**
 * @member {String} createdBy
 */
JobExecution.prototype['createdBy'] = undefined;






export default JobExecution;

