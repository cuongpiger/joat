package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()
	queue.Push(1)
	queue.Push(2)

	val, ok := queue.Pop()
	if !ok || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}

	if queue.Empty() {
		t.Errorf("Expected non-empty queue")
	}

	if queue.Size() != 1 {
		t.Errorf("Expected size 1, got %d", queue.Size())
	}

	val, ok = queue.Top()
	if !ok || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}

	val, ok = queue.Pop()
	if !ok || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}

	if !queue.Empty() {
		t.Errorf("Expected empty queue")
	}

	if queue.Size() != 0 {
		t.Errorf("Expected size 0, got %d", queue.Size())
	}
}
