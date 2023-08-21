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
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The ClusterMastersInterfaces model module.
 * @module model/ClusterMastersInterfaces
 * @version 6.2.1
 */
class ClusterMastersInterfaces {
    /**
     * Constructs a new <code>ClusterMastersInterfaces</code>.
     * @alias module:model/ClusterMastersInterfaces
     */
    constructor() { 
        
        ClusterMastersInterfaces.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterMastersInterfaces</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterMastersInterfaces} obj Optional instance to populate.
     * @return {module:model/ClusterMastersInterfaces} The populated <code>ClusterMastersInterfaces</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterMastersInterfaces();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('uniqueId')) {
                obj['uniqueId'] = ApiClient.convertToType(data['uniqueId'], 'String');
            }
            if (data.hasOwnProperty('publicIpAddress')) {
                obj['publicIpAddress'] = ApiClient.convertToType(data['publicIpAddress'], 'String');
            }
            if (data.hasOwnProperty('publicIpv6Address')) {
                obj['publicIpv6Address'] = ApiClient.convertToType(data['publicIpv6Address'], 'String');
            }
            if (data.hasOwnProperty('ipAddress')) {
                obj['ipAddress'] = ApiClient.convertToType(data['ipAddress'], 'String');
            }
            if (data.hasOwnProperty('ipv6Address')) {
                obj['ipv6Address'] = ApiClient.convertToType(data['ipv6Address'], 'String');
            }
            if (data.hasOwnProperty('ipSubnet')) {
                obj['ipSubnet'] = ApiClient.convertToType(data['ipSubnet'], 'String');
            }
            if (data.hasOwnProperty('ipv6Subnet')) {
                obj['ipv6Subnet'] = ApiClient.convertToType(data['ipv6Subnet'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('dhcp')) {
                obj['dhcp'] = ApiClient.convertToType(data['dhcp'], 'Boolean');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('poolAssigned')) {
                obj['poolAssigned'] = ApiClient.convertToType(data['poolAssigned'], 'Boolean');
            }
            if (data.hasOwnProperty('primaryInterface')) {
                obj['primaryInterface'] = ApiClient.convertToType(data['primaryInterface'], 'Boolean');
            }
            if (data.hasOwnProperty('network')) {
                obj['network'] = InlineResponse20040AppDeployInstance.constructFromObject(data['network']);
            }
            if (data.hasOwnProperty('subnet')) {
                obj['subnet'] = ApiClient.convertToType(data['subnet'], 'String');
            }
            if (data.hasOwnProperty('networkGroup')) {
                obj['networkGroup'] = ApiClient.convertToType(data['networkGroup'], 'String');
            }
            if (data.hasOwnProperty('networkPosition')) {
                obj['networkPosition'] = ApiClient.convertToType(data['networkPosition'], 'String');
            }
            if (data.hasOwnProperty('networkPool')) {
                obj['networkPool'] = ApiClient.convertToType(data['networkPool'], 'String');
            }
            if (data.hasOwnProperty('networkDomain')) {
                obj['networkDomain'] = ApiClient.convertToType(data['networkDomain'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('ipMode')) {
                obj['ipMode'] = ApiClient.convertToType(data['ipMode'], 'String');
            }
            if (data.hasOwnProperty('macAddress')) {
                obj['macAddress'] = ApiClient.convertToType(data['macAddress'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ClusterMastersInterfaces.prototype['id'] = undefined;

/**
 * @member {String} refType
 */
ClusterMastersInterfaces.prototype['refType'] = undefined;

/**
 * @member {String} refId
 */
ClusterMastersInterfaces.prototype['refId'] = undefined;

/**
 * @member {String} name
 */
ClusterMastersInterfaces.prototype['name'] = undefined;

/**
 * @member {String} internalId
 */
ClusterMastersInterfaces.prototype['internalId'] = undefined;

/**
 * @member {String} externalId
 */
ClusterMastersInterfaces.prototype['externalId'] = undefined;

/**
 * @member {String} uniqueId
 */
ClusterMastersInterfaces.prototype['uniqueId'] = undefined;

/**
 * @member {String} publicIpAddress
 */
ClusterMastersInterfaces.prototype['publicIpAddress'] = undefined;

/**
 * @member {String} publicIpv6Address
 */
ClusterMastersInterfaces.prototype['publicIpv6Address'] = undefined;

/**
 * @member {String} ipAddress
 */
ClusterMastersInterfaces.prototype['ipAddress'] = undefined;

/**
 * @member {String} ipv6Address
 */
ClusterMastersInterfaces.prototype['ipv6Address'] = undefined;

/**
 * @member {String} ipSubnet
 */
ClusterMastersInterfaces.prototype['ipSubnet'] = undefined;

/**
 * @member {String} ipv6Subnet
 */
ClusterMastersInterfaces.prototype['ipv6Subnet'] = undefined;

/**
 * @member {String} description
 */
ClusterMastersInterfaces.prototype['description'] = undefined;

/**
 * @member {Boolean} dhcp
 */
ClusterMastersInterfaces.prototype['dhcp'] = undefined;

/**
 * @member {Boolean} active
 */
ClusterMastersInterfaces.prototype['active'] = undefined;

/**
 * @member {Boolean} poolAssigned
 */
ClusterMastersInterfaces.prototype['poolAssigned'] = undefined;

/**
 * @member {Boolean} primaryInterface
 */
ClusterMastersInterfaces.prototype['primaryInterface'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} network
 */
ClusterMastersInterfaces.prototype['network'] = undefined;

/**
 * @member {String} subnet
 */
ClusterMastersInterfaces.prototype['subnet'] = undefined;

/**
 * @member {String} networkGroup
 */
ClusterMastersInterfaces.prototype['networkGroup'] = undefined;

/**
 * @member {String} networkPosition
 */
ClusterMastersInterfaces.prototype['networkPosition'] = undefined;

/**
 * @member {String} networkPool
 */
ClusterMastersInterfaces.prototype['networkPool'] = undefined;

/**
 * @member {String} networkDomain
 */
ClusterMastersInterfaces.prototype['networkDomain'] = undefined;

/**
 * @member {String} type
 */
ClusterMastersInterfaces.prototype['type'] = undefined;

/**
 * @member {String} ipMode
 */
ClusterMastersInterfaces.prototype['ipMode'] = undefined;

/**
 * @member {String} macAddress
 */
ClusterMastersInterfaces.prototype['macAddress'] = undefined;






export default ClusterMastersInterfaces;

