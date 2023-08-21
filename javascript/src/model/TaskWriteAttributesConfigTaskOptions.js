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
 * The TaskWriteAttributesConfigTaskOptions model module.
 * @module model/TaskWriteAttributesConfigTaskOptions
 * @version 6.2.1
 */
class TaskWriteAttributesConfigTaskOptions {
    /**
     * Constructs a new <code>TaskWriteAttributesConfigTaskOptions</code>.
     * @alias module:model/TaskWriteAttributesConfigTaskOptions
     */
    constructor() { 
        
        TaskWriteAttributesConfigTaskOptions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskWriteAttributesConfigTaskOptions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskWriteAttributesConfigTaskOptions} obj Optional instance to populate.
     * @return {module:model/TaskWriteAttributesConfigTaskOptions} The populated <code>TaskWriteAttributesConfigTaskOptions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskWriteAttributesConfigTaskOptions();

            if (data.hasOwnProperty('host')) {
                obj['host'] = ApiClient.convertToType(data['host'], 'String');
            }
            if (data.hasOwnProperty('sshKey')) {
                obj['sshKey'] = ApiClient.convertToType(data['sshKey'], 'String');
            }
            if (data.hasOwnProperty('localScriptGitRef')) {
                obj['localScriptGitRef'] = ApiClient.convertToType(data['localScriptGitRef'], 'String');
            }
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'String');
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
            if (data.hasOwnProperty('writeAttributes.attributes')) {
                obj['writeAttributes.attributes'] = ApiClient.convertToType(data['writeAttributes.attributes'], 'String');
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} host
 */
TaskWriteAttributesConfigTaskOptions.prototype['host'] = undefined;

/**
 * @member {String} sshKey
 */
TaskWriteAttributesConfigTaskOptions.prototype['sshKey'] = undefined;

/**
 * @member {String} localScriptGitRef
 */
TaskWriteAttributesConfigTaskOptions.prototype['localScriptGitRef'] = undefined;

/**
 * @member {String} port
 */
TaskWriteAttributesConfigTaskOptions.prototype['port'] = undefined;

/**
 * @member {String} localScriptGitId
 */
TaskWriteAttributesConfigTaskOptions.prototype['localScriptGitId'] = undefined;

/**
 * @member {String} password
 */
TaskWriteAttributesConfigTaskOptions.prototype['password'] = undefined;

/**
 * @member {String} passwordHash
 */
TaskWriteAttributesConfigTaskOptions.prototype['passwordHash'] = undefined;

/**
 * @member {String} writeAttributes.attributes
 */
TaskWriteAttributesConfigTaskOptions.prototype['writeAttributes.attributes'] = undefined;

/**
 * @member {String} username
 */
TaskWriteAttributesConfigTaskOptions.prototype['username'] = undefined;






export default TaskWriteAttributesConfigTaskOptions;

