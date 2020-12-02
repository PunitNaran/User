// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Buddhism struct {
	_tab flatbuffers.Table
}

func GetRootAsBuddhism(buf []byte, offset flatbuffers.UOffsetT) *Buddhism {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Buddhism{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Buddhism) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Buddhism) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Buddhism) Community() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func BuddhismStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BuddhismAddCommunity(builder *flatbuffers.Builder, community flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(community), 0)
}
func BuddhismEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
