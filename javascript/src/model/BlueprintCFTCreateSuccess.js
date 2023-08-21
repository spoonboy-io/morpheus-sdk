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
import BlueprintCFTCreateSuccessCloudFormation from './BlueprintCFTCreateSuccessCloudFormation';

/**
 * The BlueprintCFTCreateSuccess model module.
 * @module model/BlueprintCFTCreateSuccess
 * @version 6.2.1
 */
class BlueprintCFTCreateSuccess {
    /**
     * Constructs a new <code>BlueprintCFTCreateSuccess</code>.
     * @alias module:model/BlueprintCFTCreateSuccess
     */
    constructor() { 
        
        BlueprintCFTCreateSuccess.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BlueprintCFTCreateSuccess</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BlueprintCFTCreateSuccess} obj Optional instance to populate.
     * @return {module:model/BlueprintCFTCreateSuccess} The populated <code>BlueprintCFTCreateSuccess</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BlueprintCFTCreateSuccess();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('image')) {
                obj['image'] = ApiClient.convertToType(data['image'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('cloudFormation')) {
                obj['cloudFormation'] = BlueprintCFTCreateSuccessCloudFormation.constructFromObject(data['cloudFormation']);
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('resourcePermission')) {
                obj['resourcePermission'] = ApiClient.convertToType(data['resourcePermission'], Object);
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = ApiClient.convertToType(data['owner'], Object);
            }
            if (data.hasOwnProperty('tenant')) {
                obj['tenant'] = ApiClient.convertToType(data['tenant'], Object);
            }
        }
        return obj;
    }


}

/**
 * A name for the blueprint
 * @member {String} name
 */
BlueprintCFTCreateSuccess.prototype['name'] = undefined;

/**
 * Path to display image. Defaults to an internal Morpheus image.
 * @member {String} image
 */
BlueprintCFTCreateSuccess.prototype['image'] = undefined;

/**
 * Blueprint Type
 * @member {module:model/BlueprintCFTCreateSuccess.TypeEnum} type
 */
BlueprintCFTCreateSuccess.prototype['type'] = undefined;

/**
 * @member {module:model/BlueprintCFTCreateSuccessCloudFormation} cloudFormation
 */
BlueprintCFTCreateSuccess.prototype['cloudFormation'] = undefined;

/**
 * Private or Public Access
 * @member {module:model/BlueprintCFTCreateSuccess.VisibilityEnum} visibility
 * @default 'private'
 */
BlueprintCFTCreateSuccess.prototype['visibility'] = 'private';

/**
 * Resource Permission Block
 * @member {Object} resourcePermission
 */
BlueprintCFTCreateSuccess.prototype['resourcePermission'] = undefined;

/**
 * Owner
 * @member {Object} owner
 */
BlueprintCFTCreateSuccess.prototype['owner'] = undefined;

/**
 * Tenant
 * @member {Object} tenant
 */
BlueprintCFTCreateSuccess.prototype['tenant'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
BlueprintCFTCreateSuccess['TypeEnum'] = {

    /**
     * value: "cloudFormation"
     * @const
     */
    "cloudFormation": "cloudFormation"
};


/**
 * Allowed values for the <code>visibility</code> property.
 * @enum {String}
 * @readonly
 */
BlueprintCFTCreateSuccess['VisibilityEnum'] = {

    /**
     * value: "private"
     * @const
     */
    "private": "private",

    /**
     * value: "public"
     * @const
     */
    "public": "public"
};



export default BlueprintCFTCreateSuccess;
