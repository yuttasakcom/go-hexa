# Go Hexa

## Hexagonal Architecture or Ports and Adapters Architecture

![Go Hexa](https://github.com/yuttasakcom/go-hexa/blob/master/screenshots/go-hexa.png)

## Liveness Probe

```yaml
livenessProbe:
  exec:
    command:
      - cat
      - /tmp/live
```

## Readiness Probe

```yaml
readinessProbe:
  httpGet:
    path: /health
    port: 80
```
