// Package v1 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AWSS3UploadRequestOptions defines model for AWSS3UploadRequestOptions.
type AWSS3UploadRequestOptions map[string]interface{}

// AWSS3UploadStatus defines model for AWSS3UploadStatus.
type AWSS3UploadStatus struct {
	Url string `json:"url"`
}

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// AWSUploadStatus defines model for AWSUploadStatus.
type AWSUploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// AzureUploadRequestOptions defines model for AzureUploadRequestOptions.
type AzureUploadRequestOptions struct {

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded. This link explains how
	// to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeMetadata defines model for ComposeMetadata.
type ComposeMetadata struct {

	// ID (hash) of the built commit
	OstreeCommit *string `json:"ostree_commit,omitempty"`

	// Package list including NEVRA
	Packages *[]PackageMetadata `json:"packages,omitempty"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// ComposesResponse defines model for ComposesResponse.
type ComposesResponse struct {
	Data  []ComposesResponseItem `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// ComposesResponseItem defines model for ComposesResponseItem.
type ComposesResponseItem struct {
	CreatedAt string      `json:"created_at"`
	Id        string      `json:"id"`
	Request   interface{} `json:"request"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions []DistributionItem

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {

	// List of valid Google accounts to share the imported Compute Node image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	//     If not specified, the imported Compute Node image is not shared with any
	//     account.
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// HTTPErrorList defines model for HTTPErrorList.
type HTTPErrorList struct {
	Errors []HTTPError `json:"errors"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture  string        `json:"architecture"`
	ImageType     ImageTypes    `json:"image_type"`
	Ostree        *OSTree       `json:"ostree,omitempty"`
	UploadRequest UploadRequest `json:"upload_request"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status       string        `json:"status"`
	UploadStatus *UploadStatus `json:"upload_status,omitempty"`
}

// ImageTypes defines model for ImageTypes.
type ImageTypes string

// List of ImageTypes
const (
	ImageTypes_ami                 ImageTypes = "ami"
	ImageTypes_aws                 ImageTypes = "aws"
	ImageTypes_azure               ImageTypes = "azure"
	ImageTypes_edge_commit         ImageTypes = "edge-commit"
	ImageTypes_edge_container      ImageTypes = "edge-container"
	ImageTypes_edge_installer      ImageTypes = "edge-installer"
	ImageTypes_gcp                 ImageTypes = "gcp"
	ImageTypes_guest_image         ImageTypes = "guest-image"
	ImageTypes_image_installer     ImageTypes = "image-installer"
	ImageTypes_rhel_edge_commit    ImageTypes = "rhel-edge-commit"
	ImageTypes_rhel_edge_installer ImageTypes = "rhel-edge-installer"
	ImageTypes_vhd                 ImageTypes = "vhd"
	ImageTypes_vsphere             ImageTypes = "vsphere"
)

// OSTree defines model for OSTree.
type OSTree struct {
	Ref *string `json:"ref,omitempty"`
	Url *string `json:"url,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

// PackageMetadata defines model for PackageMetadata.
type PackageMetadata struct {
	Arch      string  `json:"arch"`
	Epoch     *string `json:"epoch,omitempty"`
	Name      string  `json:"name"`
	Release   string  `json:"release"`
	Sigmd5    string  `json:"sigmd5"`
	Signature *string `json:"signature,omitempty"`
	Type      string  `json:"type"`
	Version   string  `json:"version"`
}

// PackagesResponse defines model for PackagesResponse.
type PackagesResponse struct {
	Data  []Package `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// Readiness defines model for Readiness.
type Readiness struct {
	Readiness string `json:"readiness"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options interface{} `json:"options"`
	Type    UploadTypes `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{} `json:"options"`
	Status  string      `json:"status"`
	Type    UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// List of UploadTypes
const (
	UploadTypes_aws    UploadTypes = "aws"
	UploadTypes_aws_s3 UploadTypes = "aws.s3"
	UploadTypes_azure  UploadTypes = "azure"
	UploadTypes_gcp    UploadTypes = "gcp"
)

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetComposesParams defines parameters for GetComposes.
type GetComposesParams struct {

	// max amount of composes, default 100
	Limit *int `json:"limit,omitempty"`

	// composes page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// GetPackagesParams defines parameters for GetPackages.
type GetPackagesParams struct {

	// distribution to look up packages for
	Distribution string `json:"distribution"`

	// architecture to look up packages for
	Architecture string `json:"architecture"`

	// packages to look for
	Search string `json:"search"`

	// max amount of packages, default 100
	Limit *int `json:"limit,omitempty"`

	// packages page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// ComposeImageRequestBody defines body for ComposeImage for application/json ContentType.
type ComposeImageJSONRequestBody ComposeImageJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get the architectures and their image types available for a given distribution
	// (GET /architectures/{distribution})
	GetArchitectures(ctx echo.Context, distribution string) error
	// compose image
	// (POST /compose)
	ComposeImage(ctx echo.Context) error
	// get a collection of previous compose requests for the logged in user
	// (GET /composes)
	GetComposes(ctx echo.Context, params GetComposesParams) error
	// get status of an image compose
	// (GET /composes/{composeId})
	GetComposeStatus(ctx echo.Context, composeId string) error
	// get metadata of an image compose
	// (GET /composes/{composeId}/metadata)
	GetComposeMetadata(ctx echo.Context, composeId string) error
	// get the available distributions
	// (GET /distributions)
	GetDistributions(ctx echo.Context) error
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(ctx echo.Context) error

	// (GET /packages)
	GetPackages(ctx echo.Context, params GetPackagesParams) error
	// return the readiness
	// (GET /ready)
	GetReadiness(ctx echo.Context) error
	// get the service version
	// (GET /version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArchitectures converts echo context to params.
func (w *ServerInterfaceWrapper) GetArchitectures(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "distribution" -------------
	var distribution string

	err = runtime.BindStyledParameter("simple", false, "distribution", ctx.Param("distribution"), &distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArchitectures(ctx, distribution)
	return err
}

// ComposeImage converts echo context to params.
func (w *ServerInterfaceWrapper) ComposeImage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ComposeImage(ctx)
	return err
}

// GetComposes converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposes(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetComposesParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposes(ctx, params)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, composeId)
	return err
}

// GetComposeMetadata converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeMetadata(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeMetadata(ctx, composeId)
	return err
}

// GetDistributions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDistributions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDistributions(ctx)
	return err
}

// GetOpenapiJson converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapiJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapiJson(ctx)
	return err
}

// GetPackages converts echo context to params.
func (w *ServerInterfaceWrapper) GetPackages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPackagesParams
	// ------------- Required query parameter "distribution" -------------

	err = runtime.BindQueryParameter("form", true, true, "distribution", ctx.QueryParams(), &params.Distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// ------------- Required query parameter "architecture" -------------

	err = runtime.BindQueryParameter("form", true, true, "architecture", ctx.QueryParams(), &params.Architecture)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter architecture: %s", err))
	}

	// ------------- Required query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, true, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPackages(ctx, params)
	return err
}

// GetReadiness converts echo context to params.
func (w *ServerInterfaceWrapper) GetReadiness(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetReadiness(ctx)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVersion(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/architectures/:distribution", wrapper.GetArchitectures)
	router.POST("/compose", wrapper.ComposeImage)
	router.GET("/composes", wrapper.GetComposes)
	router.GET("/composes/:composeId", wrapper.GetComposeStatus)
	router.GET("/composes/:composeId/metadata", wrapper.GetComposeMetadata)
	router.GET("/distributions", wrapper.GetDistributions)
	router.GET("/openapi.json", wrapper.GetOpenapiJson)
	router.GET("/packages", wrapper.GetPackages)
	router.GET("/ready", wrapper.GetReadiness)
	router.GET("/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAAA//98lPYYlTMAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
