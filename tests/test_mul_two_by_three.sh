# Scenario: Two by three
# Given the library, When the clerk runs mul 2 3, Then it prints 6.
. ./lib.sh
[ "$(mul 2 3)" = 6 ]
