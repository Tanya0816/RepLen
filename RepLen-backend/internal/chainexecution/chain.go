package chainexecution
<<<<<<< HEAD
import(
	"github.com/Tanya0816/RepLen/RepLen-backend/internal/intent"
)
type ChainExecutor interface {
	ExecuteIntent(intent *intent.LenIntent) error
=======

import "github.com/Tanya0816/RepLen/RepLen-backend/internal/intent"

type ChainExecutor interface {
	ExecuteIntent(i intent.LenIntent) error
>>>>>>> 0c328719b8b6817c30e98c0027c6f8795b9d3534
}