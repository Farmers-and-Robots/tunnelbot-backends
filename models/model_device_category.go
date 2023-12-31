// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Tunnelbot - OpenAPI 3.0
//	Version: 1
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// DeviceCategory is an enum.
type DeviceCategory string

// Validate implements basic validation for this model
func (m DeviceCategory) Validate() error {
	return InKnownDeviceCategory.Validate(m)
}

var (
	DeviceCategoryCurtain    DeviceCategory = "curtain"
	DeviceCategoryFan        DeviceCategory = "fan"
	DeviceCategoryHeating    DeviceCategory = "heating"
	DeviceCategoryIrrigation DeviceCategory = "irrigation"
	DeviceCategoryLighting   DeviceCategory = "lighting"
	DeviceCategoryOther      DeviceCategory = "other"
	DeviceCategoryWeather    DeviceCategory = "weather"

	// KnownDeviceCategory is the list of valid DeviceCategory
	KnownDeviceCategory = []DeviceCategory{
		DeviceCategoryCurtain,
		DeviceCategoryFan,
		DeviceCategoryHeating,
		DeviceCategoryIrrigation,
		DeviceCategoryLighting,
		DeviceCategoryOther,
		DeviceCategoryWeather,
	}
	// KnownDeviceCategoryString is the list of valid DeviceCategory as string
	KnownDeviceCategoryString = []string{
		string(DeviceCategoryCurtain),
		string(DeviceCategoryFan),
		string(DeviceCategoryHeating),
		string(DeviceCategoryIrrigation),
		string(DeviceCategoryLighting),
		string(DeviceCategoryOther),
		string(DeviceCategoryWeather),
	}

	// InKnownDeviceCategory is an ozzo-validator for DeviceCategory
	InKnownDeviceCategory = validation.In(
		DeviceCategoryCurtain,
		DeviceCategoryFan,
		DeviceCategoryHeating,
		DeviceCategoryIrrigation,
		DeviceCategoryLighting,
		DeviceCategoryOther,
		DeviceCategoryWeather,
	)
)
