/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The ArchiveBucketCreateStorageProvider model module.
 * @module model/ArchiveBucketCreateStorageProvider
 * @version 6.1.1
 */
class ArchiveBucketCreateStorageProvider {
    /**
     * Constructs a new <code>ArchiveBucketCreateStorageProvider</code>.
     * Storage Provider
     * @alias module:model/ArchiveBucketCreateStorageProvider
     */
    constructor() { 
        
        ArchiveBucketCreateStorageProvider.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ArchiveBucketCreateStorageProvider</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ArchiveBucketCreateStorageProvider} obj Optional instance to populate.
     * @return {module:model/ArchiveBucketCreateStorageProvider} The populated <code>ArchiveBucketCreateStorageProvider</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ArchiveBucketCreateStorageProvider();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ArchiveBucketCreateStorageProvider</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ArchiveBucketCreateStorageProvider</code>.
     */
    static validateJSON(data) {

        return true;
    }


}



/**
 * @member {Number} id
 */
ArchiveBucketCreateStorageProvider.prototype['id'] = undefined;






export default ArchiveBucketCreateStorageProvider;
