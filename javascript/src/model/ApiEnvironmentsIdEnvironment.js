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
 * The ApiEnvironmentsIdEnvironment model module.
 * @module model/ApiEnvironmentsIdEnvironment
 * @version 6.2.1
 */
class ApiEnvironmentsIdEnvironment {
    /**
     * Constructs a new <code>ApiEnvironmentsIdEnvironment</code>.
     * Payload for updating a new environment
     * @alias module:model/ApiEnvironmentsIdEnvironment
     */
    constructor() { 
        
        ApiEnvironmentsIdEnvironment.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiEnvironmentsIdEnvironment</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiEnvironmentsIdEnvironment} obj Optional instance to populate.
     * @return {module:model/ApiEnvironmentsIdEnvironment} The populated <code>ApiEnvironmentsIdEnvironment</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiEnvironmentsIdEnvironment();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('sortOrder')) {
                obj['sortOrder'] = ApiClient.convertToType(data['sortOrder'], 'Number');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * A unique name for the environment
 * @member {String} name
 */
ApiEnvironmentsIdEnvironment.prototype['name'] = undefined;

/**
 * A description of the environment
 * @member {String} description
 */
ApiEnvironmentsIdEnvironment.prototype['description'] = undefined;

/**
 * private or public
 * @member {String} visibility
 * @default 'private'
 */
ApiEnvironmentsIdEnvironment.prototype['visibility'] = 'private';

/**
 * Sort order
 * @member {Number} sortOrder
 * @default 0
 */
ApiEnvironmentsIdEnvironment.prototype['sortOrder'] = 0;

/**
 * Set to false to deactivate the environment
 * @member {Boolean} active
 */
ApiEnvironmentsIdEnvironment.prototype['active'] = undefined;






export default ApiEnvironmentsIdEnvironment;

