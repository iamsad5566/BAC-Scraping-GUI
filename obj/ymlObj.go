package obj

type EnVariable struct {
	BossID  string   `yaml:"bossid"`
	AssID   string   `yaml:"assid"`
	AssPWD  string   `yaml:"asspwd"`
	OS      string   `yaml:"OS"`
	Members []string `yaml:"member"`
}

func (e *EnVariable) GetBossID() string {
	return e.BossID
}

func (e *EnVariable) GetAssID() string {
	return e.AssID
}

func (e *EnVariable) GetAssPWD() string {
	return e.AssPWD
}

func (e *EnVariable) GetOS() string {
	return e.OS
}

func (e *EnVariable) GetMembers() []string {
	return e.Members
}

func (e *EnVariable) SetBossID(newBossID string) {
	e.BossID = newBossID
}

func (e *EnVariable) SetAssID(newAssID string) {
	e.AssID = newAssID
}

func (e *EnVariable) SetAssPWD(newAssPWD string) {
	e.AssPWD = newAssPWD
}

func (e *EnVariable) SetOS(newOS string) {
	e.OS = newOS
}

func (e *EnVariable) SetMembers(newMembers []string) {
	e.Members = newMembers
}
