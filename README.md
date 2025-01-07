# Fakefish (Golang)

The purpose of this repository is propose a redraft of an already existing (and kind of amazing) [Metal3/Fakefish](https://github.com/openshift-metal3/fakefish) project. I believe that `fakefish` was originally written/proposed by [Mario Vázquez](https://github.com/mvazquezc). So why suggest a conversion to Golang? The original implementation using python is great, but it could be limited for embedded systems that don't include python. This compiled version of `fakefish` allows for easy deployment on embedded devices such as the new [JetKVM](https://jetkvm.com/), for example.

If you enable "Developer Mode" on the JetKVM, you could potentially upload this precompiled binary along with any nessesary scripts that can be used to load/unload ISOs and control the host via Redfish. This is my initial goal/proposal anyway. I've already created a [proposal](https://github.com/jetkvm/kvm/issues/36) with the JetKVM team, and have reached out to some folks to get their thoughts as well (through Discord and Email).

Once this project is a bit more ironed out, my intent would be to permanently migrate it to the [Metal3 organization](https://github.com/openshift-metal3/) for future maintenance.

## Build

Until I create a Makefile (my next objective) and start implementing gates/CI for the project, you can create a binary by running the following command:

```bash
go build -o fakefish ./cmd/fakefish
```

## Contributing

As the project matures, I will add more information on how to contribute and include a Code of Conduct. For now, feel free to open an issue or PR if you have any suggestions or feedback.