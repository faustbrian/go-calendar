.PHONY: provenance

provenance:
	test -s docs/holiday-datasets.md
	@if [ -d datasets ] && find datasets -type f -print -quit | grep -q .; then \
		echo 'bundled datasets require a dedicated deterministic provenance verifier' >&2; \
		exit 1; \
	fi
	grep -qF 'No holiday dataset is bundled' docs/holiday-datasets.md
