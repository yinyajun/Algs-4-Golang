package sort

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
	"sort"
)

func TestSorter(t *testing.T) {
	feeds := []*FeedData{}
	cop := []*FeedData{}
	for i := 0; i < 50000; i++ {
		s := rand.Float64()
		feeds = append(feeds, &FeedData{Score: s})
		cop = append(cop, &FeedData{Score: s})
	}

	// old method
	begin := time.Now()
	tk := NewTopKSorter(2000)
	topKFeeds1 := tk.GetTopK(feeds)
	fmt.Println("old topK:", time.Since(begin))

	// new method
	begin = time.Now()
	s := NewTopK(feeds, 2000, func(i, j int) bool {
		return feeds[i].Score < feeds[j].Score
	})
	topKFeeds2 := s.GetTopK().([]*FeedData)
	fmt.Println("new topK:", time.Since(begin))

	// quicksort
	begin = time.Now()
	sort.Slice(cop, func(i, j int) bool {
		return cop[i].Score > cop[j].Score
	})
	fmt.Println("quick sort:", time.Since(begin))

	for index, e := range topKFeeds1 {
		if (e.Score - cop[index].Score) > 0.00000000001 {
			t.Error("two sort are not the same")
		}
	}

	for index, e := range topKFeeds2 {
		if (e.Score - cop[index].Score) > 0.00000000001 {
			t.Error("two sort are not the same")
		}
	}
}
