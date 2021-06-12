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

// GithubWebhookPayload struct for GithubWebhookPayload
type GithubWebhookPayload struct {
	Ref        string                  `json:"ref,omitempty"`
	After      string                  `json:"after,omitempty"`
	Repository GithubWebhookRepository `json:"repository,omitempty"`
	Sender     GithubWebhookSender     `json:"sender,omitempty"`
	Pusher     GithubWebhookPusher     `json:"pusher,omitempty"`
	HeadCommit GithubCommit            `json:"head_commit,omitempty"`
}
