package repository

type Games struct {
	ID         int64  `gorm:"primarykey;autoIncrement:true;not null"`
	UserID     int64  `gorm:"default:null"`
	CountRows  int    `gorm:"default:null"`
	CountCols  int    `gorm:"default:null"`
	CountMines int    `gorm:"default:null"`
	State      string `gorm:"default:null"`
}

type Users struct {
	ID        uint   `gorm:"primarykey;autoIncrement:true;not null"`
	NickName  string `gorm:"not null"`
	AccountID int    `gorm:"not null"`
}

type Distributions struct {
	GameID    int64  `gorm:"primaryKey;autoIncrement:false"`
	RowNumber int    `gorm:"primaryKey;autoIncrement:false"`
	ColNumber int    `gorm:"primaryKey;autoIncrement:false"`
	Value     string `gorm:"default:null"`
	State     string `gorm:"default:null"`
}

type Accounts struct {
	ID    uint   `gorm:"primarykey;autoIncrement:true;not null"`
	Email string `gorm:"not null"`
}
