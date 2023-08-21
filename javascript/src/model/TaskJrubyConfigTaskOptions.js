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
 * The TaskJrubyConfigTaskOptions model module.
 * @module model/TaskJrubyConfigTaskOptions
 * @version 6.2.1
 */
class TaskJrubyConfigTaskOptions {
    /**
     * Constructs a new <code>TaskJrubyConfigTaskOptions</code>.
     * @alias module:model/TaskJrubyConfigTaskOptions
     */
    constructor() { 
        
        TaskJrubyConfigTaskOptions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskJrubyConfigTaskOptions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskJrubyConfigTaskOptions} obj Optional instance to populate.
     * @return {module:model/TaskJrubyConfigTaskOptions} The populated <code>TaskJrubyConfigTaskOptions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskJrubyConfigTaskOptions();

            if (data.hasOwnProperty('host')) {
                obj['host'] = ApiClient.convertToType(data['host'], 'String');
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('localScriptGitId')) {
                obj['localScriptGitId'] = ApiClient.convertToType(data['localScriptGitId'], 'String');
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
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'String');
            }
            if (data.hasOwnProperty('localScriptGitRef')) {
                obj['localScriptGitRef'] = ApiClient.convertToType(data['localScriptGitRef'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} host
 */
TaskJrubyConfigTaskOptions.prototype['host'] = undefined;

/**
 * @member {String} username
 */
TaskJrubyConfigTaskOptions.prototype['username'] = undefined;

/**
 * @member {String} localScriptGitId
 */
TaskJrubyConfigTaskOptions.prototype['localScriptGitId'] = undefined;

/**
 * @member {String} password
 */
TaskJrubyConfigTaskOptions.prototype['password'] = undefined;

/**
 * @member {String} passwordHash
 */
TaskJrubyConfigTaskOptions.prototype['passwordHash'] = undefined;

/**
 * @member {String} sshKey
 */
TaskJrubyConfigTaskOptions.prototype['sshKey'] = undefined;

/**
 * @member {String} port
 */
TaskJrubyConfigTaskOptions.prototype['port'] = undefined;

/**
 * @member {String} localScriptGitRef
 */
TaskJrubyConfigTaskOptions.prototype['localScriptGitRef'] = undefined;






export default TaskJrubyConfigTaskOptions;

