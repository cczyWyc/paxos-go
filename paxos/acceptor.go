package paxos

import (
	"errors"
	"fmt"
	"log"
)

type Acceptor struct {
	accepted map[string]*Proposal
	promised map[string]*Proposal
}

type AcceptorClientInterface interface {
	GetName() string
	SendPrepare(proposal *Proposal) (*Proposal, error)
	SendPropose(proposal *Proposal) (*Proposal, error)
}

// NewAcceptor create a new acceptor instance
func NewAcceptor() *Acceptor {
	return &Acceptor{
		accepted: make(map[string]*Proposal),
		promised: make(map[string]*Proposal),
	}
}

// ReceivePrepare If an acceptor receives a prepare request with number N greater than that of
// any prepare request to which it has already response, then it responds to the request with a promise not to accept
// any more proposals numbered less than N and with the highest-numbered proposal (if any) that it has accepted.
func (a *Acceptor) ReceivePrepare(proposal *Proposal) (*Proposal, error) {
	promised, ok := a.promised[proposal.Key]
	if ok && promised.Number > proposal.Number {
		return nil, errors.New(fmt.Sprintf("Already promised to accept %s which is > than requested %s",
			promised, proposal))
	}

	// Promise to accept the proposal
	a.promised[proposal.Key] = proposal
	log.Printf("Promised to accept %s", proposal)

	return proposal, nil
}

// ReceivePropose If an acceptor receive a propose request for a proposal numbered N, it accepts the proposal unless
// it has already responded to a prepare request having a number greater than N
func (a *Acceptor) ReceivePropose(proposal *Proposal) (*Proposal, error) {
	promised, ok := a.promised[proposal.Key]
	if ok && promised.Number > proposal.Number {
		return nil, errors.New(fmt.Sprintf("Already promised to accept %s which is > than requested %s",
			promised, proposal))
	}

	// accept the proposal
	a.accepted[proposal.Key] = proposal
	log.Printf("Accepted %s", proposal)

	// Truncate promises map
	a.promised = make(map[string]*Proposal)

	return proposal, nil
}
