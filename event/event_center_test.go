package event

import (
    "strconv"
    "testing"
)

func TestHandelEvent(t *testing.T) {
    for i := 0; i < 10; i++ {
        index := strconv.Itoa(i)
        AddEvent(NewModel("type_"+index, "dec_"+index, "metada"+index))
    }
    HandelEvent()
    WaitEventHandleFinished()
}
