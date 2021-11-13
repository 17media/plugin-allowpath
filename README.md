# Allow Path

Allow Path is a middleware plugin for [Traefik](https://github.com/traefik/traefik) which sends an HTTP `403 Forbidden`
response when the requested HTTP path does not matches one of the configured [regular expressions](https://github.com/google/re2/wiki/Syntax).

```deployment.yaml
command:
  - "--experimental.plugins.allowpath.modulename=github.com/17media/plugin-allowpath"
  - "--experimental.plugins.allowpath.version=v0.1.0"
```

```middleware.yaml
# Allow path prefix with /foo and exact match with /bar only
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: allowpath
spec:
  plugin:
    allowpath:
      regex:
      - ^/foo(.*)
      - ^/bar$
```
