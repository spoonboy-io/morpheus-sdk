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
 * The ControllerType model module.
 * @module model/ControllerType
 * @version 6.2.1
 */
class ControllerType {
    /**
     * Constructs a new <code>ControllerType</code>.
     * @alias module:model/ControllerType
     */
    constructor() { 
        
        ControllerType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ControllerType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ControllerType} obj Optional instance to populate.
     * @return {module:model/ControllerType} The populated <code>ControllerType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ControllerType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('displayOrder')) {
                obj['displayOrder'] = ApiClient.convertToType(data['displayOrder'], 'Number');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('creatable')) {
                obj['creatable'] = ApiClient.convertToType(data['creatable'], 'Boolean');
            }
            if (data.hasOwnProperty('maxDevices')) {
                obj['maxDevices'] = ApiClient.convertToType(data['maxDevices'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ControllerType.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ControllerType.prototype['name'] = undefined;

/**
 * @member {Number} displayOrder
 */
ControllerType.prototype['displayOrder'] = undefined;

/**
 * @member {String} category
 */
ControllerType.prototype['category'] = undefined;

/**
 * @member {Boolean} enabled
 */
ControllerType.prototype['enabled'] = undefined;

/**
 * @member {Boolean} creatable
 */
ControllerType.prototype['creatable'] = undefined;

/**
 * @member {Number} maxDevices
 */
ControllerType.prototype['maxDevices'] = undefined;






export default ControllerType;
