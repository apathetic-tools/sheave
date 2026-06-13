<!-- DECISIONS.md -->
# DECISIONS.md

A record of major design and implementation choices in **sheave** — what was considered, what was chosen, and why.

Each decision:

- Is **atomic** — focused on one clear choice.  
- Is **rationale-driven** — the “why” matters more than the “what.”  
- Should be written as if explaining it to your future self — concise, readable, and honest.  
- Includes **Context**, **Options Considered**, **Decision**, and **Consequences**.  

For formatting guidelines, see the [DECISIONS.md Style Guide](./DECISIONS_STYLE_GUIDE.md).

---

## 🐹 Switch from Python to Go
<a id="dec13"></a>*DEC 13 — 2026-06-13*

### Context

Sheave was originally written in Python (see [DEC 03](#dec03)). However, as a CLI tool that needs to be distributed to end users (who may be working in Rust, Node, or other ecosystems), Python's distribution model became a significant liability. Asking users to install Python, manage virtual environments, and deal with `pip` just to run a CLI tool introduces too much friction. Python compilation tools (like PyInstaller) often yield bloated binaries with slow startup times.

### Options Considered

| Language | Pros | Cons |
|-----------|------|------|
| **Python (Status Quo)** | ✅ Fast development<br>✅ Familiarity | ❌ Terrible distribution story<br>❌ Slow CLI startup if bundled |
| **Rust** | ✅ Single binary<br>✅ Fast execution | ❌ Slower compilation<br>❌ Steeper learning curve |
| **Go** | ✅ Statically linked single binary<br>✅ Instant startup time<br>✅ Trivial cross-compilation<br>✅ Excellent CLI ecosystem (Cobra, Charm) | ⚠️ Requires rewriting existing Python logic |

### Decision

Migrate the entire codebase from **Python to Go**.

Go compiles down to a single, dependency-free binary that can be trivially distributed across all major operating systems. This provides the ultimate "download and run" experience for end users, which is the gold standard for modern CLI tools. The rewrite is justified by the massive improvement in the distribution and installation story.

<br/><br/>

---
---

<br/><br/>

## 🤝 Adopt `Contributor Covenant 3.0` as Code of Conduct  
<a id="dec05"></a>*DEC 05 — 2025-10-10*  

### Context

The project needed a **clear, inclusive standard of behavior** for contributors and maintainers.  
As the Apathetic Tools ecosystem grows, shared norms for collaboration, respect, and conflict resolution become essential — especially for open projects that welcome community participation.  
Rather than inventing custom language, the team wanted a **widely recognized, well-maintained template** that could be easily understood, translated, and enforced.

### Options Considered

| Option | Pros | Cons |
|--------|------|------|
| **Contributor Covenant 3.0** | ✅ Industry-standard and widely adopted<br>✅ Legally sound and CC BY-SA 4.0 licensed<br>✅ Clearly defines expectations, reporting, and enforcement<br>✅ Includes inclusive language and repair-focused approach | ⚠️ Template language can feel formal or corporate |
| **Custom in-house code** | ✅ Tailored tone and structure | ❌ Risk of omissions or unclear enforcement<br>❌ Higher maintenance burden |
| **No formal code** | ✅ Less administrative work | ❌ Unclear expectations<br>❌ Difficult to moderate conflicts fairly |

### Decision

Adopt the **Contributor Covenant 3.0** as the foundation for the project’s `CODE_OF_CONDUCT.md`, adapted for the Apathetic Tools community.  
This provides a **consistent, transparent behavioral framework** while avoiding the overhead of authoring and maintaining a custom code.  
It defines reporting, enforcement, and repair processes clearly, reinforcing the community’s emphasis on accountability and respect.  

This version is lightly customized with local contact details and references to community moderation procedures, maintaining alignment with upstream guidance.

<br/><br/>

---
---

<br/><br/>

## 🧭 Choose `Python` as the Initial Implementation Language (Superseded)
<a id="dec03"></a>*DEC 03 — 2025-10-09*  

*(Note: This decision was superseded by [DEC 13](#dec13) switching the project to Go.)*

### Context

The project aims to be a **lightweight, dependency-free build tool** that runs anywhere — Linux, macOS, Windows, or CI — without setup or compilation.  
Compiled languages (e.g. Go, Rust) would require distributing multiple binaries and would prevent in-place auditing and modification.
Python 3, by contrast, is preinstalled or easily available on all major platforms, balancing universality and maintainability.

### Decision

Implement the project initially in **Python 3**. Python provided zero-dependency execution for existing Python developers and transparent source code.

<br/><br/>

---
---

<br/><br/>


## ⚖️ Choose `MIT-aNOAI` License
<a id="dec02"></a>*DEC 02 — 2025-10-09*  

### Context

This project is meant to be open, modifiable, and educational — a tool for human developers.  
The ethics and legality of AI dataset collection are still evolving, and no reliable system for consent or attribution yet exists.

The project uses AI tools but distinguishes between **using AI** and **being used by AI** without consent.

### Options Considered

- **MIT License (standard)** — simple and permissive, but allows unrestricted AI scraping.
- **MIT + “No-AI Use” rider (MIT-aNOAI)** — preserves openness while prohibiting dataset inclusion or model training; untested legally and not OSI-certified.

### Decision

Adopt the **MIT-aNOAI license** — the standard MIT license plus an explicit clause banning AI/ML training or dataset inclusion.
This keeps the project open for human collaboration while defining clear ethical boundaries.

While this may deter adopters requiring OSI-certified licenses, it can later be dual-licensed if consent-based frameworks emerge.

### Ethical Consideration

AI helped create this project but does not own it.  
The license asserts consent as a prerequisite for training use — a small boundary while the wider ecosystem matures.


<br/><br/>

---
---

<br/><br/>



## 🤖 Use `AI Assistance` for Documentation and Development  
<a id="dec01"></a>*DEC 01 — 2025-10-09*


### Context

This project started as a small internal tool. Expanding it for public release required more documentation, CLI scaffolding, and testing than available time allowed.

AI tools (notably ChatGPT) offered a practical way to draft and refine code and documentation quickly, allowing maintainers to focus on design and correctness instead of boilerplate.

### Options Considered

- **Manual authoring** — complete control but slow and repetitive.
- **Static generators (pdoc, Sphinx)** — good for APIs, poor for narrative docs.
- **AI-assisted drafting** — fast, flexible, and guided by human review.

### Decision

Use **AI-assisted authoring** (e.g. ChatGPT) for documentation and boilerplate generation, with final edits and review by maintainers.  
This balances speed and quality with limited human resources. Effort can shift from writing boilerplate to improving design and clarity.  

AI use is disclosed in headers and footers as appropriate.

### Ethical Note

AI acts as a **paid assistant**, not a data harvester.  
Its role is pragmatic and transparent — used within clear limits while the ecosystem matures.


<br/><br/>

---
---

<br/><br/>

_Written following the [Apathetic Decisions Style v1](https://apathetic-recipes.github.io/decisions-md/v1) and [ADR](https://adr.github.io/), optimized for small, evolving projects._  
_This document records **why** we build things the way we do — not just **what** we built._

> ✨ *AI was used to help draft language, formatting, and code — plus we just love em dashes.*

<p align="center">
  <sub>😐 <a href="https://apathetic-tools.github.io/">Apathetic Tools</a> © <a href="./LICENSE">MIT-aNOAI</a></sub>
</p>
