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
 * The VdiPoolConfigVolumesDisplay model module.
 * @module model/VdiPoolConfigVolumesDisplay
 * @version 6.2.1
 */
class VdiPoolConfigVolumesDisplay {
    /**
     * Constructs a new <code>VdiPoolConfigVolumesDisplay</code>.
     * @alias module:model/VdiPoolConfigVolumesDisplay
     */
    constructor() { 
        
        VdiPoolConfigVolumesDisplay.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>VdiPoolConfigVolumesDisplay</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/VdiPoolConfigVolumesDisplay} obj Optional instance to populate.
     * @return {module:model/VdiPoolConfigVolumesDisplay} The populated <code>VdiPoolConfigVolumesDisplay</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new VdiPoolConfigVolumesDisplay();

            if (data.hasOwnProperty('storage')) {
                obj['storage'] = ApiClient.convertToType(data['storage'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('controller')) {
                obj['controller'] = ApiClient.convertToType(data['controller'], 'String');
            }
            if (data.hasOwnProperty('datastore')) {
                obj['datastore'] = ApiClient.convertToType(data['datastore'], 'String');
            }
            if (data.hasOwnProperty('displayOrder')) {
                obj['displayOrder'] = ApiClient.convertToType(data['displayOrder'], 'String');
            }
            if (data.hasOwnProperty('size')) {
                obj['size'] = ApiClient.convertToType(data['size'], 'Number');
            }
            if (data.hasOwnProperty('mountPoint')) {
                obj['mountPoint'] = ApiClient.convertToType(data['mountPoint'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} storage
 */
VdiPoolConfigVolumesDisplay.prototype['storage'] = undefined;

/**
 * @member {String} name
 */
VdiPoolConfigVolumesDisplay.prototype['name'] = undefined;

/**
 * @member {String} controller
 */
VdiPoolConfigVolumesDisplay.prototype['controller'] = undefined;

/**
 * @member {String} datastore
 */
VdiPoolConfigVolumesDisplay.prototype['datastore'] = undefined;

/**
 * @member {String} displayOrder
 */
VdiPoolConfigVolumesDisplay.prototype['displayOrder'] = undefined;

/**
 * @member {Number} size
 */
VdiPoolConfigVolumesDisplay.prototype['size'] = undefined;

/**
 * @member {String} mountPoint
 */
VdiPoolConfigVolumesDisplay.prototype['mountPoint'] = undefined;






export default VdiPoolConfigVolumesDisplay;

