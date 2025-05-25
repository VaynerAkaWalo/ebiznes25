import "./styles.css"
import {Header} from "./component/Header.js";
import {Footer} from "./component/Footer.js";
import {PaymentPage} from "./component/PaymentPage.js";
import {Products} from "./component/Products.js";
import {Routes, Route} from "react-router-dom";
import {Cart} from "./component/Cart.js";

function App() {
    return (
        <>
            <div className="h-screen flex flex-col">
                <Header/>
                <div className="h-full flex justify-center items-center">
                    <Routes>
                        <Route path="/" element={<Products/>}/>
                        <Route path="/payment" element={<PaymentPage/>}/>
                        <Route path="/cart" element={<Cart/>}/>
                    </Routes>
                </div>
            </div>
            <Footer/>
        </>
    )
}

export default App
