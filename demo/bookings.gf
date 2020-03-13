type: Quantity data.Record

method: new-quantity (; Quantity) {
  data.record: (
    start     time.MIN
    end       time.MAX
    total     0
    available 0) as(Quantity)
}

type: Calendar Slice

method: new-calendar (; Calendar) {
  [new-quantity] as(Calendar)
}

type: Resource data.Record

method: new-resource (; Resource) {
  data.record: (calendar new-calendar) as(Resource)
}

type: Booking data.Record

method: new-booking (; Booking) {
  let: t time.today
  
  data.record: (
    resource NIL
    start    t
    end      (t +(1 time.days))
    quantity 1) as(Booking)
}

let: r new-resource
dump(r)

let: b new-booking
b set('resource r)
dump(b)