#!/bin/sh
# The try of workline's ADR-0051 step 4: an agent that fixes
# docs/tech/tokens.md and moves its `checked` to the commit the task names.
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
    -checked: $old
    +checked: $at
    +verified: agent:documentalist
     ---
    @@ -9,2 +10,2 @@

    -Access tokens last one hour.
    +Access tokens last two hours.
- claim:
    doc: docs/tech/tokens.md
    lines: "10"
    status: contradicted
    source: {path: src/auth/token.go, quote: "const TokenTTL = 7200"}
YAML
