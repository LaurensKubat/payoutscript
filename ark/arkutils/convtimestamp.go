package arkutils

import (
	ark "github.com/ArkEcosystem/go-client/client/two"
	"time"
)

func ParseArkTs(timestamp ark.Timestamp) time.Time {
	return time.Unix(int64(timestamp.Unix), 0)
}

func ParseTs(ts time.Time) ark.Timestamp {
	unix := int32(ts.Unix())
	human := time.Time.Format(ts, "Mon Jan 2 15:04:05 -0700 MST 2006") // formats ts according to the string format
	epoch := int32(ts.Unix() - 1490101200)
	return ark.Timestamp{Unix: unix, Epoch: epoch, Human: human}
}
