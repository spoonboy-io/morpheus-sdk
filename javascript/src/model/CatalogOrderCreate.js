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
import CatalogOrderCreateItems from './CatalogOrderCreateItems';

/**
 * The CatalogOrderCreate model module.
 * @module model/CatalogOrderCreate
 * @version 6.2.1
 */
class CatalogOrderCreate {
    /**
     * Constructs a new <code>CatalogOrderCreate</code>.
     * @alias module:model/CatalogOrderCreate
     */
    constructor() { 
        
        CatalogOrderCreate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CatalogOrderCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CatalogOrderCreate} obj Optional instance to populate.
     * @return {module:model/CatalogOrderCreate} The populated <code>CatalogOrderCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CatalogOrderCreate();

            if (data.hasOwnProperty('items')) {
                obj['items'] = ApiClient.convertToType(data['items'], [CatalogOrderCreateItems]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/CatalogOrderCreateItems>} items
 */
CatalogOrderCreate.prototype['items'] = undefined;






export default CatalogOrderCreate;

