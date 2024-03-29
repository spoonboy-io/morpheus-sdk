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
 * The SpecTemplateUpdateType model module.
 * @module model/SpecTemplateUpdateType
 * @version 6.2.1
 */
class SpecTemplateUpdateType {
    /**
     * Constructs a new <code>SpecTemplateUpdateType</code>.
     * @alias module:model/SpecTemplateUpdateType
     */
    constructor() { 
        
        SpecTemplateUpdateType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>SpecTemplateUpdateType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SpecTemplateUpdateType} obj Optional instance to populate.
     * @return {module:model/SpecTemplateUpdateType} The populated <code>SpecTemplateUpdateType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SpecTemplateUpdateType();

            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Spec Template Type. i.e. arm, cloudFormation, helm, kubernetes, oneview, terraform, ucs.
 * @member {String} code
 */
SpecTemplateUpdateType.prototype['code'] = undefined;






export default SpecTemplateUpdateType;

