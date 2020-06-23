/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SubscribeRate struct {
	SubscribeThrottlingRatePerConsumer int32 `json:"subscribeThrottlingRatePerConsumer,omitempty"`
	RatePeriodInSecond                 int32 `json:"ratePeriodInSecond,omitempty"`
}
