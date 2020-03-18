package gfoo

const (
	VersionMajor = 0
	VersionMinor = 17
)

func Init() {
	TOption.Init("Any?")
	TAny.Init("Any", &TOption)
	TNil.Init("Nil", &TOption)

	TSequence.Init("Sequence", &TAny)
	TBool.Init("Bool", &TAny)
	TWriter.Init("Writer", &TAny)
	TBuffer.Init("Buffer", &TWriter)
	TChar.Init("Char", &TAny)
	TFunction.Init("Function", &TAny)
	TId.Init("Id", &TAny)
	TNumber.Init("Number", &TAny)
	TInt.Init("Int", &TNumber, &TSequence)
	TIter.Init("Iter", &TSequence)
	TLambda.Init("Lambda", &TAny)
	TMacro.Init("Macro", &TAny)
	TMeta.Init("Type", &TAny)
	TMethod.Init("Method", &TAny)
	TPair.Init("Pair", &TSequence)
	TRecord.Init("Record", &TAny, &TSequence)
	TRgba.Init("Rgba", &TRgba)
	TScope.Init("Scope", &TAny)
	TScopeForm.Init("ScopeForm", &TAny)
	TSlice.Init("Slice", &TAny, &TSequence)
	TString.Init("String", &TAny, &TSequence)
	TTime.Init("Time", &TAny)
	TTimeDelta.Init("TimeDelta", &TAny)
	TZipWriter.Init("Writer", &TAny)

	Nil.dataType = &TNil
}
