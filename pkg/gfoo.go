package gfoo

const (
	VersionMajor = 0
	VersionMinor = 16
)

func Init() {
	TOptional.Init("Any?")
	TAny.Init("Any", &TOptional)
	TNil.Init("Nil", &TOptional)

	TSequence.Init("Sequence", &TAny)

	TBool.Init("Bool", &TAny)
	TChar.Init("Char", &TAny)
	TFunction.Init("Function", &TAny)
	TId.Init("Id", &TAny)
	TNumber.Init("Number", &TAny)
	TInt.Init("Int", &TNumber, &TSequence)
	TIterator.Init("Iterator", &TSequence)
	TLambda.Init("Lambda", &TAny)
	TMacro.Init("Macro", &TAny)
	TMeta.Init("Type", &TAny)
	TMethod.Init("Method", &TAny)
	TPair.Init("Pair", &TSequence)
	TRecord.Init("Record", &TAny, &TSequence)
	TScope.Init("Scope", &TAny)
	TScopeForm.Init("ScopeForm", &TAny)
	TSlice.Init("Slice", &TAny, &TSequence)
	TString.Init("String", &TAny, &TSequence)
	TTime.Init("Time", &TAny)
	TTimeDelta.Init("TimeDelta", &TAny)

	Nil.dataType = &TNil
}

func New() *Scope {
	s := new(Scope).Init()
	s.InitAbcModule()
	s.AddVal("data", &TScope, NewScope().InitDataModule())	
	s.AddVal("string", &TScope, NewScope().InitStringModule())	
	s.AddVal("time", &TScope, NewScope().InitTimeModule())	
	return s
}
