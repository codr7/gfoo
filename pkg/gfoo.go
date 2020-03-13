package gfoo

const (
	VersionMajor = 0
	VersionMinor = 13
)

func Init() {
	TOptional.Init("Any?")
	TAny.Init("Any", &TOptional)
	TNil.Init("Nil", &TOptional)

	TBool.Init("Bool", &TAny)
	TChar.Init("Char", &TAny)
	TFunction.Init("Function", &TAny)
	TId.Init("Id", &TAny)
	TNumber.Init("Number", &TAny)
	TInt.Init("Int", &TNumber)
	TLambda.Init("Lambda", &TAny)
	TMacro.Init("Macro", &TAny)
	TMeta.Init("Type", &TAny)
	TMethod.Init("Method", &TAny)
	TPair.Init("Pair", &TAny)
	TPairForm.Init("PairForm", &TAny)
	TRecord.Init("Record", &TAny)
	TScope.Init("Scope", &TAny)
	TScopeForm.Init("ScopeForm", &TAny)
	TSlice.Init("Slice", &TAny)
	TString.Init("String", &TAny)
	TTime.Init("Time", &TAny)
	TTimeDelta.Init("TimeDelta", &TAny)

	Nil.dataType = &TNil
}

func New() *Scope {
	s := new(Scope).Init()
	s.InitAbc()
	s.AddVal("time", &TScope, s.Clone().InitTime())	
	s.AddVal("data", &TScope, s.Clone().InitData())	
	return s
}
