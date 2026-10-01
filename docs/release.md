# Release process

1. Use a clean tree and the pinned Go 1.27.0 toolchain.
2. Update [CHANGELOG.md](../CHANGELOG.md) and version guidance.
3. Run `golib check --all`; inspect advisory NilAway output.
4. Run the hosted matrix for OS/timezone and PostgreSQL 14–18.
5. Compare the generated API baseline and benchmark history.
6. Only after candidate gates pass, create a signed semantic version tag from
   main and publish verified source assets through the maintainer release
   process. This repository has no automatic release-publisher workflow. Bind
   the module ZIP, exact go.mod, SBOM, provenance, checksum manifest, and trusted
   checksum signature to the release source, then verify the actual public
   release and a clean public consumer. Preparing v2.0.0 is not publication.

No release proceeds with a surviving curated mutation, coverage below 100.0%,
an unexplained timezone corpus change, missing provenance, or a skipped blocking
gate.
