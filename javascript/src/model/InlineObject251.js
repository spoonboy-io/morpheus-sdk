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
import ApiTasksIdExecuteJob from './ApiTasksIdExecuteJob';

/**
 * The InlineObject251 model module.
 * @module model/InlineObject251
 * @version 6.2.1
 */
class InlineObject251 {
    /**
     * Constructs a new <code>InlineObject251</code>.
     * @alias module:model/InlineObject251
     * @param job {module:model/ApiTasksIdExecuteJob} 
     */
    constructor(job) { 
        
        InlineObject251.initialize(this, job);
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
     * Constructs a <code>InlineObject251</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject251} obj Optional instance to populate.
     * @return {module:model/InlineObject251} The populated <code>InlineObject251</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject251();

            if (data.hasOwnProperty('job')) {
                obj['job'] = ApiTasksIdExecuteJob.constructFromObject(data['job']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiTasksIdExecuteJob} job
 */
InlineObject251.prototype['job'] = undefined;






export default InlineObject251;

