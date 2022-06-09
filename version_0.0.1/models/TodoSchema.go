package models

type ListItem struct {
	ItemName     string `json:"item_name"`
	ItemId       string `json:"item_id"`
	ItemStatus   bool   `json:"item_status"`
	ItemPriority int64  `json:"item_priority"`
	ItemListId   string `json:"item_list_id"`
}

type TodoList struct {
	OwnerId string     `json:"owner_id"`
	Name    string     `json:"name"`
	Id      string     `json:"id"`
	List    []ListItem `json:"items"`
	Desc    string     `json:"description"`
}
