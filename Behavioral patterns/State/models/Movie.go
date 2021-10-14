package models

type Movie struct {
	state State
}

func (m *Movie) SetState(state State){
	m.state=state
}

func (m *Movie) GetState() string{
	return m.state.GetName()
}
