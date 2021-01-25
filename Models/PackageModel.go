package Models

type Package struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	DistroID int `gorm:"not null"`
	Distro   Distro
}

func (b *Package) TableName() string {
	return "packages"
}
