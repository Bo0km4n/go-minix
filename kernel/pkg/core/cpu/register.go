package cpu

import "encoding/binary"

type reg8 interface {
	SetVal(uint8)
	GetVal() uint8
}
type reg16 interface {
	SetVal(uint16)
	GetVal() uint16
}

type AH struct {
	val uint8
}

func (ah *AH) GetVal() uint8 {
	return ah.val
}
func (ah *AH) SetVal(v uint8) {
	ah.val = v
}

type AL struct {
	val uint8
}

func (al *AL) GetVal() uint8 {
	return al.val
}
func (al *AL) SetVal(v uint8) {
	al.val = v
}

type AX struct {
	al reg8
	ah reg8
}

func (ax *AX) GetVal() uint16 {
	return binary.BigEndian.Uint16([]byte{ax.ah.GetVal(), ax.al.GetVal()})
}
func (ax *AX) SetVal(v uint16) {
	ah := uint8(v >> 8)
	al := uint8(v & 0x00ff)
	ax.ah.SetVal(ah)
	ax.al.SetVal(al)
}

type BH struct {
	val uint8
}

func (bh *BH) GetVal() uint8 {
	return bh.val
}
func (bh *BH) SetVal(v uint8) {
	bh.val = v
}

type BL struct {
	val uint8
}

func (bl *BL) GetVal() uint8 {
	return bl.val
}
func (bl *BL) SetVal(v uint8) {
	bl.val = v
}

type BX struct {
	bl reg8
	bh reg8
}

func (bx *BX) GetVal() uint16 {
	return binary.BigEndian.Uint16([]byte{bx.bh.GetVal(), bx.bl.GetVal()})
}
func (bx *BX) SetVal(v uint16) {
	bh := uint8(v >> 8)
	bl := uint8(v & 0x00ff)
	bx.bh.SetVal(bh)
	bx.bl.SetVal(bl)
}

type CH struct {
	val uint8
}

func (ch *CH) GetVal() uint8 {
	return ch.val
}
func (ch *CH) SetVal(v uint8) {
	ch.val = v
}

type CL struct {
	val uint8
}

func (cl *CL) GetVal() uint8 {
	return cl.val
}
func (cl *CL) SetVal(v uint8) {
	cl.val = v
}

type CX struct {
	cl reg8
	ch reg8
}

func (cx *CX) GetVal() uint16 {
	return binary.BigEndian.Uint16([]byte{cx.ch.GetVal(), cx.cl.GetVal()})
}
func (cx *CX) SetVal(v uint16) {
	ch := uint8(v >> 8)
	cl := uint8(v & 0x00ff)
	cx.ch.SetVal(ch)
	cx.cl.SetVal(cl)
}

type DH struct {
	val uint8
}

func (dh *DH) GetVal() uint8 {
	return dh.val
}

func (dh *DH) SetVal(v uint8) {
	dh.val = v
}

type DL struct {
	val uint8
}

func (dl *DL) GetVal() uint8 {
	return dl.val
}
func (dl *DL) SetVal(v uint8) {
	dl.val = v
}

type DX struct {
	dl reg8
	dh reg8
}

func (dx *DX) GetVal() uint16 {
	return binary.BigEndian.Uint16([]byte{dx.dh.GetVal(), dx.dl.GetVal()})
}
func (dx *DX) SetVal(v uint16) {
	dh := uint8(v >> 8)
	dl := uint8(v & 0x00ff)
	dx.dh.SetVal(dh)
	dx.dl.SetVal(dl)
}

type SP struct {
	val uint16
}

func (sp *SP) GetVal() uint16 {
	return sp.val
}

func (sp *SP) SetVal(v uint16) {
	sp.val = v
}

type BP struct {
	val uint16
}

func (bp *BP) GetVal() uint16 {
	return bp.val
}

func (bp *BP) SetVal(v uint16) {
	bp.val = v
}

type SI struct {
	val uint16
}

func (si *SI) GetVal() uint16 {
	return si.val
}

func (si *SI) SetVal(v uint16) {
	si.val = v
}

type DI struct {
	val uint16
}

func (di *DI) GetVal() uint16 {
	return di.val
}

func (di *DI) SetVal(v uint16) {
	di.val = v
}
