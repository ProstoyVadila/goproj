package models

type Survey struct {
	Skip        []string
	PackageName string
	Author      string
	Description string
	Prefix      string
	InitGit     bool
	InitVSCode  bool
}

func (s *Survey) ToSetup() *Setup {
	isSetInitGit, isSetInitVSCode, generateInFolder := true, true, true
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
		generateInFolder,
	)
}

func (s *Survey) ToGlobalConfig() *GlobalConfig {
	return NewGlobalConfig(
		s.Author,
		s.Prefix,
		s.Skip,
		s.InitGit,
		s.InitVSCode,
	)
}
