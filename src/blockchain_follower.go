package src

import (
	"fmt"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"time"
)

type startTX struct {
	done chan bool
}

//
// Saga1 worker
//
type followerworker struct {
	gen.SagaWorker
}

func (w *followerworker) HandleJobStart(process *gen.SagaWorkerProcess, job gen.SagaJob) error {
	fmt.Printf(" Worker started on Saga1 with value %q\n", job.Value)
	s := job.Value.(string) + "W"
	process.SendResult(s)
	process.State = job.TransactionID
	return nil
}
func (w *followerworker) HandleJobCancel(process *gen.SagaWorkerProcess, reason string) {
	return
}
func (w *followerworker) HandleJobCommit(process *gen.SagaWorkerProcess, final interface{}) {
	txid := process.State.(gen.SagaTransactionID)
	fmt.Printf(" Worker on Saga1 received final result for %v with value %q\n", txid, final)
}

func (w *followerworker) HandleWorkerTerminate(process *gen.SagaWorkerProcess, reason string) {
	fmt.Printf(" Worker terminated on Saga1 with reason: %q\n", reason)
}

//
// Follower
//
type Follower struct {
	gen.Saga
}

type FollowerState struct {
	done chan bool
	tmp  string
}

func (s *Follower) InitSaga(process *gen.SagaProcess, args ...etf.Term) (gen.SagaOptions, error) {
	opts := gen.SagaOptions{
		Worker: &followerworker{},
	}

	return opts, nil
}

func (s *Follower) HandleTxNew(process *gen.SagaProcess, id gen.SagaTransactionID, value interface{}) gen.SagaStatus {
	// add some sleep
	time.Sleep(300 * time.Millisecond)
	process.StartJob(id, gen.SagaJobOptions{}, value)
	return gen.SagaStatusOK
}

func (s *Follower) HandleTxCancel(process *gen.SagaProcess, id gen.SagaTransactionID, reason string) gen.SagaStatus {
	return gen.SagaStatusOK
}

func (s *Follower) HandleTxDone(process *gen.SagaProcess, id gen.SagaTransactionID, result interface{}) (interface{}, gen.SagaStatus) {
	final := result.(string) + " 🚀"
	fmt.Printf("Saga1. Final result for %v: %q\n", id, final)

	// let the main goroutine know we got finished
	state := process.State.(*FollowerState)
	state.done <- true

	return final, gen.SagaStatusOK
}

func (s *Follower) HandleTxResult(process *gen.SagaProcess, id gen.SagaTransactionID, from gen.SagaNextID, result interface{}) gen.SagaStatus {
	fmt.Printf("Saga1. Received result for %v from Saga4 (%v) with value %q\n", id, from, result)
	// to finish TX on a saga it was created we must call SendResult with the result
	process.SendResult(id, result)
	return gen.SagaStatusOK
}

func (s *Follower) HandleJobResult(process *gen.SagaProcess, id gen.SagaTransactionID, from gen.SagaJobID, result interface{}) gen.SagaStatus {
	// keep result in the process state
	state := process.State.(*FollowerState)
	state.tmp = result.(string)

	fmt.Printf("Saga1. Received result from worker with value %q\n", result)
	next := gen.SagaNext{
		Saga:       gen.ProcessID{Name: "saga2", Node: "node2@localhost"},
		TrapCancel: true,
		Value:      result.(string) + "o",
	}
	next_id, _ := process.Next(id, next)
	fmt.Printf("Saga1. ...sent %v further, to the Saga2 on Node2 (%v)\n", id, next_id)
	return gen.SagaStatusOK
}

// implement this method in order to trap TX cancelation and forward it to the Saga4
func (s *Follower) HandleSagaInfo(process *gen.SagaProcess, message etf.Term) gen.ServerStatus {
	switch m := message.(type) {
	case gen.MessageSagaCancel:
		fmt.Printf("Saga1. Trapped cancelation %v. Reason %q\n", m.TransactionID, m.Reason)
		// process state keeps the value we got from the worker
		state := process.State.(*FollowerState)
		next := gen.SagaNext{
			Saga:  gen.ProcessID{Name: "saga4", Node: "node4@localhost"},
			Value: state.tmp + "o",
		}
		next_id, _ := process.Next(m.TransactionID, next)
		fmt.Printf("Saga1. ...sent %v further, to the Saga4 on Node4 (%v)\n", m.TransactionID, next_id)
	case startTX:
		process.State = &FollowerState{
			done: m.done,
		}
		opts := gen.SagaTransactionOptions{
			TwoPhaseCommit: true,
		}
		id := process.StartTransaction(opts, "Hello ")
		fmt.Printf("Saga1. Start new transaction %v\n", id)
	}
	return gen.ServerStatusOK
}
