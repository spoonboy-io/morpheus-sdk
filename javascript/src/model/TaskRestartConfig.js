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
import TaskRestartConfigTaskOptions from './TaskRestartConfigTaskOptions';
import TaskRestartConfigTaskType from './TaskRestartConfigTaskType';

/**
 * The TaskRestartConfig model module.
 * @module model/TaskRestartConfig
 * @version 6.2.1
 */
class TaskRestartConfig {
    /**
     * Constructs a new <code>TaskRestartConfig</code>.
     * @alias module:model/TaskRestartConfig
     */
    constructor() { 
        
        TaskRestartConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskRestartConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskRestartConfig} obj Optional instance to populate.
     * @return {module:model/TaskRestartConfig} The populated <code>TaskRestartConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskRestartConfig();

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
                obj['taskType'] = TaskRestartConfigTaskType.constructFromObject(data['taskType']);
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('taskOptions')) {
                obj['taskOptions'] = TaskRestartConfigTaskOptions.constructFromObject(data['taskOptions']);
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
TaskRestartConfig.prototype['id'] = undefined;

/**
 * @member {Number} accountId
 */
TaskRestartConfig.prototype['accountId'] = undefined;

/**
 * @member {String} name
 */
TaskRestartConfig.prototype['name'] = undefined;

/**
 * @member {String} code
 */
TaskRestartConfig.prototype['code'] = undefined;

/**
 * @member {module:model/TaskRestartConfigTaskType} taskType
 */
TaskRestartConfig.prototype['taskType'] = undefined;

/**
 * @member {Array.<String>} labels
 */
TaskRestartConfig.prototype['labels'] = undefined;

/**
 * @member {String} visibility
 */
TaskRestartConfig.prototype['visibility'] = undefined;

/**
 * @member {module:model/TaskRestartConfigTaskOptions} taskOptions
 */
TaskRestartConfig.prototype['taskOptions'] = undefined;

/**
 * @member {module:model/TaskAnsiblePlaybookConfigFile} file
 */
TaskRestartConfig.prototype['file'] = undefined;

/**
 * @member {String} resultType
 */
TaskRestartConfig.prototype['resultType'] = undefined;

/**
 * @member {String} executeTarget
 */
TaskRestartConfig.prototype['executeTarget'] = undefined;

/**
 * @member {Boolean} retryable
 */
TaskRestartConfig.prototype['retryable'] = undefined;

/**
 * @member {Number} retryCount
 */
TaskRestartConfig.prototype['retryCount'] = undefined;

/**
 * @member {Number} retryDelaySeconds
 */
TaskRestartConfig.prototype['retryDelaySeconds'] = undefined;

/**
 * @member {Boolean} allowCustomConfig
 */
TaskRestartConfig.prototype['allowCustomConfig'] = undefined;

/**
 * @member {module:model/OptionTypeListCredential} credential
 */
TaskRestartConfig.prototype['credential'] = undefined;

/**
 * @member {Date} dateCreated
 */
TaskRestartConfig.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
TaskRestartConfig.prototype['lastUpdated'] = undefined;






export default TaskRestartConfig;
