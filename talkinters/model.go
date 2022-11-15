package talkinters

import "context"

type Model interface {
	CreateTalk(ctx context.Context, talkInfo *TalkInfoW) (talkID string, err error)
	OpenTalk(ctx context.Context, talkID string) (err error)
	CloseTalk(ctx context.Context, talkID string) error

	AddTalkMessage(ctx context.Context, talkID string, message *TalkMessageW) (err error)
	GetTalkMessages(ctx context.Context, talkID string, offset, count int64) (messages []*TalkMessageR, err error)

	QueryTalks(ctx context.Context, creatorID, serviceID uint64, talkID string,
		statuses []TalkStatus) (talks []*TalkInfoR, err error)
	GetPendingTalkInfos(ctx context.Context) ([]*TalkInfoR, error)
	UpdateTalkServiceID(ctx context.Context, talkID string, serviceID uint64) (err error)
}
