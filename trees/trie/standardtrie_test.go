package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStandardTrie_Insert(t *testing.T) {
	st := NewStandardTrie()
	st.Insert("words")
	st.Insert("worst")

	count := 1
	for _, child := range st.root.children {
		if child != nil {
			count++
		}
	}
	assert.Equal(t, 2, count)
}

func TestStandardTrie_Search(t *testing.T) {
	st := NewStandardTrie()
	words := []string{"the", "a", "there", "answer", "any", "by", "bye", "their"}
	for _, word := range words {
		st.Insert(word)
	}

	assert.True(t, st.Search("the"))
	assert.False(t, st.Search("this"))
	assert.True(t, st.Search("their"))
	assert.False(t, st.Search("thaw"))
}
