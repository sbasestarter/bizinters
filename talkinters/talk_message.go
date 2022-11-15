package talkinters

type TalkMessageType int

const (
	TalkMessageTypeUnknown TalkMessageType = iota
	TalkMessageTypeText
	TalkMessageTypeImage
)

type TalkMessageW struct {
	At              int64           `bson:"At" json:"at" yaml:"at"`
	CustomerMessage bool            `bson:"CustomerMessage" json:"customerMessage" yaml:"customerMessage"`
	Type            TalkMessageType `bson:"Type" json:"type" yaml:"type"`
	SenderID        uint64          `bson:"SenderID" json:"senderID" yaml:"senderID"`
	SenderUserName  string          `bson:"SenderUserName" json:"senderUserName" yaml:"senderUserName"`
	Text            string          `bson:"Text,omitempty" json:"text" yaml:"text"`
	Data            []byte          `bson:"Data,omitempty" json:"data" yaml:"data"`
}

type TalkMessageR struct {
	MessageID    string `bson:"_id" json:"messageID" yaml:"messageID"`
	TalkMessageW `bson:"inline" json:",inline" yaml:",inline"`
}
