// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/percona/pmm-managed/api/swagger/client/alerts"
	"github.com/percona/pmm-managed/api/swagger/client/base"
	"github.com/percona/pmm-managed/api/swagger/client/demo"
	"github.com/percona/pmm-managed/api/swagger/client/scrape_configs"
)

// Default pmm managed HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new pmm managed HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PmmManaged {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new pmm managed HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PmmManaged {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new pmm managed client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PmmManaged {
	cli := new(PmmManaged)
	cli.Transport = transport

	cli.Alerts = alerts.New(transport, formats)

	cli.Base = base.New(transport, formats)

	cli.Demo = demo.New(transport, formats)

	cli.ScrapeConfigs = scrape_configs.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PmmManaged is a client for pmm managed
type PmmManaged struct {
	Alerts *alerts.Client

	Base *base.Client

	Demo *demo.Client

	ScrapeConfigs *scrape_configs.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PmmManaged) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Alerts.SetTransport(transport)

	c.Base.SetTransport(transport)

	c.Demo.SetTransport(transport)

	c.ScrapeConfigs.SetTransport(transport)

}