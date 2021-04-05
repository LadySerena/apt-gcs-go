# apt-gcs-go

This used to be rust based but this repo was renamed to `apt-gcs-go` because rust isn't supported in gcp apis which is a
bit problematic. A Go based apt transport to support google cloud storage. This is a side project to learn about the
internals of debian and how packaging works. This is also part of Serena's over engineered minecraft platform.

Code adapted from the following repos

- [apt-gcs](https://github.com/dhaivat/apt-gcs)
- [apt-transport-cloudflared](https://github.com/cloudflare/apt-transport-cloudflared)

References

- [apt transport spec](http://www.fifi.org/doc/libapt-pkg-doc/method.html/ch2.html)
- [cloudflare blog post](https://blog.cloudflare.com/apt-transports/)
- [rfc 822](https://tools.ietf.org/html/rfc822)