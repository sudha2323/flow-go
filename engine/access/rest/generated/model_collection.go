/*
 * Access API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

type Collection struct {
	Id string `json:"id"`

	Transactions []string `json:"transactions"`

	Links *Links `json:"_links,omitempty"`
}
