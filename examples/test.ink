text('Hello, World!

1234567890ABCDEFGHIJK
LMNOPQRSTUVWXYZabcdef
ghijklmnopqrstuvwxyz`
~!@#$%^&*()-_=+[{]}\\|
;:\'",<.>/?', 1, 1)
a := {
	(0): 5
	(1): 10
	(2): 15
}
log(a)
log(a.(0))
a.(0) := 0
log(a)
log(a.(0))
a.b := 1
log(a)
log(a.b)
a := seq(1, 11)
log(a)
