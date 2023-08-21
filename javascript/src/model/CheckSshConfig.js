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
 * The CheckSshConfig model module.
 * @module model/CheckSshConfig
 * @version 6.2.1
 */
class CheckSshConfig {
    /**
     * Constructs a new <code>CheckSshConfig</code>.
     * @alias module:model/CheckSshConfig
     */
    constructor() { 
        
        CheckSshConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CheckSshConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CheckSshConfig} obj Optional instance to populate.
     * @return {module:model/CheckSshConfig} The populated <code>CheckSshConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CheckSshConfig();

            if (data.hasOwnProperty('sshPort')) {
                obj['sshPort'] = ApiClient.convertToType(data['sshPort'], 'Number');
            }
            if (data.hasOwnProperty('checkUser')) {
                obj['checkUser'] = ApiClient.convertToType(data['checkUser'], 'String');
            }
            if (data.hasOwnProperty('tunnelOn')) {
                obj['tunnelOn'] = ApiClient.convertToType(data['tunnelOn'], 'String');
            }
            if (data.hasOwnProperty('textCheckOn')) {
                obj['textCheckOn'] = ApiClient.convertToType(data['textCheckOn'], 'String');
            }
            if (data.hasOwnProperty('checkPassword')) {
                obj['checkPassword'] = ApiClient.convertToType(data['checkPassword'], 'String');
            }
            if (data.hasOwnProperty('sshHost')) {
                obj['sshHost'] = ApiClient.convertToType(data['sshHost'], 'String');
            }
            if (data.hasOwnProperty('sshUser')) {
                obj['sshUser'] = ApiClient.convertToType(data['sshUser'], 'String');
            }
            if (data.hasOwnProperty('webTextMatch')) {
                obj['webTextMatch'] = ApiClient.convertToType(data['webTextMatch'], 'String');
            }
            if (data.hasOwnProperty('checkPasswordHash')) {
                obj['checkPasswordHash'] = ApiClient.convertToType(data['checkPasswordHash'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} sshPort
 */
CheckSshConfig.prototype['sshPort'] = undefined;

/**
 * @member {String} checkUser
 */
CheckSshConfig.prototype['checkUser'] = undefined;

/**
 * @member {String} tunnelOn
 */
CheckSshConfig.prototype['tunnelOn'] = undefined;

/**
 * @member {String} textCheckOn
 */
CheckSshConfig.prototype['textCheckOn'] = undefined;

/**
 * @member {String} checkPassword
 */
CheckSshConfig.prototype['checkPassword'] = undefined;

/**
 * @member {String} sshHost
 */
CheckSshConfig.prototype['sshHost'] = undefined;

/**
 * @member {String} sshUser
 */
CheckSshConfig.prototype['sshUser'] = undefined;

/**
 * @member {String} webTextMatch
 */
CheckSshConfig.prototype['webTextMatch'] = undefined;

/**
 * @member {String} checkPasswordHash
 */
CheckSshConfig.prototype['checkPasswordHash'] = undefined;






export default CheckSshConfig;
