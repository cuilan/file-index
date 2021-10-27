package entity

func (Folder) TableName() string {
	return "t_folder"
}

type Folder struct {
	Id         int64  `gorm:"column:id" json:"id"`
	ParentId   int64  `gorm:"column:parent_id" json:"parent_id"`
	Name       string `gorm:"column:name" json:"name"`
	Path       string `gorm:"column:path" json:"path"`
	Type       string `gorm:"column:type" json:"type"`
	Removed    bool   `gorm:"column:removed" json:"removed"`
	Hideen     bool   `gorm:"column:hideen" json:"hideen"`
	Desc       string `gorm:"column:desc" json:"desc"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
}

func (File) TableName() string {
	return "t_file"
}

type File struct {
	Id         int64  `gorm:"column:id" json:"id"`
	FolderId   int64  `gorm:"column:folder_id" json:"folder_id"`
	Name       string `gorm:"column:name" json:"name"`
	Path       string `gorm:"column:path" json:"path"`
	Size       int64  `gorm:"column:size" json:"size"`
	SizeK      int64  `gorm:"column:size_k" json:"size_k"`
	SizeM      int    `gorm:"column:size_m" json:"size_m"`
	SizeG      int    `gorm:"column:size_g" json:"size_g"`
	Suffix     int64  `gorm:"column:suffix" json:"suffix"`
	Status     int    `gorm:"column:status" json:"status"`
	Removed    bool   `gorm:"column:removed" json:"removed"`
	Hideen     bool   `gorm:"column:hideen" json:"hideen"`
	FileTime   int64  `gorm:"column:file_time" json:"file_time"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
}

func (MinioConfig) TableName() string {
	return "t_minio_config"
}

type MinioConfig struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	Value      string `gorm:"column:value" json:"value"`
	Desc       string `gorm:"column:desc" json:"desc"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
}

func (Statistics) TableName() string {
	return "t_statistics"
}

type Statistics struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Type       string `gorm:"column:type" json:"type"`
	Count      int    `gorm:"column:count" json:"count"`
	Desc       string `gorm:"column:desc" json:"desc"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
}

func (Tag) TableName() string {
	return "t_tag"
}

type Tag struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	Desc       string `gorm:"column:desc" json:"desc"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
}

func (TagFile) TableName() string {
	return "t_tag_file"
}

type TagFile struct {
	Id         int64 `gorm:"column:id" json:"id"`
	TagId      int64 `gorm:"column:tag_id" json:"tag_id"`
	FileId     int64 `gorm:"column:file_id" json:"file_id"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
}

func (TagFolder) TableName() string {
	return "t_tag_folder"
}

type TagFolder struct {
	Id         int64 `gorm:"column:id" json:"id"`
	TagId      int64 `gorm:"column:tag_id" json:"tag_id"`
	FolderId   int64 `gorm:"column:folder_id" json:"folder_id"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
}
