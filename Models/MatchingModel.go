package Models

type Matching struct {
	ID         int `gorm:"primaryKey"`
	PackageID  int `gorm:"not null"`
	Package    Package
	DistroID   int `gorm:"not null"`
	Distro     Distro
	MatchingID int     `gorm:"not null"`
	Matching   Package `gorm:"foreignKey:MatchingID"`
}

func (b *Matching) TableName() string {
	return "matchings"
}
