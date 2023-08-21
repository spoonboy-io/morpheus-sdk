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
 * The ImageBuildConfigVolumes model module.
 * @module model/ImageBuildConfigVolumes
 * @version 6.2.1
 */
class ImageBuildConfigVolumes {
    /**
     * Constructs a new <code>ImageBuildConfigVolumes</code>.
     * @alias module:model/ImageBuildConfigVolumes
     */
    constructor() { 
        
        ImageBuildConfigVolumes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ImageBuildConfigVolumes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ImageBuildConfigVolumes} obj Optional instance to populate.
     * @return {module:model/ImageBuildConfigVolumes} The populated <code>ImageBuildConfigVolumes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ImageBuildConfigVolumes();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('size')) {
                obj['size'] = ApiClient.convertToType(data['size'], 'Number');
            }
            if (data.hasOwnProperty('maxIOPS')) {
                obj['maxIOPS'] = ApiClient.convertToType(data['maxIOPS'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('rootVolume')) {
                obj['rootVolume'] = ApiClient.convertToType(data['rootVolume'], 'Boolean');
            }
            if (data.hasOwnProperty('storageType')) {
                obj['storageType'] = ApiClient.convertToType(data['storageType'], 'Number');
            }
            if (data.hasOwnProperty('datastoreId')) {
                obj['datastoreId'] = ApiClient.convertToType(data['datastoreId'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ImageBuildConfigVolumes.prototype['id'] = undefined;

/**
 * @member {Number} size
 */
ImageBuildConfigVolumes.prototype['size'] = undefined;

/**
 * @member {String} maxIOPS
 */
ImageBuildConfigVolumes.prototype['maxIOPS'] = undefined;

/**
 * @member {String} name
 */
ImageBuildConfigVolumes.prototype['name'] = undefined;

/**
 * @member {Boolean} rootVolume
 */
ImageBuildConfigVolumes.prototype['rootVolume'] = undefined;

/**
 * @member {Number} storageType
 */
ImageBuildConfigVolumes.prototype['storageType'] = undefined;

/**
 * @member {String} datastoreId
 */
ImageBuildConfigVolumes.prototype['datastoreId'] = undefined;






export default ImageBuildConfigVolumes;
