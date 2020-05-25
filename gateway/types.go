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

type Event struct {
	Type    EventType
	Amount  float64
	Fee     float64
	Hash    string
	Coin    string
	FeeCoin string
	From    string
	To      string
	Error   error
}

type EventConsumer interface {
	Consume(event Event)
}
type EventConsumerFunc func(event Event)

func (f EventConsumerFunc) Consume(event Event) {
	f(event)
}
