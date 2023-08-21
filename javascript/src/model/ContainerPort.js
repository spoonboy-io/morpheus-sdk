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
 * The ContainerPort model module.
 * @module model/ContainerPort
 * @version 6.2.1
 */
class ContainerPort {
    /**
     * Constructs a new <code>ContainerPort</code>.
     * @alias module:model/ContainerPort
     */
    constructor() { 
        
        ContainerPort.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ContainerPort</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ContainerPort} obj Optional instance to populate.
     * @return {module:model/ContainerPort} The populated <code>ContainerPort</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ContainerPort();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('index')) {
                obj['index'] = ApiClient.convertToType(data['index'], 'Number');
            }
            if (data.hasOwnProperty('external')) {
                obj['external'] = ApiClient.convertToType(data['external'], 'Number');
            }
            if (data.hasOwnProperty('internal')) {
                obj['internal'] = ApiClient.convertToType(data['internal'], 'Number');
            }
            if (data.hasOwnProperty('displayName')) {
                obj['displayName'] = ApiClient.convertToType(data['displayName'], 'String');
            }
            if (data.hasOwnProperty('primaryPort')) {
                obj['primaryPort'] = ApiClient.convertToType(data['primaryPort'], 'Boolean');
            }
            if (data.hasOwnProperty('export')) {
                obj['export'] = ApiClient.convertToType(data['export'], 'Boolean');
            }
            if (data.hasOwnProperty('visible')) {
                obj['visible'] = ApiClient.convertToType(data['visible'], 'Boolean');
            }
            if (data.hasOwnProperty('exportName')) {
                obj['exportName'] = ApiClient.convertToType(data['exportName'], 'String');
            }
            if (data.hasOwnProperty('loadBalanceProtocol')) {
                obj['loadBalanceProtocol'] = ApiClient.convertToType(data['loadBalanceProtocol'], 'String');
            }
            if (data.hasOwnProperty('loadBalance')) {
                obj['loadBalance'] = ApiClient.convertToType(data['loadBalance'], 'Boolean');
            }
            if (data.hasOwnProperty('protocol')) {
                obj['protocol'] = ApiClient.convertToType(data['protocol'], 'String');
            }
            if (data.hasOwnProperty('link')) {
                obj['link'] = ApiClient.convertToType(data['link'], 'Boolean');
            }
            if (data.hasOwnProperty('externalIp')) {
                obj['externalIp'] = ApiClient.convertToType(data['externalIp'], 'String');
            }
            if (data.hasOwnProperty('internalIp')) {
                obj['internalIp'] = ApiClient.convertToType(data['internalIp'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ContainerPort.prototype['id'] = undefined;

/**
 * @member {Number} index
 */
ContainerPort.prototype['index'] = undefined;

/**
 * @member {Number} external
 */
ContainerPort.prototype['external'] = undefined;

/**
 * @member {Number} internal
 */
ContainerPort.prototype['internal'] = undefined;

/**
 * @member {String} displayName
 */
ContainerPort.prototype['displayName'] = undefined;

/**
 * @member {Boolean} primaryPort
 */
ContainerPort.prototype['primaryPort'] = undefined;

/**
 * @member {Boolean} export
 */
ContainerPort.prototype['export'] = undefined;

/**
 * @member {Boolean} visible
 */
ContainerPort.prototype['visible'] = undefined;

/**
 * @member {String} exportName
 */
ContainerPort.prototype['exportName'] = undefined;

/**
 * @member {String} loadBalanceProtocol
 */
ContainerPort.prototype['loadBalanceProtocol'] = undefined;

/**
 * @member {Boolean} loadBalance
 */
ContainerPort.prototype['loadBalance'] = undefined;

/**
 * @member {String} protocol
 */
ContainerPort.prototype['protocol'] = undefined;

/**
 * @member {Boolean} link
 */
ContainerPort.prototype['link'] = undefined;

/**
 * @member {String} externalIp
 */
ContainerPort.prototype['externalIp'] = undefined;

/**
 * @member {String} internalIp
 */
ContainerPort.prototype['internalIp'] = undefined;






export default ContainerPort;

