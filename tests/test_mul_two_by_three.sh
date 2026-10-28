# Scenario: Two by three
# Given the library, When I multiply 2 by 3, Then I get 6.
. ./lib.sh
[ "$(mul 2 3)" = 6 ]
