package commands

import (
	"esefexapi/sounddb"
	"fmt"
)

func fmtMetaList(metas []sounddb.SoundMeta) string {
	// log.Printf("fmtMetaList: %v", metas)
	var str string
	for _, meta := range metas {
		str += fmt.Sprintf("- %s %s `%s`\n", meta.Icon.String(), meta.Name, meta.SoundID)
	}

	// log.Println(str)
	return str
}
