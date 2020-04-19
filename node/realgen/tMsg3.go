package realgen

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

// Code generated go-ipld-prime DO NOT EDIT.

type _Msg3 struct {
	whee _Int
	woot _Int
	waga _Int
}
type Msg3 = *_Msg3

func (n _Msg3) FieldWhee() Int {
	return &n.whee
}
func (n _Msg3) FieldWoot() Int {
	return &n.woot
}
func (n _Msg3) FieldWaga() Int {
	return &n.waga
}

type _Msg3__Maybe struct {
	m schema.Maybe
	v Msg3
}
type MaybeMsg3 = *_Msg3__Maybe

func (m MaybeMsg3) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMsg3) IsUndefined() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMsg3) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMsg3) AsNode() ipld.Node {
	switch m.m {
	case schema.Maybe_Absent:
		return ipld.Undef
	case schema.Maybe_Null:
		return ipld.Null
	case schema.Maybe_Value:
		return m.v
	default:
		panic("unreachable")
	}
}
func (m MaybeMsg3) Must() Msg3 {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var (
	fieldName__Msg3_Whee = _String{"whee"}
	fieldName__Msg3_Woot = _String{"woot"}
	fieldName__Msg3_Waga = _String{"waga"}
)
var _ ipld.Node = (Msg3)(&_Msg3{})
var _ schema.TypedNode = (Msg3)(&_Msg3{})

func (Msg3) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n Msg3) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "whee":
		return &n.whee, nil
	case "woot":
		return &n.woot, nil
	case "waga":
		return &n.waga, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n Msg3) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupString(ks)
}
func (Msg3) LookupIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"realgen.Msg3"}.LookupIndex(0)
}
func (n Msg3) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (n Msg3) MapIterator() ipld.MapIterator {
	return &_Msg3__MapItr{n, 0}
}

type _Msg3__MapItr struct {
	n   Msg3
	idx int
}

func (itr *_Msg3__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 3 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__Msg3_Whee
		v = &itr.n.whee
	case 1:
		k = &fieldName__Msg3_Woot
		v = &itr.n.woot
	case 2:
		k = &fieldName__Msg3_Waga
		v = &itr.n.waga
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_Msg3__MapItr) Done() bool {
	return itr.idx >= 3
}

func (Msg3) ListIterator() ipld.ListIterator {
	return nil
}
func (Msg3) Length() int {
	return 3
}
func (Msg3) IsUndefined() bool {
	return false
}
func (Msg3) IsNull() bool {
	return false
}
func (Msg3) AsBool() (bool, error) {
	return mixins.Map{"realgen.Msg3"}.AsBool()
}
func (Msg3) AsInt() (int, error) {
	return mixins.Map{"realgen.Msg3"}.AsInt()
}
func (Msg3) AsFloat() (float64, error) {
	return mixins.Map{"realgen.Msg3"}.AsFloat()
}
func (Msg3) AsString() (string, error) {
	return mixins.Map{"realgen.Msg3"}.AsString()
}
func (Msg3) AsBytes() ([]byte, error) {
	return mixins.Map{"realgen.Msg3"}.AsBytes()
}
func (Msg3) AsLink() (ipld.Link, error) {
	return mixins.Map{"realgen.Msg3"}.AsLink()
}
func (Msg3) Style() ipld.NodeStyle {
	return _Msg3__Style{}
}

type _Msg3__Style struct{}

func (_Msg3__Style) NewBuilder() ipld.NodeBuilder {
	var nb _Msg3__Builder
	nb.Reset()
	return &nb
}

type _Msg3__Builder struct {
	_Msg3__Assembler
}

func (nb *_Msg3__Builder) Build() ipld.Node {
	if nb.state != maState_finished {
		panic("invalid state: assembler for realgen.Msg3 must be 'finished' before Build can be called!")
	}
	return nb.w
}
func (nb *_Msg3__Builder) Reset() {
	var w _Msg3
	var m schema.Maybe
	*nb = _Msg3__Builder{_Msg3__Assembler{w: &w, m: &m, state: maState_initial}}
}

