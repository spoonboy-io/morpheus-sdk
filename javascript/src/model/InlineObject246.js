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
import ApiTasksTask from './ApiTasksTask';

/**
 * The InlineObject246 model module.
 * @module model/InlineObject246
 * @version 6.2.1
 */
class InlineObject246 {
    /**
     * Constructs a new <code>InlineObject246</code>.
     * @alias module:model/InlineObject246
     * @param task {module:model/ApiTasksTask} 
     */
    constructor(task) { 
        
        InlineObject246.initialize(this, task);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, task) { 
        obj['task'] = task;
    }

    /**
     * Constructs a <code>InlineObject246</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject246} obj Optional instance to populate.
     * @return {module:model/InlineObject246} The populated <code>InlineObject246</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject246();

            if (data.hasOwnProperty('task')) {
                obj['task'] = ApiTasksTask.constructFromObject(data['task']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiTasksTask} task
 */
InlineObject246.prototype['task'] = undefined;






export default InlineObject246;

