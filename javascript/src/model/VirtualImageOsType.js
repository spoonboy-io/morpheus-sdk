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
 * The VirtualImageOsType model module.
 * @module model/VirtualImageOsType
 * @version 6.2.1
 */
class VirtualImageOsType {
    /**
     * Constructs a new <code>VirtualImageOsType</code>.
     * @alias module:model/VirtualImageOsType
     */
    constructor() { 
        
        VirtualImageOsType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>VirtualImageOsType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/VirtualImageOsType} obj Optional instance to populate.
     * @return {module:model/VirtualImageOsType} The populated <code>VirtualImageOsType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new VirtualImageOsType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('vendor')) {
                obj['vendor'] = ApiClient.convertToType(data['vendor'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('osFamily')) {
                obj['osFamily'] = ApiClient.convertToType(data['osFamily'], 'String');
            }
            if (data.hasOwnProperty('osVersion')) {
                obj['osVersion'] = ApiClient.convertToType(data['osVersion'], 'String');
            }
            if (data.hasOwnProperty('bitCount')) {
                obj['bitCount'] = ApiClient.convertToType(data['bitCount'], 'Number');
            }
            if (data.hasOwnProperty('platform')) {
                obj['platform'] = ApiClient.convertToType(data['platform'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
VirtualImageOsType.prototype['id'] = undefined;

/**
 * @member {String} code
 */
VirtualImageOsType.prototype['code'] = undefined;

/**
 * @member {String} name
 */
VirtualImageOsType.prototype['name'] = undefined;

/**
 * @member {String} description
 */
VirtualImageOsType.prototype['description'] = undefined;

/**
 * @member {String} vendor
 */
VirtualImageOsType.prototype['vendor'] = undefined;

/**
 * @member {String} category
 */
VirtualImageOsType.prototype['category'] = undefined;

/**
 * @member {String} osFamily
 */
VirtualImageOsType.prototype['osFamily'] = undefined;

/**
 * @member {String} osVersion
 */
VirtualImageOsType.prototype['osVersion'] = undefined;

/**
 * @member {Number} bitCount
 */
VirtualImageOsType.prototype['bitCount'] = undefined;

/**
 * @member {String} platform
 */
VirtualImageOsType.prototype['platform'] = undefined;






export default VirtualImageOsType;

