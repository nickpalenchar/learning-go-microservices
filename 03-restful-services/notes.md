### Takeaways

- `io.Writer` is implemented by `http.ResponseWriter` in request functions
- Any func that takes an `io.Writer` can take a `http.ResponseWriter`
- the `"encoding/json"` package is for working with all things JSON
- Using a json `Encode` and passing a `http.ResponseWriter`, you can easily encode responses in json.

### Definitions

**struct tags** are special annotations on a type of struct which can be picked up by various parsers. It
can change the appearance of data in a given struct for serialization/compatibility/etc
