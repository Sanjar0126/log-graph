## Building the Go App

```bash
go build -o build/main main.go
```
Or get from releases

---

## Running the example

```bash
python3 example.py | build/pipe --config config.yaml
```

---

## View the Dashboard
```
http://localhost:8080
```
---

## Customize patterns
Edit config.yaml to define your own listeing port, log formats, colors, and axis titles.

```yaml
port: 8080 #listening port for gui
charts:
  - name: query 
    regex: "\\[query\\]:\\s*([\\d.]+)s" #regex for parsing stdout
    x_axis: Index # x axis title for chart
    y_axis: Query Time (s) # y axis title for chart
    color: "#ff6384" # color in hex
```

---