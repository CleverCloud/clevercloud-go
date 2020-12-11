/*
 * Clever-Cloud API
 *
 * Public API for managing Clever-Cloud data and products
 *
 * API version: 1.0.1
 * Contact: support@clever-cloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package clevercloud

// WannabeApplication struct for WannabeApplication
type WannabeApplication struct {
	Name            string          `json:"name,omitempty"`
	Description     string          `json:"description,omitempty"`
	Zone            string          `json:"zone,omitempty"`
	Deploy          string          `json:"deploy,omitempty"`
	Shutdownable    bool            `json:"shutdownable,omitempty"`
	InstanceType    string          `json:"instanceType,omitempty"`
	InstanceVersion string          `json:"instanceVersion,omitempty"`
	InstanceVariant string          `json:"instanceVariant,omitempty"`
	MinInstances    int32           `json:"minInstances,omitempty"`
	MaxInstances    int32           `json:"maxInstances,omitempty"`
	MinFlavor       string          `json:"minFlavor,omitempty"`
	MaxFlavor       string          `json:"maxFlavor,omitempty"`
	Tags            []string        `json:"tags,omitempty"`
	Archived        bool            `json:"archived,omitempty"`
	StickySessions  bool            `json:"stickySessions,omitempty"`
	Homogeneous     bool            `json:"homogeneous,omitempty"`
	Favourite       bool            `json:"favourite,omitempty"`
	CancelOnPush    bool            `json:"cancelOnPush,omitempty"`
	SeparateBuild   bool            `json:"separateBuild,omitempty"`
	BuildFlavor     string          `json:"buildFlavor,omitempty"`
	OauthService    string          `json:"oauthService,omitempty"`
	OauthAppId      string          `json:"oauthAppId,omitempty"`
	OauthApp        WannabeOauthApp `json:"oauthApp,omitempty"`
	ApplianceId     string          `json:"applianceId,omitempty"`
	Branch          string          `json:"branch,omitempty"`
	ForceHttps      string          `json:"forceHttps,omitempty"`
}