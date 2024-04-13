package models

type Survey struct {
	Skip        []string
	PackageName string
	Author      string
	Description string
	InitGit     bool
	InitVSCode  bool
}

func (s *Survey) ToSetup() *Setup {
	isSetInitGit, isSetInitVSCode := true, true
	return NewSetup(
		s.PackageName,
		s.Author,
		s.Description,
		s.Skip,
		s.InitGit,
		s.InitVSCode,
		isSetInitGit,
		isSetInitVSCode,
		FromSurvey,
	)
}

func (s *Survey) ToGlobalConfig() *GlobalConfig {
	return NewGlobalConfig(
		s.Author,
		s.Skip,
		s.InitGit,
		s.InitVSCode,
	)
}
