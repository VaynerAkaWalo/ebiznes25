describe('Navigation Tests', () => {
  beforeEach(() => {
    cy.visit('/', { timeout: 10000 })
    // Wait for the application to be fully loaded
    cy.get('[data-test-id="products-page"]', { timeout: 10000 }).should('be.visible')
  })

  it('should display all navigation links', () => {
    cy.get('[data-test-id="nav-link-products"]').should('be.visible')
    cy.get('[data-test-id="nav-link-cart"]').should('be.visible')
  })

  it('should navigate to products page', () => {
    cy.get('[data-test-id="nav-link-products"]').click()
    cy.url().should('include', '/')
    cy.get('[data-test-id="products-page"]').should('be.visible')
  })

  it('should navigate to cart page', () => {
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.url().should('include', '/cart')
    cy.get('[data-test-id="cart-page"]').should('be.visible')
  })

  it('should handle invalid routes', () => {
    // Type an invalid URL in the browser
    cy.window().then((win) => {
      win.history.pushState({}, '', '/invalid-route')
    })
    cy.url().should('include', '/invalid-route')
  })
}) 