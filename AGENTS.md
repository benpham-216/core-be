# Repository instructions

## Governance

- Make changes on a working branch and submit them through a pull request.
- Never merge or push directly to `main`.
- Preserve the control-plane boundary: domain code must not import HTTP, SQL, model provider, tool, Git, or sandbox SDKs.
- Agents cannot self-approve, self-grant permissions, merge, deploy, or treat a failed validation as success.

## Required verification

Run `go test ./...` and `go vet ./...` before requesting review. Add a regression test for every corrected lifecycle or policy defect.
