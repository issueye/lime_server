package utils

import "math"

// SlicePage 切片分页
func SlicePage(page, pageSize, nums int) (sliceStart, sliceEnd int) {
	if page < 0 {
		page = 1
	}

	if pageSize < 0 {
		pageSize = 10
	}

	if pageSize > nums {
		return 0, nums
	}

	// 总页数
	pageCount := int(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}
	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize

	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd
}

// Unique 切片去重泛型函数
func Unique[T comparable](slice []T) []T {
	unique := make([]T, 0, len(slice))
	seen := make(map[T]struct{}, len(slice))

	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			unique = append(unique, v)
		}
	}

	return unique
}

func UniqueInOrder[T comparable](slice []T) []T {
	unique := make([]T, 0, len(slice))
	seen := make(map[T]struct{}, len(slice))

	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			unique = append(unique, v)
		}
	}

	return unique
}

func UniqueWithFunc[T any, K comparable](slice []T, keyFunc func(T) K) []T {
	unique := make([]T, 0, len(slice))
	seen := make(map[K]struct{}, len(slice))

	for _, v := range slice {
		key := keyFunc(v)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			unique = append(unique, v)
		}
	}

	return unique
}
