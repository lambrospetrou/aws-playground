import {expect} from "chai"
import got from "got"

import {baseUrl} from "./apisetup.mjs"

describe("Hello World Test", function() {
  // Extend the timeout since locally SAM takes time.
  this.timeout(5000)

  it("receives the hello world JSON text", async () => {
    const response = await got(`${baseUrl}/`)
    expect(response.headers["Content-Type"] ?? response.headers["content-type"]).to.eql("application/json")
    expect(response.body).to.eql(`"Hello, world!"`)
  })
})
