//nolint:all
package main

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE3ab7953DecodeGithubComFixmeMyFriendHw10ProgramOptimization(in *jlexer.Lexer, out *User) {
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
		case "Email":
			out.Email = string(in.String())
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

func easyjsonE3ab7953EncodeGithubComFixmeMyFriendHw10ProgramOptimization(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface.
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE3ab7953EncodeGithubComFixmeMyFriendHw10ProgramOptimization(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface.
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE3ab7953EncodeGithubComFixmeMyFriendHw10ProgramOptimization(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface.
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE3ab7953DecodeGithubComFixmeMyFriendHw10ProgramOptimization(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface.
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE3ab7953DecodeGithubComFixmeMyFriendHw10ProgramOptimization(l, v)
}
