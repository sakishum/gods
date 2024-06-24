// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/emirpasic/gods/v2/testutils"
)

func TestMapPut(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	if actualValue := m.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	testutils.SameElements(t, m.Keys(), []int{1, 2, 3, 4, 5, 6, 7})
	testutils.SameElements(t, m.Values(), []string{"a", "b", "c", "d", "e", "f", "g"})

	// key,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, "", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.Get(test[0].(int))
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}
}

func TestMapMin(t *testing.T) {
	m := New[int, string]()

	if k, v, ok := m.Min(); k != 0 || v != "" || ok {
		t.Errorf("Got %v->%v->%v expected %v->%v-%v", k, v, ok, 0, "", false)
	}

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	actualKey, actualValue, actualOk := m.Min()
	expectedKey, expectedValue, expectedOk := 1, "a", true
	if actualKey != expectedKey {
		t.Errorf("Got %v expected %v", actualKey, expectedKey)
	}
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualOk != expectedOk {
		t.Errorf("Got %v expected %v", actualOk, expectedOk)
	}
}

func TestMapMax(t *testing.T) {
	m := New[int, string]()

	if k, v, ok := m.Max(); k != 0 || v != "" || ok {
		t.Errorf("Got %v->%v->%v expected %v->%v-%v", k, v, ok, 0, "", false)
	}

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	actualKey, actualValue, actualOk := m.Max()
	expectedKey, expectedValue, expectedOk := 7, "g", true
	if actualKey != expectedKey {
		t.Errorf("Got %v expected %v", actualKey, expectedKey)
	}
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualOk != expectedOk {
		t.Errorf("Got %v expected %v", actualOk, expectedOk)
	}
}

func TestMapClear(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	if actualValue, expectedValue := m.Size(), 4; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	m.Clear()
	if actualValue, expectedValue := m.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapRemove(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	m.Remove(5)
	m.Remove(6)
	m.Remove(7)
	m.Remove(8)
	m.Remove(5)

	testutils.SameElements(t, m.Keys(), []int{1, 2, 3, 4})
	testutils.SameElements(t, m.Values(), []string{"a", "b", "c", "d"})

	if actualValue := m.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}

	tests2 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "", false},
		{6, "", false},
		{7, "", false},
		{8, "", false},
	}

	for _, test := range tests2 {
		actualValue, actualFound := m.Get(test[0].(int))
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	m.Remove(1)
	m.Remove(4)
	m.Remove(2)
	m.Remove(3)
	m.Remove(2)
	m.Remove(2)

	testutils.SameElements(t, m.Keys(), nil)
	testutils.SameElements(t, m.Values(), nil)
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestMapFloor(t *testing.T) {
	m := New[int, string]()
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(1, "a")

	// key,expectedKey,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{-1, 0, "", false},
		{0, 0, "", false},
		{1, 1, "a", true},
		{2, 1, "a", true},
		{3, 3, "c", true},
		{4, 3, "c", true},
		{7, 7, "g", true},
		{8, 7, "g", true},
	}

	for _, test := range tests1 {
		// retrievals
		actualKey, actualValue, actualOk := m.Floor(test[0].(int))
		if actualKey != test[1] || actualValue != test[2] || actualOk != test[3] {
			t.Errorf("Got %v, %v, %v, expected %v, %v, %v", actualKey, actualValue, actualOk, test[1], test[2], test[3])
		}
	}
}

func TestMapCeiling(t *testing.T) {
	m := New[int, string]()
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(1, "a")

	// key,expectedKey,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{-1, 1, "a", true},
		{0, 1, "a", true},
		{1, 1, "a", true},
		{2, 3, "c", true},
		{3, 3, "c", true},
		{4, 7, "g", true},
		{7, 7, "g", true},
		{8, 0, "", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualKey, actualValue, actualOk := m.Ceiling(test[0].(int))
		if actualKey != test[1] || actualValue != test[2] || actualOk != test[3] {
			t.Errorf("Got %v, %v, %v, expected %v, %v, %v", actualKey, actualValue, actualOk, test[1], test[2], test[3])
		}
	}
}

