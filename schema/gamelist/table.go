package game_list

// Table: GameList

type GameId = uint16

type BrandId = uint16

type GameName = string

type GameListTable struct {
	GameId    GameId     `gorm:"column:GameId"`
	BrandId    BrandId     `gorm:"column:BrandId"`
	GameName    GameName `gorm:"column:GameName"`
}
