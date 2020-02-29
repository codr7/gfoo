type: Quantity Record

method: new-quantity(; Quantity) {
  record: ('start time.MIN 'end time.MAX 'total 0 'available 0) as(Quantity)
}

type: Calendar Slice

method: new-calendar(; Calendar) {
  [new-quantity] as(Calendar)
}

type: Resource Record

method: new-resource (; Resource) {
  record: ('calendar new-calendar) as(Resource)
}