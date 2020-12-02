// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Gender struct {
	_tab flatbuffers.Struct
}

func (rcv *Gender) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Gender) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Gender) Male() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Gender) MutateMale(n bool) bool {
	return rcv._tab.MutateBool(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Gender) Female() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(1))
}
func (rcv *Gender) MutateFemale(n bool) bool {
	return rcv._tab.MutateBool(rcv._tab.Pos+flatbuffers.UOffsetT(1), n)
}

func CreateGender(builder *flatbuffers.Builder, male bool, female bool) flatbuffers.UOffsetT {
	builder.Prep(1, 2)
	builder.PrependBool(female)
	builder.PrependBool(male)
	return builder.Offset()
}
