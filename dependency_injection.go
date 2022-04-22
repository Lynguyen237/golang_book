// interface - defines the methods the consumers should expect
type SQSHandler interface {
	func handle(msg interface{}) error
}


// implementation - details of the implementation should be transparent to consumers
type MessagingSQSHandler struct {
	// dependencies - should be injected (i.e passed in) - allows for easier mocking/testing
	messagingClient messaging.Client
	emailClient email.IEmailClient
}

func (h *messagingSQSHandler) handle(msg interface{}) error {
	// Do message parsing
	json.ParseStuff(msg)

	// Call messaging API to get report details
	h.messagingClient.GetReport(...)
	
	// Send email to issuer
	h.emailClient.SendEmail(...)
}

// factory function - instantiates the implementation, but exposes (returns) the interface
// 
// Returning the interface ensures that the downstream consumers (i.e. the rest of the application)
// only know about the interface and do not know about the implementation. This promotes loose coupling.
func NewMessagingSQSHandler(messagingClient messaging.Client, emailClient email.IEmailClient) SQSHandler {
	return &MessagingSQSHandler{
		messagingClient: messagingClient,
		emailClient: emailClient,
	}
}
