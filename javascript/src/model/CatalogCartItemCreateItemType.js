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
 * The CatalogCartItemCreateItemType model module.
 * @module model/CatalogCartItemCreateItemType
 * @version 6.2.1
 */
class CatalogCartItemCreateItemType {
    /**
     * Constructs a new <code>CatalogCartItemCreateItemType</code>.
     * @alias module:model/CatalogCartItemCreateItemType
     */
    constructor() { 
        
        CatalogCartItemCreateItemType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CatalogCartItemCreateItemType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CatalogCartItemCreateItemType} obj Optional instance to populate.
     * @return {module:model/CatalogCartItemCreateItemType} The populated <code>CatalogCartItemCreateItemType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CatalogCartItemCreateItemType();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Catalog item name
 * @member {String} name
 */
CatalogCartItemCreateItemType.prototype['name'] = undefined;






export default CatalogCartItemCreateItemType;

