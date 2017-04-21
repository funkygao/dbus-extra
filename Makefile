present:
	@present&
	open http://127.0.0.1:3999/dbus.slide

doc:
	godoc golang.org/x/tools/present | less
