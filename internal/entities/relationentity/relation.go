package relationentity

type Relation struct {
	ID     uint64
	UserID uint64
	TeamID uint64
	RuneID uint64
}

type NamedRelation struct {
	ID       uint64
	UserID   uint64
	TeamID   uint64
	RuneID   uint64
	TeamName string
	RuneName string
}
