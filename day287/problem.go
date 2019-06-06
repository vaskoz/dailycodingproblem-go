package day287

import "container/heap"

// WebsiteVisitor represents a user visiting a website.
type WebsiteVisitor struct {
	Website string
	UserID  int
}

// WebsitePair represents a tuple of websites.
// Also contains a Similarity score based on Jaccard index.
// High similarity score is better.
type WebsitePair struct {
	A, B       string
	Similarity float64
}

type priorityQueue []WebsitePair

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Similarity < pq[j].Similarity
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	pair := x.(WebsitePair)
	*pq = append(*pq, pair)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	pair := old[n-1]
	*pq = old[0 : n-1]
	return pair
}

// TopKMostSimilarPairs returns k WebsitePairs.
// Runs in O(N^2) because it compares all pairs for similarity.
func TopKMostSimilarPairs(visits []WebsiteVisitor, k int) []WebsitePair {
	if k < 1 || len(visits) < 2 {
		return nil
	}
	sites := make(map[string][]int)
	for _, site := range visits {
		sites[site.Website] = append(sites[site.Website], site.UserID)
	}
	pq := make(priorityQueue, 0, k)
	heap.Init(&pq)
	for site1, ids1 := range sites {
		for site2, ids2 := range sites {
			if site1 >= site2 {
				continue
			}
			score := similarityScore(ids1, ids2)
			record := WebsitePair{site1, site2, score}
			if pq.Len() < k {
				heap.Push(&pq, record)
			} else {
				lowest := heap.Pop(&pq).(WebsitePair)
				if lowest.Similarity < record.Similarity {
					heap.Push(&pq, record)
				} else {
					heap.Push(&pq, lowest)
				}
			}
		}
	}
	topK := make([]WebsitePair, k)
	for i := range topK {
		topK[i] = heap.Pop(&pq).(WebsitePair)
	}
	return topK
}

func similarityScore(ids1, ids2 []int) float64 {
	seen := make(map[int]struct{}, len(ids1))
	for _, id := range ids1 {
		seen[id] = struct{}{}
	}
	var intersectionCount float64
	for _, id := range ids2 {
		if _, found := seen[id]; found {
			intersectionCount++
		}
	}
	unionCount := float64(len(ids1)+len(ids2)) - intersectionCount
	return intersectionCount / unionCount
}
