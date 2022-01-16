describe("Hello World Test", () => {
  it("receives the hello world JSON text", () => {
    cy.request("/").as("response")

    cy.get("@response").should((response) => {
      expect(response.headers["content-type"]).to.eql("application/json")
      expect(response.body).to.eql("Hello, world!")
    })
  })
})
