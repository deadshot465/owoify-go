## 2.1.0
**This version changed some mappings in presets, so the same input probably won't return the same result as version 2.0.0.**

- Add more word mappings, so it's on par with mohan-cao's latest owoify.
- Add Uwuify and Uvuify for convenience.

## 2.0.0
- Update to Go 1.18
- **(Breaking change)** Use generic function to interleave arrays.
- **(Breaking change)** Rename `InterleaveArrays` to `InterleaveSlices`
- **(Breaking change)** Use `iota` to represent owoness levels instead of strings.
- Fix `go.mod`.
- Rename regexes to match Golang's naming convention.

## 1.0.1
- Add summary.
- Update README.md.