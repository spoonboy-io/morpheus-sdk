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
 * The InstancesCloneImage model module.
 * @module model/InstancesCloneImage
 * @version 6.2.1
 */
class InstancesCloneImage {
    /**
     * Constructs a new <code>InstancesCloneImage</code>.
     * @alias module:model/InstancesCloneImage
     */
    constructor() { 
        
        InstancesCloneImage.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstancesCloneImage</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstancesCloneImage} obj Optional instance to populate.
     * @return {module:model/InstancesCloneImage} The populated <code>InstancesCloneImage</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstancesCloneImage();

            if (data.hasOwnProperty('templateName')) {
                obj['templateName'] = ApiClient.convertToType(data['templateName'], 'String');
            }
            if (data.hasOwnProperty('zoneFolder')) {
                obj['zoneFolder'] = ApiClient.convertToType(data['zoneFolder'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Image Template Name
 * @member {String} templateName
 * @default '{server.name}-{timestamp}'
 */
InstancesCloneImage.prototype['templateName'] = '{server.name}-{timestamp}';

/**
 * Zone Folder externalId. This is required for VMware
 * @member {String} zoneFolder
 */
InstancesCloneImage.prototype['zoneFolder'] = undefined;






export default InstancesCloneImage;

