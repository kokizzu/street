package wcStorage

import (
	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/L"
)

func (f *FilesMutator) DoIncStat(now int64) bool {
	const comment = `FilesMutator) DoIncStat`
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id}, A.X{
		A.X{`+`, f.AccessCount, 1},
		A.X{`=`, f.IdxLastAccessAt(), now},
	})
	return L.IsError(err, comment+` failed`)
}
