#!/bin/sh
# run-tests.sh <report>: each tests/test_*.sh a test; exit 1 if one fails.
report=$1 failed=0
mkdir -p "$(dirname "$report")"
{
  echo '<?xml version="1.0" encoding="UTF-8"?>'
  echo '<testsuites><testsuite name="tests">'
  for t in tests/test_*.sh; do
    name=$(basename "$t" .sh)
    if sh "$t" >/dev/null 2>&1; then
      echo "<testcase classname=\"tests\" name=\"$name\" time=\"0.01\"/>"
    else
      failed=1
      echo "<testcase classname=\"tests\" name=\"$name\" time=\"0.01\"><failure message=\"$t failed\"/></testcase>"
    fi
  done
  echo '</testsuite></testsuites>'
} > "$report"
exit $failed
