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
 * The ScriptCreate model module.
 * @module model/ScriptCreate
 * @version 6.2.1
 */
class ScriptCreate {
    /**
     * Constructs a new <code>ScriptCreate</code>.
     * @alias module:model/ScriptCreate
     * @param name {String} Script name
     */
    constructor(name) { 
        
        ScriptCreate.initialize(this, name);
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
     * Constructs a <code>ScriptCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ScriptCreate} obj Optional instance to populate.
     * @return {module:model/ScriptCreate} The populated <code>ScriptCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ScriptCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('scriptVersion')) {
                obj['scriptVersion'] = ApiClient.convertToType(data['scriptVersion'], 'String');
            }
            if (data.hasOwnProperty('scriptPhase')) {
                obj['scriptPhase'] = ApiClient.convertToType(data['scriptPhase'], 'String');
            }
            if (data.hasOwnProperty('scriptType')) {
                obj['scriptType'] = ApiClient.convertToType(data['scriptType'], 'String');
            }
            if (data.hasOwnProperty('script')) {
                obj['script'] = ApiClient.convertToType(data['script'], 'String');
            }
            if (data.hasOwnProperty('runAsUser')) {
                obj['runAsUser'] = ApiClient.convertToType(data['runAsUser'], 'String');
            }
            if (data.hasOwnProperty('sudoUser')) {
                obj['sudoUser'] = ApiClient.convertToType(data['sudoUser'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * Script name
 * @member {String} name
 */
ScriptCreate.prototype['name'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
ScriptCreate.prototype['labels'] = undefined;

/**
 * Script category
 * @member {String} category
 */
ScriptCreate.prototype['category'] = undefined;

/**
 * Version of the script
 * @member {String} scriptVersion
 * @default '1'
 */
ScriptCreate.prototype['scriptVersion'] = '1';

/**
 * Phase for the script, provision, start, etc.
 * @member {String} scriptPhase
 */
ScriptCreate.prototype['scriptPhase'] = undefined;

/**
 * Type for the script
 * @member {module:model/ScriptCreate.ScriptTypeEnum} scriptType
 * @default 'bash'
 */
ScriptCreate.prototype['scriptType'] = 'bash';

/**
 * Script content, that is, the code itself.
 * @member {String} script
 */
ScriptCreate.prototype['script'] = undefined;

/**
 * Run as a specific user.
 * @member {String} runAsUser
 */
ScriptCreate.prototype['runAsUser'] = undefined;

/**
 * Sudo, whether or not to run with sudo.
 * @member {Boolean} sudoUser
 * @default false
 */
ScriptCreate.prototype['sudoUser'] = false;





/**
 * Allowed values for the <code>scriptType</code> property.
 * @enum {String}
 * @readonly
 */
ScriptCreate['ScriptTypeEnum'] = {

    /**
     * value: "bash"
     * @const
     */
    "bash": "bash",

    /**
     * value: "powershell"
     * @const
     */
    "powershell": "powershell"
};



export default ScriptCreate;
