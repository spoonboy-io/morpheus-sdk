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
 * The TaskPowershellConfigTaskOptions model module.
 * @module model/TaskPowershellConfigTaskOptions
 * @version 6.2.1
 */
class TaskPowershellConfigTaskOptions {
    /**
     * Constructs a new <code>TaskPowershellConfigTaskOptions</code>.
     * @alias module:model/TaskPowershellConfigTaskOptions
     */
    constructor() { 
        
        TaskPowershellConfigTaskOptions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskPowershellConfigTaskOptions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskPowershellConfigTaskOptions} obj Optional instance to populate.
     * @return {module:model/TaskPowershellConfigTaskOptions} The populated <code>TaskPowershellConfigTaskOptions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskPowershellConfigTaskOptions();

            if (data.hasOwnProperty('host')) {
                obj['host'] = ApiClient.convertToType(data['host'], 'String');
            }
            if (data.hasOwnProperty('localScriptGitId')) {
                obj['localScriptGitId'] = ApiClient.convertToType(data['localScriptGitId'], 'String');
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
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
            if (data.hasOwnProperty('sshKey')) {
                obj['sshKey'] = ApiClient.convertToType(data['sshKey'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} host
 */
TaskPowershellConfigTaskOptions.prototype['host'] = undefined;

/**
 * @member {String} localScriptGitId
 */
TaskPowershellConfigTaskOptions.prototype['localScriptGitId'] = undefined;

/**
 * @member {String} username
 */
TaskPowershellConfigTaskOptions.prototype['username'] = undefined;

/**
 * @member {String} port
 */
TaskPowershellConfigTaskOptions.prototype['port'] = undefined;

/**
 * @member {String} localScriptGitRef
 */
TaskPowershellConfigTaskOptions.prototype['localScriptGitRef'] = undefined;

/**
 * @member {String} password
 */
TaskPowershellConfigTaskOptions.prototype['password'] = undefined;

/**
 * @member {String} passwordHash
 */
TaskPowershellConfigTaskOptions.prototype['passwordHash'] = undefined;

/**
 * @member {String} sshKey
 */
TaskPowershellConfigTaskOptions.prototype['sshKey'] = undefined;






export default TaskPowershellConfigTaskOptions;

