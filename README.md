# VeilID 

## 🛡️Privacy-Preserving Sybil Resistance via Zero-Knowledge Proofs
VeilID is a cryptographic identity bridge that allows communities to verify human uniqueness without ever exposing user data to the platform. It utilizes device-bound hardware passkeys **(WebAuthn/FIDO2)** and dynamically varying **zk-SNARKs** to prove identity.

## Cryptographic Architecture (Trustless System)
- **No Passwords**: Users authenticate via their device's secure hardware enclave (TouchID/Windows Hello). The private key never leaves the device.
- **Zero-Knowledge Circuit:** Implemented in Go using Consensys gnark. The circuit proves $Nullifier = PoseidonHash(User\_Secret + Server\_Nonce)$.
- **Anti-Tracking:** Because the Server Nonce changes per community, the output Proof Hash is entirely different every time. Server admins cannot collude to track users across platforms, yet mathematical certainty of uniqueness is maintained.

## System Flow
Discord Bot assigns a unique URL -> User authenticates via WebAuthn -> Go backend generates zk-Proof -> Proof is mathematically verified -> Bot assigns "Verified Human" role. No PII is ever transmitted to Discord.
