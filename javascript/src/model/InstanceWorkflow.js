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
import InstanceWorkflowTaskSet from './InstanceWorkflowTaskSet';

/**
 * The InstanceWorkflow model module.
 * @module model/InstanceWorkflow
 * @version 6.2.1
 */
class InstanceWorkflow {
    /**
     * Constructs a new <code>InstanceWorkflow</code>.
     * @alias module:model/InstanceWorkflow
     */
    constructor() { 
        
        InstanceWorkflow.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceWorkflow</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceWorkflow} obj Optional instance to populate.
     * @return {module:model/InstanceWorkflow} The populated <code>InstanceWorkflow</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceWorkflow();

            if (data.hasOwnProperty('taskSet')) {
                obj['taskSet'] = InstanceWorkflowTaskSet.constructFromObject(data['taskSet']);
            }
            if (data.hasOwnProperty('taskPhase')) {
                obj['taskPhase'] = ApiClient.convertToType(data['taskPhase'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/InstanceWorkflowTaskSet} taskSet
 */
InstanceWorkflow.prototype['taskSet'] = undefined;

/**
 * Task Phase to run for Provisioning workflows. The default is `provision`.
 * @member {module:model/InstanceWorkflow.TaskPhaseEnum} taskPhase
 * @default 'provision'
 */
InstanceWorkflow.prototype['taskPhase'] = 'provision';





/**
 * Allowed values for the <code>taskPhase</code> property.
 * @enum {String}
 * @readonly
 */
InstanceWorkflow['TaskPhaseEnum'] = {

    /**
     * value: "start"
     * @const
     */
    "start": "start",

    /**
     * value: "stop"
     * @const
     */
    "stop": "stop",

    /**
     * value: "preProvision"
     * @const
     */
    "preProvision": "preProvision",

    /**
     * value: "provision"
     * @const
     */
    "provision": "provision",

    /**
     * value: "postProvision"
     * @const
     */
    "postProvision": "postProvision",

    /**
     * value: "preDeploy"
     * @const
     */
    "preDeploy": "preDeploy",

    /**
     * value: "deploy"
     * @const
     */
    "deploy": "deploy",

    /**
     * value: "reconfigure"
     * @const
     */
    "reconfigure": "reconfigure",

    /**
     * value: "teardown"
     * @const
     */
    "teardown": "teardown",

    /**
     * value: "startup"
     * @const
     */
    "startup": "startup",

    /**
     * value: "shutdown"
     * @const
     */
    "shutdown": "shutdown"
};



export default InstanceWorkflow;

