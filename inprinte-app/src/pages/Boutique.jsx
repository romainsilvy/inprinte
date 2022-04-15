import React from 'react';
import '../styles/Boutique.css'
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import { ShopProduct } from '../components/ShopProduct';
import loading from '../images/loading.svg'
import bar2 from '../images/bar2.svg'



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
            <div className="B-Boutique">
                <Navbar />
                <img className="B-loading" src={loading} alt="loading" />
            </div>
        )

        return (
            <div className="B-Boutique">
                <Navbar />
                <div className="B-contentTree">
                    <div className="B-titleCat">
                        <h1>Les nouveautés</h1>
                        <img className="B-bar2" src={bar2} alt="bar" />
                    </div>
                    <div className="B-cart1">
                        <h2 className="pink">{items.BoutiqueNews[0].Name}</h2>
                        <img className="B-pictureCart" src={items.BoutiqueNews[0].Picture} alt="image" />
                        <div className="B-infos">
                            <p className="B-price">{items.BoutiqueNews[0].Price}€</p>
                        </div>
                    </div>
                    <div className="B-cart2">
                        <h2 className="purple">{items.BoutiqueNews[1].Name}</h2>
                        <img className="B-pictureCart" src={items.BoutiqueNews[1].Picture} alt="image" />
                        <div className="B-infos">
                            <p className="B-price">{items.BoutiqueNews[1].Price}€</p>
                        </div>
                    </div>
                    {/* <div className="B-cart3">
                        <h2 className="blue">{items.BoutiqueNews[2].Name}</h2>
                        <img className="B-pictureCart" src={items.BoutiqueNews[2].Picture} alt="image" />
                        <div className="B-infos">
                            <p className="B-price">{items.BoutiqueNews[2].Price}€</p>
                        </div>
                        </div> */}
                    </div>
                <div className="B-contentTwo">
                <div className="B-titleCat">
                        <h1>Les nouveautés</h1>
                        <img className="B-bar2" src={bar2} alt="bar" />
                    </div>
                    <div className="B-cart1">
                        <h2 className="pink">{items.BoutiqueMostWanted[0].Name}</h2>
                        <img className="B-pictureCart" src={items.BoutiqueMostWanted[0].Picture} alt="image" />
                        <div className="B-infos">
                            <p className="B-price">{items.BoutiqueMostWanted[0].Price}€</p>
                        </div>
                    </div>
                    {/* <div className="B-cart2">
                        <h2 className="purple">{items.BoutiqueMostWanted[1].Name}</h2>
                        <img className="B-pictureCart" src={items.BoutiqueMostWanted[1].Picture} alt="image" />
                        <div className="B-infos">
                            <p className="B-price">{items.BoutiqueMostWanted[1].Price}€</p>
                        </div>
                    </div>
                    <div className="B-cart3">
                        <h2 className="blue">{items.BoutiqueMostWanted[2].Name}</h2>
                        <img className="B-pictureCart" src={items.BoutiqueNews[2].Picture} alt="image" />
                        <div className="B-infos">
                            <p className="B-price">{items.BoutiqueMostWanted[2].Price}€</p>
                        </div>
                    </div> */}
                </div>
                <h1 className="B-shopTitle">
                    Shoppez par catégories
                </h1>
                <div className="B-allCategories">
                    {
                        items.Categories.map((item) => (
                            <p className="B-category" onClick={() => {
                                console.log(item)
                            }}>{item}</p>
                        ))
                    }
                </div>
                <div className="B-allShopProducts">
                    {
                        items.AllProducts.map((item) => (
                            <ShopProduct product={item} />
                        ))
                    }
                </div>
                <Footer />
            </div>

        )
    }
}

export default Boutique;
