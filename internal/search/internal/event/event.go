package event

const (
	SyncTopic = "sync_data_to_search"
)

type SyncEvent struct {
	Biz   string `json:"biz"`
	BizId int    `json:"bizId"`
	// 具体内容
	Data string `json:"data"`
}
