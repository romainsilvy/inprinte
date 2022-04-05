
import React from 'react';
import '../styles/Boutique.css'
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import { ShopProduct } from '../components/ShopProduct';
import loading from '../images/loading.svg'




class Boutique extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            items: [],
            DataisLoaded: false
        };
    }

    componentDidMount() {
        fetch("http://localhost:8080/boutique")
            .then((res) => res.json())
            .then((json) => {
                this.setState({
                    items: json,
                    DataisLoaded: true
                });
            })
    }

    render() {
        const { DataisLoaded, items } = this.state;
        if (!DataisLoaded) return (
            <div className="Boutique">
                <Navbar />
                <h1 className="shopTitle">
                    Shoppez par catégories
                </h1>
                <div className="allCategories">
                    <p className="category">cat1</p>
                    <p className="category">cat1</p>
                    <p className="category">cat1</p>
                    <p className="category">cat1</p>
                </div>
                <div className="allShopProducts">
                    <img className="loading" src={loading} alt="loading" />
                </div>
                <Footer />
            </div>  
            )
            

        return (
            <div className="Boutique">
                <Navbar />
                <h1 className="shopTitle">
                    Shoppez par catégories
                </h1>
                <div className="allCategories">
                    <p className="category">cat1</p>
                    <p className="category">cat1</p>
                    <p className="category">cat1</p>
                    <p className="category">cat1</p>
                </div>
                <div className="allShopProducts">
                {
                    items.AllProducts.map((item) => (
                        <ShopProduct product={item}/>
                    ))
                }
                </div>
                <Footer />
            </div> 
            
        )
    }
}

{/* <div className = "Boutique">
                <h1> Fetch data from an api in react </h1>  {
                    items.BoutiqueNews.map((item) => (
                        <ShopProduct product={item}/>
                    ))
                }
            </div> */}






{/* <div className="Boutique">
                <Navbar />
                <h1 className="shopTitle">
                    Shoppez par catégories
                </h1>
                <div className="allCategories">
                    <p className="category">cat1</p>
                    <p className="category">cat1</p>
                    <p className="category">cat1</p>
                    <p className="category">cat1</p>
                </div>
                <div className="allShopProducts">

                items.map((item) => ( 
                <ol key = { item.id } >
                    User_Name: { item.username }, 
                    Full_Name: { item.name }, 
                    User_Email: { item.email } 
                    </ol>
                ))
                    <ShopProduct product={product}/>
                    <ShopProduct/>
                    <ShopProduct/>
                    <ShopProduct/>
                </div>
                <Footer />
            </div> */}


export default Boutique;
//     let allProducts = fetch('http://localhost:8080/boutique')
//     .then((response) => {
//         return response.json()
//     })
//     .then((result) => {
//         allProducts = result.AllProducts
//         return allProducts
//     })


//     allProducts.then((a) => {
//         a.map(product => (
//             console.log(product.Name)
//         ))
//     })

//     return (
//         <div className="">
//             <Navbar />
//             <h1 className="shopTitle">
//                 Shoppez par catégories
//             </h1>
//             <div className="allCategories">
//                 <p className="category">cat1</p>
//                 <p className="category">cat1</p>
//                 <p className="category">cat1</p>
//                 <p className="category">cat1</p>
//             </div>
//             <div className="allShopProducts">

//                 {}
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>
//                 <ShopProduct/>

//             </div>
//             <Footer />
//         </div>
//     );
// }
