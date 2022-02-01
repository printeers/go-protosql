package protosql

import (
	"database/sql/driver"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Timestamp struct {
	timestamppb.Timestamp
}

func NewTimestamp(stdtime time.Time) *Timestamp {
	return &Timestamp{
		Timestamp: *timestamppb.New(stdtime),
	}
}

// Scan implements sql.Scanner for Timestamp.
func (t *Timestamp) Scan(value interface{}) error {
	stdtime, ok := value.(time.Time)
	if !ok {
		return errors.New("cannot scan unknown type into Timestamp")
	}
	t.Timestamp = *timestamppb.New(stdtime)
	return nil
}

// Value implements driver.Valuer for Timestamp.
func (t *Timestamp) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return t.Timestamp.AsTime(), nil
}

// TODO: Remove .proto and buf, just wrap-arround the generated proto methods?
