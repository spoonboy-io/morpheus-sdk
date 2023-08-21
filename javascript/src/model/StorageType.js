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
 * The StorageType model module.
 * @module model/StorageType
 * @version 6.2.1
 */
class StorageType {
    /**
     * Constructs a new <code>StorageType</code>.
     * @alias module:model/StorageType
     */
    constructor() { 
        
        StorageType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>StorageType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/StorageType} obj Optional instance to populate.
     * @return {module:model/StorageType} The populated <code>StorageType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new StorageType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('displayOrder')) {
                obj['displayOrder'] = ApiClient.convertToType(data['displayOrder'], 'Number');
            }
            if (data.hasOwnProperty('defaultType')) {
                obj['defaultType'] = ApiClient.convertToType(data['defaultType'], 'Boolean');
            }
            if (data.hasOwnProperty('customLabel')) {
                obj['customLabel'] = ApiClient.convertToType(data['customLabel'], 'Boolean');
            }
            if (data.hasOwnProperty('customSize')) {
                obj['customSize'] = ApiClient.convertToType(data['customSize'], 'Boolean');
            }
            if (data.hasOwnProperty('customSizeOptions')) {
                obj['customSizeOptions'] = ApiClient.convertToType(data['customSizeOptions'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
StorageType.prototype['id'] = undefined;

/**
 * @member {String} code
 */
StorageType.prototype['code'] = undefined;

/**
 * @member {String} name
 */
StorageType.prototype['name'] = undefined;

/**
 * @member {Number} displayOrder
 */
StorageType.prototype['displayOrder'] = undefined;

/**
 * @member {Boolean} defaultType
 */
StorageType.prototype['defaultType'] = undefined;

/**
 * @member {Boolean} customLabel
 */
StorageType.prototype['customLabel'] = undefined;

/**
 * @member {Boolean} customSize
 */
StorageType.prototype['customSize'] = undefined;

/**
 * @member {String} customSizeOptions
 */
StorageType.prototype['customSizeOptions'] = undefined;






export default StorageType;

