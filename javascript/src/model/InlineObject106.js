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
import ApiKeyPairsGenerateKeyPair from './ApiKeyPairsGenerateKeyPair';

/**
 * The InlineObject106 model module.
 * @module model/InlineObject106
 * @version 6.2.1
 */
class InlineObject106 {
    /**
     * Constructs a new <code>InlineObject106</code>.
     * @alias module:model/InlineObject106
     * @param keyPair {module:model/ApiKeyPairsGenerateKeyPair} 
     */
    constructor(keyPair) { 
        
        InlineObject106.initialize(this, keyPair);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, keyPair) { 
        obj['keyPair'] = keyPair;
    }

    /**
     * Constructs a <code>InlineObject106</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject106} obj Optional instance to populate.
     * @return {module:model/InlineObject106} The populated <code>InlineObject106</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject106();

            if (data.hasOwnProperty('keyPair')) {
                obj['keyPair'] = ApiKeyPairsGenerateKeyPair.constructFromObject(data['keyPair']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiKeyPairsGenerateKeyPair} keyPair
 */
InlineObject106.prototype['keyPair'] = undefined;






export default InlineObject106;

