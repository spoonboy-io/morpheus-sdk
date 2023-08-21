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
import OneOfstringlong from './OneOfstringlong';

/**
 * The CatalogItemTypeInstanceScribePlan model module.
 * @module model/CatalogItemTypeInstanceScribePlan
 * @version 6.2.1
 */
class CatalogItemTypeInstanceScribePlan {
    /**
     * Constructs a new <code>CatalogItemTypeInstanceScribePlan</code>.
     * @alias module:model/CatalogItemTypeInstanceScribePlan
     * @param id {module:model/OneOfstringlong} 
     */
    constructor(id) { 
        
        CatalogItemTypeInstanceScribePlan.initialize(this, id);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id) { 
        obj['id'] = id;
    }

    /**
     * Constructs a <code>CatalogItemTypeInstanceScribePlan</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CatalogItemTypeInstanceScribePlan} obj Optional instance to populate.
     * @return {module:model/CatalogItemTypeInstanceScribePlan} The populated <code>CatalogItemTypeInstanceScribePlan</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CatalogItemTypeInstanceScribePlan();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], OneOfstringlong);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/OneOfstringlong} id
 */
CatalogItemTypeInstanceScribePlan.prototype['id'] = undefined;






export default CatalogItemTypeInstanceScribePlan;
