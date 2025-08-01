package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})
}

// we want to describe to the compiler that we wish to use the == and != operators on things of type T
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	//  don't need to prove the same logic over and over

	// t.Run("string stack", func(t *testing.T) {
	// 	myStackOfStrings := new(Stack[string])

	// 	// check stack is empty
	// 	AssertTrue(t, myStackOfStrings.IsEmpty())

	// 	// add a thing, then check it's not empty
	// 	myStackOfStrings.Push("123")
	// 	AssertFalse(t, myStackOfStrings.IsEmpty())

	// 	// add another thing, pop it back again
	// 	myStackOfStrings.Push("456")
	// 	value, _ := myStackOfStrings.Pop()
	// 	AssertEqual(t, value, "456")
	// 	value, _ = myStackOfStrings.Pop()
	// 	AssertEqual(t, value, "123")
	// 	AssertTrue(t, myStackOfStrings.IsEmpty())
	// })

	t.Run("interface stack DX is horrid", func(t *testing.T) { 
		myStackOfInts := new(Stack[int])

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}
