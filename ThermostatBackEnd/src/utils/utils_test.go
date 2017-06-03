package utils

import "testing"

func TestMakeTestFail(t *testing.T){
    t.Errorf("Expected  from popping the queue, got ")
}
