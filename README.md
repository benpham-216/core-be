# core-be

Versioned, reusable backend assets for multiple tech stacks. This repository is **not** the Agentic Flow Studio control-plane runtime.

## Layout

`stacks/<stack>/<asset-kind>/<asset-slug>/` holds executable source. Every asset has `asset.yaml`, README, tests, examples and release metadata. Cross-stack interaction is limited to versioned contracts in `cross-stack/contracts/`.

The generated catalog uses stable asset IDs for discovery and provenance.