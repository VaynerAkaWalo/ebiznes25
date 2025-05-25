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
        <ul className="flex flex-wrap h-full w-full justify-evenly items-center">
            {products.map(product =>
                <li key={product.id} className="flex flex-col h-1/5 w-1/4 mx-4 border-2 py-10 text-center">
                    <p>Produkt: {product.Name}</p>
                    <p>Cena: {product.Price}zł</p>
                    <button
                        className="w-1/2 mx-auto mt-2 border-2 button"
                        onClick={() => addToCart(product)}>Dodaj do koszyka
                    </button>
                </li>
            )}
        </ul>
    )
}
