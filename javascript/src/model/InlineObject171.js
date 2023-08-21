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
 * The InlineObject171 model module.
 * @module model/InlineObject171
 * @version 6.2.1
 */
class InlineObject171 {
    /**
     * Constructs a new <code>InlineObject171</code>.
     * The parameters for update a network Edge Cluster is type dependent. The following lists the common parameters. See get a specific type to list available options for the network server type. 
     * @alias module:model/InlineObject171
     */
    constructor() { 
        
        InlineObject171.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject171</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject171} obj Optional instance to populate.
     * @return {module:model/InlineObject171} The populated <code>InlineObject171</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject171();

            if (data.hasOwnProperty('networkEdgeCluster')) {
                obj['networkEdgeCluster'] = ApiClient.convertToType(data['networkEdgeCluster'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {Object} networkEdgeCluster
 */
InlineObject171.prototype['networkEdgeCluster'] = undefined;






export default InlineObject171;

