describe('Cart Tests', () => {
  beforeEach(() => {
    cy.visit('/', { timeout: 10000 })
    // Wait for the application to be fully loaded
    cy.get('[data-test-id="products-page"]', { timeout: 10000 }).should('be.visible')
  })

  it('should show empty cart with total 0 when no products added', () => {
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.get('[data-test-id="cart-page"]').should('be.visible')
    cy.get('[data-test-id="cart-item"]').should('not.exist')
    cy.get('[data-test-id="cart-total-price"]').should('contain', '0zł')
  })

  it('should add product to cart', () => {
    // Get product details
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="product-name"]').invoke('text').then(text => {
        cy.wrap(text.replace('Produkt: ', '')).as('productName')
      })
      cy.get('[data-test-id="product-price"]').invoke('text').then(text => {
        cy.wrap(text.replace('Cena: ', '').replace('zł', '')).as('productPrice')
      })
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })

    // Go to cart and verify
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.get('[data-test-id="cart-page"]').should('be.visible')

    // Check cart contents
    cy.get('@productName').then(productName => {
      cy.get('[data-test-id="cart-item-name"]').should('contain', productName)
    })
    cy.get('@productPrice').then(productPrice => {
      cy.get('[data-test-id="cart-item-price"]').should('contain', productPrice)
    })
    cy.get('[data-test-id="cart-item-quantity"]').should('contain', '1')
  })

  it('should add two products to cart and show correct total', () => {
    // Get first product details
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="product-name"]').invoke('text').then(text => {
        cy.wrap(text.replace('Produkt: ', '')).as('firstProductName')
      })
      cy.get('[data-test-id="product-price"]').invoke('text').then(text => {
        cy.wrap(text.replace('Cena: ', '').replace('zł', '')).as('firstProductPrice')
      })
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })

    // Get second product details
    cy.get('[data-test-id="product-card"]').eq(1).within(() => {
      cy.get('[data-test-id="product-name"]').invoke('text').then(text => {
        cy.wrap(text.replace('Produkt: ', '')).as('secondProductName')
      })
      cy.get('[data-test-id="product-price"]').invoke('text').then(text => {
        cy.wrap(text.replace('Cena: ', '').replace('zł', '')).as('secondProductPrice')
      })
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })

    // Go to cart and verify
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.get('[data-test-id="cart-page"]').should('be.visible')

    // Check if both products are in cart
    cy.get('[data-test-id="cart-item"]').should('have.length', 2)
    
    // Verify first product
    cy.get('@firstProductName').then(productName => {
      cy.get('[data-test-id="cart-item-name"]').first().should('contain', productName)
    })
    cy.get('@firstProductPrice').then(productPrice => {
      cy.get('[data-test-id="cart-item-price"]').first().should('contain', productPrice)
    })

    // Verify second product
    cy.get('@secondProductName').then(productName => {
      cy.get('[data-test-id="cart-item-name"]').last().should('contain', productName)
    })
    cy.get('@secondProductPrice').then(productPrice => {
      cy.get('[data-test-id="cart-item-price"]').last().should('contain', productPrice)
    })

    // Calculate and verify total
    cy.get('@firstProductPrice').then(firstPrice => {
      cy.get('@secondProductPrice').then(secondPrice => {
        const total = Number(firstPrice) + Number(secondPrice)
        cy.get('[data-test-id="cart-total-price"]').should('contain', `${total}zł`)
      })
    })
  })

  it('should add same product multiple times and update quantity and total', () => {
    // Get product details
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="product-name"]').invoke('text').then(text => {
        cy.wrap(text.replace('Produkt: ', '')).as('productName')
      })
      cy.get('[data-test-id="product-price"]').invoke('text').then(text => {
        cy.wrap(text.replace('Cena: ', '').replace('zł', '')).as('productPrice')
      })
      // Add product 3 times
      cy.get('[data-test-id="add-to-cart-button"]').click()
      cy.get('[data-test-id="add-to-cart-button"]').click()
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })

    // Go to cart and verify
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.get('[data-test-id="cart-page"]').should('be.visible')

    // Check if only one item is in cart
    cy.get('[data-test-id="cart-item"]').should('have.length', 1)

    // Verify product details
    cy.get('@productName').then(productName => {
      cy.get('[data-test-id="cart-item-name"]').should('contain', productName)
    })
    
    // Verify quantity is 3
    cy.get('[data-test-id="cart-item-quantity"]').should('contain', '3')

    // Verify total price
    cy.get('@productPrice').then(price => {
      const total = Number(price) * 3
      cy.get('[data-test-id="cart-total-price"]').should('contain', `${total}zł`)
    })
  })

  it('should maintain cart state when navigating between pages', () => {
    // Add first product
    cy.get('[data-test-id="product-card"]').first().within(() => {
      cy.get('[data-test-id="product-name"]').invoke('text').then(text => {
        cy.wrap(text.replace('Produkt: ', '')).as('firstProductName')
      })
      cy.get('[data-test-id="product-price"]').invoke('text').then(text => {
        cy.wrap(text.replace('Cena: ', '').replace('zł', '')).as('firstProductPrice')
      })
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })

    // Go to cart and verify first product
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.get('[data-test-id="cart-page"]').should('be.visible')
    cy.get('[data-test-id="cart-item"]').should('have.length', 1)
    cy.get('@firstProductName').then(productName => {
      cy.get('[data-test-id="cart-item-name"]').should('contain', productName)
    })

    // Go back to products
    cy.get('[data-test-id="nav-link-products"]').click()
    cy.get('[data-test-id="products-page"]').should('be.visible')

    // Add second product
    cy.get('[data-test-id="product-card"]').eq(1).within(() => {
      cy.get('[data-test-id="product-name"]').invoke('text').then(text => {
        cy.wrap(text.replace('Produkt: ', '')).as('secondProductName')
      })
      cy.get('[data-test-id="product-price"]').invoke('text').then(text => {
        cy.wrap(text.replace('Cena: ', '').replace('zł', '')).as('secondProductPrice')
      })
      cy.get('[data-test-id="add-to-cart-button"]').click()
    })

    // Go back to cart and verify both products
    cy.get('[data-test-id="nav-link-cart"]').click()
    cy.get('[data-test-id="cart-page"]').should('be.visible')
    
    // Verify both products are in cart
    cy.get('[data-test-id="cart-item"]').should('have.length', 2)
    
    // Verify first product is still there
    cy.get('@firstProductName').then(productName => {
      cy.get('[data-test-id="cart-item-name"]').first().should('contain', productName)
    })
    
    // Verify second product was added
    cy.get('@secondProductName').then(productName => {
      cy.get('[data-test-id="cart-item-name"]').last().should('contain', productName)
    })

    // Verify total price
    cy.get('@firstProductPrice').then(firstPrice => {
      cy.get('@secondProductPrice').then(secondPrice => {
        const total = Number(firstPrice) + Number(secondPrice)
        cy.get('[data-test-id="cart-total-price"]').should('contain', `${total}zł`)
      })
    })
  })
}) 