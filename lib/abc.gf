macro: if: (body) {
  '(?: @body ())
}

macro: else: (body) {
  '(?: () @body)
}