func TestMapEach(t *testing.T) {
	m := New[string, int]()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	count := 0
	m.Each(func(key string, value int) {
		count++
		if actualValue, expectedValue := count, value; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		switch value {
		case 1:
			if actualValue, expectedValue := key, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := key, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 3:
			if actualValue, expectedValue := key, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	})
}

func TestMapMap(t *testing.T) {
	m := New[string, int]()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	mappedMap := m.Map(func(key1 string, value1 int) (key2 string, value2 int) {
		return key1, value1 * value1
	})
	if actualValue, _ := mappedMap.Get("a"); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, "mapped: a")
	}
	if actualValue, _ := mappedMap.Get("b"); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, "mapped: b")
	}
	if actualValue, _ := mappedMap.Get("c"); actualValue != 9 {
		t.Errorf("Got %v expected %v", actualValue, "mapped: c")
	}
	if mappedMap.Size() != 3 {
		t.Errorf("Got %v expected %v", mappedMap.Size(), 3)
	}
}

func TestMapSelect(t *testing.T) {
	m := New[string, int]()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	selectedMap := m.Select(func(key string, value int) bool {
		return key >= "a" && key <= "b"
	})
	if actualValue, _ := selectedMap.Get("a"); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, "value: a")
	}
	if actualValue, _ := selectedMap.Get("b"); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, "value: b")
	}
	if selectedMap.Size() != 2 {
		t.Errorf("Got %v expected %v", selectedMap.Size(), 2)
	}
}

func TestMapAny(t *testing.T) {
	m := New[string, int]()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	any := m.Any(func(key string, value int) bool {
		return value == 3
	})
	if any != true {
		t.Errorf("Got %v expected %v", any, true)
	}
	any = m.Any(func(key string, value int) bool {
		return value == 4
	})
	if any != false {
		t.Errorf("Got %v expected %v", any, false)
	}
}

func TestMapAll(t *testing.T) {
	m := New[string, int]()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	all := m.All(func(key string, value int) bool {
		return key >= "a" && key <= "c"
	})
	if all != true {
		t.Errorf("Got %v expected %v", all, true)
	}
	all = m.All(func(key string, value int) bool {
		return key >= "a" && key <= "b"
	})
	if all != false {
		t.Errorf("Got %v expected %v", all, false)
	}
}

func TestMapFind(t *testing.T) {
	m := New[string, int]()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	foundKey, foundValue := m.Find(func(key string, value int) bool {
		return key == "c"
	})
	if foundKey != "c" || foundValue != 3 {
		t.Errorf("Got %v -> %v expected %v -> %v", foundKey, foundValue, "c", 3)
	}
	foundKey, foundValue = m.Find(func(key string, value int) bool {
		return key == "x"
	})
	if foundKey != "" || foundValue != 0 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundKey, nil, nil)
	}
}

func TestMapChaining(t *testing.T) {
	m := New[string, int]()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	chainedMap := m.Select(func(key string, value int) bool {
		return value > 1
	}).Map(func(key string, value int) (string, int) {
		return key + key, value * value
	})
	if actualValue := chainedMap.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, found := chainedMap.Get("aa"); actualValue != 0 || found {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue, found := chainedMap.Get("bb"); actualValue != 4 || !found {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, found := chainedMap.Get("cc"); actualValue != 9 || !found {
		t.Errorf("Got %v expected %v", actualValue, 9)
	}
}

func TestMapIteratorNextOnEmpty(t *testing.T) {
	m := New[string, int]()
	it := m.Iterator()
	it = m.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty map")
	}
}

func TestMapIteratorPrevOnEmpty(t *testing.T) {
	m := New[string, int]()
	it := m.Iterator()
	it = m.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty map")
	}
}

