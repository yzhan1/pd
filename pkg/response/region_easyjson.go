// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package response

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	metapb "github.com/pingcap/kvproto/pkg/metapb"
	pdpb "github.com/pingcap/kvproto/pkg/pdpb"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson75d7afa0DecodeGithubComTikvPdServerApi(in *jlexer.Lexer, out *RegionInfo) {
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
			out.ID = uint64(in.Uint64())
		case "start_key":
			out.StartKey = string(in.String())
		case "end_key":
			out.EndKey = string(in.String())
		case "epoch":
			if in.IsNull() {
				in.Skip()
				out.RegionEpoch = nil
			} else {
				if out.RegionEpoch == nil {
					out.RegionEpoch = new(metapb.RegionEpoch)
				}
				easyjson75d7afa0DecodeGithubComPingcapKvprotoPkgMetapb(in, out.RegionEpoch)
			}
		case "peers":
			if in.IsNull() {
				in.Skip()
				out.Peers = nil
			} else {
				in.Delim('[')
				if out.Peers == nil {
					if !in.IsDelim(']') {
						out.Peers = make([]MetaPeer, 0, 2)
					} else {
						out.Peers = []MetaPeer{}
					}
				} else {
					out.Peers = (out.Peers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 MetaPeer
					easyjson75d7afa0DecodeGithubComTikvPdServerApi1(in, &v1)
					out.Peers = append(out.Peers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "leader":
			easyjson75d7afa0DecodeGithubComTikvPdServerApi1(in, &out.Leader)
		case "down_peers":
			if in.IsNull() {
				in.Skip()
				out.DownPeers = nil
			} else {
				in.Delim('[')
				if out.DownPeers == nil {
					if !in.IsDelim(']') {
						out.DownPeers = make([]PDPeerStats, 0, 1)
					} else {
						out.DownPeers = []PDPeerStats{}
					}
				} else {
					out.DownPeers = (out.DownPeers)[:0]
				}
				for !in.IsDelim(']') {
					var v2 PDPeerStats
					easyjson75d7afa0DecodeGithubComTikvPdServerApi2(in, &v2)
					out.DownPeers = append(out.DownPeers, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pending_peers":
			if in.IsNull() {
				in.Skip()
				out.PendingPeers = nil
			} else {
				in.Delim('[')
				if out.PendingPeers == nil {
					if !in.IsDelim(']') {
						out.PendingPeers = make([]MetaPeer, 0, 2)
					} else {
						out.PendingPeers = []MetaPeer{}
					}
				} else {
					out.PendingPeers = (out.PendingPeers)[:0]
				}
				for !in.IsDelim(']') {
					var v3 MetaPeer
					easyjson75d7afa0DecodeGithubComTikvPdServerApi1(in, &v3)
					out.PendingPeers = append(out.PendingPeers, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "cpu_usage":
			out.CPUUsage = uint64(in.Uint64())
		case "written_bytes":
			out.WrittenBytes = uint64(in.Uint64())
		case "read_bytes":
			out.ReadBytes = uint64(in.Uint64())
		case "written_keys":
			out.WrittenKeys = uint64(in.Uint64())
		case "read_keys":
			out.ReadKeys = uint64(in.Uint64())
		case "approximate_size":
			out.ApproximateSize = int64(in.Int64())
		case "approximate_keys":
			out.ApproximateKeys = int64(in.Int64())
		case "buckets":
			if in.IsNull() {
				in.Skip()
				out.Buckets = nil
			} else {
				in.Delim('[')
				if out.Buckets == nil {
					if !in.IsDelim(']') {
						out.Buckets = make([]string, 0, 4)
					} else {
						out.Buckets = []string{}
					}
				} else {
					out.Buckets = (out.Buckets)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Buckets = append(out.Buckets, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "replication_status":
			if in.IsNull() {
				in.Skip()
				out.ReplicationStatus = nil
			} else {
				if out.ReplicationStatus == nil {
					out.ReplicationStatus = new(ReplicationStatus)
				}
				easyjson75d7afa0DecodeGithubComTikvPdServerApi3(in, out.ReplicationStatus)
			}
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
func easyjson75d7afa0EncodeGithubComTikvPdServerApi(out *jwriter.Writer, in RegionInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"start_key\":"
		out.RawString(prefix)
		out.String(string(in.StartKey))
	}
	{
		const prefix string = ",\"end_key\":"
		out.RawString(prefix)
		out.String(string(in.EndKey))
	}
	if in.RegionEpoch != nil {
		const prefix string = ",\"epoch\":"
		out.RawString(prefix)
		easyjson75d7afa0EncodeGithubComPingcapKvprotoPkgMetapb(out, *in.RegionEpoch)
	}
	if len(in.Peers) != 0 {
		const prefix string = ",\"peers\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Peers {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson75d7afa0EncodeGithubComTikvPdServerApi1(out, v6)
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"leader\":"
		out.RawString(prefix)
		easyjson75d7afa0EncodeGithubComTikvPdServerApi1(out, in.Leader)
	}
	if len(in.DownPeers) != 0 {
		const prefix string = ",\"down_peers\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v7, v8 := range in.DownPeers {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjson75d7afa0EncodeGithubComTikvPdServerApi2(out, v8)
			}
			out.RawByte(']')
		}
	}
	if len(in.PendingPeers) != 0 {
		const prefix string = ",\"pending_peers\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v9, v10 := range in.PendingPeers {
				if v9 > 0 {
					out.RawByte(',')
				}
				easyjson75d7afa0EncodeGithubComTikvPdServerApi1(out, v10)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"cpu_usage\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.CPUUsage))
	}
	{
		const prefix string = ",\"written_bytes\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.WrittenBytes))
	}
	{
		const prefix string = ",\"read_bytes\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ReadBytes))
	}
	{
		const prefix string = ",\"written_keys\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.WrittenKeys))
	}
	{
		const prefix string = ",\"read_keys\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ReadKeys))
	}
	{
		const prefix string = ",\"approximate_size\":"
		out.RawString(prefix)
		out.Int64(int64(in.ApproximateSize))
	}
	{
		const prefix string = ",\"approximate_keys\":"
		out.RawString(prefix)
		out.Int64(int64(in.ApproximateKeys))
	}
	if len(in.Buckets) != 0 {
		const prefix string = ",\"buckets\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v11, v12 := range in.Buckets {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	if in.ReplicationStatus != nil {
		const prefix string = ",\"replication_status\":"
		out.RawString(prefix)
		easyjson75d7afa0EncodeGithubComTikvPdServerApi3(out, *in.ReplicationStatus)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v *RegionInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson75d7afa0EncodeGithubComTikvPdServerApi(&w, *v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v *RegionInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson75d7afa0EncodeGithubComTikvPdServerApi(w, *v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RegionInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson75d7afa0DecodeGithubComTikvPdServerApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RegionInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson75d7afa0DecodeGithubComTikvPdServerApi(l, v)
}
func easyjson75d7afa0DecodeGithubComTikvPdServerApi3(in *jlexer.Lexer, out *ReplicationStatus) {
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
		case "state":
			out.State = string(in.String())
		case "state_id":
			out.StateID = uint64(in.Uint64())
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
func easyjson75d7afa0EncodeGithubComTikvPdServerApi3(out *jwriter.Writer, in ReplicationStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"state\":"
		out.RawString(prefix[1:])
		out.String(string(in.State))
	}
	{
		const prefix string = ",\"state_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.StateID))
	}
	out.RawByte('}')
}
func easyjson75d7afa0DecodeGithubComTikvPdServerApi2(in *jlexer.Lexer, out *PDPeerStats) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.PeerStats = new(pdpb.PeerStats)
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
		case "peer":
			easyjson75d7afa0DecodeGithubComTikvPdServerApi1(in, &out.Peer)
		case "down_seconds":
			out.DownSeconds = uint64(in.Uint64())
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
func easyjson75d7afa0EncodeGithubComTikvPdServerApi2(out *jwriter.Writer, in PDPeerStats) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"peer\":"
		out.RawString(prefix[1:])
		easyjson75d7afa0EncodeGithubComTikvPdServerApi1(out, in.Peer)
	}
	if in.DownSeconds != 0 {
		const prefix string = ",\"down_seconds\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.DownSeconds))
	}
	out.RawByte('}')
}
func easyjson75d7afa0DecodeGithubComTikvPdServerApi1(in *jlexer.Lexer, out *MetaPeer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.Peer = new(metapb.Peer)
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
		case "role_name":
			out.RoleName = string(in.String())
		case "is_learner":
			out.IsLearner = bool(in.Bool())
		case "id":
			out.Id = uint64(in.Uint64())
		case "store_id":
			out.StoreId = uint64(in.Uint64())
		case "role":
			out.Role = metapb.PeerRole(in.Int32())
		case "is_witness":
			out.IsWitness = bool(in.Bool())
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
func easyjson75d7afa0EncodeGithubComTikvPdServerApi1(out *jwriter.Writer, in MetaPeer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"role_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.RoleName))
	}
	if in.IsLearner {
		const prefix string = ",\"is_learner\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsLearner))
	}
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Id))
	}
	if in.StoreId != 0 {
		const prefix string = ",\"store_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.StoreId))
	}
	if in.Role != 0 {
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		out.Int32(int32(in.Role))
	}
	if in.IsWitness {
		const prefix string = ",\"is_witness\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsWitness))
	}
	out.RawByte('}')
}
func easyjson75d7afa0DecodeGithubComPingcapKvprotoPkgMetapb(in *jlexer.Lexer, out *metapb.RegionEpoch) {
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
		case "conf_ver":
			out.ConfVer = uint64(in.Uint64())
		case "version":
			out.Version = uint64(in.Uint64())
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
func easyjson75d7afa0EncodeGithubComPingcapKvprotoPkgMetapb(out *jwriter.Writer, in metapb.RegionEpoch) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConfVer != 0 {
		const prefix string = ",\"conf_ver\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ConfVer))
	}
	if in.Version != 0 {
		const prefix string = ",\"version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Version))
	}
	out.RawByte('}')
}
