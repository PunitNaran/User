// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Age struct {
	_tab flatbuffers.Struct
}

func (rcv *Age) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Age) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Age) Day() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Age) MutateDay(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Age) Month() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Age) MutateMonth(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *Age) Year() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *Age) MutateYear(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func CreateAge(builder *flatbuffers.Builder, day int32, month int32, year int32) flatbuffers.UOffsetT {
	builder.Prep(4, 12)
	builder.PrependInt32(year)
	builder.PrependInt32(month)
	builder.PrependInt32(day)
	return builder.Offset()
}