type _Msg3__Assembler struct {
	w     *_Msg3
	m     *schema.Maybe
	state maState
	s     int
	f     int

	cm      schema.Maybe
	ca_whee _Int__Assembler
	ca_woot _Int__Assembler
	ca_waga _Int__Assembler
}

func (na *_Msg3__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_whee.reset()
	na.ca_woot.reset()
	na.ca_waga.reset()
}

var (
	fieldBit__Msg3_Whee        = 1 << 0
	fieldBit__Msg3_Woot        = 1 << 1
	fieldBit__Msg3_Waga        = 1 << 2
	fieldBits__Msg3_sufficient = 0 + 1<<0 + 1<<1 + 1<<2
)

func (na *_Msg3__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_Msg3{}
	}
	return na, nil
}
func (_Msg3__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"realgen.Msg3"}.BeginList(0)
}
func (na *_Msg3__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"realgen.Msg3"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Msg3__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"realgen.Msg3"}.AssignBool(false)
}
func (_Msg3__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"realgen.Msg3"}.AssignInt(0)
}
func (_Msg3__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"realgen.Msg3"}.AssignFloat(0)
}
func (_Msg3__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"realgen.Msg3"}.AssignString("")
}
func (_Msg3__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"realgen.Msg3"}.AssignBytes(nil)
}
func (_Msg3__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"realgen.Msg3"}.AssignLink(nil)
}
func (na *_Msg3__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Msg3); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.ReprKind() != ipld.ReprKind_Map {
		return ipld.ErrWrongKind{TypeName: "realgen.Msg3", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
	}
	itr := v.MapIterator()
	for !itr.Done() {
		k, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleKey().AssignNode(k); err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_Msg3__Assembler) Style() ipld.NodeStyle {
	return _Msg3__Style{}
}
func (ma *_Msg3__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_whee.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_woot.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 2:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_waga.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	default:
		panic("unreachable")
	}
}
func (ma *_Msg3__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleEntry cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleEntry cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleEntry cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}
	switch k {
	case "whee":
		if ma.s&fieldBit__Msg3_Whee != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__Msg3_Whee}
		}
		ma.s += fieldBit__Msg3_Whee
		ma.state = maState_midValue
		ma.ca_whee.w = &ma.w.whee
		ma.ca_whee.m = &ma.cm
		return &ma.ca_whee, nil
	case "woot":
		if ma.s&fieldBit__Msg3_Woot != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__Msg3_Woot}
		}
		ma.s += fieldBit__Msg3_Woot
		ma.state = maState_midValue
		ma.ca_woot.w = &ma.w.woot
		ma.ca_woot.m = &ma.cm
		return &ma.ca_woot, nil
	case "waga":
		if ma.s&fieldBit__Msg3_Waga != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__Msg3_Waga}
		}
		ma.s += fieldBit__Msg3_Waga
		ma.state = maState_midValue
		ma.ca_waga.w = &ma.w.waga
		ma.ca_waga.m = &ma.cm
		return &ma.ca_waga, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName: "realgen.Msg3", Key: &_String{k}}
	}
}
func (ma *_Msg3__Assembler) AssembleKey() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleKey cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleKey cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleKey cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midKey
	return (*_Msg3__KeyAssembler)(ma)
}
func (ma *_Msg3__Assembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	switch ma.f {
	case 0:
		ma.ca_whee.w = &ma.w.whee
		ma.ca_whee.m = &ma.cm
		return &ma.ca_whee
	case 1:
		ma.ca_woot.w = &ma.w.woot
		ma.ca_woot.m = &ma.cm
		return &ma.ca_woot
	case 2:
		ma.ca_waga.w = &ma.w.waga
		ma.ca_waga.m = &ma.cm
		return &ma.ca_waga
	default:
		panic("unreachable")
	}
}
func (ma *_Msg3__Assembler) Finish() error {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: Finish cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		panic("invalid state: Finish cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	//FIXME check if all required fields are set
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Msg3__Assembler) KeyStyle() ipld.NodeStyle {
	return _String__Style{}
}
func (ma *_Msg3__Assembler) ValueStyle(k string) ipld.NodeStyle {
	panic("todo structbuilder mapassembler valuestyle")
}

type _Msg3__KeyAssembler _Msg3__Assembler

func (_Msg3__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"realgen.Msg3.KeyAssembler"}.BeginMap(0)
}
func (_Msg3__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"realgen.Msg3.KeyAssembler"}.BeginList(0)
}
func (na *_Msg3__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"realgen.Msg3.KeyAssembler"}.AssignNull()
}
func (_Msg3__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"realgen.Msg3.KeyAssembler"}.AssignBool(false)
}
func (_Msg3__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"realgen.Msg3.KeyAssembler"}.AssignInt(0)
}
func (_Msg3__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"realgen.Msg3.KeyAssembler"}.AssignFloat(0)
}
func (ka *_Msg3__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "whee":
		if ka.s&fieldBit__Msg3_Whee != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__Msg3_Whee}
		}
		ka.s += fieldBit__Msg3_Whee
		ka.state = maState_expectValue
		ka.f = 0
	case "woot":
		if ka.s&fieldBit__Msg3_Woot != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__Msg3_Woot}
		}
		ka.s += fieldBit__Msg3_Woot
		ka.state = maState_expectValue
		ka.f = 1
	case "waga":
		if ka.s&fieldBit__Msg3_Waga != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__Msg3_Waga}
		}
		ka.s += fieldBit__Msg3_Waga
		ka.state = maState_expectValue
		ka.f = 2
	default:
		return ipld.ErrInvalidKey{TypeName: "realgen.Msg3", Key: &_String{k}}
	}
	return nil
}
func (_Msg3__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"realgen.Msg3.KeyAssembler"}.AssignBytes(nil)
}
func (_Msg3__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"realgen.Msg3.KeyAssembler"}.AssignLink(nil)
}
func (ka *_Msg3__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_Msg3__KeyAssembler) Style() ipld.NodeStyle {
	return _String__Style{}
}
func (Msg3) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Msg3) Representation() ipld.Node {
	return (*_Msg3__Repr)(n)
}

