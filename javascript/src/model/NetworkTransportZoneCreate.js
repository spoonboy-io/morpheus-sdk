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
import ApiBlueprintsIdUpdatePermissionsResourcePermissionSites from './ApiBlueprintsIdUpdatePermissionsResourcePermissionSites';

/**
 * The NetworkTransportZoneCreate model module.
 * @module model/NetworkTransportZoneCreate
 * @version 6.2.1
 */
class NetworkTransportZoneCreate {
    /**
     * Constructs a new <code>NetworkTransportZoneCreate</code>.
     * @alias module:model/NetworkTransportZoneCreate
     * @param name {String} Network transport zone name
     */
    constructor(name) { 
        
        NetworkTransportZoneCreate.initialize(this, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name) { 
        obj['name'] = name;
    }

    /**
     * Constructs a <code>NetworkTransportZoneCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkTransportZoneCreate} obj Optional instance to populate.
     * @return {module:model/NetworkTransportZoneCreate} The populated <code>NetworkTransportZoneCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkTransportZoneCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], [ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]);
            }
        }
        return obj;
    }


}

/**
 * Network transport zone name
 * @member {String} name
 */
NetworkTransportZoneCreate.prototype['name'] = undefined;

/**
 * Network transport zone description
 * @member {String} description
 */
NetworkTransportZoneCreate.prototype['description'] = undefined;

/**
 * private or public
 * @member {String} visibility
 */
NetworkTransportZoneCreate.prototype['visibility'] = undefined;

/**
 * Array of tenant account ids that are allowed access
 * @member {Array.<module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>} tenants
 */
NetworkTransportZoneCreate.prototype['tenants'] = undefined;






export default NetworkTransportZoneCreate;

