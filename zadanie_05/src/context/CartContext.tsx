import {Product} from "../client/BackendClient.ts";
import {createContext, useContext, useState} from "react";

interface CartEntry {
  product: Product
  quantity: number
}

interface CartContextType {
  entries: CartEntry[]
  addToCart: (product: Product) => void
  totalPrice: () => number
}

const CartContext = createContext<CartContextType | undefined>(undefined)

export function CartProvider({children}: { children: React.ReactNode }) {
  const [entries, setEntries] = useState<CartEntry[]>([])

  const addToCart = (product: Product) => {
    setEntries(prevEntries => {
      const existingEntry = prevEntries.find(entry => entry.product.id === product.id)
      if (existingEntry) {
        return prevEntries.map(entry =>
          entry.product.id === product.id
            ? {...entry, quantity: entry.quantity + 1}
            : entry
        )
      }
      return [...prevEntries, {product, quantity: 1}]
    })
  }

  const totalPrice = () => {
    return entries.reduce((acc, entry) => {
      return acc + entry.product.Price * entry.quantity
    }, 0)
  }

  return (
    <CartContext.Provider value={{entries, addToCart, totalPrice}}>
      {children}
    </CartContext.Provider>
  )
}

export function useCart() {
  return useContext(CartContext)
}
