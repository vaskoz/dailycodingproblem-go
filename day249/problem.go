package day249

// MaxXorPairBrute find the maximum XOR of any two elements.
// Runs in O(N^2) time and O(1) space.
func MaxXorPairBrute(nums []int64) int64 {
	max := int64(-1 << 63)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if result := nums[i] ^ nums[j]; result > max {
				max = result
			}
		}
	}
	return max
}

// MaxXorPairSort find the maximum XOR of any two elements.
// Runs in O(N log N) time and O(N) space.
func MaxXorPairSort(nums []int64) int64 {
	copied := make([]int64, len(nums))
	copy(copied, nums)
	max := int64(-1 << 63)
	for i := 1; i < len(nums); i++ {
		if result := nums[i] ^ nums[i-1]; result > max {
			max = result
		}
	}
	return max
}

// BitTrie is a trie for binary data.
type BitTrie struct {
	Value int64
	Next  [2]*BitTrie
}

// Insert a key in the BitTrie.
func (bt *BitTrie) Insert(key int64) {
	cur := bt
	for i := 63; i >= 0; i-- {
		var bit int
		if key&(1<<uint(i)) != 0 {
			bit = 1
		}
		if cur.Next[bit] == nil {
			cur.Next[bit] = &BitTrie{}
		}
		cur = cur.Next[bit]
	}
	cur.Value = key
}

// HighestXorValue returns the highest xor-ed result of a
// value that already exists in the tree.
func (bt *BitTrie) HighestXorValue(key int64) int64 {
	cur := bt
	for i := 63; i >= 0; i-- {
		var bit int
		if key&(1<<uint(i)) == 0 {
			if i != 63 {
				bit = 1
			}
		} else {
			if i == 63 {
				bit = 1
			}
		}
		if cur.Next[bit] == nil {
			bit = 1 - bit
		}
		cur = cur.Next[bit]
	}
	return key ^ cur.Value
}

// MaxXorPairLinear runs in O(N) time and O(N) space for a trie.
func MaxXorPairLinear(nums []int64) int64 {
	max := int64(-1 << 63)
	bt := &BitTrie{}
	bt.Insert(nums[0])
	for i := 1; i < len(nums); i++ {
		if result := bt.HighestXorValue(nums[i]); result > max {
			max = result
		}
		bt.Insert(nums[i])
	}
	return max
}
