package paxos

import "fmt"

type Proposal struct {
	Number int
	Key    string
	Value  []byte
}

func (p *Proposal) String() string {
	return fmt.Sprintf("(num=%d, key=\"%s\", value=\"%s\")", p.Number, p.Key, p.Value)
}
