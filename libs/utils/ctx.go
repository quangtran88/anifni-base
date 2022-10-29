package baseUtils

import (
	"context"
	"reflect"
)

func GetCtxStr(ctx context.Context, key string) string {
	v := ctx.Value(key)
	if reflect.TypeOf(v).Kind() != reflect.String {
		return ""
	}
	return v.(string)
}
