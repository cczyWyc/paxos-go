package paxos

// Proposer proposes new value to acceptors
type Proposer struct {
	acceptorClients  []AcceptorClientInterface
	acceptorPromises map[string]map[string]*Proposal
}

// NewProposer create a new Proposer instance
func NewProposer(acceptorClients []AcceptorClientInterface) *Proposer {
	acceptorPromises := make(map[string]map[string]*Proposal, len(acceptorClients))
	for _, acceptorClient := range acceptorClients {
		acceptorPromises[acceptorClient.GetName()] = make(map[string]*Proposal)
	}
	return &Proposer{
		acceptorClients:  acceptorClients,
		acceptorPromises: acceptorPromises,
	}
}

// majority return simple majority of acceptor nodes
func (p *Proposer) majority() int {
	return len(p.acceptorClients)/2 + 1
}

// majorityReached returns true if number of matching promises from acceptors is equal or greater than simple majority
// of acceptor nodes
func (p *Proposer) majorityReached(proposal *Proposal) bool {
	var marches = 0
	// Iterate over promised values for each acceptor
	for _, promiseMap := range p.acceptorPromises {
		// Skip if thr acceptor has not yet promised a proposal for this key
		promised, ok := promiseMap[proposal.Key]
		if !ok {
			continue
		}

		// If the promised and proposal number is the same, increment matches count
		if promised.Number == proposal.Number {
			marches++
		}
	}
	return marches >= p.majority()
}
