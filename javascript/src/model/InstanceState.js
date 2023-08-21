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
import AppStateOutput from './AppStateOutput';
import InstanceStateInput from './InstanceStateInput';

/**
 * The InstanceState model module.
 * @module model/InstanceState
 * @version 6.2.1
 */
class InstanceState {
    /**
     * Constructs a new <code>InstanceState</code>.
     * @alias module:model/InstanceState
     */
    constructor() { 
        
        InstanceState.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceState</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceState} obj Optional instance to populate.
     * @return {module:model/InstanceState} The populated <code>InstanceState</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceState();

            if (data.hasOwnProperty('workloads')) {
                obj['workloads'] = ApiClient.convertToType(data['workloads'], [Object]);
            }
            if (data.hasOwnProperty('iacDrift')) {
                obj['iacDrift'] = ApiClient.convertToType(data['iacDrift'], 'Boolean');
            }
            if (data.hasOwnProperty('planResources')) {
                obj['planResources'] = ApiClient.convertToType(data['planResources'], [Object]);
            }
            if (data.hasOwnProperty('specs')) {
                obj['specs'] = ApiClient.convertToType(data['specs'], [Object]);
            }
            if (data.hasOwnProperty('stateData')) {
                obj['stateData'] = ApiClient.convertToType(data['stateData'], 'String');
            }
            if (data.hasOwnProperty('planData')) {
                obj['planData'] = ApiClient.convertToType(data['planData'], 'String');
            }
            if (data.hasOwnProperty('input')) {
                obj['input'] = InstanceStateInput.constructFromObject(data['input']);
            }
            if (data.hasOwnProperty('output')) {
                obj['output'] = AppStateOutput.constructFromObject(data['output']);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<Object>} workloads
 */
InstanceState.prototype['workloads'] = undefined;

/**
 * @member {Boolean} iacDrift
 */
InstanceState.prototype['iacDrift'] = undefined;

/**
 * @member {Array.<Object>} planResources
 */
InstanceState.prototype['planResources'] = undefined;

/**
 * @member {Array.<Object>} specs
 */
InstanceState.prototype['specs'] = undefined;

/**
 * @member {String} stateData
 */
InstanceState.prototype['stateData'] = undefined;

/**
 * @member {String} planData
 */
InstanceState.prototype['planData'] = undefined;

/**
 * @member {module:model/InstanceStateInput} input
 */
InstanceState.prototype['input'] = undefined;

/**
 * @member {module:model/AppStateOutput} output
 */
InstanceState.prototype['output'] = undefined;






export default InstanceState;

