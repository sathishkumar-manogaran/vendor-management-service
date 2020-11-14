package modal

type Country struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(500)"`
}

type Service struct {
	Id      int      `orm:"auto"`
	Name    string   `orm:"size(500)"`
	Country *Country `orm:"rel(fk)"`
}

// Model Struct
type Vendor struct {
	Id      int      `orm:"auto"`
	Name    string   `orm:"size(500)"`
	Service *Service `orm:"rel(fk)"`
}
