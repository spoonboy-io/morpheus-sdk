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
import InlineResponse200106NetworkPoolIpRanges from './InlineResponse200106NetworkPoolIpRanges';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InlineResponse20094Network from './InlineResponse20094Network';

/**
 * The InlineResponse200106NetworkPool model module.
 * @module model/InlineResponse200106NetworkPool
 * @version 6.2.1
 */
class InlineResponse200106NetworkPool {
    /**
     * Constructs a new <code>InlineResponse200106NetworkPool</code>.
     * @alias module:model/InlineResponse200106NetworkPool
     */
    constructor() { 
        
        InlineResponse200106NetworkPool.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200106NetworkPool</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200106NetworkPool} obj Optional instance to populate.
     * @return {module:model/InlineResponse200106NetworkPool} The populated <code>InlineResponse200106NetworkPool</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200106NetworkPool();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = InlineResponse20094Network.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20040AppDeployInstance.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('displayName')) {
                obj['displayName'] = ApiClient.convertToType(data['displayName'], 'String');
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('dnsDomain')) {
                obj['dnsDomain'] = ApiClient.convertToType(data['dnsDomain'], 'String');
            }
            if (data.hasOwnProperty('dnsSearchPath')) {
                obj['dnsSearchPath'] = ApiClient.convertToType(data['dnsSearchPath'], 'String');
            }
            if (data.hasOwnProperty('hostPrefix')) {
                obj['hostPrefix'] = ApiClient.convertToType(data['hostPrefix'], 'String');
            }
            if (data.hasOwnProperty('httpProxy')) {
                obj['httpProxy'] = ApiClient.convertToType(data['httpProxy'], 'String');
            }
            if (data.hasOwnProperty('dnsServers')) {
                obj['dnsServers'] = ApiClient.convertToType(data['dnsServers'], ['String']);
            }
            if (data.hasOwnProperty('dnsSuffixList')) {
                obj['dnsSuffixList'] = ApiClient.convertToType(data['dnsSuffixList'], ['String']);
            }
            if (data.hasOwnProperty('dhcpServer')) {
                obj['dhcpServer'] = ApiClient.convertToType(data['dhcpServer'], 'Boolean');
            }
            if (data.hasOwnProperty('dhcpIp')) {
                obj['dhcpIp'] = ApiClient.convertToType(data['dhcpIp'], 'String');
            }
            if (data.hasOwnProperty('gateway')) {
                obj['gateway'] = ApiClient.convertToType(data['gateway'], 'String');
            }
            if (data.hasOwnProperty('netmask')) {
                obj['netmask'] = ApiClient.convertToType(data['netmask'], 'String');
            }
            if (data.hasOwnProperty('subnetAddress')) {
                obj['subnetAddress'] = ApiClient.convertToType(data['subnetAddress'], 'String');
            }
            if (data.hasOwnProperty('ipCount')) {
                obj['ipCount'] = ApiClient.convertToType(data['ipCount'], 'Number');
            }
            if (data.hasOwnProperty('freeCount')) {
                obj['freeCount'] = ApiClient.convertToType(data['freeCount'], 'Number');
            }
            if (data.hasOwnProperty('poolEnabled')) {
                obj['poolEnabled'] = ApiClient.convertToType(data['poolEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('tftpServer')) {
                obj['tftpServer'] = ApiClient.convertToType(data['tftpServer'], 'String');
            }
            if (data.hasOwnProperty('bootFile')) {
                obj['bootFile'] = ApiClient.convertToType(data['bootFile'], 'String');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'String');
            }
            if (data.hasOwnProperty('parentType')) {
                obj['parentType'] = ApiClient.convertToType(data['parentType'], 'String');
            }
            if (data.hasOwnProperty('parentId')) {
                obj['parentId'] = ApiClient.convertToType(data['parentId'], 'String');
            }
            if (data.hasOwnProperty('poolGroup')) {
                obj['poolGroup'] = ApiClient.convertToType(data['poolGroup'], 'String');
            }
            if (data.hasOwnProperty('ipRanges')) {
                obj['ipRanges'] = ApiClient.convertToType(data['ipRanges'], [InlineResponse200106NetworkPoolIpRanges]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse200106NetworkPool.prototype['id'] = undefined;

/**
 * @member {module:model/InlineResponse20094Network} type
 */
InlineResponse200106NetworkPool.prototype['type'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
InlineResponse200106NetworkPool.prototype['account'] = undefined;

/**
 * @member {String} category
 */
InlineResponse200106NetworkPool.prototype['category'] = undefined;

/**
 * @member {String} code
 */
InlineResponse200106NetworkPool.prototype['code'] = undefined;

/**
 * @member {String} name
 */
InlineResponse200106NetworkPool.prototype['name'] = undefined;

/**
 * @member {String} displayName
 */
InlineResponse200106NetworkPool.prototype['displayName'] = undefined;

/**
 * @member {String} internalId
 */
InlineResponse200106NetworkPool.prototype['internalId'] = undefined;

/**
 * @member {String} externalId
 */
InlineResponse200106NetworkPool.prototype['externalId'] = undefined;

/**
 * @member {String} dnsDomain
 */
InlineResponse200106NetworkPool.prototype['dnsDomain'] = undefined;

/**
 * @member {String} dnsSearchPath
 */
InlineResponse200106NetworkPool.prototype['dnsSearchPath'] = undefined;

/**
 * @member {String} hostPrefix
 */
InlineResponse200106NetworkPool.prototype['hostPrefix'] = undefined;

/**
 * @member {String} httpProxy
 */
InlineResponse200106NetworkPool.prototype['httpProxy'] = undefined;

/**
 * @member {Array.<String>} dnsServers
 */
InlineResponse200106NetworkPool.prototype['dnsServers'] = undefined;

/**
 * @member {Array.<String>} dnsSuffixList
 */
InlineResponse200106NetworkPool.prototype['dnsSuffixList'] = undefined;

/**
 * @member {Boolean} dhcpServer
 */
InlineResponse200106NetworkPool.prototype['dhcpServer'] = undefined;

/**
 * @member {String} dhcpIp
 */
InlineResponse200106NetworkPool.prototype['dhcpIp'] = undefined;

/**
 * @member {String} gateway
 */
InlineResponse200106NetworkPool.prototype['gateway'] = undefined;

/**
 * @member {String} netmask
 */
InlineResponse200106NetworkPool.prototype['netmask'] = undefined;

/**
 * @member {String} subnetAddress
 */
InlineResponse200106NetworkPool.prototype['subnetAddress'] = undefined;

/**
 * @member {Number} ipCount
 */
InlineResponse200106NetworkPool.prototype['ipCount'] = undefined;

/**
 * @member {Number} freeCount
 */
InlineResponse200106NetworkPool.prototype['freeCount'] = undefined;

/**
 * @member {Boolean} poolEnabled
 */
InlineResponse200106NetworkPool.prototype['poolEnabled'] = undefined;

/**
 * @member {String} tftpServer
 */
InlineResponse200106NetworkPool.prototype['tftpServer'] = undefined;

/**
 * @member {String} bootFile
 */
InlineResponse200106NetworkPool.prototype['bootFile'] = undefined;

/**
 * @member {String} refType
 */
InlineResponse200106NetworkPool.prototype['refType'] = undefined;

/**
 * @member {String} refId
 */
InlineResponse200106NetworkPool.prototype['refId'] = undefined;

/**
 * @member {String} parentType
 */
InlineResponse200106NetworkPool.prototype['parentType'] = undefined;

/**
 * @member {String} parentId
 */
InlineResponse200106NetworkPool.prototype['parentId'] = undefined;

/**
 * @member {String} poolGroup
 */
InlineResponse200106NetworkPool.prototype['poolGroup'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse200106NetworkPoolIpRanges>} ipRanges
 */
InlineResponse200106NetworkPool.prototype['ipRanges'] = undefined;






export default InlineResponse200106NetworkPool;

