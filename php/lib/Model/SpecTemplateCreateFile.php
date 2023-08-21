<?php
/**
 * SpecTemplateCreateFile
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
 * SpecTemplateCreateFile Class Doc Comment
 *
 * @category Class
 * @description File, object specifying file type and content
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class SpecTemplateCreateFile implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'specTemplateCreate_file';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'source_type' => 'string',
        'content' => 'string',
        'content_path' => 'string',
        'content_ref' => 'string',
        'repository' => '\OpenAPI\Client\Model\SpecTemplateCreateFileRepository'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'source_type' => null,
        'content' => null,
        'content_path' => null,
        'content_ref' => null,
        'repository' => null
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
        'source_type' => 'sourceType',
        'content' => 'content',
        'content_path' => 'contentPath',
        'content_ref' => 'contentRef',
        'repository' => 'repository'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'source_type' => 'setSourceType',
        'content' => 'setContent',
        'content_path' => 'setContentPath',
        'content_ref' => 'setContentRef',
        'repository' => 'setRepository'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'source_type' => 'getSourceType',
        'content' => 'getContent',
        'content_path' => 'getContentPath',
        'content_ref' => 'getContentRef',
        'repository' => 'getRepository'
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
        $this->container['source_type'] = $data['source_type'] ?? 'local';
        $this->container['content'] = $data['content'] ?? null;
        $this->container['content_path'] = $data['content_path'] ?? null;
        $this->container['content_ref'] = $data['content_ref'] ?? null;
        $this->container['repository'] = $data['repository'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['source_type'] === null) {
            $invalidProperties[] = "'source_type' can't be null";
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
     * Gets source_type
     *
     * @return string
     */
    public function getSourceType()
    {
        return $this->container['source_type'];
    }

    /**
     * Sets source_type
     *
     * @param string $source_type File Source i.e. local, repository, url.
     *
     * @return self
     */
    public function setSourceType($source_type)
    {
        $this->container['source_type'] = $source_type;

        return $this;
    }

    /**
     * Gets content
     *
     * @return string|null
     */
    public function getContent()
    {
        return $this->container['content'];
    }

    /**
     * Sets content
     *
     * @param string|null $content File content, the template text. Only required when sourceType is `local`.
     *
     * @return self
     */
    public function setContent($content)
    {
        $this->container['content'] = $content;

        return $this;
    }

    /**
     * Gets content_path
     *
     * @return string|null
     */
    public function getContentPath()
    {
        return $this->container['content_path'];
    }

    /**
     * Sets content_path
     *
     * @param string|null $content_path Content Path, the repo file location or url. Required when sourceType is repository or url.
     *
     * @return self
     */
    public function setContentPath($content_path)
    {
        $this->container['content_path'] = $content_path;

        return $this;
    }

    /**
     * Gets content_ref
     *
     * @return string|null
     */
    public function getContentRef()
    {
        return $this->container['content_ref'];
    }

    /**
     * Sets content_ref
     *
     * @param string|null $content_ref Content Ref, the branch/tag. Only used when sourceType is repo.
     *
     * @return self
     */
    public function setContentRef($content_ref)
    {
        $this->container['content_ref'] = $content_ref;

        return $this;
    }

    /**
     * Gets repository
     *
     * @return \OpenAPI\Client\Model\SpecTemplateCreateFileRepository|null
     */
    public function getRepository()
    {
        return $this->container['repository'];
    }

    /**
     * Sets repository
     *
     * @param \OpenAPI\Client\Model\SpecTemplateCreateFileRepository|null $repository repository
     *
     * @return self
     */
    public function setRepository($repository)
    {
        $this->container['repository'] = $repository;

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


