package lru

import (
	"container/list"
	"testing"
)

func TestCache_Add(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	type args struct {
		key   string
		value Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			c.Add(tt.args.key, tt.args.value)
		})
	}
}

type String string

func (d String) Len() int {
	return len(d)
}

func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestCache_Len(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			if got := c.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_RemoveOldest(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			c.RemoveOldest()
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		maxBytes  int64
		onEvicted func(string, Value)
	}
	tests := []struct {
		name string
		args args
		want *Cache
	}{
		{args: args{1, nil}, want: &Cache{maxBytes: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.maxBytes, tt.args.onEvicted); got.maxBytes != tt.args.maxBytes {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
