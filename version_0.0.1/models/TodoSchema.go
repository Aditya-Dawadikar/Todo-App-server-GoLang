package models

type ListItem struct {
	ItemName     string `json:"item_name"`
	ItemId       string `json:"item_id"`
	ItemTitle    string `json:"item_title"`
	ItemDesc     string `json:"item_desc"`
	ItemStatus   bool   `json:"item_status"`
	ItemPriority int64  `json:"item_priority"`
	ItemListId   string `json:"item_list_id"`
}

type TodoList struct {
	Owner   string     `json:"owner_id"`
	OwnerId string     `json:"owner_id"`
	Name    string     `json:"name"`
	Id      string     `json:"id"`
	List    []ListItem `json:"items"`
}
