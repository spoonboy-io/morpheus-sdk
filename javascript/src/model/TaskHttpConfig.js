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
import OptionTypeListCredential from './OptionTypeListCredential';
import TaskAnsiblePlaybookConfigFile from './TaskAnsiblePlaybookConfigFile';
import TaskHttpConfigTaskOptions from './TaskHttpConfigTaskOptions';
import TaskHttpConfigTaskType from './TaskHttpConfigTaskType';

/**
 * The TaskHttpConfig model module.
 * @module model/TaskHttpConfig
 * @version 6.2.1
 */
class TaskHttpConfig {
    /**
     * Constructs a new <code>TaskHttpConfig</code>.
     * @alias module:model/TaskHttpConfig
     */
    constructor() { 
        
        TaskHttpConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskHttpConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskHttpConfig} obj Optional instance to populate.
     * @return {module:model/TaskHttpConfig} The populated <code>TaskHttpConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskHttpConfig();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('accountId')) {
                obj['accountId'] = ApiClient.convertToType(data['accountId'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('taskType')) {
                obj['taskType'] = TaskHttpConfigTaskType.constructFromObject(data['taskType']);
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('taskOptions')) {
                obj['taskOptions'] = TaskHttpConfigTaskOptions.constructFromObject(data['taskOptions']);
            }
            if (data.hasOwnProperty('file')) {
                obj['file'] = TaskAnsiblePlaybookConfigFile.constructFromObject(data['file']);
            }
            if (data.hasOwnProperty('resultType')) {
                obj['resultType'] = ApiClient.convertToType(data['resultType'], 'String');
            }
            if (data.hasOwnProperty('executeTarget')) {
                obj['executeTarget'] = ApiClient.convertToType(data['executeTarget'], 'String');
            }
            if (data.hasOwnProperty('retryable')) {
                obj['retryable'] = ApiClient.convertToType(data['retryable'], 'Boolean');
            }
            if (data.hasOwnProperty('retryCount')) {
                obj['retryCount'] = ApiClient.convertToType(data['retryCount'], 'Number');
            }
            if (data.hasOwnProperty('retryDelaySeconds')) {
                obj['retryDelaySeconds'] = ApiClient.convertToType(data['retryDelaySeconds'], 'Number');
            }
            if (data.hasOwnProperty('allowCustomConfig')) {
                obj['allowCustomConfig'] = ApiClient.convertToType(data['allowCustomConfig'], 'Boolean');
            }
            if (data.hasOwnProperty('credential')) {
                obj['credential'] = OptionTypeListCredential.constructFromObject(data['credential']);
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
TaskHttpConfig.prototype['id'] = undefined;

/**
 * @member {Number} accountId
 */
TaskHttpConfig.prototype['accountId'] = undefined;

/**
 * @member {String} name
 */
TaskHttpConfig.prototype['name'] = undefined;

/**
 * @member {String} code
 */
TaskHttpConfig.prototype['code'] = undefined;

/**
 * @member {module:model/TaskHttpConfigTaskType} taskType
 */
TaskHttpConfig.prototype['taskType'] = undefined;

/**
 * @member {Array.<String>} labels
 */
TaskHttpConfig.prototype['labels'] = undefined;

/**
 * @member {String} visibility
 */
TaskHttpConfig.prototype['visibility'] = undefined;

/**
 * @member {module:model/TaskHttpConfigTaskOptions} taskOptions
 */
TaskHttpConfig.prototype['taskOptions'] = undefined;

/**
 * @member {module:model/TaskAnsiblePlaybookConfigFile} file
 */
TaskHttpConfig.prototype['file'] = undefined;

/**
 * @member {String} resultType
 */
TaskHttpConfig.prototype['resultType'] = undefined;

/**
 * @member {String} executeTarget
 */
TaskHttpConfig.prototype['executeTarget'] = undefined;

/**
 * @member {Boolean} retryable
 */
TaskHttpConfig.prototype['retryable'] = undefined;

/**
 * @member {Number} retryCount
 */
TaskHttpConfig.prototype['retryCount'] = undefined;

/**
 * @member {Number} retryDelaySeconds
 */
TaskHttpConfig.prototype['retryDelaySeconds'] = undefined;

/**
 * @member {Boolean} allowCustomConfig
 */
TaskHttpConfig.prototype['allowCustomConfig'] = undefined;

/**
 * @member {module:model/OptionTypeListCredential} credential
 */
TaskHttpConfig.prototype['credential'] = undefined;

/**
 * @member {Date} dateCreated
 */
TaskHttpConfig.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
TaskHttpConfig.prototype['lastUpdated'] = undefined;






export default TaskHttpConfig;
