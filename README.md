This is an abstraction layer around the excellent
cross platform UI toolkit from andlabs/ui

This abstraction layer makes it easier to write
simple interfaces for like our cloud control panel

The cross platform UI has 'quirks' due to it being
cross platform. Some of the abstraction layer here
attemps to obfuscate the ui objects so that it is
more difficult to trigger inconsistancies.

In this regard, this is an attempt to restrict
all andlabs/ui (and andlabs/libui) interaction to
the calls within this library.
