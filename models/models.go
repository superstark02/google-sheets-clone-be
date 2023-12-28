package models

type Data struct {
	ID    int64  `field:"id"`
	Data  string `field:"data"`
	Bold  bool   `field:"bold"`
	Color string `field:"color"`
}
