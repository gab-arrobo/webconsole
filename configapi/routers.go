// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2024 Canonical Ltd
// SPDX-License-Identifier: Apache-2.0

/*
 * Connectivity Service Configuration
 *
 * APIs to configure connectivity service in Aether Network
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package configapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// This function is not autogenerated
func AddConfigV1Service(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.RouterGroup {
	group := engine.Group("/config/v1")
	if len(middlewares) > 0 {
		group.Use(middlewares...)
	}
	addRoutes(group, routes)
	return group
}

func addRoutes(group *gin.RouterGroup, routes Routes) {
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			group.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			group.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			group.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},

	{
		"GetDeviceGroups",
		http.MethodGet,
		"/device-group",
		GetDeviceGroups,
	},

	{
		"GetDeviceGroupByName",
		http.MethodGet,
		"/device-group/:group-name",
		GetDeviceGroupByName,
	},

	{
		"DeviceGroupGroupNameDelete",
		http.MethodDelete,
		"/device-group/:group-name",
		DeviceGroupGroupNameDelete,
	},

	{
		"DeviceGroupGroupNamePut",
		http.MethodPut,
		"/device-group/:group-name",
		DeviceGroupGroupNamePut,
	},

	{
		"DeviceGroupGroupNamePost",
		http.MethodPost,
		"/device-group/:group-name",
		DeviceGroupGroupNamePost,
	},

	{
		"GetNetworkSlices",
		http.MethodGet,
		"/network-slice",
		GetNetworkSlices,
	},

	{
		"GetNetworkSliceByName",
		http.MethodGet,
		"/network-slice/:slice-name",
		GetNetworkSliceByName,
	},

	{
		"NetworkSliceSliceNameDelete",
		http.MethodDelete,
		"/network-slice/:slice-name",
		NetworkSliceSliceNameDelete,
	},

	{
		"NetworkSliceSliceNamePost",
		http.MethodPost,
		"/network-slice/:slice-name",
		NetworkSliceSliceNamePost,
	},

	{
		"NetworkSliceSliceNamePut",
		http.MethodPut,
		"/network-slice/:slice-name",
		NetworkSliceSliceNamePut,
	},
	{
		"GetGnbs",
		http.MethodGet,
		"/inventory/gnb",
		GetGnbs,
	},
	{
		"PostGnb",
		http.MethodPost,
		"/inventory/gnb",
		PostGnb,
	},
	{
		"PutGnb",
		http.MethodPut,
		"/inventory/gnb/:gnb-name",
		PutGnb,
	},
	{
		"DeleteGnb",
		http.MethodDelete,
		"/inventory/gnb/:gnb-name",
		DeleteGnb,
	},
	{
		"GetUpfs",
		http.MethodGet,
		"/inventory/upf",
		GetUpfs,
	},
	{
		"PostUpf",
		http.MethodPost,
		"/inventory/upf",
		PostUpf,
	},
	{
		"PutUpf",
		http.MethodPut,
		"/inventory/upf/:upf-hostname",
		PutUpf,
	},
	{
		"DeleteUpf",
		http.MethodDelete,
		"/inventory/upf/:upf-hostname",
		DeleteUpf,
	},
}
