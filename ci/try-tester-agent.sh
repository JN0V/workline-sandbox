#!/bin/sh
# A stub agent for the try: an objection's answer keeps the test; else one
# test for the scenario "Two by three".
if grep -q "objection" ; then
  echo '- note: "The scenario says 6, the test checks 6 and nothing else: kept."'
  exit 0
fi
cat <<'Y'
- patch:
    file: tests/test_mul_two_by_three.sh
    content: |
      # Scenario: Two by three
      # Given the library, When I multiply 2 by 3, Then I get 6.
      . ./lib.sh
      [ "$(mul 2 3)" = 6 ]
Y
