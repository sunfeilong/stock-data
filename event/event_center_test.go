package event

import (
    "testing"
)

func TestHandelEvent(t *testing.T) {

    eventCenter := NewECenter(1000)

    for i := 0; i < 10; i++ {
        index := string(i)
        eventCenter.AddEvent(NewModel("type_"+index, "dec_"+index, "metada"+index))
    }

    eventCenter.HandelEvent()

    eventCenter.WaitFinished()
}
