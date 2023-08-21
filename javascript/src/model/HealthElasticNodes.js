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
 * The HealthElasticNodes model module.
 * @module model/HealthElasticNodes
 * @version 6.2.1
 */
class HealthElasticNodes {
    /**
     * Constructs a new <code>HealthElasticNodes</code>.
     * @alias module:model/HealthElasticNodes
     */
    constructor() { 
        
        HealthElasticNodes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>HealthElasticNodes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/HealthElasticNodes} obj Optional instance to populate.
     * @return {module:model/HealthElasticNodes} The populated <code>HealthElasticNodes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new HealthElasticNodes();

            if (data.hasOwnProperty('ip')) {
                obj['ip'] = ApiClient.convertToType(data['ip'], 'String');
            }
            if (data.hasOwnProperty('heapPercent')) {
                obj['heapPercent'] = ApiClient.convertToType(data['heapPercent'], 'String');
            }
            if (data.hasOwnProperty('ramPercent')) {
                obj['ramPercent'] = ApiClient.convertToType(data['ramPercent'], 'String');
            }
            if (data.hasOwnProperty('cpuCount')) {
                obj['cpuCount'] = ApiClient.convertToType(data['cpuCount'], 'String');
            }
            if (data.hasOwnProperty('loadOne')) {
                obj['loadOne'] = ApiClient.convertToType(data['loadOne'], 'String');
            }
            if (data.hasOwnProperty('loadFive')) {
                obj['loadFive'] = ApiClient.convertToType(data['loadFive'], 'String');
            }
            if (data.hasOwnProperty('loadFifteen')) {
                obj['loadFifteen'] = ApiClient.convertToType(data['loadFifteen'], 'String');
            }
            if (data.hasOwnProperty('role')) {
                obj['role'] = ApiClient.convertToType(data['role'], 'String');
            }
            if (data.hasOwnProperty('master')) {
                obj['master'] = ApiClient.convertToType(data['master'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} ip
 */
HealthElasticNodes.prototype['ip'] = undefined;

/**
 * @member {String} heapPercent
 */
HealthElasticNodes.prototype['heapPercent'] = undefined;

/**
 * @member {String} ramPercent
 */
HealthElasticNodes.prototype['ramPercent'] = undefined;

/**
 * @member {String} cpuCount
 */
HealthElasticNodes.prototype['cpuCount'] = undefined;

/**
 * @member {String} loadOne
 */
HealthElasticNodes.prototype['loadOne'] = undefined;

/**
 * @member {String} loadFive
 */
HealthElasticNodes.prototype['loadFive'] = undefined;

/**
 * @member {String} loadFifteen
 */
HealthElasticNodes.prototype['loadFifteen'] = undefined;

/**
 * @member {String} role
 */
HealthElasticNodes.prototype['role'] = undefined;

/**
 * @member {String} master
 */
HealthElasticNodes.prototype['master'] = undefined;

/**
 * @member {String} name
 */
HealthElasticNodes.prototype['name'] = undefined;






export default HealthElasticNodes;

