package gfoo

type Core struct {
	Scope
	Data, String, Time, Zip Scope
	Io IoModule
}

func New() *Core {
	c := new(Core)
	c.Init()
	c.InitAbcModule()
	c.AddVal("data", &TScope, c.Data.Init().InitDataModule())
	c.AddVal("io", &TScope, c.Io.Init())
	c.AddVal("string", &TScope, c.String.Init().InitStringModule())	
	c.AddVal("time", &TScope, c.Time.Init().InitTimeModule())	
	c.AddVal("zip", &TScope, c.Zip.Init().InitZipModule())
	return c
}

