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
import InlineResponse20035Actions from './InlineResponse20035Actions';

/**
 * The InlineResponse20035 model module.
 * @module model/InlineResponse20035
 * @version 6.2.1
 */
class InlineResponse20035 {
    /**
     * Constructs a new <code>InlineResponse20035</code>.
     * @alias module:model/InlineResponse20035
     */
    constructor() { 
        
        InlineResponse20035.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20035</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20035} obj Optional instance to populate.
     * @return {module:model/InlineResponse20035} The populated <code>InlineResponse20035</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20035();

            if (data.hasOwnProperty('containerIds')) {
                obj['containerIds'] = ApiClient.convertToType(data['containerIds'], ['Number']);
            }
            if (data.hasOwnProperty('actions')) {
                obj['actions'] = ApiClient.convertToType(data['actions'], [InlineResponse20035Actions]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<Number>} containerIds
 */
InlineResponse20035.prototype['containerIds'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20035Actions>} actions
 */
InlineResponse20035.prototype['actions'] = undefined;






export default InlineResponse20035;

