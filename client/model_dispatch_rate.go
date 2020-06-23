/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DispatchRate struct {
	DispatchThrottlingRateInMsg  int32 `json:"dispatchThrottlingRateInMsg,omitempty"`
	DispatchThrottlingRateInByte int64 `json:"dispatchThrottlingRateInByte,omitempty"`
	RelativeToPublishRate        bool  `json:"relativeToPublishRate,omitempty"`
	RatePeriodInSecond           int32 `json:"ratePeriodInSecond,omitempty"`
}
