package api

type Deposit struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Amount      int64  `gorm:"size:12"`
	From        string `gorm:"type:varchar(50);unique_index"`
	Approve     int8   `gorm:"size:1"`
	AggregateID string `gorm:"type:text;unique_index"`
}
type EventStore struct {
	ID            string `gorm:"primary_key"`
	EventType     string `gorm:"type:text"`
	AggregateID   string `gorm:"type:text"`
	AggregateType string `gorm:"type:text"`
	EventData     string `gorm:"type:text"`
	Channel       string `gorm:"type:text"`
}
