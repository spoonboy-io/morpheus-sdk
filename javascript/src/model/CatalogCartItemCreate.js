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
import CatalogCartItemCreateItem from './CatalogCartItemCreateItem';

/**
 * The CatalogCartItemCreate model module.
 * @module model/CatalogCartItemCreate
 * @version 6.2.1
 */
class CatalogCartItemCreate {
    /**
     * Constructs a new <code>CatalogCartItemCreate</code>.
     * @alias module:model/CatalogCartItemCreate
     */
    constructor() { 
        
        CatalogCartItemCreate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CatalogCartItemCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CatalogCartItemCreate} obj Optional instance to populate.
     * @return {module:model/CatalogCartItemCreate} The populated <code>CatalogCartItemCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CatalogCartItemCreate();

            if (data.hasOwnProperty('item')) {
                obj['item'] = CatalogCartItemCreateItem.constructFromObject(data['item']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/CatalogCartItemCreateItem} item
 */
CatalogCartItemCreate.prototype['item'] = undefined;






export default CatalogCartItemCreate;
