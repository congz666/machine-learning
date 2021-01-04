package model

type Info struct {
	ID           int `gorm:"primaryKey"`
	TID          string
	Name         string
	City         string
	Gender       string
	Height       float64
	C1           float64
	C2           float64
	C3           float64
	C4           float64
	C5           float64
	C6           float64
	C7           float64
	C8           float64
	C9           float64
	C10          float64
	Constitution string
}
