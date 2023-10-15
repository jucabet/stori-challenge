package ports

type ITopicRepository interface {
	SendMessageToReport(fileChargeID string) error
}
