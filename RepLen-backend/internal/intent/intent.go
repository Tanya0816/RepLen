package intent
import "time"
type ActionType string
type intentStatus string
const(
	AddLiquidity    ActionType = "ADD"
	RemoveLiquidity ActionType = "REMOVE"
	Rebalance       ActionType = "REBALANCE"
)
const(
	StatusPending   intentStatus = "PENDING"
	StatusExecuted  intentStatus = "EXECUTED"
	StatusFailed    intentStatus = "FAILED"
	StatusCancelled intentStatus = "CANCELLED"
)
type LenIntent struct {
	ID        string
	Action    ActionType
	Address   string
	PoolID    string
	Amount	float64
	SignedBy  string
	Status    intentStatus
	CreatedAt time.Time      // when the intent was created
	ExecuteAt time.Time    //when the intent was executed
	ExecutedAt *time.Time	 // when the intent was executed, nil if not executed yet
}
