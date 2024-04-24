package structs

type ItemCondition struct {
	Id              int64  `json:"id"`
	ItemId          int64  `json:"item_id"`
	ImageUrl        string `json:"image_url"`
	OverallConditon string `json:"overall_conditon"`
	Status          string `json:"status"`
	Note            string `json:"note"`
	CreatedAt       string
	UpdatedAt       string
}

type ItemList struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	BoughtDate   string `json:"bought_date"`
	ImageUrl     string `json:"image_url"`
	DepartmentId int64  `json:"department_id"`
	TypeId       int64  `json:"type_id"`
	CreatedAt    string
	UpdatedAt    string
}
