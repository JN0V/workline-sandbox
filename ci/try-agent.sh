#!/bin/sh
# The try of workline's ADR-0051 step 4: an agent that adds one line to
# docs/tech/tokens.md and records `judged` at the commit the task names.
# Its patch must go in a comment, never on the branch.
prompt=$(cat)
old=$(printf '%s\n' "$prompt" | sed -n 's/^ *4 | checked: //p' | head -1)
src=$(printf '%s\n' "$prompt" | sed -n 's/^ *3 | sources: //p' | head -1)
at=$(printf '%s\n' "$prompt" | sed -n 's/.*`judged: \([0-9a-f]*\)`.*/\1/p' | head -1)
cat <<YAML
- patch: |
    --- a/docs/tech/tokens.md
    +++ b/docs/tech/tokens.md
    @@ -3,3 +3,4 @@
     sources: $src
     checked: $old
    +judged: $at
     ---
    @@ -9,1 +10,2 @@
     Access tokens last one hour.
    +The lifetime is TokenTTL, in seconds.
- note: "The lifetime's value changed: a person confirms the doc's sentence."
YAML
