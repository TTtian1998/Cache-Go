package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map contains all hashed keys
type Map struct {
	hash     Hash
	replicas int            //nums of each virtual node
	keys     []int          //hash ring
	hashMap  map[int]string //Mapping of virtual nodes and real nodes; key: hash of v-node,value:name of r-node
}

// New creates a Map instance
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add real nodes in the hash.
// key:name of real node
// For each real node, calculate the hash value by using strconv.Itoa(i) + key,
// then create Map.replicas num of virtual nodes and process mapping association
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	//sort in hash ring
	sort.Ints(m.keys)
}

// Get gets the closest r-node in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica.
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	//idx%len(m.keys) to solve situation that idx == len(m.keys), it represents 0
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
