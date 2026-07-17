#!/usr/bin/env bash
set -euo pipefail

required=(
	README.md CHANGELOG.md CONTRIBUTING.md SECURITY.md LICENSE
	docs/README.md docs/quickstart.md docs/api.md docs/arithmetic.md
	docs/gregorian-iso-vectors.md docs/timezone.md docs/timezone-corpus.md
	docs/business.md docs/exclusive-ranges.md docs/encoding.md
	docs/postgres.md docs/adapters.md docs/composition.md
	docs/holiday-datasets.md docs/versioning.md docs/carbon-migration.md
	docs/security.md docs/operations.md docs/performance.md
	docs/compatibility.md docs/troubleshooting.md docs/faq.md
	docs/hardening.md docs/release.md
)
for file in "${required[@]}"; do
	if [[ ! -s "$file" ]]; then
		printf 'required documentation is missing or empty: %s\n' "$file" >&2
		exit 1
	fi
done

ruby - <<'RUBY'
Dir.glob('**/*.md').each do |document|
  content = File.read(document, encoding: 'UTF-8')
  fenced = false
  prose = content.lines.filter_map do |line|
    if line.lstrip.start_with?('```', '~~~')
      fenced = !fenced
      nil
    elsif !fenced
      line
    end
  end.join
  prose.scan(/\[[^\]]*\]\(([^)]+)\)/).flatten.each do |target|
    next if target.start_with?('http://', 'https://', 'mailto:', '#')
    relative = target.split('#', 2).first
    next if relative.empty?
    path = File.expand_path(relative, File.dirname(document))
    abort("broken relative link in #{document}: #{target}") unless File.exist?(path)
  end
end
puts 'documentation links resolve'
RUBY

go test ./... -run '^Example'
