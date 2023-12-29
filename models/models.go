package models

type Data struct {
	Id    int64  `field:"id"`
	Data  string `field:"data"`
	Bold  bool   `field:"bold"`
	Color string `field:"color"`
	X     int64  `field:"x"`
	Y     int64  `field:"y"`
}
