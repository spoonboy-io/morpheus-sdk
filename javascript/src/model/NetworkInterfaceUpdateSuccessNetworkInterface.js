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
 * The NetworkInterfaceUpdateSuccessNetworkInterface model module.
 * @module model/NetworkInterfaceUpdateSuccessNetworkInterface
 * @version 6.2.1
 */
class NetworkInterfaceUpdateSuccessNetworkInterface {
    /**
     * Constructs a new <code>NetworkInterfaceUpdateSuccessNetworkInterface</code>.
     * @alias module:model/NetworkInterfaceUpdateSuccessNetworkInterface
     */
    constructor() { 
        
        NetworkInterfaceUpdateSuccessNetworkInterface.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NetworkInterfaceUpdateSuccessNetworkInterface</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkInterfaceUpdateSuccessNetworkInterface} obj Optional instance to populate.
     * @return {module:model/NetworkInterfaceUpdateSuccessNetworkInterface} The populated <code>NetworkInterfaceUpdateSuccessNetworkInterface</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkInterfaceUpdateSuccessNetworkInterface();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('addresses')) {
                obj['addresses'] = ApiClient.convertToType(data['addresses'], [Object]);
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('interfaceId')) {
                obj['interfaceId'] = ApiClient.convertToType(data['interfaceId'], 'String');
            }
            if (data.hasOwnProperty('displayOrder')) {
                obj['displayOrder'] = ApiClient.convertToType(data['displayOrder'], 'Number');
            }
            if (data.hasOwnProperty('networkPool')) {
                obj['networkPool'] = ApiClient.convertToType(data['networkPool'], Object);
            }
            if (data.hasOwnProperty('dhcp')) {
                obj['dhcp'] = ApiClient.convertToType(data['dhcp'], 'Boolean');
            }
            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('uniqueId')) {
                obj['uniqueId'] = ApiClient.convertToType(data['uniqueId'], 'String');
            }
            if (data.hasOwnProperty('subnet')) {
                obj['subnet'] = ApiClient.convertToType(data['subnet'], 'String');
            }
            if (data.hasOwnProperty('replaceHostRecord')) {
                obj['replaceHostRecord'] = ApiClient.convertToType(data['replaceHostRecord'], 'Boolean');
            }
            if (data.hasOwnProperty('ipMode')) {
                obj['ipMode'] = ApiClient.convertToType(data['ipMode'], 'String');
            }
            if (data.hasOwnProperty('version')) {
                obj['version'] = ApiClient.convertToType(data['version'], 'String');
            }
            if (data.hasOwnProperty('ipSubnet')) {
                obj['ipSubnet'] = ApiClient.convertToType(data['ipSubnet'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], 'String');
            }
            if (data.hasOwnProperty('publicIpAddress')) {
                obj['publicIpAddress'] = ApiClient.convertToType(data['publicIpAddress'], 'String');
            }
            if (data.hasOwnProperty('fabricId')) {
                obj['fabricId'] = ApiClient.convertToType(data['fabricId'], 'String');
            }
            if (data.hasOwnProperty('ipv6Subnet')) {
                obj['ipv6Subnet'] = ApiClient.convertToType(data['ipv6Subnet'], 'String');
            }
            if (data.hasOwnProperty('macAddress')) {
                obj['macAddress'] = ApiClient.convertToType(data['macAddress'], 'String');
            }
            if (data.hasOwnProperty('publicIpv6Address')) {
                obj['publicIpv6Address'] = ApiClient.convertToType(data['publicIpv6Address'], 'String');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('networkGroup')) {
                obj['networkGroup'] = ApiClient.convertToType(data['networkGroup'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'String');
            }
            if (data.hasOwnProperty('networkDomain')) {
                obj['networkDomain'] = ApiClient.convertToType(data['networkDomain'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('primaryInterface')) {
                obj['primaryInterface'] = ApiClient.convertToType(data['primaryInterface'], 'Boolean');
            }
            if (data.hasOwnProperty('networkPoolIPv6')) {
                obj['networkPoolIPv6'] = ApiClient.convertToType(data['networkPoolIPv6'], Object);
            }
            if (data.hasOwnProperty('network')) {
                obj['network'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['network']);
            }
            if (data.hasOwnProperty('vlanId')) {
                obj['vlanId'] = ApiClient.convertToType(data['vlanId'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('networkPosition')) {
                obj['networkPosition'] = ApiClient.convertToType(data['networkPosition'], 'String');
            }
            if (data.hasOwnProperty('poolAssigned')) {
                obj['poolAssigned'] = ApiClient.convertToType(data['poolAssigned'], 'Boolean');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('externalType')) {
                obj['externalType'] = ApiClient.convertToType(data['externalType'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['id'] = undefined;

/**
 * @member {Array.<Object>} addresses
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['addresses'] = undefined;

/**
 * @member {String} internalId
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['internalId'] = undefined;

/**
 * @member {String} interfaceId
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['interfaceId'] = undefined;

/**
 * @member {Number} displayOrder
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['displayOrder'] = undefined;

/**
 * @member {Object} networkPool
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['networkPool'] = undefined;

/**
 * @member {Boolean} dhcp
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['dhcp'] = undefined;

/**
 * @member {String} uuid
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['uuid'] = undefined;

/**
 * @member {Boolean} active
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['active'] = undefined;

/**
 * @member {String} uniqueId
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['uniqueId'] = undefined;

/**
 * @member {String} subnet
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['subnet'] = undefined;

/**
 * @member {Boolean} replaceHostRecord
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['replaceHostRecord'] = undefined;

/**
 * @member {String} ipMode
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['ipMode'] = undefined;

/**
 * @member {String} version
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['version'] = undefined;

/**
 * @member {String} ipSubnet
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['ipSubnet'] = undefined;

/**
 * @member {String} config
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['config'] = undefined;

/**
 * @member {String} publicIpAddress
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['publicIpAddress'] = undefined;

/**
 * @member {String} fabricId
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['fabricId'] = undefined;

/**
 * @member {String} ipv6Subnet
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['ipv6Subnet'] = undefined;

/**
 * @member {String} macAddress
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['macAddress'] = undefined;

/**
 * @member {String} publicIpv6Address
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['publicIpv6Address'] = undefined;

/**
 * @member {String} refType
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['refType'] = undefined;

/**
 * @member {String} networkGroup
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['networkGroup'] = undefined;

/**
 * @member {String} refId
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['refId'] = undefined;

/**
 * @member {String} networkDomain
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['networkDomain'] = undefined;

/**
 * @member {String} name
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['name'] = undefined;

/**
 * @member {Boolean} primaryInterface
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['primaryInterface'] = undefined;

/**
 * @member {Object} networkPoolIPv6
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['networkPoolIPv6'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} network
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['network'] = undefined;

/**
 * @member {String} vlanId
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['vlanId'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} type
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['type'] = undefined;

/**
 * @member {String} networkPosition
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['networkPosition'] = undefined;

/**
 * @member {Boolean} poolAssigned
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['poolAssigned'] = undefined;

/**
 * @member {String} description
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['description'] = undefined;

/**
 * @member {String} externalType
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['externalType'] = undefined;

/**
 * @member {String} externalId
 */
NetworkInterfaceUpdateSuccessNetworkInterface.prototype['externalId'] = undefined;






export default NetworkInterfaceUpdateSuccessNetworkInterface;
