<?php
/**
 * ScriptUpdate
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
 * ScriptUpdate Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ScriptUpdate implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'scriptUpdate';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'name' => 'string',
        'labels' => 'string[]',
        'category' => 'string',
        'script_version' => 'string',
        'script_phase' => 'string',
        'script_type' => 'string',
        'script' => 'string',
        'run_as_user' => 'string',
        'sudo_user' => 'bool'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'name' => null,
        'labels' => null,
        'category' => null,
        'script_version' => null,
        'script_phase' => null,
        'script_type' => null,
        'script' => null,
        'run_as_user' => null,
        'sudo_user' => null
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
        'name' => 'name',
        'labels' => 'labels',
        'category' => 'category',
        'script_version' => 'scriptVersion',
        'script_phase' => 'scriptPhase',
        'script_type' => 'scriptType',
        'script' => 'script',
        'run_as_user' => 'runAsUser',
        'sudo_user' => 'sudoUser'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'name' => 'setName',
        'labels' => 'setLabels',
        'category' => 'setCategory',
        'script_version' => 'setScriptVersion',
        'script_phase' => 'setScriptPhase',
        'script_type' => 'setScriptType',
        'script' => 'setScript',
        'run_as_user' => 'setRunAsUser',
        'sudo_user' => 'setSudoUser'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'name' => 'getName',
        'labels' => 'getLabels',
        'category' => 'getCategory',
        'script_version' => 'getScriptVersion',
        'script_phase' => 'getScriptPhase',
        'script_type' => 'getScriptType',
        'script' => 'getScript',
        'run_as_user' => 'getRunAsUser',
        'sudo_user' => 'getSudoUser'
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

    const SCRIPT_TYPE_BASH = 'bash';
    const SCRIPT_TYPE_POWERSHELL = 'powershell';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getScriptTypeAllowableValues()
    {
        return [
            self::SCRIPT_TYPE_BASH,
            self::SCRIPT_TYPE_POWERSHELL,
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
        $this->container['name'] = $data['name'] ?? null;
        $this->container['labels'] = $data['labels'] ?? null;
        $this->container['category'] = $data['category'] ?? null;
        $this->container['script_version'] = $data['script_version'] ?? '1';
        $this->container['script_phase'] = $data['script_phase'] ?? null;
        $this->container['script_type'] = $data['script_type'] ?? 'bash';
        $this->container['script'] = $data['script'] ?? null;
        $this->container['run_as_user'] = $data['run_as_user'] ?? null;
        $this->container['sudo_user'] = $data['sudo_user'] ?? false;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        $allowedValues = $this->getScriptTypeAllowableValues();
        if (!is_null($this->container['script_type']) && !in_array($this->container['script_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'script_type', must be one of '%s'",
                $this->container['script_type'],
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
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name Script name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets labels
     *
     * @return string[]|null
     */
    public function getLabels()
    {
        return $this->container['labels'];
    }

    /**
     * Sets labels
     *
     * @param string[]|null $labels Array of label strings, can be used for filtering.
     *
     * @return self
     */
    public function setLabels($labels)
    {
        $this->container['labels'] = $labels;

        return $this;
    }

    /**
     * Gets category
     *
     * @return string|null
     */
    public function getCategory()
    {
        return $this->container['category'];
    }

    /**
     * Sets category
     *
     * @param string|null $category Script category
     *
     * @return self
     */
    public function setCategory($category)
    {
        $this->container['category'] = $category;

        return $this;
    }

    /**
     * Gets script_version
     *
     * @return string|null
     */
    public function getScriptVersion()
    {
        return $this->container['script_version'];
    }

    /**
     * Sets script_version
     *
     * @param string|null $script_version Version of the script
     *
     * @return self
     */
    public function setScriptVersion($script_version)
    {
        $this->container['script_version'] = $script_version;

        return $this;
    }

    /**
     * Gets script_phase
     *
     * @return string|null
     */
    public function getScriptPhase()
    {
        return $this->container['script_phase'];
    }

    /**
     * Sets script_phase
     *
     * @param string|null $script_phase Phase for the script, provision, start, etc.
     *
     * @return self
     */
    public function setScriptPhase($script_phase)
    {
        $this->container['script_phase'] = $script_phase;

        return $this;
    }

    /**
     * Gets script_type
     *
     * @return string|null
     */
    public function getScriptType()
    {
        return $this->container['script_type'];
    }

    /**
     * Sets script_type
     *
     * @param string|null $script_type Type for the script
     *
     * @return self
     */
    public function setScriptType($script_type)
    {
        $allowedValues = $this->getScriptTypeAllowableValues();
        if (!is_null($script_type) && !in_array($script_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'script_type', must be one of '%s'",
                    $script_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['script_type'] = $script_type;

        return $this;
    }

    /**
     * Gets script
     *
     * @return string|null
     */
    public function getScript()
    {
        return $this->container['script'];
    }

    /**
     * Sets script
     *
     * @param string|null $script Script content, that is, the code itself.
     *
     * @return self
     */
    public function setScript($script)
    {
        $this->container['script'] = $script;

        return $this;
    }

    /**
     * Gets run_as_user
     *
     * @return string|null
     */
    public function getRunAsUser()
    {
        return $this->container['run_as_user'];
    }

    /**
     * Sets run_as_user
     *
     * @param string|null $run_as_user Run as a specific user.
     *
     * @return self
     */
    public function setRunAsUser($run_as_user)
    {
        $this->container['run_as_user'] = $run_as_user;

        return $this;
    }

    /**
     * Gets sudo_user
     *
     * @return bool|null
     */
    public function getSudoUser()
    {
        return $this->container['sudo_user'];
    }

    /**
     * Sets sudo_user
     *
     * @param bool|null $sudo_user Sudo, whether or not to run with sudo.
     *
     * @return self
     */
    public function setSudoUser($sudo_user)
    {
        $this->container['sudo_user'] = $sudo_user;

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

