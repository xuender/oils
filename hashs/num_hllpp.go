package hashs

import (
	"github.com/retailnext/hllpp"
	"github.com/xuender/oils/base"
	"golang.org/x/exp/constraints"
)

type NumHLLPP[T constraints.Integer | constraints.Float] struct {
	Hllpp *hllpp.HLLPP
}

func NewNumHLLPP[T constraints.Integer | constraints.Float]() *NumHLLPP[T] {
	return &NumHLLPP[T]{Hllpp: hllpp.New()}
}

func (p *NumHLLPP[T]) Add(num T) {
	p.Hllpp.Add(base.Number2Bytes(num))
}

func (p *NumHLLPP[T]) Count() uint64 {
	return p.Hllpp.Count()
}

func (p *NumHLLPP[T]) Marshal() []byte {
	return p.Hllpp.Marshal()
}

// nolint
func Unmarshal[T constraints.Integer | constraints.Float](data []byte) (*NumHLLPP[T], error) {
	hll, err := hllpp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	return &NumHLLPP[T]{
		Hllpp: hll,
	}, nil
}
