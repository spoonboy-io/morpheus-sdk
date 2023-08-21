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
 * The TaskAnsiblePlaybookConfigTaskOptions model module.
 * @module model/TaskAnsiblePlaybookConfigTaskOptions
 * @version 6.2.1
 */
class TaskAnsiblePlaybookConfigTaskOptions {
    /**
     * Constructs a new <code>TaskAnsiblePlaybookConfigTaskOptions</code>.
     * @alias module:model/TaskAnsiblePlaybookConfigTaskOptions
     */
    constructor() { 
        
        TaskAnsiblePlaybookConfigTaskOptions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskAnsiblePlaybookConfigTaskOptions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskAnsiblePlaybookConfigTaskOptions} obj Optional instance to populate.
     * @return {module:model/TaskAnsiblePlaybookConfigTaskOptions} The populated <code>TaskAnsiblePlaybookConfigTaskOptions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskAnsiblePlaybookConfigTaskOptions();

            if (data.hasOwnProperty('ansibleOptions')) {
                obj['ansibleOptions'] = ApiClient.convertToType(data['ansibleOptions'], 'String');
            }
            if (data.hasOwnProperty('ansiblePlaybook')) {
                obj['ansiblePlaybook'] = ApiClient.convertToType(data['ansiblePlaybook'], 'String');
            }
            if (data.hasOwnProperty('sshKey')) {
                obj['sshKey'] = ApiClient.convertToType(data['sshKey'], 'String');
            }
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'String');
            }
            if (data.hasOwnProperty('localScriptGitRef')) {
                obj['localScriptGitRef'] = ApiClient.convertToType(data['localScriptGitRef'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
            if (data.hasOwnProperty('passwordHash')) {
                obj['passwordHash'] = ApiClient.convertToType(data['passwordHash'], 'String');
            }
            if (data.hasOwnProperty('localScriptGitId')) {
                obj['localScriptGitId'] = ApiClient.convertToType(data['localScriptGitId'], 'String');
            }
            if (data.hasOwnProperty('ansibleGitId')) {
                obj['ansibleGitId'] = ApiClient.convertToType(data['ansibleGitId'], 'String');
            }
            if (data.hasOwnProperty('host')) {
                obj['host'] = ApiClient.convertToType(data['host'], 'String');
            }
            if (data.hasOwnProperty('ansibleSkipTags')) {
                obj['ansibleSkipTags'] = ApiClient.convertToType(data['ansibleSkipTags'], 'String');
            }
            if (data.hasOwnProperty('ansibleTags')) {
                obj['ansibleTags'] = ApiClient.convertToType(data['ansibleTags'], 'String');
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('ansibleGitRef')) {
                obj['ansibleGitRef'] = ApiClient.convertToType(data['ansibleGitRef'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} ansibleOptions
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['ansibleOptions'] = undefined;

/**
 * @member {String} ansiblePlaybook
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['ansiblePlaybook'] = undefined;

/**
 * @member {String} sshKey
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['sshKey'] = undefined;

/**
 * @member {String} port
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['port'] = undefined;

/**
 * @member {String} localScriptGitRef
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['localScriptGitRef'] = undefined;

/**
 * @member {String} password
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['password'] = undefined;

/**
 * @member {String} passwordHash
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['passwordHash'] = undefined;

/**
 * @member {String} localScriptGitId
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['localScriptGitId'] = undefined;

/**
 * @member {String} ansibleGitId
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['ansibleGitId'] = undefined;

/**
 * @member {String} host
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['host'] = undefined;

/**
 * @member {String} ansibleSkipTags
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['ansibleSkipTags'] = undefined;

/**
 * @member {String} ansibleTags
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['ansibleTags'] = undefined;

/**
 * @member {String} username
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['username'] = undefined;

/**
 * @member {String} ansibleGitRef
 */
TaskAnsiblePlaybookConfigTaskOptions.prototype['ansibleGitRef'] = undefined;






export default TaskAnsiblePlaybookConfigTaskOptions;
