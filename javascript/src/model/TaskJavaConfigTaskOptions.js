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
 * The TaskJavaConfigTaskOptions model module.
 * @module model/TaskJavaConfigTaskOptions
 * @version 6.2.1
 */
class TaskJavaConfigTaskOptions {
    /**
     * Constructs a new <code>TaskJavaConfigTaskOptions</code>.
     * @alias module:model/TaskJavaConfigTaskOptions
     */
    constructor() { 
        
        TaskJavaConfigTaskOptions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskJavaConfigTaskOptions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskJavaConfigTaskOptions} obj Optional instance to populate.
     * @return {module:model/TaskJavaConfigTaskOptions} The populated <code>TaskJavaConfigTaskOptions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskJavaConfigTaskOptions();

            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'String');
            }
            if (data.hasOwnProperty('jsScript')) {
                obj['jsScript'] = ApiClient.convertToType(data['jsScript'], 'String');
            }
            if (data.hasOwnProperty('host')) {
                obj['host'] = ApiClient.convertToType(data['host'], 'String');
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
            if (data.hasOwnProperty('localScriptGitId')) {
                obj['localScriptGitId'] = ApiClient.convertToType(data['localScriptGitId'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} username
 */
TaskJavaConfigTaskOptions.prototype['username'] = undefined;

/**
 * @member {String} port
 */
TaskJavaConfigTaskOptions.prototype['port'] = undefined;

/**
 * @member {String} jsScript
 */
TaskJavaConfigTaskOptions.prototype['jsScript'] = undefined;

/**
 * @member {String} host
 */
TaskJavaConfigTaskOptions.prototype['host'] = undefined;

/**
 * @member {String} localScriptGitRef
 */
TaskJavaConfigTaskOptions.prototype['localScriptGitRef'] = undefined;

/**
 * @member {String} password
 */
TaskJavaConfigTaskOptions.prototype['password'] = undefined;

/**
 * @member {String} passwordHash
 */
TaskJavaConfigTaskOptions.prototype['passwordHash'] = undefined;

/**
 * @member {String} sshKey
 */
TaskJavaConfigTaskOptions.prototype['sshKey'] = undefined;

/**
 * @member {String} localScriptGitId
 */
TaskJavaConfigTaskOptions.prototype['localScriptGitId'] = undefined;






export default TaskJavaConfigTaskOptions;

