<?php
/**
 * GuidanceAzureReservationsConfigServicesAzureVmsSummary
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
 * GuidanceAzureReservationsConfigServicesAzureVmsSummary Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class GuidanceAzureReservationsConfigServicesAzureVmsSummary implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'guidanceAzureReservations_config_services_azureVms_summary';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'total_savings' => 'float',
        'currency_code' => 'string',
        'total_savings_percent' => 'float',
        'term' => 'string',
        'payment_option' => 'string',
        'service' => 'string',
        'on_demand_count' => 'int',
        'on_demand_cost' => 'float',
        'reserved_count' => 'int',
        'reserved_cost' => 'int',
        'recommended_count' => 'int',
        'recommended_cost' => 'float'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'total_savings' => null,
        'currency_code' => null,
        'total_savings_percent' => null,
        'term' => null,
        'payment_option' => null,
        'service' => null,
        'on_demand_count' => 'int64',
        'on_demand_cost' => null,
        'reserved_count' => 'int64',
        'reserved_cost' => 'int64',
        'recommended_count' => 'int64',
        'recommended_cost' => null
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
        'total_savings' => 'totalSavings',
        'currency_code' => 'currencyCode',
        'total_savings_percent' => 'totalSavingsPercent',
        'term' => 'term',
        'payment_option' => 'paymentOption',
        'service' => 'service',
        'on_demand_count' => 'onDemandCount',
        'on_demand_cost' => 'onDemandCost',
        'reserved_count' => 'reservedCount',
        'reserved_cost' => 'reservedCost',
        'recommended_count' => 'recommendedCount',
        'recommended_cost' => 'recommendedCost'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'total_savings' => 'setTotalSavings',
        'currency_code' => 'setCurrencyCode',
        'total_savings_percent' => 'setTotalSavingsPercent',
        'term' => 'setTerm',
        'payment_option' => 'setPaymentOption',
        'service' => 'setService',
        'on_demand_count' => 'setOnDemandCount',
        'on_demand_cost' => 'setOnDemandCost',
        'reserved_count' => 'setReservedCount',
        'reserved_cost' => 'setReservedCost',
        'recommended_count' => 'setRecommendedCount',
        'recommended_cost' => 'setRecommendedCost'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'total_savings' => 'getTotalSavings',
        'currency_code' => 'getCurrencyCode',
        'total_savings_percent' => 'getTotalSavingsPercent',
        'term' => 'getTerm',
        'payment_option' => 'getPaymentOption',
        'service' => 'getService',
        'on_demand_count' => 'getOnDemandCount',
        'on_demand_cost' => 'getOnDemandCost',
        'reserved_count' => 'getReservedCount',
        'reserved_cost' => 'getReservedCost',
        'recommended_count' => 'getRecommendedCount',
        'recommended_cost' => 'getRecommendedCost'
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
        $this->container['total_savings'] = $data['total_savings'] ?? null;
        $this->container['currency_code'] = $data['currency_code'] ?? null;
        $this->container['total_savings_percent'] = $data['total_savings_percent'] ?? null;
        $this->container['term'] = $data['term'] ?? null;
        $this->container['payment_option'] = $data['payment_option'] ?? null;
        $this->container['service'] = $data['service'] ?? null;
        $this->container['on_demand_count'] = $data['on_demand_count'] ?? null;
        $this->container['on_demand_cost'] = $data['on_demand_cost'] ?? null;
        $this->container['reserved_count'] = $data['reserved_count'] ?? null;
        $this->container['reserved_cost'] = $data['reserved_cost'] ?? null;
        $this->container['recommended_count'] = $data['recommended_count'] ?? null;
        $this->container['recommended_cost'] = $data['recommended_cost'] ?? null;
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
     * Gets total_savings
     *
     * @return float|null
     */
    public function getTotalSavings()
    {
        return $this->container['total_savings'];
    }

    /**
     * Sets total_savings
     *
     * @param float|null $total_savings total_savings
     *
     * @return self
     */
    public function setTotalSavings($total_savings)
    {
        $this->container['total_savings'] = $total_savings;

        return $this;
    }

    /**
     * Gets currency_code
     *
     * @return string|null
     */
    public function getCurrencyCode()
    {
        return $this->container['currency_code'];
    }

    /**
     * Sets currency_code
     *
     * @param string|null $currency_code currency_code
     *
     * @return self
     */
    public function setCurrencyCode($currency_code)
    {
        $this->container['currency_code'] = $currency_code;

        return $this;
    }

    /**
     * Gets total_savings_percent
     *
     * @return float|null
     */
    public function getTotalSavingsPercent()
    {
        return $this->container['total_savings_percent'];
    }

    /**
     * Sets total_savings_percent
     *
     * @param float|null $total_savings_percent total_savings_percent
     *
     * @return self
     */
    public function setTotalSavingsPercent($total_savings_percent)
    {
        $this->container['total_savings_percent'] = $total_savings_percent;

        return $this;
    }

    /**
     * Gets term
     *
     * @return string|null
     */
    public function getTerm()
    {
        return $this->container['term'];
    }

    /**
     * Sets term
     *
     * @param string|null $term term
     *
     * @return self
     */
    public function setTerm($term)
    {
        $this->container['term'] = $term;

        return $this;
    }

    /**
     * Gets payment_option
     *
     * @return string|null
     */
    public function getPaymentOption()
    {
        return $this->container['payment_option'];
    }

    /**
     * Sets payment_option
     *
     * @param string|null $payment_option payment_option
     *
     * @return self
     */
    public function setPaymentOption($payment_option)
    {
        $this->container['payment_option'] = $payment_option;

        return $this;
    }

    /**
     * Gets service
     *
     * @return string|null
     */
    public function getService()
    {
        return $this->container['service'];
    }

    /**
     * Sets service
     *
     * @param string|null $service service
     *
     * @return self
     */
    public function setService($service)
    {
        $this->container['service'] = $service;

        return $this;
    }

    /**
     * Gets on_demand_count
     *
     * @return int|null
     */
    public function getOnDemandCount()
    {
        return $this->container['on_demand_count'];
    }

    /**
     * Sets on_demand_count
     *
     * @param int|null $on_demand_count on_demand_count
     *
     * @return self
     */
    public function setOnDemandCount($on_demand_count)
    {
        $this->container['on_demand_count'] = $on_demand_count;

        return $this;
    }

    /**
     * Gets on_demand_cost
     *
     * @return float|null
     */
    public function getOnDemandCost()
    {
        return $this->container['on_demand_cost'];
    }

    /**
     * Sets on_demand_cost
     *
     * @param float|null $on_demand_cost on_demand_cost
     *
     * @return self
     */
    public function setOnDemandCost($on_demand_cost)
    {
        $this->container['on_demand_cost'] = $on_demand_cost;

        return $this;
    }

    /**
     * Gets reserved_count
     *
     * @return int|null
     */
    public function getReservedCount()
    {
        return $this->container['reserved_count'];
    }

    /**
     * Sets reserved_count
     *
     * @param int|null $reserved_count reserved_count
     *
     * @return self
     */
    public function setReservedCount($reserved_count)
    {
        $this->container['reserved_count'] = $reserved_count;

        return $this;
    }

    /**
     * Gets reserved_cost
     *
     * @return int|null
     */
    public function getReservedCost()
    {
        return $this->container['reserved_cost'];
    }

    /**
     * Sets reserved_cost
     *
     * @param int|null $reserved_cost reserved_cost
     *
     * @return self
     */
    public function setReservedCost($reserved_cost)
    {
        $this->container['reserved_cost'] = $reserved_cost;

        return $this;
    }

    /**
     * Gets recommended_count
     *
     * @return int|null
     */
    public function getRecommendedCount()
    {
        return $this->container['recommended_count'];
    }

    /**
     * Sets recommended_count
     *
     * @param int|null $recommended_count recommended_count
     *
     * @return self
     */
    public function setRecommendedCount($recommended_count)
    {
        $this->container['recommended_count'] = $recommended_count;

        return $this;
    }

    /**
     * Gets recommended_cost
     *
     * @return float|null
     */
    public function getRecommendedCost()
    {
        return $this->container['recommended_cost'];
    }

    /**
     * Sets recommended_cost
     *
     * @param float|null $recommended_cost recommended_cost
     *
     * @return self
     */
    public function setRecommendedCost($recommended_cost)
    {
        $this->container['recommended_cost'] = $recommended_cost;

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

