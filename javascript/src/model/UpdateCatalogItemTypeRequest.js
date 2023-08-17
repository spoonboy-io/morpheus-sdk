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
import UpdateCatalogItemTypeRequestCatalogItemType from './UpdateCatalogItemTypeRequestCatalogItemType';

/**
 * The UpdateCatalogItemTypeRequest model module.
 * @module model/UpdateCatalogItemTypeRequest
 * @version 6.1.1
 */
class UpdateCatalogItemTypeRequest {
    /**
     * Constructs a new <code>UpdateCatalogItemTypeRequest</code>.
     * @alias module:model/UpdateCatalogItemTypeRequest
     */
    constructor() { 
        
        UpdateCatalogItemTypeRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>UpdateCatalogItemTypeRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UpdateCatalogItemTypeRequest} obj Optional instance to populate.
     * @return {module:model/UpdateCatalogItemTypeRequest} The populated <code>UpdateCatalogItemTypeRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UpdateCatalogItemTypeRequest();

            if (data.hasOwnProperty('catalogItemType')) {
                obj['catalogItemType'] = UpdateCatalogItemTypeRequestCatalogItemType.constructFromObject(data['catalogItemType']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>UpdateCatalogItemTypeRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>UpdateCatalogItemTypeRequest</code>.
     */
    static validateJSON(data) {
        // validate the optional field `catalogItemType`
        if (data['catalogItemType']) { // data not null
          UpdateCatalogItemTypeRequestCatalogItemType.validateJSON(data['catalogItemType']);
        }

        return true;
    }


}



/**
 * @member {module:model/UpdateCatalogItemTypeRequestCatalogItemType} catalogItemType
 */
UpdateCatalogItemTypeRequest.prototype['catalogItemType'] = undefined;






export default UpdateCatalogItemTypeRequest;

