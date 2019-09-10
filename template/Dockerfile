{{ if DockerHub }}FROM scratch

COPY {{Name}} /bin/{{Name}}

ENTRYPOINT ["/bin/{{Name}}"]
{{ end }}