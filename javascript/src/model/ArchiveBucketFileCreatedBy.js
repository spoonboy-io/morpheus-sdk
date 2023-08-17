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
 * The ArchiveBucketFileCreatedBy model module.
 * @module model/ArchiveBucketFileCreatedBy
 * @version 6.1.1
 */
class ArchiveBucketFileCreatedBy {
    /**
     * Constructs a new <code>ArchiveBucketFileCreatedBy</code>.
     * @alias module:model/ArchiveBucketFileCreatedBy
     */
    constructor() { 
        
        ArchiveBucketFileCreatedBy.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ArchiveBucketFileCreatedBy</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ArchiveBucketFileCreatedBy} obj Optional instance to populate.
     * @return {module:model/ArchiveBucketFileCreatedBy} The populated <code>ArchiveBucketFileCreatedBy</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ArchiveBucketFileCreatedBy();

            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ArchiveBucketFileCreatedBy</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ArchiveBucketFileCreatedBy</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['username'] && !(typeof data['username'] === 'string' || data['username'] instanceof String)) {
            throw new Error("Expected the field `username` to be a primitive type in the JSON string but got " + data['username']);
        }

        return true;
    }


}



/**
 * @member {String} username
 */
ArchiveBucketFileCreatedBy.prototype['username'] = undefined;






export default ArchiveBucketFileCreatedBy;
