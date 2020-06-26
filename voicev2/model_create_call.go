/*
 * Voice API BETA
 *
 * This is the *beta* version of the Voice API. Calls created with v2 must be managed using [v1 endpoints](/api/voice).  Voice v2 is provided to allow users to create IP calls. If you do not have this requirement we recommend that you stay on v1 for now.  > This API may break backwards compatibility at short notice (60 days) 
 *
 * API version: 2.1.1
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package voicev2

// CreateCall struct for CreateCall
type CreateCall struct {
	// Temporarily removed, this causes syntax invalidation
	// To []OneOfobjectobjectobjectobjectobject `json:"to"`
	From CreateCallFrom `json:"from"`
	// **Required** unless `answer_url` is provided.  The [Nexmo Call Control Object](/voice/voice-api/ncco-reference) to use for this call.
	Ncco []map[string]interface{} `json:"ncco,omitempty"`
	// **Required** unless `ncco` is provided.  The webhook endpoint where you provide the [Nexmo Call Control Object](/voice/voice-api/ncco-reference) that governs this call.
	AnswerUrl []string `json:"answer_url,omitempty"`
	// The HTTP method used to send event information to answer_url.
	AnswerMethod string `json:"answer_method,omitempty"`
	// **Required** unless `event_url` is configured at the application level, see [Create an Application](/api/application.v2#createApplication)  The webhook endpoint where call progress events are sent to. For more information about the values sent, see [Event webhook](/voice/voice-api/webhook-reference#event-webhook).
	EventUrl []string `json:"event_url"`
	// The HTTP method used to send event information to event_url.
	EventMethod string `json:"event_method,omitempty"`
	// Configure the behavior when Nexmo detects that the call is answered by voicemail. If Continue Nexmo sends an HTTP request to event_url with the Call event machine. hangup  end the call
	MachineDetection string `json:"machine_detection,omitempty"`
	// Set the number of seconds that elapse before Nexmo hangs up after the call state changes to in_progress.
	LengthTimer int32 `json:"length_timer,omitempty"`
	// Set the number of seconds that elapse before Nexmo hangs up after the call state changes to ‘ringing’.
	RingingTimer int32 `json:"ringing_timer,omitempty"`
}
