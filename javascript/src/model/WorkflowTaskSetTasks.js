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
import WorkflowTask from './WorkflowTask';

/**
 * The WorkflowTaskSetTasks model module.
 * @module model/WorkflowTaskSetTasks
 * @version 6.2.1
 */
class WorkflowTaskSetTasks {
    /**
     * Constructs a new <code>WorkflowTaskSetTasks</code>.
     * @alias module:model/WorkflowTaskSetTasks
     */
    constructor() { 
        
        WorkflowTaskSetTasks.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>WorkflowTaskSetTasks</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/WorkflowTaskSetTasks} obj Optional instance to populate.
     * @return {module:model/WorkflowTaskSetTasks} The populated <code>WorkflowTaskSetTasks</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new WorkflowTaskSetTasks();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('taskPhase')) {
                obj['taskPhase'] = ApiClient.convertToType(data['taskPhase'], 'String');
            }
            if (data.hasOwnProperty('taskOrder')) {
                obj['taskOrder'] = ApiClient.convertToType(data['taskOrder'], 'Number');
            }
            if (data.hasOwnProperty('task')) {
                obj['task'] = WorkflowTask.constructFromObject(data['task']);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
WorkflowTaskSetTasks.prototype['id'] = undefined;

/**
 * @member {String} taskPhase
 */
WorkflowTaskSetTasks.prototype['taskPhase'] = undefined;

/**
 * @member {Number} taskOrder
 */
WorkflowTaskSetTasks.prototype['taskOrder'] = undefined;

/**
 * @member {module:model/WorkflowTask} task
 */
WorkflowTaskSetTasks.prototype['task'] = undefined;






export default WorkflowTaskSetTasks;

