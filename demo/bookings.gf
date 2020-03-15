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

method: update (in Quantity (start end) Time (total available) Int; Quantity) {
  in.end <=(start) or: in.start >=(end) ?: in {
    say(["match: " in.start in.end])
  }
}

type: Calendar Slice

method: new-calendar (; Calendar) {
  [new-quantity] as(Calendar)
}

method: update-quantity (in Calendar (start end) Time (total available) Int; Calendar) {
  in map(/: (in) {in update(start end total available)}) as(Calendar)
}

type: Resource Record

method: new-resource (; Resource) {
  record: (calendar new-calendar) as(Resource)
}

method: update-calendar (in Resource (start end) Time (total available) Int;) {
  in set('calendar in.calendar update-quantity(start end total available))
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

method: store (in Booking) {
  in.resource update-calendar(in.start in.end 0 !in.quantity)
}

let: r new-resource
dump(r)

r update-calendar(time.MIN time.MAX 10 10)

let: b new-booking
b set('resource r)
dump(b)

b store