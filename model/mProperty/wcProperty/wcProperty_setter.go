package wcProperty

import (
	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/L"
)

func (p *PropLikeCountMutator) DoInc(delta int) bool {
	x, err := p.Adapter.Upsert(p.SpaceName(), A.X{p.PropId, 1}, A.X{
		A.X{`+`, p.IdxCount(), delta},
	})
	_ = x
	return L.IsError(err, `PropLikeCountMutator.DoInc`)
}
