package common

import (
	"fmt"
	"time"
)

type TimestampStyle string

const (
	StyleShortTime     TimestampStyle = "t"
	StyleLongTime      TimestampStyle = "T"
	StyleShortDate     TimestampStyle = "d"
	StyleLongDate      TimestampStyle = "D"
	StyleShortDateTime TimestampStyle = "f"
	StyleLongDateTime  TimestampStyle = "F"
	StyleRelative      TimestampStyle = "R"
)

func NewDiscordTime(t time.Time) *DiscordTime {
	return &DiscordTime{t}
}

type DiscordTime struct {
	time.Time
}

func (t *DiscordTime) Format(style TimestampStyle) string {
	return fmt.Sprintf("<t:%d:%s>", t.Unix(), style)
}

func (t *DiscordTime) String() string {
	return t.ShortDateTime()
}

func (t *DiscordTime) ShortTime() string {
	return t.Format(StyleShortTime)
}

func (t *DiscordTime) LongTime() string {
	return t.Format(StyleLongTime)
}

func (t *DiscordTime) ShortDate() string {
	return t.Format(StyleShortDate)
}

func (t *DiscordTime) LongDate() string {
	return t.Format(StyleLongDate)
}

func (t *DiscordTime) ShortDateTime() string {
	return t.Format(StyleShortDateTime)
}

func (t *DiscordTime) LongDateTime() string {
	return t.Format(StyleLongDateTime)
}

func (t *DiscordTime) Relative() string {
	return t.Format(StyleRelative)
}
