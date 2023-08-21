<?php
/**
 * InstanceWorkflow
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * InstanceWorkflow Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InstanceWorkflow implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'instanceWorkflow';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'task_set' => '\OpenAPI\Client\Model\InstanceWorkflowTaskSet',
        'task_phase' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'task_set' => null,
        'task_phase' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'task_set' => 'taskSet',
        'task_phase' => 'taskPhase'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'task_set' => 'setTaskSet',
        'task_phase' => 'setTaskPhase'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'task_set' => 'getTaskSet',
        'task_phase' => 'getTaskPhase'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    const TASK_PHASE_START = 'start';
    const TASK_PHASE_STOP = 'stop';
    const TASK_PHASE_PRE_PROVISION = 'preProvision';
    const TASK_PHASE_PROVISION = 'provision';
    const TASK_PHASE_POST_PROVISION = 'postProvision';
    const TASK_PHASE_PRE_DEPLOY = 'preDeploy';
    const TASK_PHASE_DEPLOY = 'deploy';
    const TASK_PHASE_RECONFIGURE = 'reconfigure';
    const TASK_PHASE_TEARDOWN = 'teardown';
    const TASK_PHASE_STARTUP = 'startup';
    const TASK_PHASE_SHUTDOWN = 'shutdown';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getTaskPhaseAllowableValues()
    {
        return [
            self::TASK_PHASE_START,
            self::TASK_PHASE_STOP,
            self::TASK_PHASE_PRE_PROVISION,
            self::TASK_PHASE_PROVISION,
            self::TASK_PHASE_POST_PROVISION,
            self::TASK_PHASE_PRE_DEPLOY,
            self::TASK_PHASE_DEPLOY,
            self::TASK_PHASE_RECONFIGURE,
            self::TASK_PHASE_TEARDOWN,
            self::TASK_PHASE_STARTUP,
            self::TASK_PHASE_SHUTDOWN,
        ];
    }
    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['task_set'] = $data['task_set'] ?? null;
        $this->container['task_phase'] = $data['task_phase'] ?? 'provision';
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        $allowedValues = $this->getTaskPhaseAllowableValues();
        if (!is_null($this->container['task_phase']) && !in_array($this->container['task_phase'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'task_phase', must be one of '%s'",
                $this->container['task_phase'],
                implode("', '", $allowedValues)
            );
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets task_set
     *
     * @return \OpenAPI\Client\Model\InstanceWorkflowTaskSet|null
     */
    public function getTaskSet()
    {
        return $this->container['task_set'];
    }

    /**
     * Sets task_set
     *
     * @param \OpenAPI\Client\Model\InstanceWorkflowTaskSet|null $task_set task_set
     *
     * @return self
     */
    public function setTaskSet($task_set)
    {
        $this->container['task_set'] = $task_set;

        return $this;
    }

    /**
     * Gets task_phase
     *
     * @return string|null
     */
    public function getTaskPhase()
    {
        return $this->container['task_phase'];
    }

    /**
     * Sets task_phase
     *
     * @param string|null $task_phase Task Phase to run for Provisioning workflows. The default is `provision`.
     *
     * @return self
     */
    public function setTaskPhase($task_phase)
    {
        $allowedValues = $this->getTaskPhaseAllowableValues();
        if (!is_null($task_phase) && !in_array($task_phase, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'task_phase', must be one of '%s'",
                    $task_phase,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['task_phase'] = $task_phase;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


