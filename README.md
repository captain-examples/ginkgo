# Getting Captain working with Ginkgo

1. 🧪 Ensure Ginkgo produces JSON output

`ginkgo --json-report ./report.json ./...` will produce Captain-compatible JSON output.

2. 🔐 Create an Access Token

Create an Access Token for your organization within [Captain][captain] ([more documentation here][create-access-token]).

Add the new token as an action secret to your repository. Conventionally, we call this secret `RWX_ACCESS_TOKEN`.

3. 💌 Install the Captain CLI, and then call it when running tests

See the [full documentation on test suite integration][test-suite-integration].

```yaml
  - run: |
      captain run --suite-id captain-examples-ginkgo --test-results report.json -- \
        ginkgo --keep-going --json-report ./report.json ./...
    env:
      RWX_ACCESS_TOKEN: ${{ secrets.RWX_ACCESS_TOKEN }}
```

4. 🎉 See your test results in Captain!

Take a look at the [final workflow!][workflow-with-captain]

[captain]: https://account.rwx.com/deep_link/manage/access_tokens
[create-access-token]: https://www.rwx.com/docs/access-tokens
[workflow-with-captain]: https://github.com/captain-examples/ginkgo/blob/main/.github/workflows/ci.yml
[test-suite-integration]: https://www.rwx.com/captain/docs/test-suite-integration
