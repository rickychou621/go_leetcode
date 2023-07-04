package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("TwoSum1", func(t *testing.T) {

		nums := []int{3, 3}
		target := 6

		expectResult := []int{0, 1}
		result := twoSum(nums, target)
		if !reflect.DeepEqual(result, expectResult) {
			t.Errorf(
				"TwoSumErr Test Fail: ExpectResult: %+v got: %+v",
				expectResult, result,
			)
		}
	})

	t.Run("TwoSum2", func(t *testing.T) {

		nums := []int{2, 7, 11, 15}
		target := 9

		expectResult := []int{0, 1}
		result := twoSum(nums, target)
		if !reflect.DeepEqual(result, expectResult) {
			t.Errorf(
				"TwoSumErr Test Fail: ExpectResult: %+v got: %+v",
				expectResult, result,
			)
		}
	})

	t.Run("TwoSum2", func(t *testing.T) {

		nums := []int{3, 2, 4}
		target := 6

		expectResult := []int{1, 2}
		result := twoSum(nums, target)
		if !reflect.DeepEqual(result, expectResult) {
			t.Errorf(
				"TwoSumErr Test Fail: ExpectResult: %+v got: %+v",
				expectResult, result,
			)
		}
	})
}
