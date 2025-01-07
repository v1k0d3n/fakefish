# Fakefish (Golang)

Why recreate this project? The purpose is really a redraft of the already existing (and kind of amazing) [Metal3/Fakefish](https://github.com/openshift-metal3/fakefish) project, which was originally written/proposed by [Mario Vázquez](https://github.com/mvazquezc). But the original is limited for embedded systems, and I need a binary version of Fakefish for things like the new [JetKVM](https://jetkvm.com/). 

For the JetKVM, once you go into "Developer Mode", you can upload the binary and a collection of scripts that can be used to load/unload ISOs and control the host via Redfish. That's my goal anyway, and I've already rreached out to some folks on the team to get their thoughts (over at JetKVM).

Once this project is a bit more ironed out, I will permanently migrate it to GitHub and propose it to the Metal3 team for further maintenance.

## Build

Until I create a Makefile (my next objective), you can create a binary by running the following command:

```bash
go build -o fakefish ./cmd/fakefish
```