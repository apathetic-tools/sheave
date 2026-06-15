package registry

import (
	"strings"
)

// Resolve applies include and exclude patterns to filter the registry for a specific item type.
// Custom local items are included by default. Built-in items must be explicitly included.
func (r *Registry) Resolve(itemType string, includes, excludes []string) []*Item {
	// 1. Base selection: all custom items of the given type
	active := make(map[string]*Item)
	for key, item := range r.customs {
		if item.Type == itemType {
			active[key] = item
		}
	}

	// 2. Process includes
	for _, pattern := range includes {
		matches := r.MatchPattern(pattern, itemType)
		for _, item := range matches {
			key := item.ID
			if item.Family != "" {
				key = item.Family + "/" + item.ID
			}
			active[key] = item
		}
	}

	// 3. Process excludes
	for _, pattern := range excludes {
		matches := r.MatchPattern(pattern, itemType)
		for _, item := range matches {
			key := item.ID
			if item.Family != "" {
				key = item.Family + "/" + item.ID
			}
			delete(active, key)
		}
	}

	result := make([]*Item, 0, len(active))
	for _, item := range active {
		result = append(result, item)
	}
	return result
}

// MatchPattern finds items matching a pattern string.
// Patterns can be:
// - `*` or `family/*` for wildcards
// - exact `family/id` or `id`
// Prefix `#` means builtin only.
// Prefix `~` means custom local only.
// No prefix means both (with custom overriding builtin on match).
func (r *Registry) MatchPattern(pattern, itemType string) []*Item {
	var checkBuiltins, checkCustoms bool

	if strings.HasPrefix(pattern, "#") {
		pattern = pattern[1:]
		checkBuiltins = true
	} else if strings.HasPrefix(pattern, "~") {
		pattern = pattern[1:]
		checkCustoms = true
	} else {
		checkBuiltins = true
		checkCustoms = true
	}

	var matches []*Item
	seen := make(map[string]bool)

	isMatch := func(item *Item) bool {
		if item.Type != itemType {
			return false
		}
		key := item.ID
		if item.Family != "" {
			key = item.Family + "/" + item.ID
		}

		if strings.HasSuffix(pattern, "/*") {
			prefix := strings.TrimSuffix(pattern, "/*")
			return item.Family == prefix
		} else if pattern == "*" {
			return true
		}
		return key == pattern || item.ID == pattern
	}

	// Order matters for fallback: custom overrides builtin.
	if checkCustoms {
		for key, item := range r.customs {
			if isMatch(item) {
				matches = append(matches, item)
				seen[key] = true
			}
		}
	}

	if checkBuiltins {
		for key, item := range r.builtins {
			// If we are checking both, skip builtin if custom already provided it
			if checkCustoms && checkBuiltins && seen[key] {
				continue
			}
			if isMatch(item) {
				matches = append(matches, item)
			}
		}
	}

	return matches
}
