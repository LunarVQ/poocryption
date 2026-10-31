# poocryption

Public-key encryption where the ciphertext is 💩 and the keys are emoji.
Standard library only.

```
keygen ─► (pub, priv) ─emojiset─► 🦊🍕🚀…  /  🐙🎲🌵…
S ─RSA-OAEP(pub)─► wrapped      S ─HKDF─► K ─AES-256-GCM─► ct      header ‖ ct ─poop─► 💩💩💩
```

The asymmetric core is a cycle: `x → x^e mod N`. Anyone can step forward;
the private key holder knows the cycle's period (`λ(N)`, from the factors of
`N`) and jumps the rest of the way around in one exponentiation (`d = e⁻¹ mod λ`).

## Layout

| Path | Role |
| --- | --- |
| `poocryption.go` | Public API: `GenerateKey`, `Encrypt(pt, pub)`, `Decrypt(poop, priv)` |
| `internal/trapdoor/trapdoor.go` | Keypair, `Step` / `Walk` / `Shortcut` / `Unlock` (the cycle), `Wrap` / `Unwrap` (OAEP) |
| `internal/trapdoor/encoding.go` | Key serialisation (private record is just P, Q, E) |
| `internal/emojiset/emojiset.go` | bytes ⇄ emoji alphabet, used for keys only |
| `internal/kdf/kdf.go` | HKDF: session secret → payload key |
| `internal/aead/aead.go` | AES-256-GCM seal/open |
| `internal/format/header.go` | Container header (key ID, wrapped key, nonce), authenticated as AAD |
| `internal/poop/poop.go` | bytes ⇄ 💩 via Unicode variation selectors, used for ciphertext only |
| `cmd/poocrypt/main.go` | CLI: `keygen`, `enc`, `dec`, `walk` |

## Dependencies

None outside the Go standard library. Needs Go 1.24+ for `crypto/hkdf`.

## Status

Skeleton. Types, signatures and the pipeline are in place; bodies are `TODO`.
