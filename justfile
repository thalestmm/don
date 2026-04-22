_help:
	@just --list

watch:
    @air -c config/air.toml

push message="chore: update":
    @git add .
    @git commit -m "{{message}}"
    @git push
