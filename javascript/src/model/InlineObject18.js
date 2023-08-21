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
import ApiBackupsJobsJob from './ApiBackupsJobsJob';

/**
 * The InlineObject18 model module.
 * @module model/InlineObject18
 * @version 6.2.1
 */
class InlineObject18 {
    /**
     * Constructs a new <code>InlineObject18</code>.
     * @alias module:model/InlineObject18
     * @param job {module:model/ApiBackupsJobsJob} 
     */
    constructor(job) { 
        
        InlineObject18.initialize(this, job);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, job) { 
        obj['job'] = job;
    }

    /**
     * Constructs a <code>InlineObject18</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject18} obj Optional instance to populate.
     * @return {module:model/InlineObject18} The populated <code>InlineObject18</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject18();

            if (data.hasOwnProperty('job')) {
                obj['job'] = ApiBackupsJobsJob.constructFromObject(data['job']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiBackupsJobsJob} job
 */
InlineObject18.prototype['job'] = undefined;






export default InlineObject18;

