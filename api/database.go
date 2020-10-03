package api

type Deposit struct {
	ID      uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Amount  int64  `gorm:"size:12"`
	From    string `gorm:"type:varchar(50);unique_index"`
	Approve int8   `gorm:"size:1"`
}
