// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package events

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF27e5b1aDecodeGithubComGui774umeKriePkgKrieEvents(in *jlexer.Lexer, out *BPFFilterEventSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.BPFFilterEvent = new(BPFFilterEvent)
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cmd":
			out.Cmd = BPFFilterCmd(in.Uint32())
		case "family":
			out.Family = AddressFamily(in.Uint16())
		case "type":
			out.Type = SocketType(in.Uint32())
		case "protocol":
			out.Protocol = L3Protocol(in.Uint16())
		case "prog_len":
			out.ProgLen = uint16(in.Uint16())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF27e5b1aEncodeGithubComGui774umeKriePkgKrieEvents(out *jwriter.Writer, in BPFFilterEventSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Cmd != 0 {
		const prefix string = ",\"cmd\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.Cmd).MarshalJSON())
	}
	if in.Family != 0 {
		const prefix string = ",\"family\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Family).MarshalJSON())
	}
	if in.Type != 0 {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Type).MarshalJSON())
	}
	if in.Protocol != 0 {
		const prefix string = ",\"protocol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Protocol).MarshalJSON())
	}
	if in.ProgLen != 0 {
		const prefix string = ",\"prog_len\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint16(uint16(in.ProgLen))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BPFFilterEventSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF27e5b1aEncodeGithubComGui774umeKriePkgKrieEvents(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BPFFilterEventSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF27e5b1aDecodeGithubComGui774umeKriePkgKrieEvents(l, v)
}
func easyjsonF27e5b1aDecodeGithubComGui774umeKriePkgKrieEvents1(in *jlexer.Lexer, out *BPFEventSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "map":
			if in.IsNull() {
				in.Skip()
				out.Map = nil
			} else {
				if out.Map == nil {
					out.Map = new(BPFMap)
				}
				easyjsonF27e5b1aDecodeGithubComGui774umeKriePkgKrieEvents2(in, out.Map)
			}
		case "program":
			if in.IsNull() {
				in.Skip()
				out.Program = nil
			} else {
				if out.Program == nil {
					out.Program = new(BPFProgram)
				}
				easyjsonF27e5b1aDecodeGithubComGui774umeKriePkgKrieEvents3(in, out.Program)
			}
		case "cmd":
			out.Cmd = BPFCmd(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF27e5b1aEncodeGithubComGui774umeKriePkgKrieEvents1(out *jwriter.Writer, in BPFEventSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Map != nil {
		const prefix string = ",\"map\":"
		first = false
		out.RawString(prefix[1:])
		easyjsonF27e5b1aEncodeGithubComGui774umeKriePkgKrieEvents2(out, *in.Map)
	}
	if in.Program != nil {
		const prefix string = ",\"program\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonF27e5b1aEncodeGithubComGui774umeKriePkgKrieEvents3(out, *in.Program)
	}
	{
		const prefix string = ",\"cmd\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Cmd).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BPFEventSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF27e5b1aEncodeGithubComGui774umeKriePkgKrieEvents1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BPFEventSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF27e5b1aDecodeGithubComGui774umeKriePkgKrieEvents1(l, v)
}
func easyjsonF27e5b1aDecodeGithubComGui774umeKriePkgKrieEvents3(in *jlexer.Lexer, out *BPFProgram) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint32(in.Uint32())
		case "type":
			out.Type = BPFProgramType(in.Uint32())
		case "attach_type":
			out.AttachType = BPFAttachType(in.Uint32())
		case "helpers":
			if in.IsNull() {
				in.Skip()
				out.Helpers = nil
			} else {
				in.Delim('[')
				if out.Helpers == nil {
					if !in.IsDelim(']') {
						out.Helpers = make(BPFHelperFuncList, 0, 16)
					} else {
						out.Helpers = BPFHelperFuncList{}
					}
				} else {
					out.Helpers = (out.Helpers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 BPFHelperFunc
					v1 = BPFHelperFunc(in.Uint32())
					out.Helpers = append(out.Helpers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "tag":
			out.Tag = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF27e5b1aEncodeGithubComGui774umeKriePkgKrieEvents3(out *jwriter.Writer, in BPFProgram) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.ID))
	}
	if in.Type != 0 {
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.Raw((in.Type).MarshalJSON())
	}
	if in.AttachType != 0 {
		const prefix string = ",\"attach_type\":"
		out.RawString(prefix)
		out.Raw((in.AttachType).MarshalJSON())
	}
	if len(in.Helpers) != 0 {
		const prefix string = ",\"helpers\":"
		out.RawString(prefix)
		out.Raw((in.Helpers).MarshalJSON())
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.Tag != "" {
		const prefix string = ",\"tag\":"
		out.RawString(prefix)
		out.String(string(in.Tag))
	}
	out.RawByte('}')
}
func easyjsonF27e5b1aDecodeGithubComGui774umeKriePkgKrieEvents2(in *jlexer.Lexer, out *BPFMap) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint32(in.Uint32())
		case "type":
			out.Type = BPFMapType(in.Uint32())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF27e5b1aEncodeGithubComGui774umeKriePkgKrieEvents2(out *jwriter.Writer, in BPFMap) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.ID))
	}
	if in.Type != 0 {
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.Raw((in.Type).MarshalJSON())
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}