package models

type Game struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	CountRows  int    `json:"count_rows"`
	CountCols  int    `json:"count_cols"`
	CountMines int    `json:"count_mines"`
	State      string `json:"state"`
}

type Distribution struct {
	GameID    int64  `json:"game_id"`
	RowNumber int    `json:"row_number"`
	ColNumber int    `json:"col_number"`
	Value     string `json:"value"`
	State     string `json:"state"`
}

type User struct {
	ID       int64  `json:"id"`
	NickName string `json:"nick_name"`
}

type Accounts struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}
