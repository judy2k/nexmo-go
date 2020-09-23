/*
 * Number Insight API
 *
 * Nexmo's Number Insight API delivers real-time intelligence about the validity, reachability and roaming status of a phone number and tells you how to format the number correctly in your application. There are three levels of Number Insight API available: [Basic, Standard and Advanced](https://developer.nexmo.com/number-insight/overview#basic-standard-and-advanced-apis). The advanced API is available asynchronously as well as synchronously.
 *
 * API version: 1.0.7
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package numberInsight
// NiResponseJsonStandardAllOf struct for NiResponseJsonStandardAllOf
type NiResponseJsonStandardAllOf struct {
	// The amount in EUR charged to your account.
	RequestPrice string `json:"request_price,omitempty"`
	// If there is an internal lookup error, the `refund_price` will reflect the lookup price. If `cnam` is requested for a non-US number the `refund_price` will reflect the `cnam` price. If both of these conditions occur, `refund_price` is the sum of the lookup price and `cnam` price.
	RefundPrice string `json:"refund_price,omitempty"`
	// Your account balance in EUR after this request.
	RemainingBalance float32 `json:"remaining_balance,omitempty"`
	CurrentCarrier NiCurrentCarrierProperties `json:"current_carrier,omitempty"`
	OriginalCarrier NiInitialCarrierProperties `json:"original_carrier,omitempty"`
	// If the user has changed carrier for `number`. The assumed status means that the information supplier has replied to the request but has not said explicitly that the number is ported.
	Ported string `json:"ported,omitempty"`
	Roaming NiRoaming `json:"roaming,omitempty"`
	CallerIdentity NiCallerIdentity `json:"caller_identity,omitempty"`
	// Full name of the person or business who owns the phone number. `unknown` if this information is not available. This parameter is only present if `cnam` had a value of `true` within the request.
	CallerName string `json:"caller_name,omitempty"`
	// Last name of the person who owns the phone number if the owner is an individual. This parameter is only present if `cnam` had a value of `true` within the request.
	LastName string `json:"last_name,omitempty"`
	// First name of the person who owns the phone number if the owner is an individual. This parameter is only present if `cnam` had a value of `true` within the request.
	FirstName string `json:"first_name,omitempty"`
	// The value will be `business` if the owner of a phone number is a business. If the owner is an individual the value will be `consumer`. The value will be `unknown` if this information is not available. This parameter is only present if `cnam` had a value of `true` within the request.
	CallerType string `json:"caller_type,omitempty"`
}