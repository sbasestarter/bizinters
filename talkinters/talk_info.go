package talkinters

type TalkStatus int

const (
	TalkStatusNone TalkStatus = iota
	TalkStatusOpened
	TalkStatusClosed
)

type TalkInfoW struct {
	Status          TalkStatus `bson:"Status" json:"status" yaml:"status"`
	Title           string     `bson:"Title" json:"title" yaml:"title"`
	StartAt         int64      `bson:"StartAt" json:"startAt" yaml:"startAt"`
	FinishedAt      int64      `bson:"FinishedAt" json:"finishedAt" yaml:"finishedAt"`
	CreatorID       uint64     `bson:"CreatorID" json:"creatorID" yaml:"creatorID"`
	ServiceID       uint64     `bson:"ServiceID" json:"serviceID" yaml:"serviceID"`
	CreatorUserName string     `bson:"CreatorUserName" json:"creatorUserName" yaml:"creatorUserName"`
	ActID           string     `bson:"ActID" json:"ActID,omitempty" yaml:"ActID,omitempty"`
	BizID           string     `bson:"BizID" json:"BizID,omitempty" yaml:"BizID,omitempty"`
}

func (w *TalkInfoW) Valid() (f bool) {
	if w == nil {
		return
	}

	if w.Title == "" || w.CreatorID == 0 || w.CreatorUserName == "" || w.ActID == "" || w.BizID == "" {
		return
	}

	f = true

	return
}

type TalkInfoR struct {
	TalkID    string `bson:"_id" json:"talkID" yaml:"talkID"`
	TalkInfoW `bson:"inline" json:",inline" yaml:",inline"`
}
