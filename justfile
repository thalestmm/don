_help:
  @echo ""
  @just --list

p msg="chore: update":
  @git add .
  @git commit -m "{{ msg }}"
  @git push
