package conflicts_handler

import (
	conflicts_page "github.com/Uh-little-less-dum/build/pkg/conflicts/page"
	conflicts_slot "github.com/Uh-little-less-dum/build/pkg/conflicts/slot"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	conflict_types "github.com/Uh-little-less-dum/go-utils/pkg/constants/conflictTypes"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

// VERSION_NEXT: Ths should also check for conflicts amongst embeddable components to make sure there are no name clashes.
type BuildConflictsManager struct {
	Page      []*conflicts_page.Conflict
	Slot      []*conflicts_slot.Conflict
	slotIndex int
	pageIndex int
	// Index of the conflict type as it pertains to the conflict_types.ConflictResolveOrder method.
	conflictTypeIndex int
}

func (b BuildConflictsManager) UnresolvedPageConflicts() int {
	n := 0
	for _, p := range b.Page {
		if !p.Resolved {
			n++
		}
	}
	return n
}

func (b BuildConflictsManager) UnresolvedSlotConflicts() int {
	n := 0
	for _, p := range b.Slot {
		if !p.Resolved {
			n++
		}
	}
	return n
}

func (b *BuildConflictsManager) AllValid() bool {
	return len(b.Page) == 0 && len(b.Slot) == 0
}

func (b *BuildConflictsManager) NextSlotConflict() (nextConflict *conflicts_slot.Conflict, ok bool) {
	n := len(b.Slot)
	if b.slotIndex >= n-1 {
		return &conflicts_slot.Conflict{}, false
	}
	return b.Slot[n+1], true
}

func (b *BuildConflictsManager) NextPageConflict() (nextConflict *conflicts_page.Conflict, ok bool) {
	n := len(b.Slot)
	if b.pageIndex >= n-1 {
		return &conflicts_page.Conflict{}, false
	}
	return b.Page[n+1], true
}

func (b *BuildConflictsManager) GatherPluginConflicts(items []*ulld_plugin.Plugin) {
	for i, p := range items {
		for i2, j := range items {
			if i2 != i {
				b.Slot = append(b.Slot, p.HasSlotConflict(j)...)
				b.Page = append(b.Page, p.HasPageConflict(j)...)
			}
		}
	}
}

func (b *BuildConflictsManager) resolveActiveSlotConfig() bool {
	var newItems []*conflicts_slot.Conflict
	if b.slotIndex >= len(b.Slot)-1 {
		log.Warn("Attempted to remove a slot conflict that no longer exists.")
		return false
	}
	id := b.Slot[b.slotIndex].Id()
	for _, item := range b.Slot {
		if item.Id() != id {
			newItems = append(newItems, item)
		}
	}
	hasMore := b.slotIndex < len(b.Slot)
	if hasMore {
		b.slotIndex++
	}
	return hasMore
}

func (b *BuildConflictsManager) resolveActivePageConfig() bool {
	var newItems []*conflicts_page.Conflict
	if b.pageIndex >= len(b.Page)-1 {
		log.Warn("Attempted to remove a page conflict that no longer exists.")
		return false
	}
	id := b.Page[b.pageIndex].Id()
	for _, item := range b.Page {
		if item.Id() != id {
			newItems = append(newItems, item)
		}
	}
	hasMore := b.pageIndex < len(b.Page)
	if hasMore {
		b.pageIndex++
	}
	return hasMore
}

func (b *BuildConflictsManager) CurrentConflictTypes() (current conflict_types.ConflictType) {
	ordered := conflict_types.ConflictResolveOrder()
	return ordered[b.conflictTypeIndex]
}

// Returns a boolean indicating whehter more conflicts of that type exist.
func (b *BuildConflictsManager) resolveByConflictByType(ctype conflict_types.ConflictType) (hasMoreOfType bool) {
	switch ctype {
	case conflict_types.Page:
		return b.resolveActivePageConfig()
	default:
		return b.resolveActiveSlotConfig()
	}
}

func (b BuildConflictsManager) conflictTypeToBuildStage(ctype conflict_types.ConflictType) build_stages.BuildStage {
	switch ctype {
	case conflict_types.Page:
		return build_stages.ResolvePageConflicts
	case conflict_types.Slot:
		return build_stages.ResolveSlotConflicts
	default:
		return build_stages.PostConflictResolveBuild
	}
}

func (b BuildConflictsManager) withBuildStageMsg(newConflictType conflict_types.ConflictType, c tea.Cmd) tea.Cmd {
	return tea.Batch(c, signals.SetStage(b.conflictTypeToBuildStage(newConflictType)))
}

func (b *BuildConflictsManager) ResolveAndContinue() tea.Cmd {
	ctype := b.CurrentConflictTypes()
	hasMore := b.resolveByConflictByType(ctype)
	if hasMore {
		return func() tea.Msg {
			return signals.ResolveConflictMsg{ConflictType: ctype}
		}
	} else {
		newType := conflict_types.NextConflictType(ctype)
		if newType == conflict_types.AllValid {
			return signals.SetStage(build_stages.PostConflictResolveBuild)
		}
		return b.withBuildStageMsg(newType, func() tea.Msg {
			return signals.ResolveConflictMsg{ConflictType: newType}
		})
	}
}