type _Msg3__Repr _Msg3

var (
	fieldName__Msg3_Whee_serial = _String{"whee"}
	fieldName__Msg3_Woot_serial = _String{"woot"}
	fieldName__Msg3_Waga_serial = _String{"waga"}
)
var _ ipld.Node = &_Msg3__Repr{}

func (_Msg3__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n *_Msg3__Repr) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "whee":
		return n.whee.Representation(), nil
	case "woot":
		return n.woot.Representation(), nil
	case "waga":
		return n.waga.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n *_Msg3__Repr) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupString(ks)
}
func (_Msg3__Repr) LookupIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"realgen.Msg3.Repr"}.LookupIndex(0)
}
func (n _Msg3__Repr) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (n *_Msg3__Repr) MapIterator() ipld.MapIterator {
	return &_Msg3__ReprMapItr{n, 0}
}

type _Msg3__ReprMapItr struct {
	n   *_Msg3__Repr
	idx int
}

func (itr *_Msg3__ReprMapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 3 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__Msg3_Whee_serial
		v = itr.n.whee.Representation()
	case 1:
		k = &fieldName__Msg3_Woot_serial
		v = itr.n.woot.Representation()
	case 2:
		k = &fieldName__Msg3_Waga_serial
		v = itr.n.waga.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_Msg3__ReprMapItr) Done() bool {
	return itr.idx >= 3
}
func (_Msg3__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_Msg3__Repr) Length() int {
	l := 3
	return l
}
func (_Msg3__Repr) IsUndefined() bool {
	return false
}
func (_Msg3__Repr) IsNull() bool {
	return false
}
func (_Msg3__Repr) AsBool() (bool, error) {
	return mixins.Map{"realgen.Msg3.Repr"}.AsBool()
}
func (_Msg3__Repr) AsInt() (int, error) {
	return mixins.Map{"realgen.Msg3.Repr"}.AsInt()
}
func (_Msg3__Repr) AsFloat() (float64, error) {
	return mixins.Map{"realgen.Msg3.Repr"}.AsFloat()
}
func (_Msg3__Repr) AsString() (string, error) {
	return mixins.Map{"realgen.Msg3.Repr"}.AsString()
}
func (_Msg3__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"realgen.Msg3.Repr"}.AsBytes()
}
func (_Msg3__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"realgen.Msg3.Repr"}.AsLink()
}
func (_Msg3__Repr) Style() ipld.NodeStyle {
	return _Msg3__ReprStyle{}
}

