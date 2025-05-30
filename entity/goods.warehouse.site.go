package entity

type SiteWarehouses struct {
	SiteId             int    `json:"siteId"`   // 站点 Id
	SiteName           string `json:"siteName"` // 站点名称
	ValidWarehouseList []struct {
		WarehouseId      string `json:"warehouseId"`
		WarehouseName    string `json:"warehouseName"`
		WarehouseDisable bool   `json:"warehouseDisable"`
	} `json:"validWarehouseList"` // 有效仓库列表
}
