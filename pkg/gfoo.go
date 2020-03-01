package gfoo

const (
	VersionMajor = 0
	VersionMinor = 9
)

func Init() {
	TAny.Init("Any")
	TBool.Init("Bool", &TAny)
	TFunction.Init("Function", &TAny)
	TId.Init("Id", &TAny)
	TNumber.Init("Number", &TAny)
	TInt.Init("Int", &TNumber)
	TLambda.Init("Lambda", &TAny)
	TMacro.Init("Macro", &TAny)
	TMeta.Init("Type", &TAny)
	TMethod.Init("Method", &TAny)
	TNil.Init("Nil")
	TPair.Init("Pair", &TAny)
	TRecord.Init("Record", &TAny)
	TScope.Init("Scope", &TAny)
	TScopeForm.Init("ScopeForm", &TAny)
	TSlice.Init("Slice", &TAny)
	TString.Init("String", &TAny)
	TTime.Init("Time", &TAny)

	Nil.dataType = &TNil
}

func New() *Scope {
	s := new(Scope).Init()
	s.InitAbc()
	s.AddVal("data", &TScope, s.Clone().InitData())	
	s.AddVal("time", &TScope, s.Clone().InitTime())	
	return s
}