func TestMapIteratorNext(t *testing.T) {
	m := New[string, int]()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)

	it := m.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		value := it.Value()
		switch key {
		case "a":
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "b":
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "c":
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := value, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapIteratorPrev(t *testing.T) {
	m := New[string, int]()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)

	it := m.Iterator()
	for it.Next() {
	}
	countDown := m.Size()
	for it.Prev() {
		key := it.Key()
		value := it.Value()
		switch key {
		case "a":
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "b":
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "c":
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := value, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapIteratorBegin(t *testing.T) {
	m := New[int, string]()
	it := m.Iterator()
	it.Begin()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestMapIteratorEnd(t *testing.T) {
	m := New[int, string]()
	it := m.Iterator()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it.End()
	it.Prev()
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestMapIteratorFirst(t *testing.T) {
	m := New[int, string]()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it := m.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestMapIteratorLast(t *testing.T) {
	m := New[int, string]()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it := m.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestMapIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// NextTo (empty)
	{
		m := New[int, string]()
		it := m.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
	}

	// NextTo (not found)
	{
		m := New[int, string]()
		m.Put(0, "xx")
		m.Put(1, "yy")
		it := m.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
	}

	// NextTo (found)
	{
		m := New[int, string]()
		m.Put(0, "aa")
		m.Put(1, "bb")
		m.Put(2, "cc")
		it := m.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
		if index, value := it.Key(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Next() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Key(), it.Value(); index != 2 || value != "cc" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "cc")
		}
		if it.Next() {
			t.Errorf("Should not go past last element")
		}
	}
}

func TestMapIteratorPrevTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// PrevTo (empty)
	{
		m := New[int, string]()
		it := m.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
	}

	// PrevTo (not found)
	{
		m := New[int, string]()
		m.Put(0, "xx")
		m.Put(1, "yy")
		it := m.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
	}

	// PrevTo (found)
	{
		m := New[int, string]()
		m.Put(0, "aa")
		m.Put(1, "bb")
		m.Put(2, "cc")
		it := m.Iterator()
		it.End()
		if !it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
		if index, value := it.Key(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Prev() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Key(), it.Value(); index != 0 || value != "aa" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "aa")
		}
		if it.Prev() {
			t.Errorf("Should not go before first element")
		}
	}
}

func TestMapSerialization(t *testing.T) {
	for i := 0; i < 10; i++ {
		original := New[string, string]()
		original.Put("d", "4")
		original.Put("e", "5")
		original.Put("c", "3")
		original.Put("b", "2")
		original.Put("a", "1")

		assertSerialization(original, "A", t)

		serialized, err := original.ToJSON()
		if err != nil {
			t.Errorf("Got error %v", err)
		}
		assertSerialization(original, "B", t)

		deserialized := New[string, string]()
		err = deserialized.FromJSON(serialized)
		if err != nil {
			t.Errorf("Got error %v", err)
		}
		assertSerialization(deserialized, "C", t)
	}

	m := New[string, float64]()
	m.Put("a", 1.0)
	m.Put("b", 2.0)
	m.Put("c", 3.0)

	_, err := json.Marshal([]interface{}{"a", "b", "c", m})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	err = json.Unmarshal([]byte(`{"a":1,"b":2}`), &m)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
}

func TestMapString(t *testing.T) {
	c := New[string, int]()
	c.Put("a", 1)
	if !strings.HasPrefix(c.String(), "TreeMap") {
		t.Errorf("String should start with container name")
	}
}

// noinspection GoBoolExpressions
func assertSerialization(m *Map[string, string], txt string, t *testing.T) {
	if actualValue := m.Keys(); false ||
		actualValue[0] != "a" ||
		actualValue[1] != "b" ||
		actualValue[2] != "c" ||
		actualValue[3] != "d" ||
		actualValue[4] != "e" {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, "[a,b,c,d,e]")
	}
	if actualValue := m.Values(); false ||
		actualValue[0] != "1" ||
		actualValue[1] != "2" ||
		actualValue[2] != "3" ||
		actualValue[3] != "4" ||
		actualValue[4] != "5" {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, "[1,2,3,4,5]")
	}
	if actualValue, expectedValue := m.Size(), 5; actualValue != expectedValue {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, expectedValue)
	}
}

func benchmarkGet(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func BenchmarkTreeMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, struct{}]()
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
