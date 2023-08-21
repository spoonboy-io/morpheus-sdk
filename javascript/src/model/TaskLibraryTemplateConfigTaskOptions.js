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
 * The TaskLibraryTemplateConfigTaskOptions model module.
 * @module model/TaskLibraryTemplateConfigTaskOptions
 * @version 6.2.1
 */
class TaskLibraryTemplateConfigTaskOptions {
    /**
     * Constructs a new <code>TaskLibraryTemplateConfigTaskOptions</code>.
     * @alias module:model/TaskLibraryTemplateConfigTaskOptions
     */
    constructor() { 
        
        TaskLibraryTemplateConfigTaskOptions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskLibraryTemplateConfigTaskOptions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskLibraryTemplateConfigTaskOptions} obj Optional instance to populate.
     * @return {module:model/TaskLibraryTemplateConfigTaskOptions} The populated <code>TaskLibraryTemplateConfigTaskOptions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskLibraryTemplateConfigTaskOptions();

            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('localScriptGitId')) {
                obj['localScriptGitId'] = ApiClient.convertToType(data['localScriptGitId'], 'String');
            }
            if (data.hasOwnProperty('sshKey')) {
                obj['sshKey'] = ApiClient.convertToType(data['sshKey'], 'String');
            }
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
            if (data.hasOwnProperty('passwordHash')) {
                obj['passwordHash'] = ApiClient.convertToType(data['passwordHash'], 'String');
            }
            if (data.hasOwnProperty('containerTemplateId')) {
                obj['containerTemplateId'] = ApiClient.convertToType(data['containerTemplateId'], 'String');
            }
            if (data.hasOwnProperty('containerTemplate')) {
                obj['containerTemplate'] = ApiClient.convertToType(data['containerTemplate'], 'String');
            }
            if (data.hasOwnProperty('localScriptGitRef')) {
                obj['localScriptGitRef'] = ApiClient.convertToType(data['localScriptGitRef'], 'String');
            }
            if (data.hasOwnProperty('host')) {
                obj['host'] = ApiClient.convertToType(data['host'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} username
 */
TaskLibraryTemplateConfigTaskOptions.prototype['username'] = undefined;

/**
 * @member {String} localScriptGitId
 */
TaskLibraryTemplateConfigTaskOptions.prototype['localScriptGitId'] = undefined;

/**
 * @member {String} sshKey
 */
TaskLibraryTemplateConfigTaskOptions.prototype['sshKey'] = undefined;

/**
 * @member {String} port
 */
TaskLibraryTemplateConfigTaskOptions.prototype['port'] = undefined;

/**
 * @member {String} password
 */
TaskLibraryTemplateConfigTaskOptions.prototype['password'] = undefined;

/**
 * @member {String} passwordHash
 */
TaskLibraryTemplateConfigTaskOptions.prototype['passwordHash'] = undefined;

/**
 * @member {String} containerTemplateId
 */
TaskLibraryTemplateConfigTaskOptions.prototype['containerTemplateId'] = undefined;

/**
 * @member {String} containerTemplate
 */
TaskLibraryTemplateConfigTaskOptions.prototype['containerTemplate'] = undefined;

/**
 * @member {String} localScriptGitRef
 */
TaskLibraryTemplateConfigTaskOptions.prototype['localScriptGitRef'] = undefined;

/**
 * @member {String} host
 */
TaskLibraryTemplateConfigTaskOptions.prototype['host'] = undefined;






export default TaskLibraryTemplateConfigTaskOptions;

