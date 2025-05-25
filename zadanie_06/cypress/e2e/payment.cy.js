describe('Payment Tests', () => {
  beforeEach(() => {
    cy.visit('/', { timeout: 10000 })
    cy.get('[data-test-id="products-page"]', { timeout: 10000 }).should('be.visible')
  })

  it('should not show payment button when cart is empty', () => {
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.get('[data-test-id="cart-page"]').should('be.visible')
    cy.get('[data-test-id="cart-item"]').should('not.exist')
    cy.contains('Przejdź do płatności').should('not.exist')
  })

  it('should show payment button when cart has items', () => {
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.get('[data-test-id="cart-page"]').should('be.visible')
    cy.get('[data-test-id="cart-item"]').should('exist')
    cy.contains('Przejdź do płatności').should('be.visible')
  })

  it('should show MasterCard payment option', () => {
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.contains('Przejdź do płatności').click()
    cy.get('[data-test-id="payment-page"]').should('be.visible')
    cy.contains('MasterCard').should('be.visible')
  })

  it('should show Visa payment option', () => {
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.contains('Przejdź do płatności').click()
    cy.get('[data-test-id="payment-page"]').should('be.visible')
    cy.contains('Visa').should('be.visible')
  })

  it('should display error for missing first name', () => {
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.contains('Przejdź do płatności').click()
    cy.get('[data-test-id="payment-page"]').should('be.visible')
    cy.contains('MasterCard').click()
    cy.get('[data-test-id="card-input-lastname"]').type('Kowalski')
    cy.contains('Zapłać').click()
    cy.contains('Podaj imie').should('be.visible')
  })

  it('should display error for missing last name', () => {
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.contains('Przejdź do płatności').click()
    cy.get('[data-test-id="payment-page"]').should('be.visible')
    cy.contains('Visa').click()
    cy.get('[data-test-id="card-input-firstname"]').type('Jan')
    cy.contains('Zapłać').click()
    cy.contains('Podaj nazwisko').should('be.visible')
  })

  it('should display error for invalid card number', () => {
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.contains('Przejdź do płatności').click()
    cy.get('[data-test-id="payment-page"]').should('be.visible')
    cy.contains('MasterCard').click()
    cy.get('[data-test-id="card-input-firstname"]').type('Jan')
    cy.get('[data-test-id="card-input-lastname"]').type('Kowalski')
    cy.get('[data-test-id="card-input-number"]').type('1234')
    cy.contains('Zapłać').click()
    cy.contains('Nieprawidłowy numer karty').should('be.visible')
  })

  it('should display error for invalid CVV', () => {
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.contains('Przejdź do płatności').click()
    cy.get('[data-test-id="payment-page"]').should('be.visible')
    cy.contains('Visa').click()
    cy.get('[data-test-id="card-input-firstname"]').type('Jan')
    cy.get('[data-test-id="card-input-lastname"]').type('Kowalski')
    cy.get('[data-test-id="card-input-number"]').type('1111222233334444')
    cy.get('[data-test-id="card-input-expire"]').type('11/27')
    cy.get('[data-test-id="card-input-cvv"]').type('1')
    cy.contains('Zapłać').click()
    cy.contains('Nieprawidłowy kod CVV').should('be.visible')
  })

  it('should accept valid card details without errors', () => {
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.contains('Przejdź do płatności').click()
    cy.get('[data-test-id="payment-page"]').should('be.visible')
    cy.contains('MasterCard').click()
    cy.get('[data-test-id="card-input-firstname"]').type('Jan')
    cy.get('[data-test-id="card-input-lastname"]').type('Kowalski')
    cy.get('[data-test-id="card-input-number"]').type('1111222233334444')
    cy.get('[data-test-id="card-input-expire"]').type('11/27')
    cy.get('[data-test-id="card-input-cvv"]').type('123')
    cy.contains('Zapłać').click()
    cy.contains('Nieprawidłowy numer karty').should('not.exist')
    cy.contains('Nieprawidłowy kod CVV').should('not.exist')
    cy.contains('Nieprawidłowy format daty ważności').should('not.exist')
  })
})

