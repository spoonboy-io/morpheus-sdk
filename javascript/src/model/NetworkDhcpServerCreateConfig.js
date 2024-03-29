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
 * The NetworkDhcpServerCreateConfig model module.
 * @module model/NetworkDhcpServerCreateConfig
 * @version 6.2.1
 */
class NetworkDhcpServerCreateConfig {
    /**
     * Constructs a new <code>NetworkDhcpServerCreateConfig</code>.
     * Configuration object with parameters that vary by pool type.
     * @alias module:model/NetworkDhcpServerCreateConfig
     */
    constructor() { 
        
        NetworkDhcpServerCreateConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NetworkDhcpServerCreateConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkDhcpServerCreateConfig} obj Optional instance to populate.
     * @return {module:model/NetworkDhcpServerCreateConfig} The populated <code>NetworkDhcpServerCreateConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkDhcpServerCreateConfig();

            if (data.hasOwnProperty('edgeCluster')) {
                obj['edgeCluster'] = ApiClient.convertToType(data['edgeCluster'], 'String');
            }
            if (data.hasOwnProperty('preferredEdgeNode1')) {
                obj['preferredEdgeNode1'] = ApiClient.convertToType(data['preferredEdgeNode1'], 'String');
            }
            if (data.hasOwnProperty('preferredEdgeNode2')) {
                obj['preferredEdgeNode2'] = ApiClient.convertToType(data['preferredEdgeNode2'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Edge Cluster
 * @member {String} edgeCluster
 */
NetworkDhcpServerCreateConfig.prototype['edgeCluster'] = undefined;

/**
 * Active Edge Node Options obtained by calling option source with :optionSource = nsxtEdgeNodes and networkServerId param
 * @member {String} preferredEdgeNode1
 */
NetworkDhcpServerCreateConfig.prototype['preferredEdgeNode1'] = undefined;

/**
 * Standby Edge Node Options obtained by calling option source with optionSource = nsxtEdgeNodes and networkServerId param
 * @member {String} preferredEdgeNode2
 */
NetworkDhcpServerCreateConfig.prototype['preferredEdgeNode2'] = undefined;






export default NetworkDhcpServerCreateConfig;

