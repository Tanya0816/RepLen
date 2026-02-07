package store
import (
"log"
//"sync"
"time"
//"github.com/Tanya0816/RepLen/RepLen-backend/internal/intent"
)

func(s *IntentStore) StartExecutor() {   // This function starts a background goroutine and has ticker for checking every 3 seconds.
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			log.Println("Executing intents...")
			s.executeReadyIntents()
		}
	}()
}
//scan filter and execute
func (s *IntentStore) executeReadyIntents() {   //This part says "Do not reveal your secrets to strangers" but in code form
	s.mu.Lock()
	defer s.mu.Unlock()      //Lock your doors!!
	now := time.Now()
	for id, intent := range s.intents {
		if intent.Status!="PENDING" {
			continue             // Savdhan rhe, Satark rhe (This is matter of safety!!)
		}
		if intent.ExecuteAt.Before(now) || intent.ExecuteAt.Equal(now) {  //delaybased privay check
			log.Printf("Executing intent ID: %s, Action: %s, Amount: %f\n", intent.ID, intent.Action, intent.Amount) //simulated execution by logging
			executedTime := time.Now()  // Updated intent status and executed time
			intent.Status = "EXECUTED"
			s.lastCheckedAt = time.Now()
			intent.ExecutedAt = &executedTime
			s.intents[id] = intent       //Its the matter of figure☆*: .｡. o(≧▽≦), I mean values update in the store.
		}
	}
}