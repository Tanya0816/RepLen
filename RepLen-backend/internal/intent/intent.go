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
<<<<<<< HEAD
	Amount	  int64
	CreatedAt time.Time      // when the intent was created
	ExecutedAt time.Time    //when the intent was executed
    Executed   bool           // whether the intent has been executed

=======
	Amount	float64
	SignedBy  string
	Status    intentStatus
	CreatedAt time.Time      // when the intent was created
	ExecuteAt time.Time    //when the intent was executed
	ExecutedAt *time.Time	 // when the intent was executed, nil if not executed yet
>>>>>>> 0c328719b8b6817c30e98c0027c6f8795b9d3534
}
