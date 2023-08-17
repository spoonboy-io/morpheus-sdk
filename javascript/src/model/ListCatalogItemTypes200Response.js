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
import CatalogItemType from './CatalogItemType';
import Meta from './Meta';
import MetaMeta from './MetaMeta';

/**
 * The ListCatalogItemTypes200Response model module.
 * @module model/ListCatalogItemTypes200Response
 * @version 6.1.1
 */
class ListCatalogItemTypes200Response {
    /**
     * Constructs a new <code>ListCatalogItemTypes200Response</code>.
     * @alias module:model/ListCatalogItemTypes200Response
     * @implements module:model/Meta
     */
    constructor() { 
        Meta.initialize(this);
        ListCatalogItemTypes200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ListCatalogItemTypes200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ListCatalogItemTypes200Response} obj Optional instance to populate.
     * @return {module:model/ListCatalogItemTypes200Response} The populated <code>ListCatalogItemTypes200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ListCatalogItemTypes200Response();
            Meta.constructFromObject(data, obj);

            if (data.hasOwnProperty('meta')) {
                obj['meta'] = MetaMeta.constructFromObject(data['meta']);
            }
            if (data.hasOwnProperty('catalogItemTypes')) {
                obj['catalogItemTypes'] = ApiClient.convertToType(data['catalogItemTypes'], [CatalogItemType]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ListCatalogItemTypes200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ListCatalogItemTypes200Response</code>.
     */
    static validateJSON(data) {
        // validate the optional field `meta`
        if (data['meta']) { // data not null
          MetaMeta.validateJSON(data['meta']);
        }
        if (data['catalogItemTypes']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['catalogItemTypes'])) {
                throw new Error("Expected the field `catalogItemTypes` to be an array in the JSON data but got " + data['catalogItemTypes']);
            }
            // validate the optional field `catalogItemTypes` (array)
            for (const item of data['catalogItemTypes']) {
                CatalogItemType.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {module:model/MetaMeta} meta
 */
ListCatalogItemTypes200Response.prototype['meta'] = undefined;

/**
 * @member {Array.<module:model/CatalogItemType>} catalogItemTypes
 */
ListCatalogItemTypes200Response.prototype['catalogItemTypes'] = undefined;


// Implement Meta interface:
/**
 * @member {module:model/MetaMeta} meta
 */
Meta.prototype['meta'] = undefined;




export default ListCatalogItemTypes200Response;
