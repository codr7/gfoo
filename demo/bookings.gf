type: Quantity data.Record

method: new-quantity (; Quantity) {
  data.record: ('start time.MIN 'end time.MAX 'total 0 'available 0) as(Quantity)
}

type: Calendar Slice

method: new-calendar (; Calendar) {
  [new-quantity] as(Calendar)
}

type: Resource data.Record

method: new-resource (; Resource) {
  data.record: ('calendar new-calendar) as(Resource)
}

say(new-resource)