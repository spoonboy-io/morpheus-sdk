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
import ApiTaskSetsTaskSetTasks from './ApiTaskSetsTaskSetTasks';

/**
 * The ApiTaskSetsTaskSet model module.
 * @module model/ApiTaskSetsTaskSet
 * @version 6.2.1
 */
class ApiTaskSetsTaskSet {
    /**
     * Constructs a new <code>ApiTaskSetsTaskSet</code>.
     * @alias module:model/ApiTaskSetsTaskSet
     * @param name {String} A unique name for the workflow
     */
    constructor(name) { 
        
        ApiTaskSetsTaskSet.initialize(this, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name) { 
        obj['name'] = name;
    }

    /**
     * Constructs a <code>ApiTaskSetsTaskSet</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiTaskSetsTaskSet} obj Optional instance to populate.
     * @return {module:model/ApiTaskSetsTaskSet} The populated <code>ApiTaskSetsTaskSet</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiTaskSetsTaskSet();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('optionTypes')) {
                obj['optionTypes'] = ApiClient.convertToType(data['optionTypes'], ['Number']);
            }
            if (data.hasOwnProperty('tasks')) {
                obj['tasks'] = ApiTaskSetsTaskSetTasks.constructFromObject(data['tasks']);
            }
        }
        return obj;
    }


}

/**
 * A unique name for the workflow
 * @member {String} name
 */
ApiTaskSetsTaskSet.prototype['name'] = undefined;

/**
 * A description of the workflow
 * @member {String} description
 */
ApiTaskSetsTaskSet.prototype['description'] = undefined;

/**
 * An array of Category labels for filtering
 * @member {Array.<String>} labels
 */
ApiTaskSetsTaskSet.prototype['labels'] = undefined;

/**
 * Workflow type
 * @member {module:model/ApiTaskSetsTaskSet.TypeEnum} type
 * @default 'provision'
 */
ApiTaskSetsTaskSet.prototype['type'] = 'provision';

/**
 * private or public
 * @member {module:model/ApiTaskSetsTaskSet.VisibilityEnum} visibility
 * @default 'private'
 */
ApiTaskSetsTaskSet.prototype['visibility'] = 'private';

/**
 * List of option type IDs for use with operational workflow configuration.
 * @member {Array.<Number>} optionTypes
 */
ApiTaskSetsTaskSet.prototype['optionTypes'] = undefined;

/**
 * @member {module:model/ApiTaskSetsTaskSetTasks} tasks
 */
ApiTaskSetsTaskSet.prototype['tasks'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
ApiTaskSetsTaskSet['TypeEnum'] = {

    /**
     * value: "provision"
     * @const
     */
    "provision": "provision",

    /**
     * value: "operation"
     * @const
     */
    "operation": "operation"
};


/**
 * Allowed values for the <code>visibility</code> property.
 * @enum {String}
 * @readonly
 */
ApiTaskSetsTaskSet['VisibilityEnum'] = {

    /**
     * value: "private"
     * @const
     */
    "private": "private",

    /**
     * value: "public"
     * @const
     */
    "public": "public"
};



export default ApiTaskSetsTaskSet;
