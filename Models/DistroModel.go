package Models

type Distro struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	ParentID int
}

func (b *Distro) TableName() string {
	return "distros"
}
