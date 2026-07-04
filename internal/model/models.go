package model

import "time"

// DictData 对应 sys_dict_data 表
type DictData struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	TypeCode  string    `gorm:"column:type_code;not null"`
	DictLabel string    `gorm:"column:dict_label;not null"`
	DictValue string    `gorm:"column:dict_value;not null"`
	SortOrder int32     `gorm:"column:sort_order;default:0"`
	IsEnabled bool      `gorm:"column:is_enabled;default:true"`
	CssClass  string    `gorm:"column:css_class"`
	Remark    string    `gorm:"column:remark"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}

// TableName 指定表名和 schema
func (DictData) TableName() string {
	return "dict_area.sys_dict_data"
}

// Area 对应 sys_area 表
type Area struct {
	AreaCode   string    `gorm:"column:area_code;primaryKey"`
	ParentCode string    `gorm:"column:parent_code;not null"`
	AreaName   string    `gorm:"column:area_name;not null"`
	ShortName  string    `gorm:"column:short_name"`
	AreaLevel  int32     `gorm:"column:area_level;not null"`
	Pinyin     string    `gorm:"column:pinyin"`
	MergerName string    `gorm:"column:merger_name"`
	IsEnabled  bool      `gorm:"column:is_enabled;default:true"`
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}

func (Area) TableName() string {
	return "dict_area.sys_area"
}
