//// This file was auto-generated using github.com/clipperhouse/gen
//// Modifying this file is not recommended as it will likely be overwritten in the future
//
//// Sort (if included below) is a modification of http://golang.org/pkg/sort/#Sort
//// List (if included below) is a modification of http://golang.org/pkg/container/list/
//// Ring (if included below) is a modification of http://golang.org/pkg/container/ring/
//// Copyright 2009 The Go Authors. All rights reserved.
//// Use of this source code is governed by a BSD-style
//// license that can be found in the LICENSE file.
//
//// Set (if included below) is a modification of https://github.com/deckarep/golang-set
//// The MIT License (MIT)
//// Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)
//
package services

//
//import (
//	"errors"
//	"github.com/justinsb/gova/collections"
//)
//
//// ServiceConfigs is a slice of type ServiceConfig, for use with gen methods below. Use this type where you would use []ServiceConfig. (This is required because slices cannot be method receivers.)
//type ServiceConfigs []ServiceConfig
//
//// All verifies that all elements of ServiceConfigs return true for the passed func. See: http://clipperhouse.github.io/gen/#All
//func (rcv ServiceConfigs) All(fn func(ServiceConfig) bool) bool {
//	for _, v := range rcv {
//		if !fn(v) {
//			return false
//		}
//	}
//	return true
//}
//
//// Any verifies that one or more elements of ServiceConfigs return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
//func (rcv ServiceConfigs) Any(fn func(ServiceConfig) bool) bool {
//	for _, v := range rcv {
//		if fn(v) {
//			return true
//		}
//	}
//	return false
//}
//
//func (rcv ServiceConfigs) At(index int) interface{} {
//	return rcv[index]
//}
//
//// Count gives the number elements of ServiceConfigs that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
//func (rcv ServiceConfigs) Count(fn func(ServiceConfig) bool) (result int) {
//	for _, v := range rcv {
//		if fn(v) {
//			result++
//		}
//	}
//	return
//}
//
//// Distinct returns a new ServiceConfigs slice whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
//func (rcv ServiceConfigs) Distinct() (result ServiceConfigs) {
//	appended := make(map[ServiceConfig]bool)
//	for _, v := range rcv {
//		if !appended[v] {
//			result = append(result, v)
//			appended[v] = true
//		}
//	}
//	return result
//}
//
//// DistinctBy returns a new ServiceConfigs slice whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
//func (rcv ServiceConfigs) DistinctBy(equal func(ServiceConfig, ServiceConfig) bool) (result ServiceConfigs) {
//	for _, v := range rcv {
//		eq := func(_app ServiceConfig) bool {
//			return equal(v, _app)
//		}
//		if !result.Any(eq) {
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//// Each iterates over ServiceConfigs and executes the passed func against each element. See: http://clipperhouse.github.io/gen/#Each
//func (rcv ServiceConfigs) Each(fn func(ServiceConfig)) {
//	for _, v := range rcv {
//		fn(v)
//	}
//}
//
//// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
//func (rcv ServiceConfigs) First(fn func(ServiceConfig) bool) (result ServiceConfig, err error) {
//	for _, v := range rcv {
//		if fn(v) {
//			result = v
//			return
//		}
//	}
//	err = errors.New("no ServiceConfigs elements return true for passed func")
//	return
//}
//
//// IsSortedBy reports whether an instance of ServiceConfigs is sorted, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
//func (rcv ServiceConfigs) IsSortedBy(less func(ServiceConfig, ServiceConfig) bool) bool {
//	n := len(rcv)
//	for i := n - 1; i > 0; i-- {
//		if less(rcv[i], rcv[i-1]) {
//			return false
//		}
//	}
//	return true
//}
//
//// IsSortedDesc reports whether an instance of ServiceConfigs is sorted in descending order, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
//func (rcv ServiceConfigs) IsSortedByDesc(less func(ServiceConfig, ServiceConfig) bool) bool {
//	greater := func(a, b ServiceConfig) bool {
//		return a != b && !less(a, b)
//	}
//	return rcv.IsSortedBy(greater)
//}
//
//func (rcv ServiceConfigs) Iterator() collections.Iterator {
//	it := collections.NewSequenceIterator(rcv)
//	return &it
//}
//
//// MaxBy returns an element of ServiceConfigs containing the maximum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MaxBy
//func (rcv ServiceConfigs) MaxBy(less func(ServiceConfig, ServiceConfig) bool) (result ServiceConfig, err error) {
//	l := len(rcv)
//	if l == 0 {
//		err = errors.New("cannot determine the MaxBy of an empty slice")
//		return
//	}
//	m := 0
//	for i := 1; i < l; i++ {
//		if rcv[i] != rcv[m] && !less(rcv[i], rcv[m]) {
//			m = i
//		}
//	}
//	result = rcv[m]
//	return
//}
//
//// MinBy returns an element of ServiceConfigs containing the minimum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MinBy
//func (rcv ServiceConfigs) MinBy(less func(ServiceConfig, ServiceConfig) bool) (result ServiceConfig, err error) {
//	l := len(rcv)
//	if l == 0 {
//		err = errors.New("cannot determine the Min of an empty slice")
//		return
//	}
//	m := 0
//	for i := 1; i < l; i++ {
//		if less(rcv[i], rcv[m]) {
//			m = i
//		}
//	}
//	result = rcv[m]
//	return
//}
//
//// Creates a new array, including type casting if needed
//func NewServiceConfigs(items collections.Sequence) ServiceConfigs {
//	n := items.Size()
//	if n == 0 {
//		return ServiceConfigs{}
//	}
//
//	ret := make([]ServiceConfig, 0, n)
//	for it := items.Iterator(); it.HasNext(); {
//		item := it.Next().(ServiceConfig)
//		ret = append(ret, item)
//	}
//
//	return ret
//}
//
//// Single returns exactly one element of ServiceConfigs that returns true for the passed func. Returns error if no or multiple elements return true. See: http://clipperhouse.github.io/gen/#Single
//func (rcv ServiceConfigs) Single(fn func(ServiceConfig) bool) (result ServiceConfig, err error) {
//	var candidate ServiceConfig
//	found := false
//	for _, v := range rcv {
//		if fn(v) {
//			if found {
//				err = errors.New("multiple ServiceConfigs elements return true for passed func")
//				return
//			}
//			candidate = v
//			found = true
//		}
//	}
//	if found {
//		result = candidate
//	} else {
//		err = errors.New("no ServiceConfigs elements return true for passed func")
//	}
//	return
//}
//
//// Returns the number of items in the collection
//func (rcv ServiceConfigs) Size() int {
//	return len(rcv)
//}
//
//// SortBy returns a new ordered ServiceConfigs slice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
//func (rcv ServiceConfigs) SortBy(less func(ServiceConfig, ServiceConfig) bool) ServiceConfigs {
//	result := make(ServiceConfigs, len(rcv))
//	copy(result, rcv)
//	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
//	n := len(result)
//	maxDepth := 0
//	for i := n; i > 0; i >>= 1 {
//		maxDepth++
//	}
//	maxDepth *= 2
//	quickSortServiceConfigs(result, less, 0, n, maxDepth)
//	return result
//}
//
//// SortByDesc returns a new, descending-ordered ServiceConfigs slice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
//func (rcv ServiceConfigs) SortByDesc(less func(ServiceConfig, ServiceConfig) bool) ServiceConfigs {
//	greater := func(a, b ServiceConfig) bool {
//		return a != b && !less(a, b)
//	}
//	return rcv.SortBy(greater)
//}
//
//// Where returns a new ServiceConfigs slice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
//func (rcv ServiceConfigs) Where(fn func(ServiceConfig) bool) (result ServiceConfigs) {
//	for _, v := range rcv {
//		if fn(v) {
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//// Sort support methods
//
//func swapServiceConfigs(rcv ServiceConfigs, a, b int) {
//	rcv[a], rcv[b] = rcv[b], rcv[a]
//}
//
//// Insertion sort
//func insertionSortServiceConfigs(rcv ServiceConfigs, less func(ServiceConfig, ServiceConfig) bool, a, b int) {
//	for i := a + 1; i < b; i++ {
//		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
//			swapServiceConfigs(rcv, j, j-1)
//		}
//	}
//}
//
//// siftDown implements the heap property on rcv[lo, hi).
//// first is an offset into the array where the root of the heap lies.
//func siftDownServiceConfigs(rcv ServiceConfigs, less func(ServiceConfig, ServiceConfig) bool, lo, hi, first int) {
//	root := lo
//	for {
//		child := 2*root + 1
//		if child >= hi {
//			break
//		}
//		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
//			child++
//		}
//		if !less(rcv[first+root], rcv[first+child]) {
//			return
//		}
//		swapServiceConfigs(rcv, first+root, first+child)
//		root = child
//	}
//}
//
//func heapSortServiceConfigs(rcv ServiceConfigs, less func(ServiceConfig, ServiceConfig) bool, a, b int) {
//	first := a
//	lo := 0
//	hi := b - a
//
//	// Build heap with greatest element at top.
//	for i := (hi - 1) / 2; i >= 0; i-- {
//		siftDownServiceConfigs(rcv, less, i, hi, first)
//	}
//
//	// Pop elements, largest first, into end of rcv.
//	for i := hi - 1; i >= 0; i-- {
//		swapServiceConfigs(rcv, first, first+i)
//		siftDownServiceConfigs(rcv, less, lo, i, first)
//	}
//}
//
//// Quicksort, following Bentley and McIlroy,
//// Engineering a Sort Function, SP&E November 1993.
//
//// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
//func medianOfThreeServiceConfigs(rcv ServiceConfigs, less func(ServiceConfig, ServiceConfig) bool, a, b, c int) {
//	m0 := b
//	m1 := a
//	m2 := c
//	// bubble sort on 3 elements
//	if less(rcv[m1], rcv[m0]) {
//		swapServiceConfigs(rcv, m1, m0)
//	}
//	if less(rcv[m2], rcv[m1]) {
//		swapServiceConfigs(rcv, m2, m1)
//	}
//	if less(rcv[m1], rcv[m0]) {
//		swapServiceConfigs(rcv, m1, m0)
//	}
//	// now rcv[m0] <= rcv[m1] <= rcv[m2]
//}
//
//func swapRangeServiceConfigs(rcv ServiceConfigs, a, b, n int) {
//	for i := 0; i < n; i++ {
//		swapServiceConfigs(rcv, a+i, b+i)
//	}
//}
//
//func doPivotServiceConfigs(rcv ServiceConfigs, less func(ServiceConfig, ServiceConfig) bool, lo, hi int) (midlo, midhi int) {
//	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
//	if hi-lo > 40 {
//		// Tukey's Ninther, median of three medians of three.
//		s := (hi - lo) / 8
//		medianOfThreeServiceConfigs(rcv, less, lo, lo+s, lo+2*s)
//		medianOfThreeServiceConfigs(rcv, less, m, m-s, m+s)
//		medianOfThreeServiceConfigs(rcv, less, hi-1, hi-1-s, hi-1-2*s)
//	}
//	medianOfThreeServiceConfigs(rcv, less, lo, m, hi-1)
//
//	// Invariants are:
//	//	rcv[lo] = pivot (set up by ChoosePivot)
//	//	rcv[lo <= i < a] = pivot
//	//	rcv[a <= i < b] < pivot
//	//	rcv[b <= i < c] is unexamined
//	//	rcv[c <= i < d] > pivot
//	//	rcv[d <= i < hi] = pivot
//	//
//	// Once b meets c, can swap the "= pivot" sections
//	// into the middle of the slice.
//	pivot := lo
//	a, b, c, d := lo+1, lo+1, hi, hi
//	for {
//		for b < c {
//			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
//				b++
//			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
//				swapServiceConfigs(rcv, a, b)
//				a++
//				b++
//			} else {
//				break
//			}
//		}
//		for b < c {
//			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
//				c--
//			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
//				swapServiceConfigs(rcv, c-1, d-1)
//				c--
//				d--
//			} else {
//				break
//			}
//		}
//		if b >= c {
//			break
//		}
//		// rcv[b] > pivot; rcv[c-1] < pivot
//		swapServiceConfigs(rcv, b, c-1)
//		b++
//		c--
//	}
//
//	min := func(a, b int) int {
//		if a < b {
//			return a
//		}
//		return b
//	}
//
//	n := min(b-a, a-lo)
//	swapRangeServiceConfigs(rcv, lo, b-n, n)
//
//	n = min(hi-d, d-c)
//	swapRangeServiceConfigs(rcv, c, hi-n, n)
//
//	return lo + b - a, hi - (d - c)
//}
//
//func quickSortServiceConfigs(rcv ServiceConfigs, less func(ServiceConfig, ServiceConfig) bool, a, b, maxDepth int) {
//	for b-a > 7 {
//		if maxDepth == 0 {
//			heapSortServiceConfigs(rcv, less, a, b)
//			return
//		}
//		maxDepth--
//		mlo, mhi := doPivotServiceConfigs(rcv, less, a, b)
//		// Avoiding recursion on the larger subproblem guarantees
//		// a stack depth of at most lg(b-a).
//		if mlo-a < b-mhi {
//			quickSortServiceConfigs(rcv, less, a, mlo, maxDepth)
//			a = mhi // i.e., quickSortServiceConfigs(rcv, mhi, b)
//		} else {
//			quickSortServiceConfigs(rcv, less, mhi, b, maxDepth)
//			b = mlo // i.e., quickSortServiceConfigs(rcv, a, mlo)
//		}
//	}
//	if b-a > 1 {
//		insertionSortServiceConfigs(rcv, less, a, b)
//	}
//}
