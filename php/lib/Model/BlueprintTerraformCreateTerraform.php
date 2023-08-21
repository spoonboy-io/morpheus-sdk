<?php
/**
 * BlueprintTerraformCreateTerraform
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
 * BlueprintTerraformCreateTerraform Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class BlueprintTerraformCreateTerraform implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'blueprintTerraformCreate_terraform';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'config_type' => 'string',
        'json' => 'string',
        'tf' => 'string',
        'git' => '\OpenAPI\Client\Model\BlueprintTerraformCreateTerraformGit',
        'tfvar_secret' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'config_type' => null,
        'json' => null,
        'tf' => null,
        'git' => null,
        'tfvar_secret' => null
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
        'config_type' => 'configType',
        'json' => 'json',
        'tf' => 'tf',
        'git' => 'git',
        'tfvar_secret' => 'tfvarSecret'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'config_type' => 'setConfigType',
        'json' => 'setJson',
        'tf' => 'setTf',
        'git' => 'setGit',
        'tfvar_secret' => 'setTfvarSecret'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'config_type' => 'getConfigType',
        'json' => 'getJson',
        'tf' => 'getTf',
        'git' => 'getGit',
        'tfvar_secret' => 'getTfvarSecret'
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

    const CONFIG_TYPE_TF = 'tf';
    const CONFIG_TYPE_SPEC = 'spec';
    const CONFIG_TYPE_GIT = 'git';
    const CONFIG_TYPE_JSON = 'json';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getConfigTypeAllowableValues()
    {
        return [
            self::CONFIG_TYPE_TF,
            self::CONFIG_TYPE_SPEC,
            self::CONFIG_TYPE_GIT,
            self::CONFIG_TYPE_JSON,
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
        $this->container['config_type'] = $data['config_type'] ?? null;
        $this->container['json'] = $data['json'] ?? null;
        $this->container['tf'] = $data['tf'] ?? null;
        $this->container['git'] = $data['git'] ?? null;
        $this->container['tfvar_secret'] = $data['tfvar_secret'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['config_type'] === null) {
            $invalidProperties[] = "'config_type' can't be null";
        }
        $allowedValues = $this->getConfigTypeAllowableValues();
        if (!is_null($this->container['config_type']) && !in_array($this->container['config_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'config_type', must be one of '%s'",
                $this->container['config_type'],
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
     * Gets config_type
     *
     * @return string
     */
    public function getConfigType()
    {
        return $this->container['config_type'];
    }

    /**
     * Sets config_type
     *
     * @param string $config_type Configuration Type
     *
     * @return self
     */
    public function setConfigType($config_type)
    {
        $allowedValues = $this->getConfigTypeAllowableValues();
        if (!in_array($config_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'config_type', must be one of '%s'",
                    $config_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['config_type'] = $config_type;

        return $this;
    }

    /**
     * Gets json
     *
     * @return string|null
     */
    public function getJson()
    {
        return $this->container['json'];
    }

    /**
     * Sets json
     *
     * @param string|null $json Terraform definition in JSON for `configType` `json`
     *
     * @return self
     */
    public function setJson($json)
    {
        $this->container['json'] = $json;

        return $this;
    }

    /**
     * Gets tf
     *
     * @return string|null
     */
    public function getTf()
    {
        return $this->container['tf'];
    }

    /**
     * Sets tf
     *
     * @param string|null $tf Terraform definition for `configType` `tf`
     *
     * @return self
     */
    public function setTf($tf)
    {
        $this->container['tf'] = $tf;

        return $this;
    }

    /**
     * Gets git
     *
     * @return \OpenAPI\Client\Model\BlueprintTerraformCreateTerraformGit|null
     */
    public function getGit()
    {
        return $this->container['git'];
    }

    /**
     * Sets git
     *
     * @param \OpenAPI\Client\Model\BlueprintTerraformCreateTerraformGit|null $git git
     *
     * @return self
     */
    public function setGit($git)
    {
        $this->container['git'] = $git;

        return $this;
    }

    /**
     * Gets tfvar_secret
     *
     * @return string|null
     */
    public function getTfvarSecret()
    {
        return $this->container['tfvar_secret'];
    }

    /**
     * Sets tfvar_secret
     *
     * @param string|null $tfvar_secret tfvar secret from Morpheus Cypher
     *
     * @return self
     */
    public function setTfvarSecret($tfvar_secret)
    {
        $this->container['tfvar_secret'] = $tfvar_secret;

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

