import {useEffect, useState} from "react";
import {BackendClient, Product} from "../client/BackendClient.ts";
import {useCart} from "../context/CartContext.tsx";

export function Products() {
    const [products, setProducts] = useState<Product[]>([])
    const {addToCart} = useCart()

    useEffect(() => {
        BackendClient.getProducts().then(response => setProducts(response.data))
    }, [])

    return (
        <ul data-test-id="products-page" className="flex flex-wrap h-full w-full justify-evenly items-center">
            {products.map(product =>
                <li key={product.id} data-test-id="product-card" className="flex flex-col h-1/5 w-1/4 mx-4 border-2 py-10 text-center">
                    <p data-test-id="product-name">Produkt: {product.Name}</p>
                    <p data-test-id="product-price">Cena: {product.Price}zł</p>
                    <button
                        data-test-id="add-to-cart-button"
                        className="w-1/2 mx-auto mt-2 border-2 button"
                        onClick={() => addToCart(product)}>Dodaj do koszyka
                    </button>
                </li>
            )}
        </ul>
    )
}
