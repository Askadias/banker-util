package gateway

type Wallet struct {
	BaseCoin   string
	Address    string
	PrivateKey string
}

type EventType string

const (
	TypeBuy       EventType = "buy"
	TypeSend      EventType = "send"
	TypeMultisend EventType = "multisend"
)

type SendEvent struct {
	Amount float64 `json:"amount,omitempty"`
	Coin   string  `json:"coin,omitempty"`
	To     string  `json:"to,omitempty"`
	ToCoin string  `json:"toCoin,omitempty"`
}

type Event struct {
	SendEvent
	Hash    string      `json:"hash,omitempty"`
	From    string      `json:"from,omitempty"`
	Fee     float64     `json:"fee,omitempty"`
	FeeCoin string      `json:"feeCoin,omitempty"`
	Type    EventType   `json:"type,omitempty"`
	Items   []SendEvent `json:"items,omitempty"`
	Error   error       `json:"error,omitempty"`
}

type EventConsumer interface {
	Consume(event Event)
}
type EventConsumerFunc func(event Event)

func (f EventConsumerFunc) Consume(event Event) {
	f(event)
}
