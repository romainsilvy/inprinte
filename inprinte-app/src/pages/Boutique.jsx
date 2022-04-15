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
            <div className="Boutique">
                <Navbar />
                <img className="loading" src={loading} alt="loading" />
            </div>
        )

        return (
            <div className="Boutique">
                <Navbar />
                <div className="firstcontent">
                    <div className="titleCat">
                        <h1>Nouveauté</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                    <div className="cartP1">
                        <h2>{items.BoutiqueNews[0].Name}</h2>
                        <img className="pictureCart" src={items.BoutiqueNews[0].Picture} alt="image" />
                        <div className="infos">
                            <p clasName="price">{items.BoutiqueNews[0].Price}€</p>
                        </div>
                    </div>
                    <div className="cartP2">
                        <h2>{items.BoutiqueNews[1].Name}</h2>
                        <img className="pictureCart" src={items.BoutiqueNews[1].Picture} alt="image" />
                        <div className="infos">
                            <p clasName="price">{items.BoutiqueNews[1].Price}€</p>
                        </div>
                    </div>
                    <div className="cartP3">
                        <h2>Name</h2>
                        <img className="pictureCart" src="" alt="image" />
                        <div className="infos">
                            <p clasName="price">25€</p>
                        </div>
                    </div>
                    <div className="cartP4">
                        <h2>Name</h2>
                        <img className="pictureCart" src="" alt="image" />
                        <div className="infos">
                            <p clasName="price">25€</p>
                        </div>
                    </div>
                </div>
                <div className="contentTwo">
                    <div className="titleCat2">
                        <h1>Les plus demandés</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                    <div className="cardTop2">
                        <h2>{items.BoutiqueMostWanted[0].Name}</h2>
                        <img className="pictureCardTop2" src={items.BoutiqueMostWanted[0].Picture} alt="image" />
                        <div className="infos">
                            <p clasName="price">{items.BoutiqueMostWanted[0].Price}€</p>
                        </div>
                    </div>
                    <div className="cardTop1">
                    <h2>Name</h2>
                        <img className="pictureCardTop1" src="" alt="image" />
                        <div className="infos">
                            <p clasName="price">25€</p>
                        </div>
                    </div>
                    <div className="cardTop3">
                    <h2>Name</h2>
                        <img className="pictureCardTop3" src="" alt="image" />
                        <div className="infos">
                            <p clasName="price">25€</p>
                        </div>
                    </div>
                </div>
                <h1 className="shopTitle">
                    Shoppez par catégories
                </h1>
                <div className="allCategories">
                    {
                        items.Categories.map((item) => (
                            <p className="category" onClick={() => {
                                console.log(item)
                            }}>{item}</p>
                        ))
                    }
                </div>
                <div className="allShopProducts">
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
