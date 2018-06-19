Microservice of symmetric encryption(Blowfish) on golang.
=========================

>
>
>

At first you should generate encryption key (if you don't have it)
```bash
pwgen 24 1
```

This microservice supports `Content-Type` :

- `application/json`
- `application/octet-stream` (msgpack)


API
=========================


>

<a name="encryption">[POST] /encryption</a>
--------

Encrypt your plaintext by key

Request:
- key:(string)                    # Encryption key (string with 24 length)
- text:(string)                   # Text to encode

```json
{"key": "phahf7woh8auvooJiebeong2", "text": "I want to encode this text"}
```

200:
```json
{
    "result": "000000000000000020cad804dd2b6c272498052709fd58f669a670ddd0901acbe726224ccb2ed9d8"
}
```

>

<a name="decryption">[POST] /decryption</a>
--------

Decryption your ciphertext by key

Request:
- key:(string)                     # Encryption key (string with 24 length)
- text:(string)                    # Text to decode

```json
{"key": "phahf7woh8auvooJiebeong2", "text": "000000000000000020cad804dd2b6c272498052709fd58f669a670ddd0901acbe726224ccb2ed9d8"}
```

200:
```json
{
    "result": "I want to encode this text"
}
```

>

<a name="multi_encryption">[POST] /multi_encryption</a>
--------

Multi Encryption of your plaintexts by key

Request:
- key:(string)                   # Encryption key (string with 24 length)
- texts:[]string                 # Texts to encode

```json
{"key": "phahf7woh8auvooJiebeong2", "texts": ["text 1", "text 2", "text 3"]}
```

200:
```json
{
    "result": [
        "0000000000000000badb5ed98e106434",
        "00000000000000004b475e7f3ca46ec5",
        "00000000000000006bb2993acc134f0f"
    ]
}
```

>

<a name="liveness">[GET] /_liveness</a>
--------
liveness Probe

>

<a name="readiness">[GET] /_readiness</a>
--------
readiness Probe