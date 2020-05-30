package gateway

type Wallet struct {
	BaseCoin   string
	Address    string
	PrivateKey string
}

type EventType string

const (
	TypeSend      EventType = "send"
	TypeMultisend EventType = "multisend"
)

type SendEvent struct {
	Amount  float64 `json:"amount,omitempty"`
	Fee     float64 `json:"fee,omitempty"`
	Hash    string  `json:"hash,omitempty"`
	Coin    string  `json:"coin,omitempty"`
	FeeCoin string  `json:"feeCoin,omitempty"`
	From    string  `json:"from,omitempty"`
	To      string  `json:"to,omitempty"`
}

type Event struct {
	SendEvent
	Type  EventType   `json:"type,omitempty"`
	Items []SendEvent `json:"items,omitempty"`
	Error error       `json:"error,omitempty"`
}

type EventConsumer interface {
	Consume(event Event)
}
type EventConsumerFunc func(event Event)

func (f EventConsumerFunc) Consume(event Event) {
	f(event)
}
