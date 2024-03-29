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
import ApiZonesZoneIdDataStoresIdDatastore from './ApiZonesZoneIdDataStoresIdDatastore';

/**
 * The InlineObject43 model module.
 * @module model/InlineObject43
 * @version 6.2.1
 */
class InlineObject43 {
    /**
     * Constructs a new <code>InlineObject43</code>.
     * @alias module:model/InlineObject43
     * @param datastore {module:model/ApiZonesZoneIdDataStoresIdDatastore} 
     */
    constructor(datastore) { 
        
        InlineObject43.initialize(this, datastore);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, datastore) { 
        obj['datastore'] = datastore;
    }

    /**
     * Constructs a <code>InlineObject43</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject43} obj Optional instance to populate.
     * @return {module:model/InlineObject43} The populated <code>InlineObject43</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject43();

            if (data.hasOwnProperty('datastore')) {
                obj['datastore'] = ApiZonesZoneIdDataStoresIdDatastore.constructFromObject(data['datastore']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiZonesZoneIdDataStoresIdDatastore} datastore
 */
InlineObject43.prototype['datastore'] = undefined;






export default InlineObject43;

