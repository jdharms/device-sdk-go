// -*- mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
)

const (
	ClientData     = "Data"
	ClientMetadata = "Metadata"
	ClientLogging  = "Logging"

	APIv1Prefix = "/api/v1"
	Colon       = ":"
	HttpScheme  = "http://"
	HttpProto   = "HTTP"

	ConfigDirectory    = "./res"
	ConfigFileName     = "configuration.toml"
	ConfigRegistryStem = "edgex/devices/1.0/"
	WritableKey        = "/Writable"
	RegistryFailLimit  = 3

	APICallbackRoute        = APIv1Prefix + "/callback"
	APIValueDescriptorRoute = APIv1Prefix + "/valuedescriptor"
	APIDiscoveryRoute       = APIv1Prefix + "/discovery"
	APIPingRoute            = APIv1Prefix + "/ping"

	IdVar        string = "id"
	NameVar      string = "name"
	CommandVar   string = "command"
	GetCmdMethod string = "get"
	SetCmdMethod string = "set"

	CorrelationHeader = clients.CorrelationHeader
	URLRawQuery = "urlRawQuery"
)
