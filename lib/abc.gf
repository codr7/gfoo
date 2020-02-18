macro: if: (body) {
  '(branch: @body ())
}

macro: else: (body) {
  '(branch: () @body)
}