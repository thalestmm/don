_help:
    @echo ""
    @just --list

p msg="update" mode="chore":
    @git add .
    @git commit -m "{{ mode }}: {{ msg }}"
    @git push
