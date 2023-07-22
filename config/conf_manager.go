package config

type Manager struct {
	ConfDatabase ConfDatabase
}

var ConfManager Manager

func (m *Manager) Init(confPath string) error {
	if err := m.InitConf(confPath); err != nil {
		return err
	}
	return nil
}

func (m *Manager) InitConf(confPath string) error {
	if err := m.ConfDatabase.Init(confPath + "/database.yaml"); err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetPgSqlConf() *ConfPgSQL {
	return m.ConfDatabase.ConfPgSQL
}
