package skilltrees

// Tehnologies types
const (
	FRONTEND = "frontend"
	BACKEND = "backend"
)

// Skill ...
type Skill struct {
	Name string
	Active bool
	Buyed bool
	Locked bool
	Price int
	Img string
	CodeIncome float32
	ReliabilityIncome float32
	Type string
	MaxCPS int
	DecreaseCode int
	BaseSkill bool
	NextSkills []Skill
	AdditionalSkill bool
}
