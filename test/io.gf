use: io (new-buffer to-byte to-int write)

42 to-byte to-int check: =(42)

[
  new-buffer
  .. write(to-byte(1))
  .. write(to-byte(2))
  .. write(to-byte(3))
  ...
] check: =([to-byte(1) to-byte(2) to-byte(3)])