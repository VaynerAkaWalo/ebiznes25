import "./styles.css"
import {Header} from "./component/Header.js";
import {Footer} from "./component/Footer.js";
import {PaymentPage} from "./component/PaymentPage.js";
import {Products} from "./component/Products.js";

function App() {
  return (
    <>
        <div className="h-screen flex flex-col">
            <Header/>
            <div className="h-full flex justify-center items-center">
                <Products/>
                {/*<PaymentPage/>*/}
            </div>
        </div>
        <Footer/>
    </>
  )
}

export default App
