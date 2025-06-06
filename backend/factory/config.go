// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2019 free5GC.org
// SPDX-FileCopyrightText: 2024 Canonical Ltd
//
// SPDX-License-Identifier: Apache-2.0
//

/*

 * WebUi Configuration Factory

 */

package factory

import (
	"github.com/omec-project/util/logger"
)

type Config struct {
	Info          *Info          `yaml:"info"`
	Configuration *Configuration `yaml:"configuration"`
	Logger        *logger.Logger `yaml:"logger"`
}

type Info struct {
	Version     string `yaml:"version,omitempty"`
	Description string `yaml:"description,omitempty"`
	HttpVersion int    `yaml:"http-version,omitempty"`
}

type Configuration struct {
	Mongodb                 *Mongodb    `yaml:"mongodb"`
	WebuiTLS                *TLS        `yaml:"webui-tls"`
	NfConfigTLS             *TLS        `yaml:"nfconfig-tls"`
	RocEnd                  *RocEndpt   `yaml:"managedByConfigPod,omitempty"` // fetch config during bootup
	LteEnd                  []*LteEndpt `yaml:"endpoints,omitempty"`          // LTE endpoints are configured and not auto-detected
	Mode5G                  bool        `yaml:"mode5G,omitempty"`
	SdfComp                 bool        `yaml:"spec-compliant-sdf"`
	EnableAuthentication    bool        `yaml:"enableAuthentication,omitempty"`
	SendPebbleNotifications bool        `yaml:"send-pebble-notifications,omitempty"`
	CfgPort                 int         `yaml:"cfgport,omitempty"`
}

type TLS struct {
	PEM string `yaml:"pem,omitempty"`
	Key string `yaml:"key,omitempty"`
}

type Mongodb struct {
	Name           string `yaml:"name,omitempty"`
	Url            string `yaml:"url,omitempty"`
	AuthKeysDbName string `yaml:"authKeysDbName"`
	AuthUrl        string `yaml:"authUrl"`
	WebuiDBName    string `yaml:"webuiDbName,omitempty"`
	WebuiDBUrl     string `yaml:"webuiDbUrl,omitempty"`
}

type RocEndpt struct {
	SyncUrl string `yaml:"syncUrl,omitempty"`
	Enabled bool   `yaml:"enabled,omitempty"`
}

type LteEndpt struct {
	NodeType       string `yaml:"type,omitempty"`
	ConfigPushUrl  string `yaml:"configPushUrl,omitempty"`
	ConfigCheckUrl string `yaml:"configCheckUrl,omitempty"` // only for 4G components
}
