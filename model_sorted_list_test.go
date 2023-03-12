package goga

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModelSortedListAll(t *testing.T) {

	cases := modelSortedList{
		{nil, 3},
		{nil, 1},
		{nil, 2},
		{nil, 0.1},
		{nil, 55},
		{nil, 5.5},
		{nil, 1000},
	}
	sort.Sort(cases)
	require.Equal(t, 0.1, cases[0].cost)

}
