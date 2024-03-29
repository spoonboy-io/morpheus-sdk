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
import NetworkRoutersCreateNetworkServer from './NetworkRoutersCreateNetworkServer';
import NetworkRoutersCreateSite from './NetworkRoutersCreateSite';
import NetworkRoutersCreateType from './NetworkRoutersCreateType';
import NetworkRoutersCreateZone from './NetworkRoutersCreateZone';

/**
 * The NetworkRoutersCreate model module.
 * @module model/NetworkRoutersCreate
 * @version 6.2.1
 */
class NetworkRoutersCreate {
    /**
     * Constructs a new <code>NetworkRoutersCreate</code>.
     * @alias module:model/NetworkRoutersCreate
     * @param name {String} Name
     * @param type {module:model/NetworkRoutersCreateType} 
     * @param site {module:model/NetworkRoutersCreateSite} 
     */
    constructor(name, type, site) { 
        
        NetworkRoutersCreate.initialize(this, name, type, site);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, type, site) { 
        obj['name'] = name;
        obj['type'] = type;
        obj['site'] = site;
    }

    /**
     * Constructs a <code>NetworkRoutersCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkRoutersCreate} obj Optional instance to populate.
     * @return {module:model/NetworkRoutersCreate} The populated <code>NetworkRoutersCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkRoutersCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = NetworkRoutersCreateType.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('site')) {
                obj['site'] = NetworkRoutersCreateSite.constructFromObject(data['site']);
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = NetworkRoutersCreateZone.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('networkServer')) {
                obj['networkServer'] = NetworkRoutersCreateNetworkServer.constructFromObject(data['networkServer']);
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
NetworkRoutersCreate.prototype['name'] = undefined;

/**
 * @member {module:model/NetworkRoutersCreateType} type
 */
NetworkRoutersCreate.prototype['type'] = undefined;

/**
 * @member {module:model/NetworkRoutersCreateSite} site
 */
NetworkRoutersCreate.prototype['site'] = undefined;

/**
 * Can be used to enable / disable the network router (true, false). Default is on
 * @member {Boolean} enabled
 */
NetworkRoutersCreate.prototype['enabled'] = undefined;

/**
 * @member {module:model/NetworkRoutersCreateZone} zone
 */
NetworkRoutersCreate.prototype['zone'] = undefined;

/**
 * @member {module:model/NetworkRoutersCreateNetworkServer} networkServer
 */
NetworkRoutersCreate.prototype['networkServer'] = undefined;






export default NetworkRoutersCreate;

