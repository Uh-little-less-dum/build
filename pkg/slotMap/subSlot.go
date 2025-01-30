package slot_map

import (
	"fmt"

	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	"github.com/tidwall/gjson"
)

type SubSlot struct {
	subSlot string
}

func NewSubSlot(s string) SubSlot {
	return SubSlot{subSlot: s}
}

func (s SubSlot) OutputType(slotKey SlotKey, data gjson.Result) string {
	pathString := fmt.Sprintf("%s.%s.type", slotKey, s.subSlot)
	return data.Get(pathString).Str
}

func (s SubSlot) ClientOnly(slotKey SlotKey, data gjson.Result) bool {
	pathString := fmt.Sprintf("%s.%s.clientOnly", slotKey, s.subSlot)
	return data.Get(pathString).Bool()
}

func (s SubSlot) PropsExtends(slotKey SlotKey, data gjson.Result) string {
	pathString := fmt.Sprintf("%s.%s.propsExtends", slotKey, s.subSlot)
	return data.Get(pathString).Str
}

func (s SubSlot) OutputPath(slotKey SlotKey, data gjson.Result) string {
	pathString := fmt.Sprintf("%s.%s.path", slotKey, s.subSlot)
	return data.Get(pathString).Str
}

func (s SubSlot) writeOutput(paths target_paths.TargetPaths) {

}
