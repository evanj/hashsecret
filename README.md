# Reversing Secret's phone numbers

**Question**: How could you reverse engineer [Secret](https://www.secret.ly/)'s phone number database, if you had a copy?  
**Answer**: You can try every possible phone number in about 2.5 hours.

(Thanks to [this Twitter thread](https://twitter.com/coda/status/436267472639897600) for inspiring me to actually write this)

## Details

Here is [what is known about how Secret stores phone numbers](https://medium.com/secret-den/12ab82fda29f):

> we locally hash the contact details first (with shared salt) [...] Therefore, [+15552786005] becomes [a22d75c92a630725f4] and the original number never leaves your phone.

I wrote a quick-and-dirty Go program to try and answer this question. I'm assuming Secret uses a standard cryptographic hash. I'll give them the benefit of the doubt and say SHA256, since its one of the slower standard choices. Since the shared salt must be in the client somewhere, I've added an 8 character constant string to the hash to simulate it.

**Time to hash 8 billion numbers**: 2:34:15 (154 minutes, or 9255 seconds)  
**CPU**: Intel(R) Core(TM) i5-2500K CPU @ 3.30GHz


## Conclusion

Your phone number isn't secret to the Secret developers, or anyone that hacks them.


## Make it faster

Ordered approximately from least to most effort

* Use a faster computer.
* Generate numbers with a ["stateful iterator"](http://ewencp.org/blog/golang-iterators/) instead of a Go routine.
* Increment phone number bytes in place instead of creating a string for each number.
* Divide the number space across multiple cores.
* Only generate valid phone numbers (see the [North American Dialing plan](http://www.nanpa.com/)). I generate all numbers from 200-000-0000 through 999-999-9999, many of which are not actually numbers.
* Use optimized C/C++, particularly for SHA256.
* Use GPUs.

I used a Go routine and a channel because its a natural way to write a generator in Go. This program runs slightly *slower* if you run it with GOMAXPROCS=2 (10782 seconds with 2 procs, 9255 with 1). This is a good example of how *not* to parallelize this problem.


## Run it yourself

* `go build hashsecret.go`
* `time ./hashsecret`
