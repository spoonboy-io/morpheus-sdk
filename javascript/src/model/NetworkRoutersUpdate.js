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
import NetworkRoutersUpdateNetworkServer from './NetworkRoutersUpdateNetworkServer';
import NetworkRoutersUpdateSite from './NetworkRoutersUpdateSite';
import NetworkRoutersUpdateType from './NetworkRoutersUpdateType';
import NetworkRoutersUpdateZone from './NetworkRoutersUpdateZone';

/**
 * The NetworkRoutersUpdate model module.
 * @module model/NetworkRoutersUpdate
 * @version 6.2.1
 */
class NetworkRoutersUpdate {
    /**
     * Constructs a new <code>NetworkRoutersUpdate</code>.
     * @alias module:model/NetworkRoutersUpdate
     */
    constructor() { 
        
        NetworkRoutersUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NetworkRoutersUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkRoutersUpdate} obj Optional instance to populate.
     * @return {module:model/NetworkRoutersUpdate} The populated <code>NetworkRoutersUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkRoutersUpdate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = NetworkRoutersUpdateType.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('site')) {
                obj['site'] = NetworkRoutersUpdateSite.constructFromObject(data['site']);
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = NetworkRoutersUpdateZone.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('networkServer')) {
                obj['networkServer'] = NetworkRoutersUpdateNetworkServer.constructFromObject(data['networkServer']);
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
NetworkRoutersUpdate.prototype['name'] = undefined;

/**
 * @member {module:model/NetworkRoutersUpdateType} type
 */
NetworkRoutersUpdate.prototype['type'] = undefined;

/**
 * @member {module:model/NetworkRoutersUpdateSite} site
 */
NetworkRoutersUpdate.prototype['site'] = undefined;

/**
 * Can be used to enable / disable the network router (true, false). Default is on
 * @member {Boolean} enabled
 */
NetworkRoutersUpdate.prototype['enabled'] = undefined;

/**
 * @member {module:model/NetworkRoutersUpdateZone} zone
 */
NetworkRoutersUpdate.prototype['zone'] = undefined;

/**
 * @member {module:model/NetworkRoutersUpdateNetworkServer} networkServer
 */
NetworkRoutersUpdate.prototype['networkServer'] = undefined;






export default NetworkRoutersUpdate;

