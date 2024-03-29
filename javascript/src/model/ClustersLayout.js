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
 * The ClustersLayout model module.
 * @module model/ClustersLayout
 * @version 6.2.1
 */
class ClustersLayout {
    /**
     * Constructs a new <code>ClustersLayout</code>.
     * @alias module:model/ClustersLayout
     */
    constructor() { 
        
        ClustersLayout.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClustersLayout</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClustersLayout} obj Optional instance to populate.
     * @return {module:model/ClustersLayout} The populated <code>ClustersLayout</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClustersLayout();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('provisionTypeCode')) {
                obj['provisionTypeCode'] = ApiClient.convertToType(data['provisionTypeCode'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ClustersLayout.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ClustersLayout.prototype['name'] = undefined;

/**
 * @member {String} provisionTypeCode
 */
ClustersLayout.prototype['provisionTypeCode'] = undefined;






export default ClustersLayout;

