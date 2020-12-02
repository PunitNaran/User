// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Religon struct {
	_tab flatbuffers.Table
}

func GetRootAsReligon(buf []byte, offset flatbuffers.UOffsetT) *Religon {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Religon{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Religon) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Religon) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Religon) Any() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateAny(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *Religon) Christian(obj *Christian) *Christian {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Christian)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Religon) Islam(obj *Islam) *Islam {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Islam)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Religon) Agnostic() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateAgnostic(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func (rcv *Religon) Nonreligious() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateNonreligious(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func (rcv *Religon) Hinduism(obj *Hinduism) *Hinduism {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Hinduism)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Religon) Confucianism() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateConfucianism(n bool) bool {
	return rcv._tab.MutateBoolSlot(16, n)
}

func (rcv *Religon) Taoism() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateTaoism(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

func (rcv *Religon) Buddhism(obj *Buddhism) *Buddhism {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Buddhism)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Religon) Sikhism() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateSikhism(n bool) bool {
	return rcv._tab.MutateBoolSlot(22, n)
}

func (rcv *Religon) Juche() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateJuche(n bool) bool {
	return rcv._tab.MutateBoolSlot(24, n)
}

func (rcv *Religon) Spiritism() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateSpiritism(n bool) bool {
	return rcv._tab.MutateBoolSlot(26, n)
}

func (rcv *Religon) Judaism() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateJudaism(n bool) bool {
	return rcv._tab.MutateBoolSlot(28, n)
}

func (rcv *Religon) Bahai() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateBahai(n bool) bool {
	return rcv._tab.MutateBoolSlot(30, n)
}

func (rcv *Religon) Jainism() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateJainism(n bool) bool {
	return rcv._tab.MutateBoolSlot(32, n)
}

func (rcv *Religon) Shinto() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateShinto(n bool) bool {
	return rcv._tab.MutateBoolSlot(34, n)
}

func (rcv *Religon) Caodai() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateCaodai(n bool) bool {
	return rcv._tab.MutateBoolSlot(36, n)
}

func (rcv *Religon) Other() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Religon) MutateOther(n bool) bool {
	return rcv._tab.MutateBoolSlot(38, n)
}

func ReligonStart(builder *flatbuffers.Builder) {
	builder.StartObject(18)
}
func ReligonAddAny(builder *flatbuffers.Builder, any bool) {
	builder.PrependBoolSlot(0, any, false)
}
func ReligonAddChristian(builder *flatbuffers.Builder, christian flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(christian), 0)
}
func ReligonAddIslam(builder *flatbuffers.Builder, islam flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(islam), 0)
}
func ReligonAddAgnostic(builder *flatbuffers.Builder, agnostic bool) {
	builder.PrependBoolSlot(3, agnostic, false)
}
func ReligonAddNonreligious(builder *flatbuffers.Builder, nonreligious bool) {
	builder.PrependBoolSlot(4, nonreligious, false)
}
func ReligonAddHinduism(builder *flatbuffers.Builder, hinduism flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(hinduism), 0)
}
func ReligonAddConfucianism(builder *flatbuffers.Builder, confucianism bool) {
	builder.PrependBoolSlot(6, confucianism, false)
}
func ReligonAddTaoism(builder *flatbuffers.Builder, taoism bool) {
	builder.PrependBoolSlot(7, taoism, false)
}
func ReligonAddBuddhism(builder *flatbuffers.Builder, buddhism flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(buddhism), 0)
}
func ReligonAddSikhism(builder *flatbuffers.Builder, sikhism bool) {
	builder.PrependBoolSlot(9, sikhism, false)
}
func ReligonAddJuche(builder *flatbuffers.Builder, juche bool) {
	builder.PrependBoolSlot(10, juche, false)
}
func ReligonAddSpiritism(builder *flatbuffers.Builder, spiritism bool) {
	builder.PrependBoolSlot(11, spiritism, false)
}
func ReligonAddJudaism(builder *flatbuffers.Builder, judaism bool) {
	builder.PrependBoolSlot(12, judaism, false)
}
func ReligonAddBahai(builder *flatbuffers.Builder, bahai bool) {
	builder.PrependBoolSlot(13, bahai, false)
}
func ReligonAddJainism(builder *flatbuffers.Builder, jainism bool) {
	builder.PrependBoolSlot(14, jainism, false)
}
func ReligonAddShinto(builder *flatbuffers.Builder, shinto bool) {
	builder.PrependBoolSlot(15, shinto, false)
}
func ReligonAddCaodai(builder *flatbuffers.Builder, caodai bool) {
	builder.PrependBoolSlot(16, caodai, false)
}
func ReligonAddOther(builder *flatbuffers.Builder, other bool) {
	builder.PrependBoolSlot(17, other, false)
}
func ReligonEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