type _Msg3__ReprStyle struct{}

func (_Msg3__ReprStyle) NewBuilder() ipld.NodeBuilder {
	var nb _Msg3__ReprBuilder
	nb.Reset()
	return &nb
}

type _Msg3__ReprBuilder struct {
	_Msg3__ReprAssembler
}

func (nb *_Msg3__ReprBuilder) Build() ipld.Node {
	if nb.state != maState_finished {
		panic("invalid state: assembler for realgen.Msg3.Repr must be 'finished' before Build can be called!")
	}
	return nb.w
}
func (nb *_Msg3__ReprBuilder) Reset() {
	var w _Msg3
	var m schema.Maybe
	*nb = _Msg3__ReprBuilder{_Msg3__ReprAssembler{w: &w, m: &m, state: maState_initial}}
}

type _Msg3__ReprAssembler struct {
	w     *_Msg3
	m     *schema.Maybe
	state maState
	s     int
	f     int

	cm      schema.Maybe
	ca_whee _Int__ReprAssembler
	ca_woot _Int__ReprAssembler
	ca_waga _Int__ReprAssembler
}

func (na *_Msg3__ReprAssembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_whee.reset()
	na.ca_woot.reset()
	na.ca_waga.reset()
}
func (na *_Msg3__ReprAssembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_Msg3{}
	}
	return na, nil
}
func (_Msg3__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"realgen.Msg3.Repr"}.BeginList(0)
}
func (na *_Msg3__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"realgen.Msg3.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Msg3__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"realgen.Msg3.Repr"}.AssignBool(false)
}
func (_Msg3__ReprAssembler) AssignInt(int) error {
	return mixins.MapAssembler{"realgen.Msg3.Repr"}.AssignInt(0)
}
func (_Msg3__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"realgen.Msg3.Repr"}.AssignFloat(0)
}
func (_Msg3__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"realgen.Msg3.Repr"}.AssignString("")
}
func (_Msg3__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"realgen.Msg3.Repr"}.AssignBytes(nil)
}
func (_Msg3__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"realgen.Msg3.Repr"}.AssignLink(nil)
}
func (na *_Msg3__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Msg3); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.ReprKind() != ipld.ReprKind_Map {
		return ipld.ErrWrongKind{TypeName: "realgen.Msg3.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
	}
	itr := v.MapIterator()
	for !itr.Done() {
		k, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleKey().AssignNode(k); err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_Msg3__ReprAssembler) Style() ipld.NodeStyle {
	return _Msg3__ReprStyle{}
}
func (ma *_Msg3__ReprAssembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 2:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	default:
		panic("unreachable")
	}
}
func (ma *_Msg3__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleEntry cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleEntry cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleEntry cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}
	switch k {
	case "whee":
		if ma.s&fieldBit__Msg3_Whee != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__Msg3_Whee_serial}
		}
		ma.s += fieldBit__Msg3_Whee
		ma.state = maState_midValue
		ma.ca_whee.w = &ma.w.whee
		ma.ca_whee.m = &ma.cm
		return &ma.ca_whee, nil
	case "woot":
		if ma.s&fieldBit__Msg3_Woot != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__Msg3_Woot_serial}
		}
		ma.s += fieldBit__Msg3_Woot
		ma.state = maState_midValue
		ma.ca_woot.w = &ma.w.woot
		ma.ca_woot.m = &ma.cm
		return &ma.ca_woot, nil
	case "waga":
		if ma.s&fieldBit__Msg3_Waga != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__Msg3_Waga_serial}
		}
		ma.s += fieldBit__Msg3_Waga
		ma.state = maState_midValue
		ma.ca_waga.w = &ma.w.waga
		ma.ca_waga.m = &ma.cm
		return &ma.ca_waga, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName: "realgen.Msg3.Repr", Key: &_String{k}}
	}
}
func (ma *_Msg3__ReprAssembler) AssembleKey() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleKey cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleKey cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleKey cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midKey
	return (*_Msg3__ReprKeyAssembler)(ma)
}
func (ma *_Msg3__ReprAssembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	switch ma.f {
	case 0:
		ma.ca_whee.w = &ma.w.whee
		ma.ca_whee.m = &ma.cm
		return &ma.ca_whee
	case 1:
		ma.ca_woot.w = &ma.w.woot
		ma.ca_woot.m = &ma.cm
		return &ma.ca_woot
	case 2:
		ma.ca_waga.w = &ma.w.waga
		ma.ca_waga.m = &ma.cm
		return &ma.ca_waga
	default:
		panic("unreachable")
	}
}
func (ma *_Msg3__ReprAssembler) Finish() error {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: Finish cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		panic("invalid state: Finish cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	//FIXME check if all required fields are set
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Msg3__ReprAssembler) KeyStyle() ipld.NodeStyle {
	return _String__Style{}
}
func (ma *_Msg3__ReprAssembler) ValueStyle(k string) ipld.NodeStyle {
	panic("todo structbuilder mapassembler repr valuestyle")
}

type _Msg3__ReprKeyAssembler _Msg3__ReprAssembler

func (_Msg3__ReprKeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"realgen.Msg3.Repr.ReprKeyAssembler"}.BeginMap(0)
}
func (_Msg3__ReprKeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"realgen.Msg3.Repr.ReprKeyAssembler"}.BeginList(0)
}
func (na *_Msg3__ReprKeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"realgen.Msg3.Repr.ReprKeyAssembler"}.AssignNull()
}
func (_Msg3__ReprKeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"realgen.Msg3.Repr.ReprKeyAssembler"}.AssignBool(false)
}
func (_Msg3__ReprKeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"realgen.Msg3.Repr.ReprKeyAssembler"}.AssignInt(0)
}
func (_Msg3__ReprKeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"realgen.Msg3.Repr.ReprKeyAssembler"}.AssignFloat(0)
}
func (ka *_Msg3__ReprKeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "whee":
		if ka.s&fieldBit__Msg3_Whee != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__Msg3_Whee_serial}
		}
		ka.s += fieldBit__Msg3_Whee
		ka.state = maState_expectValue
		ka.f = 0
	case "woot":
		if ka.s&fieldBit__Msg3_Woot != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__Msg3_Woot_serial}
		}
		ka.s += fieldBit__Msg3_Woot
		ka.state = maState_expectValue
		ka.f = 1
	case "waga":
		if ka.s&fieldBit__Msg3_Waga != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__Msg3_Waga_serial}
		}
		ka.s += fieldBit__Msg3_Waga
		ka.state = maState_expectValue
		ka.f = 2
	default:
		return ipld.ErrInvalidKey{TypeName: "realgen.Msg3.Repr", Key: &_String{k}}
	}
	return nil
}
func (_Msg3__ReprKeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"realgen.Msg3.Repr.ReprKeyAssembler"}.AssignBytes(nil)
}
func (_Msg3__ReprKeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"realgen.Msg3.Repr.ReprKeyAssembler"}.AssignLink(nil)
}
func (ka *_Msg3__ReprKeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_Msg3__ReprKeyAssembler) Style() ipld.NodeStyle {
	return _String__Style{}
}
