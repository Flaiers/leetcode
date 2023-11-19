package main

import "math/rand"

const size = 100

type RandomizedSet struct {
	store  map[int]int
	values []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		store:  make(map[int]int, size),
		values: make([]int, 0, size),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.store[val]; !ok {
		this.values = append(this.values, val)
		this.store[val] = len(this.values) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.store[val]; ok {
		lastIdx := len(this.values) - 1
		lastValue := this.values[lastIdx]

		this.values[idx] = lastValue
		this.values[lastIdx] = val
		this.store[lastValue] = idx

		this.values = this.values[:lastIdx]
		delete(this.store, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.values))
	return this.values[idx]
}

func main() {
	rs := Constructor()
	rs.Insert(1)
	rs.Remove(2)
	rs.GetRandom()
}
