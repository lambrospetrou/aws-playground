import {expect} from "chai"
import {handler} from "../../src/fn-apimain.mjs"

describe("Main Function Handler", function() {
  it("returns the hello world JSON response", async () => {
    const response = await handler({}, {functionName: "dummy"})
    expect(response.headers["Content-Type"]).to.eql("application/json")
    expect(response.body).to.eql(`"Hello, world!"`)
  })
})
