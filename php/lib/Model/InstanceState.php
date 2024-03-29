<?php
/**
 * InstanceState
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
 * InstanceState Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InstanceState implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'instanceState';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'workloads' => 'object[]',
        'iac_drift' => 'bool',
        'plan_resources' => 'object[]',
        'specs' => 'object[]',
        'state_data' => 'string',
        'plan_data' => 'string',
        'input' => '\OpenAPI\Client\Model\InstanceStateInput',
        'output' => '\OpenAPI\Client\Model\AppStateOutput'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'workloads' => null,
        'iac_drift' => null,
        'plan_resources' => null,
        'specs' => null,
        'state_data' => null,
        'plan_data' => null,
        'input' => null,
        'output' => null
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
        'workloads' => 'workloads',
        'iac_drift' => 'iacDrift',
        'plan_resources' => 'planResources',
        'specs' => 'specs',
        'state_data' => 'stateData',
        'plan_data' => 'planData',
        'input' => 'input',
        'output' => 'output'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'workloads' => 'setWorkloads',
        'iac_drift' => 'setIacDrift',
        'plan_resources' => 'setPlanResources',
        'specs' => 'setSpecs',
        'state_data' => 'setStateData',
        'plan_data' => 'setPlanData',
        'input' => 'setInput',
        'output' => 'setOutput'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'workloads' => 'getWorkloads',
        'iac_drift' => 'getIacDrift',
        'plan_resources' => 'getPlanResources',
        'specs' => 'getSpecs',
        'state_data' => 'getStateData',
        'plan_data' => 'getPlanData',
        'input' => 'getInput',
        'output' => 'getOutput'
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
        $this->container['workloads'] = $data['workloads'] ?? null;
        $this->container['iac_drift'] = $data['iac_drift'] ?? null;
        $this->container['plan_resources'] = $data['plan_resources'] ?? null;
        $this->container['specs'] = $data['specs'] ?? null;
        $this->container['state_data'] = $data['state_data'] ?? null;
        $this->container['plan_data'] = $data['plan_data'] ?? null;
        $this->container['input'] = $data['input'] ?? null;
        $this->container['output'] = $data['output'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

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
     * Gets workloads
     *
     * @return object[]|null
     */
    public function getWorkloads()
    {
        return $this->container['workloads'];
    }

    /**
     * Sets workloads
     *
     * @param object[]|null $workloads workloads
     *
     * @return self
     */
    public function setWorkloads($workloads)
    {
        $this->container['workloads'] = $workloads;

        return $this;
    }

    /**
     * Gets iac_drift
     *
     * @return bool|null
     */
    public function getIacDrift()
    {
        return $this->container['iac_drift'];
    }

    /**
     * Sets iac_drift
     *
     * @param bool|null $iac_drift iac_drift
     *
     * @return self
     */
    public function setIacDrift($iac_drift)
    {
        $this->container['iac_drift'] = $iac_drift;

        return $this;
    }

    /**
     * Gets plan_resources
     *
     * @return object[]|null
     */
    public function getPlanResources()
    {
        return $this->container['plan_resources'];
    }

    /**
     * Sets plan_resources
     *
     * @param object[]|null $plan_resources plan_resources
     *
     * @return self
     */
    public function setPlanResources($plan_resources)
    {
        $this->container['plan_resources'] = $plan_resources;

        return $this;
    }

    /**
     * Gets specs
     *
     * @return object[]|null
     */
    public function getSpecs()
    {
        return $this->container['specs'];
    }

    /**
     * Sets specs
     *
     * @param object[]|null $specs specs
     *
     * @return self
     */
    public function setSpecs($specs)
    {
        $this->container['specs'] = $specs;

        return $this;
    }

    /**
     * Gets state_data
     *
     * @return string|null
     */
    public function getStateData()
    {
        return $this->container['state_data'];
    }

    /**
     * Sets state_data
     *
     * @param string|null $state_data state_data
     *
     * @return self
     */
    public function setStateData($state_data)
    {
        $this->container['state_data'] = $state_data;

        return $this;
    }

    /**
     * Gets plan_data
     *
     * @return string|null
     */
    public function getPlanData()
    {
        return $this->container['plan_data'];
    }

    /**
     * Sets plan_data
     *
     * @param string|null $plan_data plan_data
     *
     * @return self
     */
    public function setPlanData($plan_data)
    {
        $this->container['plan_data'] = $plan_data;

        return $this;
    }

    /**
     * Gets input
     *
     * @return \OpenAPI\Client\Model\InstanceStateInput|null
     */
    public function getInput()
    {
        return $this->container['input'];
    }

    /**
     * Sets input
     *
     * @param \OpenAPI\Client\Model\InstanceStateInput|null $input input
     *
     * @return self
     */
    public function setInput($input)
    {
        $this->container['input'] = $input;

        return $this;
    }

    /**
     * Gets output
     *
     * @return \OpenAPI\Client\Model\AppStateOutput|null
     */
    public function getOutput()
    {
        return $this->container['output'];
    }

    /**
     * Sets output
     *
     * @param \OpenAPI\Client\Model\AppStateOutput|null $output output
     *
     * @return self
     */
    public function setOutput($output)
    {
        $this->container['output'] = $output;

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


