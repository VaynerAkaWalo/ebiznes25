describe('Products Page', () => {
  beforeEach(() => {
    cy.visit('/', { timeout: 10000 })
    // Wait for the application to be fully loaded
    cy.get('[data-test-id="products-page"]', { timeout: 10000 }).should('be.visible')
  })

  it('should display list of products', () => {
    cy.get('[data-test-id="product-card"]').should('have.length.at.least', 1)
  })

  it('should display product details', () => {
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="product-name"]').should('be.visible')
      cy.get('[data-test-id="product-price"]').should('be.visible')
      cy.get('[data-test-id="add-to-cart-button"]').should('be.visible')
    })
  })
}) 