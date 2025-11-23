package helper

import (
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func ParseID(params httprouter.Params) uint {
	idStr := params.ByName("id")
	idUint, _ := strconv.ParseUint(idStr, 10, 64)
	return uint(idUint)
}