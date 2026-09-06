// Copyright 2024 Tamás Gulácsi.
//
// SPDX-License-Identifier: EUPL-1.2

package client

import "github.com/go-openapi/runtime"
import "github.com/go-openapi/strfmt"

type Client struct {
	*MantisBTRESTAPI
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) Client {
	return Client{
		MantisBTRESTAPI: New(transport, formats),
		AuthInfo:        authInfo,
	}
}
