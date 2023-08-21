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
 * The TaskJrubyConfigTaskType model module.
 * @module model/TaskJrubyConfigTaskType
 * @version 6.2.1
 */
class TaskJrubyConfigTaskType {
    /**
     * Constructs a new <code>TaskJrubyConfigTaskType</code>.
     * @alias module:model/TaskJrubyConfigTaskType
     */
    constructor() { 
        
        TaskJrubyConfigTaskType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskJrubyConfigTaskType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskJrubyConfigTaskType} obj Optional instance to populate.
     * @return {module:model/TaskJrubyConfigTaskType} The populated <code>TaskJrubyConfigTaskType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskJrubyConfigTaskType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
TaskJrubyConfigTaskType.prototype['id'] = undefined;

/**
 * @member {module:model/TaskJrubyConfigTaskType.CodeEnum} code
 */
TaskJrubyConfigTaskType.prototype['code'] = undefined;

/**
 * @member {String} name
 */
TaskJrubyConfigTaskType.prototype['name'] = undefined;





/**
 * Allowed values for the <code>code</code> property.
 * @enum {String}
 * @readonly
 */
TaskJrubyConfigTaskType['CodeEnum'] = {

    /**
     * value: "jrubyTask"
     * @const
     */
    "jrubyTask": "jrubyTask"
};



export default TaskJrubyConfigTaskType;

