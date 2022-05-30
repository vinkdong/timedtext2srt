package timedtext2srt

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

type TimedText struct {
	WireMagic string  `json:"wire_magic"`
	Events    []Event `json:"events"`
}

type Event struct {
	TStartMs    int64 `json:"tStartMs"`
	DDurationMs int64 `json:"dDurationMs"`
	Segs        []Seg `json:"segs"`
}

type Seg struct {
	Utf8 string `json:"utf8"`
}

func (t *TimedText) ToSrt() []byte {
	to := bytes.Buffer{}
	for i, e := range t.Events {
		to.WriteString(strconv.Itoa(i + 1))
		to.WriteByte('\n')
		st := time.Unix(0, e.TStartMs*1e6)
		endMs := e.TStartMs + e.DDurationMs
		et := time.Unix(0, endMs*1e6)
		to.WriteString(fmt.Sprintf("%s.%d --> %s.%d\n",
			st.UTC().Format("15:04:05"), e.TStartMs%1000,
			et.UTC().Format("15:04:05"), endMs%1000,
		))
		to.WriteString(st.Format(e.Segs[0].Utf8) + "\n\n")
	}
	return to.Bytes()
}
