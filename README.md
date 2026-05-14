# VeilID: Zero-Knowledge Identity Provider for Discord

**VeilID** is a privacy-first identity bridge that verifies real-world humanity without compromising digital anonymity. By combining hardware-bound passkeys (**WebAuthn**) with mathematical proofs (**zk-SNARKs**), VeilID allows Discord communities to verify their members are unique, real humans while ensuring that neither the platform nor the server administrators ever see PII (Personally Identifiable Information).

---

## Core Philosophy: "Verification without Vulnerability"

Traditional Discord verification relies on "Honeypots"—databases of IDs or phone numbers that are prime targets for hackers. VeilID replaces trust in people with trust in mathematics.

* **Trustless Architecture:** We don't ask you to trust our server; we ask you to trust the laws of cryptography.
* **Hardware-Bound:** Authentication is anchored to the Secure Enclave of your physical device (TouchID, FaceID, Windows Hello).
* **Non-Linkable:** Using cryptographic **Nullifiers**, your identity on Server A cannot be linked to your identity on Server B.

---

## Key Features

### 1. Hybrid Trust Model (The "Anchor")

We support two tiers of human verification to solve the Oracle Problem:

* **Tier 1 (Standard):** PII-based verification (e.g., Stripe Identity/ID Scan) for high-convenience.
* **Tier 2 (Paranoid/Safe):** **zTLS (Zero-Knowledge TLS)**. Prove you own a bank account or university email via an encrypted web session without revealing your login credentials or the contents of the session.

### 2. Privacy-Preserving Proofs

* **Poseidon Hashing:** Uses ZK-friendly arithmetic hashing to keep computational costs low and verification speeds high.
* **Dynamic zk-SNARKs:** Generates a unique proof for every verification request. Even if our database is breached, hackers find only "Commitments"—mathematical signatures that are useless without the user's private hardware key.

### 3. Full Discord TOS Compliance

* **Official OAuth2 Flow:** No self-bots, no scraping.
* **External Verification:** All sensitive actions happen on `veilid.app`, keeping PII entirely off Discord's infrastructure.

---

## Tech Stack

* **Backend:** [Go (Golang)](https://go.dev/) (High-concurrency systems programming)
* **ZK Framework:** [gnark](https://gnark.consensys.io/) by Consensys (Native Go zk-SNARKs)
* **Auth Protocol:** [WebAuthn / FIDO2](https://webauthn.guide/) (Hardware-level biometrics)
* **Bot Library:** [DiscordGo](https://github.com/bwmarrin/discordgo)
* **Database:** [PostgreSQL](https://www.postgresql.org/) with `pgx/v5`
* **Hashing:** Poseidon (Arithmetization-oriented hash function)

---

## Project Structure

```text
.
├── cmd/veilid/             # Application entry point
├── internal/
│   ├── bot/                # Discord bot handlers and slash commands
│   ├── circuit/            # gnark ZK-SNARK circuits (The Math)
│   ├── db/                 # PostgreSQL schema and repository logic
│   ├── server/             # HTTP & OAuth2 server
│   └── webauthn/           # FIDO2 Passkey implementation
├── pkg/utils/              # Shared cryptographic helpers
├── ui/                     # Frontend verification interface
├── context.md              # AI Agent context for development
└── agent-guidelines.md     # Development standards and security rules
```
## Getting Started

### Prerequisites
- Go 1.23+

- Docker & Docker Compose (for PostgreSQL)

- A Discord Developer Application (Bot Token & Client Secret)

### Setup

Clone the Repo:

```Bash
git clone [https://github.com/PundarikakshNTripathi/VeilID.git](https://github.com/PundarikakshNTripathi/VeilID.git)
cd VeilID
```


2. **Environment Variables:**
```bash
cp .env.example .env
```
(Fill in your Discord and DB credentials)

3. **Spin up Infrastructure:**
```bash
    make up
```
Run the Application:

```bash
    make run
```

---

## Collaborative Workflow

This project follows an open-source standard for contribution:

1. **Branching:** All features go through `develop` before being merged into `main`. 
   - `feat/feature-name`
   - `fix/bug-name`
2. **Commits:** Use [Conventional Commits](https://www.conventionalcommits.org/).
   - `feat(circuit): add nullifier constraint`
   - `fix(bot): resolve state mismatch in oauth`
3. **PRs:** Every PR must include unit tests for new logic and pass the CI/CD pipeline.

---

## Roadmap
- [ ] Phase 1: Go/Discord/DB Boilerplate & OAuth2
- [ ] Phase 2: WebAuthn Registration/Login Ceremony
- [ ] Phase 3: Poseidon commitment storage
- [ ] Phase 4: `gnark` Circuit implementation & Role granting
- [ ] Phase 5: zTLS (TLSNotary) Alpha integration

---