package gfoo

type Core struct {
	Scope
	Abc AbcModule
	Data DataModule
	String StringModule
	Time TimeModule
	Zip ZipModule
	Image ImageModule
	Io IoModule
	Iter IterModule
	Math MathModule
	Png PngModule
}

func New() *Core {
	c := new(Core)
	c.Init()
	
	c.AddModule("abc", c.Abc.Init())
	c.AddModule("data", c.Data.Init())
	c.AddModule("image", c.Image.Init())
	c.AddModule("io", c.Io.Init())
	c.AddModule("iter", c.Iter.Init())
	c.AddModule("math", c.Math.Init())
	c.AddModule("string", c.String.Init())	
	c.AddModule("png", c.Png.Init())
	c.AddModule("time", c.Time.Init())	
	c.AddModule("zip", c.Zip.Init())

	c.Use(NewVal(&TModule, &c.Abc.Module), []string{"use:"}, NilPos)
	return c
}

