package sys

import (
	"fmt"
	"iads/server/internals/pkg/models/basemodel"
)

func TableName(name string) string {
	return fmt.Sprintf("%s%s%s", basemodel.GetTablePrefix(), "sys_", name)
}
