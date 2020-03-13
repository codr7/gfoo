use: data (Record record:)
use: time (days today)

type: Quantity Record

method: new-quantity (; Quantity) {
  record: (
    start     time.MIN
    end       time.MAX
    total     0
    available 0) as(Quantity)
}

type: Calendar Slice

method: new-calendar (; Calendar) {
  [new-quantity] as(Calendar)
}

type: Resource Record

method: new-resource (; Resource) {
  record: (calendar new-calendar) as(Resource)
}

type: Booking Record

method: new-booking (; Booking) {
  let: t today
  
  record: (
    resource NIL
    start    t
    end      (t +(1 days))
    quantity 1) as(Booking)
}

let: r new-resource
dump(r)

let: b new-booking
b set('resource r)
dump(b)