# Shannon Cipher Engine
The Shanon Cipher Engine uses a one-time pad mechanism to encipher, not encrypt, credit cards. The library is focused on core functionality of grooming numbers, generating one time pads, padded pans, tokens and the core elements necessary to encipher and decipher credit card PANs.

Named in honor of Claude Sahnon who proved that this mechanism is a way to securely store information. To be clear, a padded pan can't be "cracked". There isn't an encryption algorithm. 

## Design
The library is designed using a number of first class functions along with pipelines and stages that compose them with channels in between the stages. For example, a pipeline might look like this:

**pan -> compact and strip spaces -> generate random pad -> create padded pan by XOR**

or

**pan ->compact and strip spaces -> parse BIN -> parse last 4 or last 2 -> generate token as BIN + random + last 4 -> verify not-Luhn valid token**

Because these pipelines are composable, one could decide to make the token Luhn valid. In that case, the token would be indistinguishable from a real credit card number (probably not desirable but still possible)

## Perfect Secrecy

"This is a very strong notion of security first developed during WWII by Claude Shannon and proved, mathematically, to be true for the one-time pad by Shannon about the same time. His result was published in the Bell Labs Technical Journal in 1949. Properly used, one-time pads are secure in this sense even against adversaries with infinite computational power.

Claude Shannon proved, using information theory, that the one-time pad has a property he termed perfect secrecy; that is, the ciphertext gives absolutely no additional information about the plaintext. This is because, given a truly random key that is used only once, a ciphertext can be translated into any plaintext of the same length, and all are equally likely. Thus, the a priori probability of a plaintext message M is the same as the a posteriori probability of a plaintext message M given the corresponding ciphertext"

https://en.wikipedia.org/wiki/One-time_pad

While the current implementation uses the crypto secure implementation provided by the Go language, it is important to note that there are a few other mechanisms that contribute to the overall entropy of the encipherment system.

1. Mulitple channels of different lengths create different sized integers in different locations from the same secure PRNG. While this is primarly in place to ensure high performance and the ability to handle peaky loads, it also means that the random number pumps are constructing random numbers of differing lengths and putting them in channels. That's all asynchronous and the sie of the various channels is configurable.
2. After tokens are generated, they are Luhn checked. If the token is found to be Luhn valid, indicating a valid credit card PAN, the token is thrown away and regenerated.This should occur about 1 in 10 attempts. Even if villain knew the security algorithm, the seed, and the expected randomm number stream, the incoming PANs, used to generate the tokens, act as a source of entropy in this case. And those are unknowable. 
 
## Fast!

Encryption is often problematic due to the computational load it puts on a system (in addition to being vulnerable to attack). In the case of encipherment via OTP,  the actual computation is very simple both to encipher and decipher. Basically it looks like this:

Pan (credit card number) ^ pad (random number) = padded pan

In plain English, do an exclusive OR of the credit card number with a random number and get an enciphered value that can only be retrieved by once again XOR'ing the pad with the padded pan.

## What's the catch?

One must use three different databases to store the values necessary to reconstruct the pan. This is also part of the security mechanics. Gain access to two of the three databases and it provides little value. One database might store the token and expiration date. The second database might store the padded pan keyed by token. The third would store the OTP keyed by the token. 

To use the values for a credit card transaction, the token, expiration date, pad and padded pan would be reassembled and the credit card number(pan) would be restored for use. 

## Why?

In addition to the obvious security of such a system, the credit card industry now requires that credit card numbers not be stored even if encrypted. At first glance that seems like an impossible task. The OTP encipherment solves the problem. 


