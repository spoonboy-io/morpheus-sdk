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
import ArchiveBucketCreateStorageProvider from './ArchiveBucketCreateStorageProvider';
import UpdateBlueprintPermissionsRequestResourcePermissionSitesInner from './UpdateBlueprintPermissionsRequestResourcePermissionSitesInner';

/**
 * The ArchiveBucketCreate model module.
 * @module model/ArchiveBucketCreate
 * @version 6.1.1
 */
class ArchiveBucketCreate {
    /**
     * Constructs a new <code>ArchiveBucketCreate</code>.
     * @alias module:model/ArchiveBucketCreate
     */
    constructor() { 
        
        ArchiveBucketCreate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ArchiveBucketCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ArchiveBucketCreate} obj Optional instance to populate.
     * @return {module:model/ArchiveBucketCreate} The populated <code>ArchiveBucketCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ArchiveBucketCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('storageProvider')) {
                obj['storageProvider'] = ArchiveBucketCreateStorageProvider.constructFromObject(data['storageProvider']);
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('isPublic')) {
                obj['isPublic'] = ApiClient.convertToType(data['isPublic'], 'Boolean');
            }
            if (data.hasOwnProperty('accounts')) {
                obj['accounts'] = UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.constructFromObject(data['accounts']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ArchiveBucketCreate</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ArchiveBucketCreate</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // ensure the json data is a string
        if (data['description'] && !(typeof data['description'] === 'string' || data['description'] instanceof String)) {
            throw new Error("Expected the field `description` to be a primitive type in the JSON string but got " + data['description']);
        }
        // validate the optional field `storageProvider`
        if (data['storageProvider']) { // data not null
          ArchiveBucketCreateStorageProvider.validateJSON(data['storageProvider']);
        }
        // ensure the json data is a string
        if (data['visibility'] && !(typeof data['visibility'] === 'string' || data['visibility'] instanceof String)) {
            throw new Error("Expected the field `visibility` to be a primitive type in the JSON string but got " + data['visibility']);
        }
        // validate the optional field `accounts`
        if (data['accounts']) { // data not null
          UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.validateJSON(data['accounts']);
        }

        return true;
    }


}



/**
 * A name for the archive bucket. Must be globally unique.
 * @member {String} name
 */
ArchiveBucketCreate.prototype['name'] = undefined;

/**
 * A description for the archive bucket
 * @member {String} description
 */
ArchiveBucketCreate.prototype['description'] = undefined;

/**
 * @member {module:model/ArchiveBucketCreateStorageProvider} storageProvider
 */
ArchiveBucketCreate.prototype['storageProvider'] = undefined;

/**
 * Visibility - Set to public to allow all tenants
 * @member {module:model/ArchiveBucketCreate.VisibilityEnum} visibility
 * @default 'private'
 */
ArchiveBucketCreate.prototype['visibility'] = 'private';

/**
 * Public URL - Set to true to allow anonymous access
 * @member {Boolean} isPublic
 * @default false
 */
ArchiveBucketCreate.prototype['isPublic'] = false;

/**
 * @member {module:model/UpdateBlueprintPermissionsRequestResourcePermissionSitesInner} accounts
 */
ArchiveBucketCreate.prototype['accounts'] = undefined;





/**
 * Allowed values for the <code>visibility</code> property.
 * @enum {String}
 * @readonly
 */
ArchiveBucketCreate['VisibilityEnum'] = {

    /**
     * value: "public"
     * @const
     */
    "public": "public",

    /**
     * value: "private"
     * @const
     */
    "private": "private"
};



export default ArchiveBucketCreate;
