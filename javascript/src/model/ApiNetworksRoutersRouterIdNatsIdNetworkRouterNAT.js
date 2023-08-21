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
 * The ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT model module.
 * @module model/ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT
 * @version 6.2.1
 */
class ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT {
    /**
     * Constructs a new <code>ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT</code>.
     * For a full list of available NAT options, see natOptionTypes in the specific Network Router Type
     * @alias module:model/ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT
     */
    constructor() { 
        
        ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT} obj Optional instance to populate.
     * @return {module:model/ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT} The populated <code>ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], Object);
            }
        }
        return obj;
    }


}

/**
 * Sets name of NAT
 * @member {Object} name
 */
ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT.prototype['name'] = undefined;






export default ApiNetworksRoutersRouterIdNatsIdNetworkRouterNAT;

