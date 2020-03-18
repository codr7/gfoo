package gfoo

type Core struct {
	Scope
	Data, String, Time, Zip Scope
	Image ImageModule
	Io IoModule
	Math MathModule
	Png PngModule
}

func New() *Core {
	c := new(Core)
	c.Init()
	c.InitAbcModule()
	c.AddVal("data", &TScope, c.Data.Init().InitDataModule())
	c.AddVal("image", &TScope, c.Image.Init())
	c.AddVal("io", &TScope, c.Io.Init())
	c.AddVal("math", &TScope, c.Math.Init())
	c.AddVal("string", &TScope, c.String.Init().InitStringModule())	
	c.AddVal("png", &TScope, c.Png.Init())
	c.AddVal("time", &TScope, c.Time.Init().InitTimeModule())	
	c.AddVal("zip", &TScope, c.Zip.Init().InitZipModule())
	return c
}

