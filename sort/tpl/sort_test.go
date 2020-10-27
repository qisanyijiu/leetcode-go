package tpl

import "testing"

func TestEqualSlice(t *testing.T) {
	var (
		input  = []int{1, 3, 5, 2, 4, 7, 6, 8, 9}
		except = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	)
	if !EqualSlice(except, except) {
		t.Error("error equal func")
	}
	if !EqualSlice(input, input) {
		t.Error("error equal func")
	}
	if EqualSlice(input, except) {
		t.Error("error equal func")
	}
}

func TestSort(t *testing.T) {
	t.Run("bubble", testBubbleSort)
	t.Run("heap", testHeapSort)
	t.Run("insertion", testInsertionSort)
	t.Run("merge", testMergeSort)
	t.Run("quick", testQuickSort)
	t.Run("selection", testSelectionSort)
	t.Run("shell", testShellSort)

}

func testBubbleSort(t *testing.T) {
	var (
		input  = []int{1, 3, 5, 2, 4, 7, 6, 8, 9}
		except = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	)

	BubbleSort(input)

	if !EqualSlice(input, except) {
		t.Errorf("\nexcept: %v\ngot:    %v", except, input)
	}
}

func testHeapSort(t *testing.T) {
	var (
		input  = []int{1, 3, 5, 2, 4, 7, 6, 8, 9}
		except = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	)

	HeapSort(input)

	if !EqualSlice(input, except) {
		t.Errorf("\nexcept: %v\ngot:    %v", except, input)
	}
}

func testInsertionSort(t *testing.T) {
	var (
		input  = []int{1, 3, 5, 2, 4, 7, 6, 8, 9}
		except = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	)

	InsertionSort(input)

	if !EqualSlice(input, except) {
		t.Errorf("\nexcept: %v\ngot:    %v", except, input)
	}
}

func testMergeSort(t *testing.T) {
	var (
		input  = []int{1, 3, 5, 2, 4, 7, 6, 8, 9}
		except = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	)

	MergeSort(input)

	if !EqualSlice(input, except) {
		t.Errorf("\nexcept: %v\ngot:    %v", except, input)
	}
}

func testQuickSort(t *testing.T) {
	var (
		input  = []int{1, 3, 5, 2, 4, 7, 6, 8, 9}
		except = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	)

	QuickSort(input)

	if !EqualSlice(input, except) {
		t.Errorf("\nexcept: %v\ngot:    %v", except, input)
	}
}

func testSelectionSort(t *testing.T) {
	var (
		input  = []int{1, 3, 5, 2, 4, 7, 6, 8, 9}
		except = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	)

	SelectionSort(input)

	if !EqualSlice(input, except) {
		t.Errorf("\nexcept: %v\ngot:    %v", except, input)
	}
}

func testShellSort(t *testing.T) {
	var (
		input  = []int{1, 3, 5, 2, 4, 7, 6, 8, 9}
		except = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	)

	ShellSort(input)

	if !EqualSlice(input, except) {
		t.Errorf("\nexcept: %v\ngot: %v", except, input)
	}
}
